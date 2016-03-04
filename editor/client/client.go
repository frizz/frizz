package client // import "kego.io/editor/client"

// ke: {"package": {"jstest": true}}

import (
	"bytes"
	"net/rpc"

	"fmt"

	"encoding/gob"

	"encoding/base64"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/websocket"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/editor"
	"kego.io/editor/client/components"
	"kego.io/editor/client/connection"
	"kego.io/editor/shared"
	"kego.io/kerr"
	"kego.io/process/parser"
	"kego.io/system"
)

type appData struct {
	fail    chan error
	spinner *dom.HTMLDivElement
	ctx     context.Context
	env     *envctx.Env
	conn    *connection.Conn
}

var app appData
var window dom.Window
var doc dom.HTMLDocument
var body *dom.HTMLBodyElement

func Start() error {

	// We parse the json info attribute from the body tag
	info, err := getInfo()
	if err != nil {
		return kerr.Wrap("MGLVIQIDDY", err)
	}

	app.env = &envctx.Env{
		Path:    info.Path,
		Aliases: info.Aliases,
	}
	app.fail = make(chan error)
	app.ctx = envctx.NewContext(context.Background(), app.env)
	app.ctx = sysctx.NewContext(app.ctx)
	app.ctx = jsonctx.AutoContext(app.ctx)

	pcache, err := registerTypes(app.ctx, app.env.Path, info.Imports)
	if err != nil {
		return kerr.Wrap("MMJDDOBAUK", err)
	}

	types := map[string][]byte{}
	for _, name := range pcache.Files.Keys() {
		if b, ok := pcache.Files.Get(name); ok {
			types[name] = b
		}
	}

	packageNode, err := editor.UnmarshalNode(app.ctx, []byte(info.Package))
	if err != nil {
		return kerr.Wrap("KXIKEWOKJI", err)
	}

	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/_rpc", window.Location().Hostname, window.Location().Port))
	if err != nil {
		return kerr.Wrap("HNQFLPFAJD", err)
	}

	app.conn = connection.New(rpc.NewClient(ws))

	p := &components.PageView{
		Environment: app.env,
		Package:     packageNode,
	}
	vecty.RenderAsBody(p)
	js.Global.Get("window").Call("eval", "Split(['#tree', '#main'], {sizes:[25, 75]});")
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

		window.AddEventListener("keydown", true, func(e dom.Event) {
			k := e.(*dom.KeyboardEvent)
			switch doc.ActiveElement().TagName() {
			case "INPUT", "TEXTAREA":
				if k.KeyCode == 27 {
					doc.ActiveElement().Blur()
				}
				return
			default:
				t.KeyboardEvent(k)
			}
		})

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

func getInfo() (info shared.Info, err error) {
	body = dom.
		GetWindow().
		Document().(dom.HTMLDocument).
		GetElementByID("body").(*dom.HTMLBodyElement)
	infoBase64 := body.GetAttribute("info")
	infoBytes, err := base64.StdEncoding.DecodeString(infoBase64)
	if err != nil {
		return shared.Info{}, kerr.Wrap("UTKDDLYKKH", err)
	}
	buf := bytes.NewBuffer(infoBytes)
	if err := gob.NewDecoder(buf).Decode(&info); err != nil {
		return shared.Info{}, kerr.Wrap("AAFXLQRUEW", err)
	}
	return info, nil
}
