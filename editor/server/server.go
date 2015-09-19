package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	kejson "kego.io/json"
	"kego.io/kerr"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"

	"fmt"

	"encoding/json"

	"net/url"

	"github.com/gopherjs/gopherjs/build"
	"kego.io/editor/shared"
	"kego.io/process"
	"kego.io/system"
)

func Start(path string, verbose bool) error {

	if verbose {
		fmt.Println("Starting editor... ")
	}

	aliases, err := initialise(path)
	if err != nil {
		return kerr.New("SWSQDFXIEV", err, "initialise")
	}

	// This contains the source map that will be persisted between requests
	var mapping []byte

	finished := make(chan bool)
	timeout := make(chan bool)
	ping := make(chan bool)
	fail := make(chan error)

	http.HandleFunc("/script.js", func(w http.ResponseWriter, req *http.Request) {
		if err := script(w, req, path, aliases, false, &mapping); err != nil {
			fail <- kerr.New("XPVTVKDWHJ", err, "script (js)")
		}
	})
	http.HandleFunc("/script.js.map", func(w http.ResponseWriter, req *http.Request) {
		if err := script(w, req, path, aliases, true, &mapping); err != nil {
			fail <- kerr.New("JRLOMPOHRQ", err, "script (map)")
		}
	})
	http.HandleFunc("/_ping", func(http.ResponseWriter, *http.Request) {
		ping <- true
	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := wildcard(w, req, path, aliases); err != nil {
			fail <- kerr.New("QOMJGNOCQF", err, "wildcard")
		}
	})

	go func() { serve(finished, fail) }()
	go func() { pong(ping, timeout, true) }()

	select {
	case <-finished:
		if verbose {
			fmt.Println("Exiting editor server (finished)... ")
		}
	case <-timeout:
		if verbose {
			fmt.Println("Exiting editor server (client closed)... ")
		}
	case err := <-fail:
		return kerr.New("WKHPTVJBIL", err, "Fail channel receive")
	}

	return nil
}

func initialise(path string) (aliases map[string]string, err error) {

	set, err := process.InitialiseManually(true, true, false, false, path)
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

	if len(req.URL.Path) > 1 {
		name := req.URL.Path[1:]
		hashed, ok := system.GetGlobal(path, name)
		if ok {
			b, err := kejson.Marshal(hashed.Object)
			if err != nil {
				return kerr.New("WTXNLMABAS", err, "kejson.Marshal")
			}
			if _, err := w.Write(b); err != nil {
				return kerr.New("OQVYRENYUX", err, "w.Write (data)")
			}
		} else {
			w.WriteHeader(404)
		}
		return nil
	}

	globals := system.GetAllGlobalsInPackage(path, nil)

	names := []string{}
	for _, hashed := range globals {
		name := hashed.Object.GetBase().Id.Name
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

func serve(finished chan bool, fail chan error) {

	// Starting with port zero chooses a random open port
	listner, err := net.Listen("tcp", ":0")
	if err != nil {
		fail <- kerr.New("QGLXHWPWQW", err, "net.Listen")
	}
	defer listner.Close()

	// Here we get the address we're serving on
	address, ok := listner.Addr().(*net.TCPAddr)
	if !ok {
		fail <- kerr.New("CBLPYVGGUR", nil, "Can't find address (l.Addr() is not *net.TCPAddr)")
	}

	// We open the default browser and navigate to the address we're serving from.
	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d/", address.Port)); err != nil {
		fail <- kerr.New("AEJLAXGVVA", err, "browser.OpenUrl")
	}

	if err := http.Serve(listner, nil); err != nil {
		fail <- kerr.New("TUCBTWMRNN", err, "http.Serve")
	}

	finished <- true
}

// pong starts a timeout. If a ping is not received before the timeout, the
// timeout channel is signaled. If the ping is received, the timeout starts
// again. The initial timeout is 20 seconds, to give the page time to
// initialize. Subsequent requests have a 3 second timeout.
func pong(ping chan bool, timeout chan bool, first bool) {

	duration := time.Second * 3
	if first {
		duration = time.Second * 20
	}

	select {
	case <-ping:
		pong(ping, timeout, false)
	case <-time.After(duration):
		timeout <- true
	}
}
