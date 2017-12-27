package edit

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"frizz.io/edit/assets"
	"frizz.io/edit/common"

	"time"

	"net"

	pkg_path "path"

	"bytes"
	"go/ast"
	"go/parser"
	"go/token"

	"encoding/gob"

	"io/ioutil"

	"path/filepath"

	"mime"

	gobuild "go/build"

	"os"

	"encoding/json"

	"frizz.io/config"
	"frizz.io/edit/auther"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/leemcloughlin/gofarmhash"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpgzip"
	"golang.org/x/tools/go/loader"
)

const writeTimeout = time.Second * 2

func Open(ctx context.Context, cancel context.CancelFunc, env vos.Env, out io.Writer, path string) error {

	fail := make(chan error)

	auth := auther.New()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch {
		case strings.HasSuffix(req.URL.Path, "/favicon.ico"):
			if err := serveStatic("favicon.ico", w, req); err != nil {
				fail <- err
			}
		case config.DEV && (req.URL.Path == "/static/bootstrap.js" || req.URL.Path == "/static/bootstrap.js.map"):
			// Only generate bootstrap dynamically in debug mode. Non debug mode delivers
			// bootstrap.js from static files (minified, without source map).
			if err := script(ctx, env, w, req.URL.Path == "/static/bootstrap.js.map"); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/static/"):
			if err := static(w, req); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/blob/"):
			if err := blob(ctx, env, auth, w, req); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/src/"):
			if err := src(ctx, env, w, req); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/data/"):
			if err := data(ctx, env, w, req); err != nil {
				fail <- err
			}
		default:
			if err := root(ctx, auth, w, req); err != nil {
				fail <- err
			}
		}
	})

	//if err := rpc.Register(&Server{ctx: ctx, auth: auth}); err != nil {
	//	return errors.WithStack(err)
	//}

	//http.Handle("/_rpc", websocket.Handler(func(ws *websocket.Conn) {
	//	ws.PayloadType = websocket.BinaryFrame
	//	rpc.ServeConn(ws)
	//}))

	go func() {
		if err := serve(ctx, path, out); err != nil {
			fail <- err
		}
	}()

	done := ctx.Done()

	for {
		select {
		case <-done:
			fmt.Fprintln(out, "Exiting editor server (interupted)... ")
			return nil
		case err, open := <-fail:
			if !open {
				// Channel has been closed, so app should gracefully exit.
				cancel()
				fmt.Fprintln(out, "Exiting editor server (finished)... ")
			} else {
				fmt.Fprintln(out, err)
			}
		}
	}
}

func sourceFilter(fpath string) bool {
	return strings.HasSuffix(fpath, ".go") && !strings.HasSuffix(fpath, "_test.go")
}

func dataFilter(fpath string) bool {
	return strings.HasSuffix(fpath, ".frizz.json")
}

func src(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request) error {
	pkg := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/src/"), ".bin")
	err := bundle(ctx, env, w, req, pkg, sourceFilter)
	if err != nil {
		return err
	}
	return nil
}

func data(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request) error {
	pkg := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/data/"), ".bin")
	err := bundle(ctx, env, w, req, pkg, dataFilter)
	if err != nil {
		return err
	}
	return nil
}

func bundle(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request, path string, filter func(string) bool) error {

	hashFromRequest, err := strconv.ParseUint(req.URL.Query().Get("hash"), 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing hash for %s", req.URL), 500)
		return errors.WithStack(err)
	}

	dir, err := patsy.Dir(env, path)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting dir for %s", path), 500)
		return err
	}

	bundle, err := getBundle(dir, filter, false)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting bundle for %s", path), 500)
		return errors.WithStack(err)
	}

	if bundle.Hash != hashFromRequest {
		http.Error(w, fmt.Sprintf("incorrect hash %d for %s", hashFromRequest, path), 401)
		return fmt.Errorf("incorrect hash %d for %s", hashFromRequest, path)
	}

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(bundle); err != nil {
		http.Error(w, fmt.Sprintf("error encoding buddle for %s", path), 500)
		return errors.WithStack(err)
	}

	w.Header().Set("Cache-Control", "max-age=31536000")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := io.Copy(w, buf); err != nil {
		http.Error(w, fmt.Sprintf("error streaming bundle for %s", path), 500)
		return errors.WithStack(err)
	}

	return nil
}

