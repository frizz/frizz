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
	"go/types"

	"encoding/gob"

	"io/ioutil"

	"path/filepath"

	"mime"

	gobuild "go/build"

	"os"

	"encoding/json"

	"frizz.io/config"
	"frizz.io/edit/auther"
	"frizz.io/edit/hasher"
	"frizz.io/edit/jscompiler"
	"github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/compiler/prelude"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpgzip"
	"golang.org/x/tools/go/loader"
)

const writeTimeout = time.Second * 2
const JS_VER = "17"

func Open(ctx context.Context, cancel context.CancelFunc, env vos.Env, out io.Writer, path string) error {

	fail := make(chan error)

	auth := auther.New()

	jsc := jscompiler.New(vos.Os())

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
		case strings.HasPrefix(req.URL.Path, "/js/"):
			if err := js(ctx, env, jsc, w, req); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/main/"):
			if err := jsMain(ctx, env, w, req); err != nil {
				fail <- err
			}
		case strings.HasPrefix(req.URL.Path, "/root/"):
			if err := rootjs(ctx, env, auth, jsc, w, req); err != nil {
				fail <- err
			}
		case req.URL.Path == "/prelude.js":
			if err := servePrelude(w); err != nil {
				fail <- err
			}
		default:
			if err := root(ctx, env, auth, jsc, w, req); err != nil {
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

func jsMain(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request) error {
	path := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/main/"), ".js")

	dir, err := patsy.Dir(env, path)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting dir for for %s", path), 500)
		return err
	}

	files, err := getDataFiles(dir)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting data files for %s", path), 500)
		return errors.WithStack(err)
	}

	hash, err := hasher.Hash(files)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting data files hash for %s", path), 500)
		return errors.WithStack(err)
	}

	/*
		package main

		func main() {
			editor.Edit(<pkg>.Package, <hash>)
		}
	*/
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("frizz.io/edit/editor", "Edit").Call(jen.Qual(path, "Package"), jen.Lit(hash)),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "main.go", buf.Bytes(), parser.ParseComments)
	if err != nil {
		return err
	}

	conf := loader.Config{}
	conf.Fset = fset
	conf.CreateFromFiles("main", parsed)
	conf.ParserMode = parser.ParseComments
	conf.Build = func() *gobuild.Context { c := gobuild.Default; return &c }() // make a copy of gobuild.Default
	conf.Build.GOPATH = env.Getenv("GOPATH")
	conf.Build.BuildTags = []string{"js"}
	conf.AllowErrors = true
	conf.TypeChecker.Error = func(e error) {}
	loaded, err := conf.Load()
	if err != nil {
		return errors.WithStack(err)
	}

	archives := make(map[string]*compiler.Archive)
	var importContext *compiler.ImportContext
	importContext = &compiler.ImportContext{
		Packages: make(map[string]*types.Package),
		Import: func(path string) (*compiler.Archive, error) {

			// find in local cache
			if a, ok := archives[path]; ok {
				return a, nil
			}

			pi := loaded.Package(path)
			importContext.Packages[path] = pi.Pkg

			// find in standard library cache
			archive, err := openArchive(path)
			if err != nil {
				return nil, err
			}
			if archive != nil {
				archives[path] = archive
				return archive, nil
			}

			// compile package
			archive, err = compiler.Compile(path, pi.Files, loaded.Fset, importContext, false)
			if err != nil {
				return nil, err
			}
			archives[path] = archive
			return archive, nil
		},
	}

	// compile package
	archive, err := compiler.Compile("main", []*ast.File{parsed}, loaded.Fset, importContext, false)
	if err != nil {
		return err
	}

	if err := servePackage(w, archive); err != nil {
		return err
	}

	return nil
}

