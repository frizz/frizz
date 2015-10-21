package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"path"

	"golang.org/x/net/websocket"

	"kego.io/kerr"
	"kego.io/process/scan"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"

	"fmt"

	"encoding/json"

	"net/url"

	"strings"

	"github.com/gopherjs/gopherjs/build"
	"kego.io/editor/server/static"
	"kego.io/editor/shared"
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/process"
	"kego.io/system"
)

type appData struct {
	path    string
	aliases map[string]string
	pkg     *system.Package
	verbose bool
	fail    chan error
	debug   bool // in debug mode we don't exit the server on connection close
}

var app appData

func Start(path string, verbose bool, debug bool) error {

	app.debug = debug
	app.path = path
	app.verbose = verbose
	pkg, err := initialise(path)
	if err != nil {
		return kerr.New("SWSQDFXIEV", err, "initialise")
	}
	app.aliases = pkg.Aliases
	app.pkg = pkg

	if verbose {
		fmt.Println("Starting editor server... ")
	}

	app.fail = make(chan error)

	// This contains the source map that will be persisted between requests
	var mapping []byte
	http.HandleFunc("/script.js", func(w http.ResponseWriter, req *http.Request) {
		if err := script(w, req, path, false, &mapping); err != nil {
			app.fail <- kerr.New("XPVTVKDWHJ", err, "script (js)")
			return
		}
	})
	// TODO: work out how to use script mapping.
	//http.HandleFunc("/script.js.map", func(w http.ResponseWriter, req *http.Request) {
	//	if err := script(w, req, path, aliases, true, &mapping); err != nil {
	//		app.fail <- kerr.New("JRLOMPOHRQ", err, "script (map)")
	//		return
	//	}
	//})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := root(w, req, app.pkg, app.path, app.aliases); err != nil {
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
		c := connection.New(ws, app.fail, app.debug, app.path, app.aliases)

		sourceRequestsChannel := c.Subscribe(system.NewReference("kego.io/editor/shared/messages", "sourceRequest"))
		go handle(func() error { return getSource(sourceRequestsChannel, c) })

		if err := c.Receive(); err != nil {
			app.fail <- err
			return
		}
	}))

	go handle(func() error { return serve() })

	for {
		err, open := <-app.fail
		if !open {
			// Channel has been closed, so app should gracefully exit.
			if verbose {
				fmt.Println("Exiting editor server (finished)... ")
			}
		} else {
			// Error received, so app should display error.
			//return kerr.New("WKHPTVJBIL", err, "Fail channel receive")
			fmt.Println(err)
		}
		if !app.debug {
			return nil
		}
	}

	return nil
}

func handle(f func() error) {
	if err := f(); err != nil {
		app.fail <- err
	}
}

func getSource(in chan messages.MessageInterface, conn *connection.Conn) error {
	for {
		m := <-in
		request, ok := m.(*messages.SourceRequest)
		if !ok {
			return kerr.New("VOXPGGLWTT", nil, "Message %T is not a *messages.sourceRequest", m)
		}
		hashed, ok := system.GetSource(system.NewReference(app.path, request.Name.Value))
		data := ""
		if ok {
			data = string(hashed.Source)
		}
		response := messages.NewSourceResponse(request.Name.Value, ok, data)
		conn.Respond(response, request.Guid.Value)
	}
}

func initialise(path string) (pkg *system.Package, err error) {
	set, err := process.InitialiseManually(true, false, false, false, path)
	if err != nil {
		return nil, kerr.New("LFRIFXNHUY", err, "process.InitialiseManually")
	}
	if err := scan.ScanForPackage(set); err != nil {
		return nil, kerr.New("ASQLIYWNLN", err, "scan.ScanForPackage")
	}
	if err := scan.ScanForTypes(false, set); err != nil {
		return nil, kerr.New("BIVHXIAIKJ", err, "scan.ScanForTypes")
	}
	if err := scan.ScanForSource(set); err != nil {
		return nil, kerr.New("DLUESVWHXO", err, "scan.ScanForSource")
	}
	p, ok := system.GetPackage(path)
	if !ok {
		return nil, kerr.New("IHIYKRRYWL", nil, "package not found")
	}
	return p.Package, nil
}

func script(w http.ResponseWriter, req *http.Request, packagePath string, mapper bool, mapping *[]byte) error {

	options := &build.Options{CreateMapFile: true}
	s := build.NewSession(options)

	buildPkg, err := build.Import(path.Join(packagePath, "editor"), 0, s.InstallSuffix(), options.BuildTags)
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
	if err := s.BuildPackage(pkg); err != nil {
		return kerr.New("TXUYQOUNQS", err, "s.BuildPackage")
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

func root(w http.ResponseWriter, req *http.Request, pkg *system.Package, path string, aliases map[string]string) error {

	sources := system.GetAllSourceInPackage(path)

	if b, err := static.Asset(req.URL.Path[1:]); err == nil {
		if strings.HasSuffix(req.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		w.Write(b)
		return nil
	}

	data := []string{}
	types := []string{}
	for _, hashed := range sources {
		if hashed.Type == system.NewReference("kego.io/system", "type") {
			types = append(types, hashed.Id.Name)
		} else {
			data = append(data, hashed.Id.Name)
		}
	}

	b, err := ke.MarshalContext(pkg, path, aliases)
	if err != nil {
		return kerr.New("OUBOTYGPKU", err, "MarshalContext")
	}

	info := shared.Info{
		Path:    path,
		Aliases: aliases,
		Data:    data,
		Types:   types,
		Package: string(b),
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
					.content {
						padding: 2px 5px 2px 5px;
						cursor: pointer;
					}
					.content.selected {
						background-color: #eeeeee;
					}
					.mdl-spinner {
						position: absolute;
						left: 50%;
						top: 50%;
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

func serve() error {

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

	// We open the default browser and navigate to the address we're serving from.
	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d/", address.Port)); err != nil {
		return kerr.New("AEJLAXGVVA", err, "browser.OpenUrl")
	}

	if err := http.Serve(listner, nil); err != nil {
		return kerr.New("TUCBTWMRNN", err, "http.Serve")
	}

	if app.debug {
		return kerr.New("ATUTBOICGJ", nil, "Connection closed")
	}
	close(app.fail)
	return nil

}