func getBundle(dir string, filter func(string) bool, justhash bool) (*common.Bundle, error) {
	fia, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	files := make(map[string][]byte)
	for _, fi := range fia {
		if !filter(fi.Name()) {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(dir, fi.Name()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		files[fi.Name()] = b
	}
	hash, err := hashFiles(files)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if justhash {
		return &common.Bundle{Hash: hash}, nil
	}
	return &common.Bundle{Hash: hash, Files: files}, nil
}

func blob(ctx context.Context, env vos.Env, auth auther.Auther, w http.ResponseWriter, req *http.Request) error {

	path := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/blob/"), ".bin")

	a, err := common.DecodeAuth(req.URL.Query().Get("auth"))
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding auth for %s", path), 500)
		return errors.WithStack(err)
	}
	if !auth.Auth(a.Id, a.Hash) {
		http.Error(w, fmt.Sprintf("auth error for %s", path), 401)
		return fmt.Errorf("auth error for %s", path)
	}

	dir, err := patsy.Dir(env, path)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting dir for for %s", path), 500)
		return err
	}

	bundle, err := getBundle(dir, dataFilter, true)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting bundle for %s", path), 500)
		return errors.WithStack(err)
	}

	blob := &common.Blob{
		Hash:   bundle.Hash,
		Source: map[string]uint64{},
	}

	conf := loader.Config{}
	conf.Import(path)
	conf.Import("frizz.io/edit/editor")
	conf.ParserMode = parser.ImportsOnly
	conf.Build = func() *gobuild.Context { c := gobuild.Default; return &c }() // make a copy of gobuild.Default
	conf.Build.GOPATH = env.Getenv("GOPATH")
	conf.Build.BuildTags = []string{"js"}
	conf.AllowErrors = true
	conf.TypeChecker.Error = func(e error) {}
	loaded, err := conf.Load()
	if err != nil {
		http.Error(w, fmt.Sprintf("error loading code for %s", path), 500)
		return errors.WithStack(err)
	}
	for p, pi := range loaded.AllPackages {

		// Some system packages don't have any files (e.g. unsafe) - we can skip them.
		if len(pi.Files) == 0 {
			continue
		}

		// Packages under GOROOT are in the standard library - we skip them
		dir, _ := filepath.Split(loaded.Fset.File(pi.Files[0].Pos()).Name())
		if strings.HasPrefix(dir, conf.Build.GOROOT) {
			blob.Lib = append(blob.Lib, p.Path())
			continue
		}

		bundle, err := getBundle(dir, sourceFilter, true)
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting bundle hash for %s", p.Path()), 500)
			return errors.WithStack(err)
		}

		blob.Source[p.Path()] = bundle.Hash

	}

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(blob); err != nil {
		http.Error(w, fmt.Sprintf("error encoding blob for %s", path), 500)
		return errors.WithStack(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := io.Copy(w, buf); err != nil {
		http.Error(w, fmt.Sprintf("error streaming blob for %s", path), 500)
		return err
	}

	return nil
}

func hashFiles(in map[string][]byte) (uint64, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	return farmhash.Hash64(b), nil
}

func static(w http.ResponseWriter, req *http.Request) error {
	return serveStatic(strings.TrimPrefix(req.URL.Path, "/static/"), w, req)
}

func serveStatic(name string, w http.ResponseWriter, req *http.Request) error {
	var file http.File
	var err error
	file, err = assets.Assets.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			// Special case: in /static/pkg/ we don't want 404 errors because we can't stop them from
			// popping up in the js console. Instead, deiver a 200 with a zero lenth body.
			if strings.HasPrefix(req.URL.Path, "/static/pkg/") {
				if err := writeWithTimeout(w, []byte{}); err != nil {
					return err
				}
				return nil
			}
			http.NotFound(w, req)
			return nil
		}
		http.Error(w, fmt.Sprintf("error opening %s", name), 500)
		return nil
	}
	defer file.Close()

	w.Header().Set("Cache-Control", "max-age=31536000")
	w.Header().Set("Content-Type", mime.TypeByExtension(pkg_path.Ext(req.URL.Path)))

	_, noCompress := file.(httpgzip.NotWorthGzipCompressing)
	gzb, isGzb := file.(httpgzip.GzipByter)

	if isGzb && !noCompress && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		if err := writeWithTimeout(w, gzb.GzipBytes()); err != nil {
			http.Error(w, fmt.Sprintf("error streaming gzipped %s", name), 500)
			return err
		}
	} else {
		if err := streamWithTimeout(w, file); err != nil {
			http.Error(w, fmt.Sprintf("error streaming %s", name), 500)
			return err
		}
	}
	return nil

}

