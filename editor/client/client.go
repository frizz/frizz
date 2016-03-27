package client // import "kego.io/editor/client"

// ke: {"package": {"jstest": true}}

import (
	"bytes"
	"net/rpc"

	"fmt"

	"encoding/gob"

	"encoding/base64"

	"github.com/davelondon/vecty"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/editor/shared"
	"kego.io/kerr"
	"kego.io/process/parser"
	"kego.io/system"
)

/*
type appData struct {
	fail    chan error
	spinner *dom.HTMLDivElement
	ctx     context.Context
	env     *envctx.Env
	conn    *connection.Conn
}

var app appData
*/

func Start() error {

	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/_rpc", dom.GetWindow().Location().Hostname, dom.GetWindow().Location().Port))
	if err != nil {
		return kerr.Wrap("HNQFLPFAJD", err)
	}

	// We parse the json info attribute from the body tag
	info, err := getInfo()
	if err != nil {
		return kerr.Wrap("MGLVIQIDDY", err)
	}

	env := &envctx.Env{
		Path:    info.Path,
		Aliases: info.Aliases,
	}
	app := &stores.App{
		Conn: connection.New(rpc.NewClient(ws)),
		Fail: make(chan error),
	}

	var ctx context.Context
	ctx = context.Background()
	ctx = envctx.NewContext(ctx, env)
	ctx = sysctx.NewContext(ctx)
	ctx = jsonctx.AutoContext(ctx)
	ctx = stores.NewContext(ctx, app)

	app.Init(ctx)

	pcache, err := registerTypes(ctx, env.Path, info.Imports)
	if err != nil {
		return kerr.Wrap("MMJDDOBAUK", err)
	}

	types := map[string][]byte{}
	for _, name := range pcache.Files.Keys() {
		if b, ok := pcache.Files.Get(name); ok {
			types[name] = b
		}
	}

	p := views.NewPage(ctx, env)
	vecty.RenderAsBody(p)

	// TODO: work out why I can't seem to call this without using eval
	js.Global.Get("window").Call("eval", "Split(['#tree', '#main'], {sizes:[25, 75]});")

	addKeyboardEvents(app)

	go func() {
		app.Dispatch(&actions.InitialState{
			Info: info,
		})
	}()

	go func() {
		err, open := <-app.Fail
		if !open {
			// Channel has been closed, so app should gracefully exit.
			fmt.Println("Server disconnected")
		} else {
			// Error received, so app should display error.
			fmt.Println(err.Error())
		}
	}()

	return nil
}
func StartOld() error {
	/*
		// We create a new root tree element
		t := tree.New(app.ctx, content, app.conn, app.fail)
		root := tree.NewRoot(t, nav)

		if err := root.AddPackage(editorNode, info.Data, types); err != nil {
			return kerr.Wrap("EAIHJLNBFA", err)
		}

		go func() {
			err, open := <-app.fail
			if !open {
				// Channel has been closed, so app should gracefully exit.
				body.SetInnerHTML("<pre>Server disconnected.</pre>")
			} else {
				// Error received, so app should display error.
				console.Error(err.Error())
				body.SetInnerHTML("<pre>Error:" + err.Error() + "</pre>")
			}
			app.conn.Close()
		}()

		app.spinner.Style().Set("display", "none")
	*/
	return nil
}

func registerTypes(ctx context.Context, path string, imports map[string]*shared.ImportInfo) (*sysctx.SysPackageInfo, error) {
	system.RegisterJsonTypes(ctx)
	scache := sysctx.FromContext(ctx)
	var current *sysctx.SysPackageInfo
	for _, info := range imports {
		env := &envctx.Env{Path: info.Path, Aliases: info.Aliases}
		pcache := scache.SetEnv(env)
		for _, typeBytes := range info.Types {
			if err := parser.ProcessTypeSourceBytes(ctx, env, typeBytes, pcache, nil); err != nil {
				return nil, kerr.Wrap("UJLXYWCVUC", err)
			}
		}
		if path == info.Path {
			current = pcache
		}
	}
	return current, nil
}

func getInfo() (info *shared.Info, err error) {
	info = &shared.Info{}
	infoBase64 := dom.
		GetWindow().
		Document().(dom.HTMLDocument).
		GetElementByID("body").(*dom.HTMLBodyElement).
		GetAttribute("info")
	infoBytes, err := base64.StdEncoding.DecodeString(infoBase64)
	if err != nil {
		return nil, kerr.Wrap("UTKDDLYKKH", err)
	}
	buf := bytes.NewBuffer(infoBytes)
	if err := gob.NewDecoder(buf).Decode(info); err != nil {
		return nil, kerr.Wrap("AAFXLQRUEW", err)
	}
	return info, nil
}

func addKeyboardEvents(app *stores.App) {
	window := dom.GetWindow()
	document := window.Document().(dom.HTMLDocument)
	window.AddEventListener("keydown", true, func(e dom.Event) {
		k := e.(*dom.KeyboardEvent)
		switch document.ActiveElement().TagName() {
		case "INPUT", "TEXTAREA":
			if k.KeyCode == 27 {
				// escape
				document.ActiveElement().Blur()
			}
			return
		default:
			if k.KeyCode >= 37 && k.KeyCode <= 40 {
				// up, down, left, right
				k.PreventDefault()
				go func() {
					app.Dispatch(&actions.KeyboardEvent{KeyCode: k.KeyCode})
				}()
			}
		}
	})
}
