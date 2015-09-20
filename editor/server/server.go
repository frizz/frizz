package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"path"
	"strings"

	"golang.org/x/net/websocket"

	"kego.io/ke"
	"kego.io/kerr"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"

	"fmt"

	"encoding/json"

	"net/url"

	"github.com/gopherjs/gopherjs/build"
	"kego.io/editor/shared"
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/process"
	"kego.io/system"
)

type appData struct {
	path     string
	aliases  map[string]string
	verbose  bool
	fail     chan error
	finished chan bool
}

var app appData

func Start(path string, verbose bool) error {

	app.path = path
	app.verbose = verbose
	aliases, err := initialise(path)
	if err != nil {
		return kerr.New("SWSQDFXIEV", err, "initialise")
	}
	app.aliases = aliases

	if verbose {
		fmt.Println("Starting editor server... ")
	}

	app.finished = make(chan bool)
	app.fail = make(chan error)

	// This contains the source map that will be persisted between requests
	var mapping []byte
	http.HandleFunc("/script.js", func(w http.ResponseWriter, req *http.Request) {
		if err := script(w, req, path, aliases, false, &mapping); err != nil {
			app.fail <- kerr.New("XPVTVKDWHJ", err, "script (js)")
			return
		}
	})
	http.HandleFunc("/script.js.map", func(w http.ResponseWriter, req *http.Request) {
		if err := script(w, req, path, aliases, true, &mapping); err != nil {
			app.fail <- kerr.New("JRLOMPOHRQ", err, "script (map)")
			return
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := wildcard(w, req, path, aliases); err != nil {
			app.fail <- kerr.New("QOMJGNOCQF", err, "wildcard")
			return
		}
	})

	http.Handle("/_socket", websocket.Handler(func(ws *websocket.Conn) {
		c := connection.New(ws, app.fail, app.finished, path, aliases)

		globalRequestsChannel := c.Subscribe(system.NewReference("kego.io/editor/shared/messages", "globalRequest"))
		go handle(func() error { return getGlobals(globalRequestsChannel, c) })

		if err := c.Receive(); err != nil {
			app.fail <- err
			return
		}
	}))

	go handle(func() error { return serve(app.finished) })

	select {
	case <-app.finished:
		if verbose {
			fmt.Println("Exiting editor server (finished)... ")
		}
	case err := <-app.fail:
		return kerr.New("WKHPTVJBIL", err, "Fail channel receive")
	}
	return nil
}

func handle(f func() error) {
	if err := f(); err != nil {
		app.fail <- err
	}
}

func getGlobals(in chan messages.Message, conn *connection.Conn) error {
	for {
		m := <-in
		request, ok := m.(*messages.GlobalRequest)
		if !ok {
			return kerr.New("BKGSQAWUJY", nil, "Message %T is not a *messages.Request", m)
		}
		hashed, ok := system.GetGlobal(app.path, request.Name.Value)
		data := ""
		if ok {
			b, err := ke.Marshal(hashed.Object)
			if err != nil {
				return kerr.New("WTXNLMABAS", err, "ke.Marshal")
			}
			data = string(b)
		}
		response := messages.NewGlobalResponse(request.Name.Value, ok, data)
		conn.Respond(response, request.Guid.Value)
	}
}

func initialise(path string) (aliases map[string]string, err error) {
	set, err := process.InitialiseManually(true, false, false, false, path)
	if err != nil {
		return nil, kerr.New("LFRIFXNHUY", err, "process.InitialiseManually")
	}
	aliases, err = process.ScanForAliases(set)
	if err != nil {
		return nil, kerr.New("ASQLIYWNLN", err, "process.ScanForAliases")
	}
	if err := process.ScanForTypes(false, set); err != nil {
		return nil, kerr.New("BIVHXIAIKJ", err, "process.ScanForTypes")
	}
	if err = process.ScanForGlobals(set); err != nil {
		return nil, kerr.New("MDUFNJBVGQ", err, "process.ScanForGlobals")
	}
	return aliases, nil
}

func script(w http.ResponseWriter, req *http.Request, packagePath string, aliases map[string]string, mapper bool, mapping *[]byte) error {

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
	pkg := &build.PackageData{Package: buildPkg}
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

func wildcard(w http.ResponseWriter, req *http.Request, path string, aliases map[string]string) error {

	globals := system.GetAllGlobalsInPackage(path, nil)

	names := []string{}
	for _, hashed := range globals {
		name := hashed.Object.Object().Id.Name
		if strings.HasPrefix(name, "@") {
			continue
		}
		names = append(names, name)
	}

	info := shared.Info{
		Path:    path,
		Aliases: aliases,
		Globals: names,
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

				<link rel="stylesheet"
					href="https://storage.googleapis.com/code.getmdl.io/1.0.0/material.indigo-pink.min.css">
				<script src="https://storage.googleapis.com/code.getmdl.io/1.0.0/material.min.js"></script>
				<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">

				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
				<style>
					.node {
						position: relative;
						margin-left: 30px;
					}
					.root {
						margin-left: 0px;
					}
					.toggle {
						position: absolute;
						left: -25px;
						top: 0px;
						width: 15px;
					}
					.toggle svg {
						opacity: 0.5;
						height: 20px;
						width: 20px;
					}
				</style>
			</head>
			<body id="body" info="` + attribute + `">
			</body>
			<script src="script.js"></script>
		</html>`)

	if _, err := w.Write(source); err != nil {
		return kerr.New("ICJSAIMDRF", err, "w.Write (root)")
	}

	return nil
}

func serve(finished chan bool) error {

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

	finished <- true
	return nil

}