func root(ctx context.Context, auth auther.Auther, w http.ResponseWriter, req *http.Request) error {

	id := make([]byte, 64)
	if _, err := rand.Read(id); err != nil {
		http.Error(w, "error reading rand", 500)
		return errors.WithStack(err)
	}
	attribute, err := common.Auth{Id: id, Hash: auth.Sign(id)}.Encode()
	if err != nil {
		http.Error(w, "error encoding auth", 500)
		return errors.WithStack(err)
	}

	packagePath := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/"), "/")

	source := []byte(`
		<html>
			<head>
				<meta charset="utf-8">
				<script src="/static/jquery-2.2.4.min.js"></script>
				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
			</head>
			<body id="body" auth="` + attribute + `">
				<input type="text" id="path" value="` + packagePath + `">
				<button id="load">Load</button>
				<span id="log">Loading...</span>
				<div id="editor"></div>
			</body>
			<script src="/static/bootstrap.js"></script>
		</html>`)

	if err := writeWithTimeout(w, source); err != nil {
		http.Error(w, "error streaming homepage", 500)
		return err
	}

	return nil
}

func streamWithTimeout(w io.Writer, r io.Reader) error {
	c := make(chan error, 1)
	go func() {
		_, err := io.Copy(w, r)
		c <- err
	}()
	select {
	case err := <-c:
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	case <-time.After(writeTimeout):
		return errors.New("timeout")
	}
}

func writeWithTimeout(w io.Writer, b []byte) error {
	return streamWithTimeout(w, bytes.NewBuffer(b))
}

func serve(ctx context.Context, path string, out io.Writer) error {

	// Try with port 8080, but use a random open port (0) if that doesn't work.
	var listener net.Listener
	var err error
	if listener, err = net.Listen("tcp", ":8080"); err != nil {
		if listener, err = net.Listen("tcp", ":0"); err != nil {
			return errors.WithStack(err)
		}
	}
	defer listener.Close()

	// Here we get the address we're serving on
	address, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return errors.New("Can't find address (l.Addr() is not *net.TCPAddr)")
	}

	url := fmt.Sprintf("http://localhost:%d/%s", address.Port, path)

	fmt.Fprintf(out, "Server now running on localhost:%d\n", address.Port)

	// We open the default browser and navigate to the address we're serving
	// from. This will error if the system doesn't have a browser, so we can
	// ignore the error.
	browser.OpenURL(url)

	withCancel(ctx, func() {
		err = http.Serve(listener, nil)
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil

}

func withCancel(ctx context.Context, op func()) {
	c := make(chan struct{}, 1)
	go func() {
		op()
		close(c)
	}()
	select {
	case <-c:
	case <-ctx.Done():
	}
}

func script(ctx context.Context, env vos.Env, w http.ResponseWriter, sourceMap bool) error {

	var out []byte
	var err error

	withCancel(ctx, func() {
		out, err = Compile(env, sourceMap, true, false)
	})
	if err != nil {
		return err
	}

	withCancel(ctx, func() {
		err = writeWithTimeout(w, out)
	})
	if err != nil {
		return err
	}

	return nil
}

func Compile(env vos.Env, mapping bool, addMappingComment bool, minify bool) ([]byte, error) {

	dir, err := patsy.Dir(env, "frizz.io/edit/bootstrap/main")
	if err != nil {
		return nil, err
	}

	options := &build.Options{
		GOROOT:        env.Getenv("GOROOT"),
		GOPATH:        env.Getenv("GOPATH"),
		CreateMapFile: mapping,
	}
	s := build.NewSession(options)
	packages := make(map[string]*compiler.Archive)
	importContext := &compiler.ImportContext{
		Packages: s.Types,
		Import:   s.BuildImportPath,
	}

	fset := token.NewFileSet()

	parsed, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var files []*ast.File
	for _, f := range parsed["main"].Files {
		files = append(files, f)
	}

	mainPkg, err := compiler.Compile("main", files, fset, importContext, minify)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	packages["main"] = mainPkg

	bufCode := bytes.NewBuffer(nil)
	filter := &compiler.SourceMapFilter{Writer: bufCode}
	allPkgs, err := compiler.ImportDependencies(mainPkg, importContext.Import)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if mapping {
		bufMap := bytes.NewBuffer(nil)
		smap := &sourcemap.Map{File: "script.js"}
		filter.MappingCallback = build.NewMappingCallback(smap, options.GOROOT, options.GOPATH, false)
		if err := compiler.WriteProgramCode(allPkgs, filter); err != nil {
			return nil, errors.WithStack(err)
		}
		if err := smap.WriteTo(bufMap); err != nil {
			return nil, errors.WithStack(err)
		}
		return bufMap.Bytes(), nil
	}

	if err := compiler.WriteProgramCode(allPkgs, filter); err != nil {
		return nil, errors.WithStack(err)
	}
	if addMappingComment {
		if _, err := bufCode.WriteString("//# sourceMappingURL=/static/bootstrap.js.map\n"); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return bufCode.Bytes(), nil
}
