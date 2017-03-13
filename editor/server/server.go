package server // import "kego.io/editor/server"

import (
	"bytes"
	"net"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/dave/gopackages"
	"github.com/dave/kerr"

	"fmt"

	"strings"

	"net/rpc"

	"encoding/base64"
	"encoding/gob"

	"time"

	"io"

	"context"

	"github.com/pkg/browser"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/context/wgctx"
	"kego.io/editor/server/auther"
	"kego.io/editor/server/static"
	"kego.io/editor/shared"
	"kego.io/process"
	"kego.io/process/generate"
	"kego.io/process/parser"
	"kego.io/system"
)

const writeTimeout = time.Second * 2

func Start(ctx context.Context, cancel context.CancelFunc) error {

	wgctx.Add(ctx, "Start")
	defer wgctx.Done(ctx, "Start")

	auth := auther.New()

	cmd := cmdctx.FromContext(ctx)

	fail := make(chan error)

	cmd.Println("Starting editor server... ")

	// This contains the source map that will be persisted between requests
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/favicon.ico") {
			w.WriteHeader(404)
			return
		}
		if strings.HasSuffix(req.URL.Path, "/script.js") {
			if err := script(ctx, w, req, false); err != nil {
				fail <- kerr.Wrap("XPVTVKDWHJ", err)
				return
			}
			return
		}
		if strings.HasSuffix(req.URL.Path, "/script.js.map") {
			if err := script(ctx, w, req, true); err != nil {
				fail <- kerr.Wrap("JAIBRHULSI", err)
				return
			}
			return
		}
		if err := root(ctx, w, req, auth); err != nil {
			fail <- kerr.Wrap("QOMJGNOCQF", err)
			return
		}
	})

	if err := rpc.Register(&Server{ctx: ctx, auth: auth}); err != nil {
		return kerr.Wrap("RESBVVGRMH", err)
	}

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
}

func script(ctx context.Context, w http.ResponseWriter, req *http.Request, sourceMap bool) error {

	wgctx.Add(ctx, "script")
	defer wgctx.Done(ctx, "script")

	path := ""
	if strings.HasSuffix(req.URL.Path, ".map") {
		path = req.URL.Path[1 : len(req.URL.Path)-14]
	} else {
		path = req.URL.Path[1 : len(req.URL.Path)-10]
	}

	env, err := parser.ScanForEnv(ctx, path)

	// This is the client code for the editor which we will compile to
	// Javascript using GopherJs below.
	source, err := generate.Editor(ctx, env)
	if err != nil {
		return kerr.Wrap("UWPDBQXURR", err)
	}

	var out []byte
	withCancel(ctx, func() {
		out, err = Compile(ctx, source, sourceMap)
	})
	if err != nil {
		return kerr.Wrap("LSUXHJMSSX", err)
	}

	withCancel(ctx, func() {
		err = writeWithTimeout(w, out)
	})
	if err != nil {
		return kerr.Wrap("GAOWWBAIAL", err)
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

func writeWithTimeout(w io.Writer, b []byte) error {
	c := make(chan error, 1)
	go func() {
		_, err := w.Write(b)
		c <- err
	}()
	select {
	case err := <-c:
		if err != nil {
			return kerr.Wrap("MJDFBJLCIT", err)
		}
		return nil
	case <-time.After(writeTimeout):
		return kerr.New("RTEATWMHDT", "Timeout")
	}
}

func root(ctx context.Context, w http.ResponseWriter, req *http.Request, auth auther.Auther) error {

	wgctx.Add(ctx, "root")
	defer wgctx.Done(ctx, "root")

	if b, err := static.Asset(req.URL.Path[1:]); err == nil {
		if strings.HasSuffix(req.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		if err := writeWithTimeout(w, b); err != nil {
			return kerr.Wrap("PVPCXBIUJT", err)
		}
		return nil
	}

	path := req.URL.Path[1:]
	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}

	// use a new context with a blank sysctx for the duration of this function
	// to prevent caching
	ctx = sysctx.NewContext(ctx)

	scache := sysctx.FromContext(ctx)

	env, err := parser.ScanForEnv(ctx, path)
	if err != nil {
		if _, ok := kerr.Source(err).(gopackages.NotFoundError); ok {
			w.WriteHeader(404)
			return nil
		}
		return kerr.Wrap("ALINBMKDRP", err)
	}

	pcache, err := parser.Parse(ctx, path)
	if err != nil {
		return kerr.Wrap("HIHWJRPUKE", err)
	}

	if err := process.GenerateAll(ctx, env.Path, map[string]bool{}); err != nil {
		return kerr.Wrap("LVGHABDYNQ", err)
	}

	data := map[string]string{}
	for _, name := range pcache.Globals.Keys() {
		if g, ok := pcache.Globals.Get(name); ok {
			data[name] = g.File
		}
	}

	pkgBytes := pcache.PackageBytes
	pkgFilename := pcache.PackageFilename
	if pkgBytes == nil {
		b, err := system.Marshal(ctx, system.EmptyPackage())
		if err != nil {
			return kerr.Wrap("OUBOTYGPKU", err)
		}
		pkgBytes = b
		pkgFilename = "package.ke.json"
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
			Types:   map[string]shared.TypeInfo{},
		}
		for _, name := range syspi.Files.Keys() {
			if b, ok := syspi.Files.Get(name); ok {
				info.Types[name] = shared.TypeInfo{
					File:  b.File,
					Bytes: b.Bytes,
				}
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
		Path:            env.Path,
		Aliases:         env.Aliases,
		Data:            data,
		Package:         pkgBytes,
		PackageFilename: pkgFilename,
		Imports:         imports,
		Hash:            auth.Sign([]byte(env.Path)),
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
				<link rel="stylesheet" href="/editors.css">
				<link rel="stylesheet" href="/tree.css">
				<script src="/jquery-2.2.4.min.js"></script>
				<script src="/jquery-ui/jquery-ui.min.js"></script>
				<script src="/split.min.js"></script>
				<script src="/bootstrap/js/bootstrap.min.js"></script>
				<link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo=">
			</head>
			<body id="body" info="` + attrib + `"></body>
			<script src="/` + env.Path + `/script.js"></script>
		</html>`)

	if err := writeWithTimeout(w, source); err != nil {
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
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", cmd.Port))
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

	cmd.Printf("Server now running on %s\n", url)

	// We open the default browser and navigate to the address we're serving
	// from. This will error if the system doesn't have a browser, so we can
	// ignore the error.
	browser.OpenURL(url)

	withCancel(ctx, func() {
		err = http.Serve(listner, nil)
	})
	if err != nil {
		return kerr.Wrap("TUCBTWMRNN", err)
	}

	if cmd.Debug {
		return kerr.New("ATUTBOICGJ", "Connection closed")
	}
	return nil

}
