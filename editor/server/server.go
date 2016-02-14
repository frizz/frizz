package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/websocket"

	"kego.io/kerr"

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
	"kego.io/process/packages"
	"kego.io/process/parser"
	"kego.io/process/scanner"
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

	cmd.Println("Starting editor server... ")

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
				app.fail <- kerr.Wrap("XPVTVKDWHJ", err)
				return
			}
			return
		}
		if err := root(ctx, w, req); err != nil {
			app.fail <- kerr.Wrap("QOMJGNOCQF", err)
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
				cmd.Println("Exiting editor server (finished)... ")
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
			return kerr.New("VOXPGGLWTT", "Message %T is not a *messages.DataRequest", m)
		}

		env, err := parser.ScanForEnv(ctx, request.Package.Value())
		if err != nil {
			return kerr.Wrap("EPCOFHDMBP", err)
		}

		file := filepath.Join(env.Dir, request.File.Value())

		bytes, err := scanner.ProcessFile(file)

		localContext := envctx.NewContext(ctx, env)
		o := &system.Object{}
		if err := ke.UnmarshalUntyped(localContext, bytes, o); err != nil {
			return kerr.Wrap("WHMKLGRVKV", err)
		}
		if o.Id.Name != request.Name.Value() {
			return kerr.New("GDLLEGNJOP", "Id does not match")
		}

		response := messages.NewDataResponse(request.Package.Value(), request.Name.Value(), true, string(bytes))
		conn.Respond(response, request.Guid.Value())

	}
}

func script(ctx context.Context, w http.ResponseWriter, req *http.Request, mapper bool, mapping *[]byte) error {

	wgctx.Add(ctx, "script")
	defer wgctx.Done(ctx, "script")

	path := req.URL.Path[1 : len(req.URL.Path)-10]

	env, err := parser.ScanForEnv(ctx, path)

	// This is the client code for the editor which we will compile to Javascript using GopherJs
	// below. GopherJs doesn't make it easy to compile directly from a string, so we write the
	// source file to a temporary location and delete after we have compiled it to Javascript.
	source, err := generate.Editor(ctx, env)
	if err != nil {
		return kerr.Wrap("UWPDBQXURR", err)
	}

	editorPath, _, namespaceDir, err := CreateTemporaryPackage(ctx, "a", map[string]string{"a.go": string(source)})
	if err != nil {
		return kerr.Wrap("RDRIUFUOFY", err)
	}
	defer os.RemoveAll(namespaceDir)

	options := &build.Options{CreateMapFile: true}
	s := build.NewSession(options)

	buildPkg, err := build.Import(editorPath, 0, s.InstallSuffix(), options.BuildTags)
	if err != nil {
		return kerr.Wrap("DNQVYTHSPT", err)
	}
	if buildPkg.Name != "main" {
		return kerr.New("ADHPBPSKYV", "Package name %s should be main", buildPkg.Name)
	}

	if mapper {
		if _, err := w.Write(*mapping); err != nil {
			return kerr.Wrap("WFVDCWDVWL", err)
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
			return kerr.Wrap("TXUYQOUNQS", err)
		}
	case <-ctx.Done():
		return nil
	}

	sourceMapFilter := &compiler.SourceMapFilter{Writer: buf}
	m := &sourcemap.Map{File: "script.js"}
	sourceMapFilter.MappingCallback = build.NewMappingCallback(m, options.GOROOT, options.GOPATH)

	deps, err := compiler.ImportDependencies(pkg.Archive, s.BuildImportPath)
	if err != nil {
		return kerr.Wrap("OVDUPSTRNR", err)
	}
	if err := compiler.WriteProgramCode(deps, sourceMapFilter); err != nil {
		return kerr.Wrap("YVHQEJXQGP", err)
	}

	mapBuf := bytes.NewBuffer(nil)
	m.WriteTo(mapBuf)
	*mapping = mapBuf.Bytes()

	buf.WriteString("//# sourceMappingURL=script.js.map\n")

	if _, err := w.Write(buf.Bytes()); err != nil {
		return kerr.Wrap("SSUPQDWLIV", err)
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

	env, err := parser.ScanForEnv(ctx, path)
	if err != nil {
		if _, ok := kerr.Source(err).(packages.NotFoundError); !ok {
			return kerr.Wrap("ALINBMKDRP", err)
		}
		w.WriteHeader(404)
		return nil
	}

	pcache, ok := scache.Get(path)
	if !ok {
		var err error
		pcache, err = parser.Parse(ctx, path)
		if err != nil {
			return kerr.Wrap("HIHWJRPUKE", err)
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
			return kerr.Wrap("OUBOTYGPKU", err)
		}
		pkgBytes = b
	}

	imports := map[string]*shared.ImportInfo{}

	var scan func(string) error
	scan = func(path string) error {
		if _, ok := imports[path]; ok {
			return nil
		}
		syspi, ok := scache.Get(path)
		if !ok {
			return kerr.New("VIGKIUPNCF", "%s not found in sys ctx")
		}
		info := &shared.ImportInfo{
			Path:    path,
			Aliases: syspi.Aliases,
			Types:   map[string][]byte{},
		}
		for _, name := range syspi.Files.Keys() {
			if b, ok := syspi.Files.Get(name); ok {
				info.Types[name] = b
			}
		}
		imports[path] = info

		for child, _ := range syspi.Aliases {
			if err := scan(child); err != nil {
				return kerr.Wrap("NCULMUUUOT", err)
			}
		}
		return nil
	}
	// First we always import system
	if err := scan("kego.io/system"); err != nil {
		return kerr.Wrap("KRXSLOJKWV", err)
	}
	if err := scan(env.Path); err != nil {
		return kerr.Wrap("EELKQDCJGN", err)
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
		return kerr.Wrap("OHBYTULHUQ", err)
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
		return kerr.Wrap("ICJSAIMDRF", err)
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
		return kerr.Wrap("QGLXHWPWQW", err)
	}
	defer listner.Close()

	// Here we get the address we're serving on
	address, ok := listner.Addr().(*net.TCPAddr)
	if !ok {
		return kerr.New("CBLPYVGGUR", "Can't find address (l.Addr() is not *net.TCPAddr)")
	}

	url := fmt.Sprintf("http://localhost:%d/%s/", address.Port, env.Path)

	// We open the default browser and navigate to the address we're serving from.
	if err := browser.OpenURL(url); err != nil {
		return kerr.Wrap("AEJLAXGVVA", err)
	}

	cmd.Printf("Server now running on %s\n", url)

	c := make(chan error, 1)
	go func() {
		c <- http.Serve(listner, nil)
	}()

	select {
	case err := <-c:
		if err != nil {
			return kerr.Wrap("TUCBTWMRNN", err)
		}
	case <-ctx.Done():
		// continue
	}

	if cmd.Debug {
		return kerr.New("ATUTBOICGJ", "Connection closed")
	}
	return nil

}
