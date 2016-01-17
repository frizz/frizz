package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/websocket"

	"kego.io/kerr"
	"kego.io/parse"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"

	"fmt"

	"encoding/json"

	"net/url"

	"strings"

	"path/filepath"

	"github.com/gopherjs/gopherjs/build"
	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/context/wgctx"
	"kego.io/editor/server/static"
	"kego.io/editor/shared"
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/process/generate"
	"kego.io/process/pkgutils"
	"kego.io/process/scanutils"
	"kego.io/process/tests"
	"kego.io/system"
)

type appData struct {
	fail chan error
	ctx  context.Context
}

var app appData

func Start(ctx context.Context, cancel context.CancelFunc) error {

	wgctx.Add(ctx, "Start")
	defer wgctx.Done(ctx, "Start")

	app.ctx = ctx

	cmd := cmdctx.FromContext(ctx)

	app.fail = make(chan error)

	if cmd.Log {
		fmt.Println("Starting editor server... ")
	}

	// This contains the source map that will be persisted between requests
	var mapping []byte
	// TODO: work out how to use script mapping.
	//http.HandleFunc("/script.js.map", func(w http.ResponseWriter, req *http.Request) {
	//	if err := script(w, req, path, aliases, true, &mapping); err != nil {
	//		app.fail <- kerr.New("JRLOMPOHRQ", err, "script (map)")
	//		return
	//	}
	//})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/favicon.ico") {
			w.WriteHeader(404)
			return
		}
		if strings.HasSuffix(req.URL.Path, "/script.js") {
			if err := script(ctx, w, req, false, &mapping); err != nil {
				app.fail <- kerr.New("XPVTVKDWHJ", err, "script (js)")
				return
			}
			return
		}
		if err := root(ctx, w, req); err != nil {
			app.fail <- kerr.New("QOMJGNOCQF", err, "root")
			return
		}
	})

	http.Handle("/_socket", websocket.Handler(func(ws *websocket.Conn) {
		if connection.MESSAGE_TYPE == connection.M_BINARY {
			ws.PayloadType = 0x2 // binary messages
		} else {
			ws.PayloadType = 0x1 // string messages
		}
		c := connection.New(ctx, ws, app.fail, cmd.Debug)

		dataRequestsChannel := c.Subscribe(*system.NewReference("kego.io/editor/shared/messages", "dataRequest"))
		go handle(func() error { return getData(ctx, dataRequestsChannel, c) })

		if err := c.Receive(); err != nil {
			app.fail <- err
			return
		}
	}))

	go handle(func() error { return serve(ctx) })

	done := app.ctx.Done()
	for {
		select {
		case <-done:
			fmt.Println("Exiting editor server (interupted)... ")
			return nil
		case err, open := <-app.fail:
			if !open {
				cancel()
				// Channel has been closed, so app should gracefully exit.
				if cmd.Log {
					fmt.Println("Exiting editor server (finished)... ")
				}
			} else {
				// Error received, so app should display error.
				//return kerr.New("WKHPTVJBIL", err, "Fail channel receive")
				fmt.Println(err)
			}
			if !cmd.Debug {
				return nil
			}
		}
	}

	return nil
}

func handle(f func() error) {
	if err := f(); err != nil {
		app.fail <- err
	}
}

func getData(ctx context.Context, in chan messages.MessageInterface, conn *connection.Conn) error {
	//env := envctx.FromContext(ctx)
	for {
		m := <-in
		request, ok := m.(*messages.DataRequest)
		if !ok {
			return kerr.New("VOXPGGLWTT", nil, "Message %T is not a *messages.DataRequest", m)
		}

		env, err := parse.ScanForEnv(ctx, request.Package.Value())
		if err != nil {
			return kerr.New("EPCOFHDMBP", err, "parse.ScanForEnv")
		}

		file := filepath.Join(env.Dir, request.File.Value())

		bytes, err := scanutils.ProcessFile(file)

		localContext := envctx.NewContext(ctx, env)
		o := &system.Object{}
		if err := ke.UnmarshalUntyped(localContext, bytes, o); err != nil {
			return kerr.New("WHMKLGRVKV", err, "ke.UnmarshalUntyped")
		}
		if o.Id.Name != request.Name.Value() {
			return kerr.New("GDLLEGNJOP", nil, "Id does not match")
		}

		response := messages.NewDataResponse(request.Package.Value(), request.Name.Value(), true, string(bytes))
		conn.Respond(response, request.Guid.Value())

	}
}

