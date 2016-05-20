package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/websocket"

	"github.com/davelondon/kerr"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pkg/browser"

	"fmt"

	"strings"

	"path/filepath"

	"net/rpc"

	"encoding/base64"
	"encoding/gob"

	"github.com/gopherjs/gopherjs/build"
	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/context/wgctx"
	"kego.io/editor/server/static"
	"kego.io/editor/shared"
	"kego.io/ke"
	"kego.io/process/generate"
	"kego.io/process/packages"
	"kego.io/process/parser"
	"kego.io/process/scanner"
	"kego.io/system"
)

func Start(ctx context.Context, cancel context.CancelFunc) error {

	wgctx.Add(ctx, "Start")
	defer wgctx.Done(ctx, "Start")

	cmd := cmdctx.FromContext(ctx)

	fail := make(chan error)

	cmd.Println("Starting editor server... ")

	// This contains the source map that will be persisted between requests
	var mapping []byte
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/favicon.ico") {
			w.WriteHeader(404)
			return
		}
		if strings.HasSuffix(req.URL.Path, "/script.js") {
			if err := script(ctx, w, req, false, &mapping); err != nil {
				fail <- kerr.Wrap("XPVTVKDWHJ", err)
				return
			}
			return
		}
		if strings.HasSuffix(req.URL.Path, "/script.js.map") {
			if err := script(ctx, w, req, true, &mapping); err != nil {
				fail <- kerr.Wrap("JAIBRHULSI", err)
				return
			}
			return
		}
		if err := root(ctx, w, req); err != nil {
			fail <- kerr.Wrap("QOMJGNOCQF", err)
			return
		}
	})

	rpc.Register(&Server{ctx: ctx})

	http.Handle("/_rpc", websocket.Handler(func(ws *websocket.Conn) {
		ws.PayloadType = websocket.BinaryFrame
		rpc.ServeConn(ws)
	}))

	go func() {
		if err := serve(ctx); err != nil {
			fail <- err
		}
	}()

	done := ctx.Done()
	for {
		select {
		case <-done:
			fmt.Println("Exiting editor server (interupted)... ")
			return nil
		case err, open := <-fail:
			if !open {
				// Channel has been closed, so app should gracefully exit.
				cancel()
				cmd.Println("Exiting editor server (finished)... ")
			} else {
				// Error received, so app should display error.
				// return kerr.New("WKHPTVJBIL", err, "Fail channel receive")
				fmt.Println(err)
			}
			if !cmd.Debug {
				return nil
			}
		}
	}

	return nil
}

type Server struct {
	ctx context.Context
}

func (s *Server) Data(request *shared.DataRequest, response *shared.DataResponse) error {

	env, err := parser.ScanForEnv(s.ctx, request.Package)
	if err != nil {
		return kerr.Wrap("PNAGGKHDYL", err)
	}

	file := filepath.Join(env.Dir, request.File)

	bytes, err := scanner.ProcessFile(file)

	localContext := envctx.NewContext(s.ctx, env)
	o := &system.Object{}
	if err := ke.UnmarshalUntyped(localContext, bytes, o); err != nil {
		return kerr.Wrap("SVINFEMKBG", err)
	}
	if o.Id.Name != request.Name {
		return kerr.New("TNJSLMPMLB", "Id does not match")
	}

	response.Package = request.Package
	response.Name = request.Name
	response.Found = true
	response.Data = bytes
	return nil
}

func script(ctx context.Context, w http.ResponseWriter, req *http.Request, mapper bool, mapping *[]byte) error {

	wgctx.Add(ctx, "script")
	defer wgctx.Done(ctx, "script")

	path := ""
	if strings.HasSuffix(req.URL.Path, ".map") {
		path = req.URL.Path[1 : len(req.URL.Path)-14]
	} else {
		path = req.URL.Path[1 : len(req.URL.Path)-10]
	}

	env, err := parser.ScanForEnv(ctx, path)

	// This is the client code for the editor which we will compile to Javascript using GopherJs
	// below. GopherJs doesn't make it easy to compile directly from a string, so we write the
	// source file to a temporary location and delete after we have compiled it to Javascript.
	source, err := generate.Editor(ctx, env)
	if err != nil {
		return kerr.Wrap("UWPDBQXURR", err)
	}

	editorPath, namespaceDir, err := CreateTemporaryPackage(ctx, "kego.temporary", "a", map[string]string{"a.go": string(source)})
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

	if mapper && mapping != nil {
		if _, err := w.Write(*mapping); err != nil {
			return kerr.Wrap("WFVDCWDVWL", err)
		}
		return nil
	}

	buf := bytes.NewBuffer(nil)
	pkg := &build.PackageData{Package: buildPkg.Package}

	type ret struct {
		archive *compiler.Archive
		err     error
	}

	c := make(chan ret, 1)
	go func() {
		a, err := s.BuildPackage(pkg)
		c <- ret{a, err}
	}()

	var r ret
	select {
	case r = <-c:
		if r.err != nil {
			return kerr.Wrap("TXUYQOUNQS", r.err)
		}
	case <-ctx.Done():
		return nil
	}

	filter := &compiler.SourceMapFilter{Writer: buf}
	smap := &sourcemap.Map{File: "script.js"}
	filter.MappingCallback = build.NewMappingCallback(smap, options.GOROOT, options.GOPATH)

	deps, err := compiler.ImportDependencies(r.archive, s.BuildImportPath)
	if err != nil {
		return kerr.Wrap("OVDUPSTRNR", err)
	}
	if err := compiler.WriteProgramCode(deps, filter); err != nil {
		return kerr.Wrap("YVHQEJXQGP", err)
	}

	mapBuf := bytes.NewBuffer(nil)
	if err := smap.WriteTo(mapBuf); err != nil {
		return kerr.Wrap("VYQGYAAADG", err)
	}
	*mapping = mapBuf.Bytes()

	if mapper {
		if _, err := w.Write(*mapping); err != nil {
			return kerr.Wrap("WKAMMWFXAN", err)
		}
		return nil
	}

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

	imports := map[string]shared.ImportInfo{}

	var scan func(string) error
	scan = func(path string) error {
		if _, ok := imports[path]; ok {
			return nil
		}
		syspi, ok := scache.Get(path)
		if !ok {
			return kerr.New("VIGKIUPNCF", "%s not found in sys ctx", path)
		}
		info := shared.ImportInfo{
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

		for _, child := range syspi.Aliases {
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
	buf := bytes.NewBuffer([]byte{})
	err = gob.NewEncoder(buf).Encode(info)
	if err != nil {
		return kerr.Wrap("OHBYTULHUQ", err)
	}
	base64EncodedString := base64.StdEncoding.EncodeToString(buf.Bytes())
	attrib := base64EncodedString

	source := []byte(`
		<html>
			<head>
				<meta charset="utf-8">
				<link rel="stylesheet" href="/bootstrap/css/bootstrap.min.css">
				<link rel="stylesheet" href="/bootstrap/css/bootstrap-theme.min.css">
				<link rel="stylesheet" href="/split.css">
				<link rel="stylesheet" href="/tree.css">
				<script src="/jquery-2.2.1.min.js"></script>
				<script src="/split.min.js"></script>
				<script src="/bootstrap/js/bootstrap.min.js"></script>
				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
			</head>
			<body id="body" info="` + attrib + `"></body>
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
	//listner, err := net.Listen("tcp", ":0")
	listner, err := net.Listen("tcp", ":8099")
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
