package editor // import "kego.io/editor"

import (
	"go/build"
	"io"
	"net"
	"net/http"
	"path/filepath"

	"github.com/pkg/browser"

	"fmt"

	gbuild "github.com/gopherjs/gopherjs/build"
	"kego.io/kerr"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func StartServer(path string) error {
	http.HandleFunc("/", hello)

	options := &gbuild.Options{CreateMapFile: true}
	dirs := append(filepath.SplitList(build.Default.GOPATH), build.Default.GOROOT)
	sourceFiles := http.FileServer(serveCommandFileSystem{options: options, dirs: dirs, sourceMaps: make(map[string][]byte)})

	//fmt.Printf("serving at http://localhost:%d\n", port)
	//fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), sourceFiles))

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return kerr.New("ALTNEBOFGB", err, "net.Listen")
	}
	defer l.Close()

	addr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return kerr.New("HSWSQFNQQH", nil, "Can't find address (l.Addr() is not *net.TCPAddr)")
	}

	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d/github.com/davelondon/ke_site/editor/", addr.Port)); err != nil {
		return kerr.New("MMMFFQDKQP", err, "browser.OpenUrl")
	}

	if err := http.Serve(l, sourceFiles); err != nil {
		return kerr.New("BXWMMJPTFR", err, "http.Serve")
	}

	return nil
}