func script(ctx context.Context, w http.ResponseWriter, req *http.Request, mapper bool, mapping *[]byte) error {

	wgctx.Add(ctx, "script")
	defer wgctx.Done(ctx, "script")

	path := req.URL.Path[1 : len(req.URL.Path)-10]

	env, err := parse.ScanForEnv(ctx, path)

	// This is the client code for the editor which we will compile to Javascript using GopherJs
	// below. GopherJs doesn't make it easy to compile directly from a string, so we write the
	// source file to a temporary location and delete after we have compiled it to Javascript.
	source, err := generate.Editor(ctx, env)
	if err != nil {
		return kerr.New("UWPDBQXURR", err, "generate.Editor")
	}

	namespace, err := tests.CreateTemporaryNamespace()
	if err != nil {
		return kerr.New("AFRRXCMOCM", err, "CreateTemporaryNamespace")
	}
	defer os.RemoveAll(namespace)

	editorPath, _, _, err := tests.CreateTemporaryPackage(namespace, "a", map[string]string{"a.go": string(source)})
	if err != nil {
		return kerr.New("RDRIUFUOFY", err, "CreateTemporaryPackage")
	}

	options := &build.Options{CreateMapFile: true}
	s := build.NewSession(options)

	buildPkg, err := build.Import(editorPath, 0, s.InstallSuffix(), options.BuildTags)
	if err != nil {
		return kerr.New("DNQVYTHSPT", err, "build.Import")
	}
	if buildPkg.Name != "main" {
		return kerr.New("ADHPBPSKYV", nil, "Package name %s should be main", buildPkg.Name)
	}

	if mapper {
		if _, err := w.Write(*mapping); err != nil {
			return kerr.New("WFVDCWDVWL", err, "w.Write (mapping)")
		}
		return nil
	}

	buf := bytes.NewBuffer(nil)
	pkg := &build.PackageData{Package: buildPkg.Package}

	c := make(chan error, 1)
	go func() {
		c <- s.BuildPackage(pkg)
	}()

	select {
	case err := <-c:
		if err != nil {
			return kerr.New("TXUYQOUNQS", err, "s.BuildPackage")
		}
	case <-ctx.Done():
		return nil
	}

	sourceMapFilter := &compiler.SourceMapFilter{Writer: buf}
	m := &sourcemap.Map{File: "script.js"}
	sourceMapFilter.MappingCallback = build.NewMappingCallback(m, options.GOROOT, options.GOPATH)

	deps, err := compiler.ImportDependencies(pkg.Archive, s.ImportContext.Import)
	if err != nil {
		return kerr.New("OVDUPSTRNR", err, "compiler.ImportDependencies")
	}
	if err := compiler.WriteProgramCode(deps, sourceMapFilter); err != nil {
		return kerr.New("YVHQEJXQGP", err, "compiler.WriteProgramCode")
	}

	mapBuf := bytes.NewBuffer(nil)
	m.WriteTo(mapBuf)
	*mapping = mapBuf.Bytes()

	buf.WriteString("//# sourceMappingURL=script.js.map\n")

	if _, err := w.Write(buf.Bytes()); err != nil {
		return kerr.New("SSUPQDWLIV", err, "w.Write (js)")
	}
	return nil
}

