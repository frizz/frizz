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

	"bytes"
	"go/ast"
	"go/parser"
	"go/token"

	"encoding/base64"
	"encoding/gob"

	"io/ioutil"

	"path/filepath"

	"mime"
	"path"

	"frizz.io/config"
	"frizz.io/edit/auther"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpgzip"
)

const writeTimeout = time.Second * 2

func Open(ctx context.Context, cancel context.CancelFunc, env vos.Env, out io.Writer, path string) error {

	fail := make(chan error)

	auth := auther.New()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/favicon.ico") {
			w.WriteHeader(404)
			return
		}
		if config.DEV {
			// Only generate bootstrap dynamically in debug mode. Non debug mode delivers
			// bootstrap.js from static files (minified, without source map).
			if strings.HasSuffix(req.URL.Path, "/bootstrap.js") {
				if err := script(ctx, env, w, false); err != nil {
					fail <- err
					return
				}
				return
			}
			if strings.HasSuffix(req.URL.Path, "/bootstrap.js.map") {
				if err := script(ctx, env, w, true); err != nil {
					fail <- err
					return
				}
				return
			}
		}
		if strings.HasSuffix(req.URL.Path, "/blob.bin") {
			if err := blob(ctx, env, auth, w, req); err != nil {
				fail <- err
				return
			}
			return
		}
		if err := root(ctx, auth, w, req); err != nil {
			fail <- err
			return
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

func blob(ctx context.Context, env vos.Env, auth auther.Auther, w http.ResponseWriter, req *http.Request) error {

	info, err := common.NewRequestInfo(req.URL.Query().Get("info"))
	if err != nil {
		return errors.WithStack(err)
	}
	if !auth.Auth(info.Id, info.Hash) {
		return errors.New("auth error")
	}

	pathAndName := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/"), "/blob.bin")

	dir, err := patsy.Dir(env, pathAndName)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	blob := &common.Blob{Files: map[string][]byte{}}
	for _, file := range files {
		if (strings.HasSuffix(file.Name(), ".go") && !strings.HasSuffix(file.Name(), "_test.go")) || strings.HasSuffix(file.Name(), ".frizz.json") {
			b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			blob.Files[file.Name()] = b
		}
	}

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(blob); err != nil {
		return errors.WithStack(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := io.Copy(w, buf); err != nil {
		return err
	}

	return nil
}

func root(ctx context.Context, auth auther.Auther, w http.ResponseWriter, req *http.Request) error {

	pathAndName := strings.TrimPrefix(req.URL.Path, "/")

	if file, err := assets.Assets.Open(pathAndName); err == nil {
		defer file.Close()

		w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(pathAndName)))

		_, noCompress := file.(httpgzip.NotWorthGzipCompressing)
		gzb, isGzb := file.(httpgzip.GzipByter)

		if isGzb && !noCompress && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			if err := writeWithTimeout(w, gzb.GzipBytes()); err != nil {
				return err
			}
		} else {
			if err := streamWithTimeout(w, file); err != nil {
				return err
			}
		}
		return nil
	}

	id := make([]byte, 64)
	if _, err := rand.Read(id); err != nil {
		return errors.WithStack(err)
	}
	info := common.RequestInfo{
		Id:   id,
		Hash: auth.Sign(id),
	}
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(info); err != nil {
		return errors.WithStack(err)
	}
	attrib := base64.StdEncoding.EncodeToString(buf.Bytes())

	source := []byte(`
		<html>
			<head>
				<meta charset="utf-8">
				<script src="/jquery-2.2.4.min.js"></script>
				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
			</head>
			<body id="body" info="` + attrib + `"></body>
			<script src="/bootstrap.js"></script>
		</html>`)

	if err := writeWithTimeout(w, source); err != nil {
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

	// Starting with port zero chooses a random open port
	listner, err := net.Listen("tcp", ":0")
	if err != nil {
		return errors.WithStack(err)
	}
	defer listner.Close()

	// Here we get the address we're serving on
	address, ok := listner.Addr().(*net.TCPAddr)
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
		err = http.Serve(listner, nil)
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

	dir, err := patsy.Dir(env, "frizz.io/edit/bootstrap")
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
		if _, err := bufCode.WriteString("//# sourceMappingURL=/bootstrap.js.map\n"); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return bufCode.Bytes(), nil
}