func js(ctx context.Context, env vos.Env, jsc *jscompiler.JsCompiler, w http.ResponseWriter, req *http.Request) error {

	// TODO: Add caching
	/*
		name := fmt.Sprintf("v4/%s?%d", path, source.Hash)
		found, resp, err := b.Cache.Get(name)
		if err != nil {
			return nil, err
		}
		if found {
			cached++
			archive, err := compiler.ReadArchive(path+".a", path, bytes.NewBuffer(resp), b.Packages)
			if err != nil {
				return nil, err
			}
			b.Archives[path] = archive
			return archive, nil
		}
	*/

	path := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/js/"), ".js")

	archive, err := openArchive(path)
	if err != nil {
		return err
	}
	if archive != nil {
		if err := servePackage(w, archive); err != nil {
			return err
		}
		return nil
	}

	/*
		loaded, err := loadProg(env, path, parser.ParseComments)
		if err != nil {
			return err
		}

		archives := make(map[string]*compiler.Archive)
		var importContext *compiler.ImportContext
		importContext = &compiler.ImportContext{
			Packages: make(map[string]*types.Package),
			Import: func(path string) (*compiler.Archive, error) {

				// find in local cache
				if a, ok := archives[path]; ok {
					return a, nil
				}

				pi := loaded.Package(path)
				importContext.Packages[path] = pi.Pkg

				// find in standard library cache
				archive, err := openArchive(path)
				if err != nil {
					return nil, err
				}
				if archive != nil {
					archives[path] = archive
					return archive, nil
				}

				// compile package
				archive, err = compiler.Compile(path, pi.Files, loaded.Fset, importContext, false)
				if err != nil {
					return nil, err
				}
				archives[path] = archive
				return archive, nil
			},
		}

		archive, err = importContext.Import(path)
		if err != nil {
			return err
		}*/

	archive = jsc.Get(path)

	if err := servePackage(w, archive); err != nil {
		return err
	}

	return nil
}

func openArchive(path string) (*compiler.Archive, error) {
	f, err := assets.Assets.Open(fmt.Sprintf("pkg/%s.a", path))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}

	defer f.Close()

	archive := new(compiler.Archive)
	if err := gob.NewDecoder(f).Decode(archive); err != nil {
		return nil, errors.WithStack(err)
	}
	return archive, nil

}

func servePackage(w http.ResponseWriter, archive *compiler.Archive) error {

	selection := make(map[*compiler.Decl]struct{})
	for _, d := range archive.Declarations {
		selection[d] = struct{}{}
	}

	buf := &bytes.Buffer{}

	buf.WriteString("$initialisers[\"" + archive.ImportPath + "\"] = function () {")

	if err := compiler.WritePkgCode(archive, selection, false, &compiler.SourceMapFilter{Writer: buf}); err != nil {
		return err
	}

	buf.WriteString("};")

	w.Header().Set("Cache-Control", "max-age=31536000")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := io.Copy(w, buf); err != nil {
		http.Error(w, fmt.Sprintf("error streaming bundle for %s", archive.ImportPath), 500)
		return errors.WithStack(err)
	}
	return nil
}