func root(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	wgctx.Add(ctx, "root")
	defer wgctx.Done(ctx, "root")

	scache := sysctx.FromContext(ctx)

	if b, err := static.Asset(req.URL.Path[1:]); err == nil {
		if strings.HasSuffix(req.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		w.Write(b)
		return nil
	}

	path := req.URL.Path[1:]
	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}

	env, err := parse.ScanForEnv(ctx, path)
	if err != nil {
		if _, ok := kerr.Source(err).(pkgutils.NotFoundError); ok {
			w.WriteHeader(404)
			return nil
		}
		return kerr.New("ALINBMKDRP", err, "scanForEnv")
	}

	pcache, ok := scache.Get(path)
	if !ok {
		var err error
		pcache, err = parse.Parse(ctx, path)
		if err != nil {
			return kerr.New("HIHWJRPUKE", err, "parse.Parse")
		}
	}

	data := map[string]string{}
	for _, name := range pcache.Globals.Keys() {
		if g, ok := pcache.Globals.Get(name); ok {
			data[name] = g.File
		}
	}

	pkgBytes := pcache.PackageBytes
	if pkgBytes == nil {
		b, err := ke.MarshalContext(ctx, system.EmptyPackage())
		if err != nil {
			return kerr.New("OUBOTYGPKU", err, "MarshalContext")
		}
		pkgBytes = b
	}

	getImport := func(importPath string) (*shared.ImportInfo, error) {
		importPackageInfo, ok := scache.Get(importPath)
		if !ok {
			return nil, kerr.New("VIGKIUPNCF", nil, "%s not found in sys ctx")
		}
		info := &shared.ImportInfo{
			Path:    importPath,
			Aliases: importPackageInfo.Environment.Aliases,
			Types:   map[string][]byte{},
		}
		for _, name := range importPackageInfo.TypeSource.Keys() {
			if b, ok := importPackageInfo.TypeSource.Get(name); ok {
				info.Types[name] = b
			}
		}
		return info, nil
	}

	systemImport, err := getImport("kego.io/system")
	if err != nil {
		return kerr.New("SAFQNEKXAP", err, "getImport")
	}
	imports := map[string]*shared.ImportInfo{"kego.io/system": systemImport}
	var scan func(string) error
	scan = func(p string) error {
		if _, ok := imports[p]; ok {
			return nil
		}
		aliasImport, err := getImport(p)
		if err != nil {
			return kerr.New("YTHDSRBMBW", err, "getImport")
		}
		imports[p] = aliasImport

		for child, _ := range aliasImport.Aliases {
			if err := scan(child); err != nil {
				return kerr.New("NCULMUUUOT", err, "scan")
			}
		}
		return nil
	}
	if err := scan(env.Path); err != nil {
		return kerr.New("EELKQDCJGN", err, "scan")
	}

	info := shared.Info{
		Path:    env.Path,
		Aliases: env.Aliases,
		Data:    data,
		Package: pkgBytes,
		Imports: imports,
	}
	marshalled, err := json.Marshal(info)
	if err != nil {
		return kerr.New("OHBYTULHUQ", err, "json.Marshal (info)")
	}
	attribute := url.QueryEscape(fmt.Sprintf("%s", marshalled))

	source := []byte(`
		<html>
			<head>
				<meta charset="utf-8">

				<link rel="stylesheet" href="/material.min.css">
				<script src="/material.min.js"></script>
				<link rel="stylesheet" href="/icon.css">

				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
				<style>
					.node {
						position: relative;
						margin-left: 15px;
					}
					.root {
						margin-top: -10px;
						margin-left: 20px;
					}
					.toggle {
						cursor: pointer;
						position: absolute;
						left: -20px;
						top: 2px;
						width: 15px;
					}
					.toggle svg {
						opacity: 0.5;
						height: 20px;
						width: 20px;
					}
					.label {
						cursor: pointer;
					}
					.content {
						padding: 2px 5px 2px 5px;
					}
					.content.selected {
						background-color: #eeeeee;
					}
					.mdl-spinner {
						position: absolute;
						left: 50%;
						top: 50%;
					}
					.badge {
						position: absolute;
						margin-top: -2px;
						margin-left: -3px;
						opacity: 0.5;
					}
				</style>
			</head>
			<body id="body" info="` + attribute + `">
				<div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div>
				<div class="mdl-layout mdl-js-layout mdl-layout--fixed-drawer">
					<div class="mdl-layout__drawer">
						<span class="mdl-layout-title"></span>
						<nav class="mdl-navigation"></nav>
					</div>
					<main class="mdl-layout__content mdl-color--grey-100">
						<div class="page-content"></div>
					</main>
				</div>
			</body>
			<script src="script.js"></script>
		</html>`)

	if _, err := w.Write(source); err != nil {
		return kerr.New("ICJSAIMDRF", err, "w.Write (root)")
	}

	return nil
}

func serve(ctx context.Context) error {

	wgctx.Add(ctx, "serve")
	defer wgctx.Done(ctx, "serve")

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	// Starting with port zero chooses a random open port
	listner, err := net.Listen("tcp", ":0")
	if err != nil {
		return kerr.New("QGLXHWPWQW", err, "net.Listen")
	}
	defer listner.Close()

	// Here we get the address we're serving on
	address, ok := listner.Addr().(*net.TCPAddr)
	if !ok {
		return kerr.New("CBLPYVGGUR", nil, "Can't find address (l.Addr() is not *net.TCPAddr)")
	}

	url := fmt.Sprintf("http://localhost:%d/%s/", address.Port, env.Path)

	// We open the default browser and navigate to the address we're serving from.
	if err := browser.OpenURL(url); err != nil {
		return kerr.New("AEJLAXGVVA", err, "browser.OpenUrl")
	}

	if cmd.Log {
		fmt.Printf("Server now running on %s\n", url)
	}

	c := make(chan error, 1)
	go func() {
		c <- http.Serve(listner, nil)
	}()

	select {
	case err := <-c:
		if err != nil {
			return kerr.New("TUCBTWMRNN", err, "http.Serve")
		}
	case <-ctx.Done():
		// continue
	}

	if cmd.Debug {
		return kerr.New("ATUTBOICGJ", nil, "Connection closed")
	}
	return nil

}