func servePrelude(w http.ResponseWriter) error {

	buf := bytes.NewBufferString(prelude.Prelude)

	w.Header().Set("Cache-Control", "max-age=31536000")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := io.Copy(w, buf); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func src(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request) error {
	pkg := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/src/"), ".bin")
	files, err := getSourceFiles(env, pkg)
	if err != nil {
		return err
	}
	if err := bundle(ctx, env, w, req, pkg, files); err != nil {
		return err
	}
	return nil
}

func data(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request) error {
	pkg := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/data/"), ".bin")
	dir, err := patsy.Dir(env, pkg)
	if err != nil {
		return err
	}
	files, err := getDataFiles(dir)
	if err != nil {
		return err
	}
	if err := bundle(ctx, env, w, req, pkg, files); err != nil {
		return err
	}
	return nil
}

func bundle(ctx context.Context, env vos.Env, w http.ResponseWriter, req *http.Request, path string, files map[string][]byte) error {

	hashFromRequest, err := strconv.ParseUint(req.URL.Query().Get("hash"), 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing hash for %s", req.URL), 500)
		return errors.WithStack(err)
	}

	bundle, err := getBundle(files)
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

func getBundle(files map[string][]byte) (*common.Bundle, error) {
	hash, err := hasher.Hash(files)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &common.Bundle{Hash: hash, Files: files}, nil
}

func getDataFiles(dir string) (map[string][]byte, error) {
	fia, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	files := make(map[string][]byte)
	for _, fi := range fia {
		if !strings.HasSuffix(fi.Name(), ".frizz.json") {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(dir, fi.Name()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		files[fi.Name()] = b
	}
	return files, nil
}

func loadProg(env vos.Env, path string, mode parser.Mode) (*loader.Program, error) {
	conf := loader.Config{}
	conf.Import(path)
	conf.Import("frizz.io/edit/editor")
	conf.ParserMode = mode
	conf.Build = func() *gobuild.Context { c := gobuild.Default; return &c }() // make a copy of gobuild.Default
	//conf.Build.GOARCH = "js"
	conf.Build.GOPATH = env.Getenv("GOPATH")
	conf.Build.BuildTags = []string{"js"}
	conf.AllowErrors = true
	conf.TypeChecker.Error = func(e error) {}
	loaded, err := conf.Load()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return loaded, nil
}

func getSourceFiles(env vos.Env, path string) (map[string][]byte, error) {
	loaded, err := loadProg(env, path, parser.PackageClauseOnly)
	if err != nil {
		return nil, err
	}
	pi := loaded.Package(path)

	files := map[string][]byte{}
	for _, f := range pi.Files {
		fpath := loaded.Fset.File(f.Pos()).Name()
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, name := filepath.Split(fpath)
		files[name] = b
	}
	return files, nil
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

	files, err := getDataFiles(dir)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting data files for %s", path), 500)
		return errors.WithStack(err)
	}

	hash, err := hasher.Hash(files)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting data files hash for %s", path), 500)
		return errors.WithStack(err)
	}

	blob := &common.Blob{
		Hash:   hash,
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

		files := map[string][]byte{}
		for _, f := range pi.Files {
			fpath := loaded.Fset.File(f.Pos()).Name()
			_, fname := filepath.Split(fpath)
			b, err := ioutil.ReadFile(fpath)
			if err != nil {
				http.Error(w, fmt.Sprintf("error reading %s in %s", fname, p.Path()), 500)
				return errors.WithStack(err)
			}
			files[fname] = b
		}

		hash, err := hasher.Hash(files)
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting source files hash for %s", p.Path()), 500)
			return errors.WithStack(err)
		}

		blob.Source[p.Path()] = hash

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

func root(ctx context.Context, env vos.Env, auth auther.Auther, jsc *jscompiler.JsCompiler, w http.ResponseWriter, req *http.Request) error {

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

	path := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/"), "/")

	source := []byte(`
		<html>
			<head>
				<meta charset="utf-8">
				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
			</head>
			<body id="wrapper" auth="` + attribute + `">
				<span id="log">Parsing project...</span>
			</body>
			<script>
				var $progressCount = 0;
			</script>
			<script src="/root/` + path + `.js"></script>
		</html>`)

	if err := writeWithTimeout(w, source); err != nil {
		http.Error(w, "error streaming homepage", 500)
		return err
	}

	return nil
}

func rootjs(ctx context.Context, env vos.Env, auth auther.Auther, jsc *jscompiler.JsCompiler, w http.ResponseWriter, req *http.Request) error {
	path := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/root/"), ".js")

	prog, err := jsc.Hint(path, "frizz.io/edit/editor")
	if err != nil {
		return err
	}

	orderedPackages := []string{
		"github.com/gopherjs/gopherjs/js",
		"github.com/gopherjs/gopherjs/nosync",
	}
	done := map[string]struct{}{
		"github.com/gopherjs/gopherjs/js":     {},
		"github.com/gopherjs/gopherjs/nosync": {},
	}
	hashes := map[string]uint64{}
	var orderImports func(local string) error
	orderImports = func(local string) error {
		pi := prog.Package(local)
		if len(pi.Files) == 0 {
			done[local] = struct{}{}
			return nil
		}
		std := strings.HasPrefix(prog.Fset.File(pi.Files[0].Pos()).Name(), gobuild.Default.GOROOT)
		if std {
			// stdlib package. make sure it's valid:
			if _, err := assets.Assets.Open(fmt.Sprintf("pkg/%s.a", local)); os.IsNotExist(err) {
				done[local] = struct{}{}
				return nil
			}
		}

		for _, child := range pi.Pkg.Imports() {
			if _, ok := done[child.Path()]; ok {
				continue
			}
			orderImports(child.Path())
		}

		if !std {
			files := map[string][]byte{}
			imports := map[string]uint64{}
			for _, f := range pi.Files {
				fpath := prog.Fset.File(f.Pos()).Name()
				b, err := ioutil.ReadFile(fpath)
				if err != nil {
					return err
				}
				_, name := filepath.Split(fpath)
				files[name] = b
			}
			for _, child := range pi.Pkg.Imports() {
				if hash, ok := hashes[child.Path()]; ok {
					imports[child.Path()] = hash
				}
			}
			hash, err := hasher.HashWithImports(files, imports)
			if err != nil {
				return err
			}
			hashes[local] = hash
		}

		orderedPackages = append(orderedPackages, local)
		done[local] = struct{}{}
		return nil
	}
	if err := orderImports(path); err != nil {
		return err
	}
	if err := orderImports("frizz.io/edit/editor"); err != nil {
		return err
	}

	packagesJson, err := json.Marshal(orderedPackages)
	if err != nil {
		return err
	}

	hashesJson, err := json.Marshal(hashes)
	if err != nil {
		return err
	}

	src := &bytes.Buffer{}
	fmt.Fprint(src, "\"use strict\";\n")
	fmt.Fprint(src, "var $initialisers = {};\n")
	fmt.Fprint(src, "var $mainPkg;\n")
	fmt.Fprintf(src, "var $path = \"%s\";\n", path)
	fmt.Fprintf(src, "var $pkgs = %s;\n", packagesJson)
	fmt.Fprintf(src, "var $hashes = %s;\n", hashesJson)
	fmt.Fprintf(src, "var $ver = \"%s\";\n", JS_VER)
	fmt.Fprint(src, "var $progressCount = 0;\n")
	fmt.Fprint(src, "var $progressTotal = 0;\n")
	fmt.Fprint(src, `
var finished = function() {
	document.getElementById("log").innerHTML = "Initialising...";
	$pkgs.forEach(function(path){
		$initialisers[path]();
	});
	$initialisers["main"]();
	$mainPkg = $packages["main"];
	$synthesizeMethods();
	$packages["runtime"].$init();
	$go($mainPkg.$init, []);
	$flushConsole();
}
var load = function(url) {
	$progressTotal++;
    var tag = document.createElement('script');
    tag.src = url;
	var done = function() {
		$progressCount++;
		document.getElementById("log").innerHTML = "Loading " + $progressCount + " / " + $progressTotal;
		if ($progressCount == $progressTotal) {
			finished();
		}
	}
    tag.onload = done;
    tag.onreadystatechange = done;
    document.body.appendChild(tag);
}
load("/prelude.js");
load("/main/" + $path + ".js?v="+$ver);
$pkgs.forEach(function(path){
	var hash = $hashes[path] ? "&hash="+$hashes[path] : "";
	load("/js/" + path + ".js?v="+$ver+hash);
});
`)

	if err := writeWithTimeout(w, src.Bytes()); err != nil {
		http.Error(w, "error streaming data", 500)
		return err
	}

	return nil

	/*
			initialisers := "\n"
			imports := "\n"
			for _, p := range orderedPackages {
				hash := ""
				if h, ok := hashes[p]; ok {
					hash = fmt.Sprintf("&hash=%d", h)
				}
				imports += fmt.Sprintf("\t\t\t<script src=\"/js/%s.js?v="+JS_VER+hash+"\"></script>\n", p)
				initialisers += fmt.Sprintf("\t\t\t\t$initialisers[\"%s\"]();\n", p)
			}

			source := []byte(`
				<html>
					<head>
						<meta charset="utf-8">
						<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
					</head>
					<body id="wrapper" auth="` + attribute + `">
						<span id="log">Loading...</span>
					</body>
					<script>
						var $progressCount = 0;
						var $progressTotal = ` + fmt.Sprint(len(orderedPackages)) + `;
					</script>
					<script src="/prelude.js?v=` + PRELUDE_VER + `"></script>
					` + imports + `
					<script src="/main/` + path + `.js?v=` + MAIN_VER + `"></script>
					<script>
						"use strict";
						` + initialisers + `
						$initialisers["main"]();
						var $mainPkg = $packages["main"];
						$synthesizeMethods();
						$packages["runtime"].$init();
						$go($mainPkg.$init, []);
						$flushConsole();
					</script>
				</html>`)

		if err := writeWithTimeout(w, source); err != nil {
			http.Error(w, "error streaming homepage", 500)
			return err
		}

		return nil*/
}

func rootOld(ctx context.Context, auth auther.Auther, w http.ResponseWriter, req *http.Request) error {

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
			<body id="wrapper" auth="` + attribute + `">
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

	fmt.Fprintf(out, "Server now running on localhost:%d\n", address.Port)

	// We open the default browser and navigate to the address we're serving
	// from. This will error if the system doesn't have a browser, so we can
	// ignore the error.
	browser.OpenURL(fmt.Sprintf("http://localhost:%d/%s", address.Port, path))

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
