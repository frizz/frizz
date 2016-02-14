package client // import "kego.io/editor/client"

// ke: {"package": {"jstest": true}}

import (
	"net/url"

	"fmt"

	"github.com/gopherjs/websocket"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/editor"
	"kego.io/editor/client/console"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared"
	"kego.io/editor/shared/connection"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/process/parser"
	"kego.io/system"
)

type appData struct {
	fail    chan error
	conn    *connection.Conn
	spinner *dom.HTMLDivElement
	ctx     context.Context
	env     *envctx.Env
}

var app appData
var window dom.Window
var doc dom.HTMLDocument
var body *dom.HTMLBodyElement

func Start() error {

	window = dom.GetWindow()
	doc = window.Document().(dom.HTMLDocument)
	body = doc.GetElementByID("body").(*dom.HTMLBodyElement)

	// We parse the json info attribute from the body tag
	info, err := getInfo(body)
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

	scache := sysctx.FromContext(app.ctx)

	system.RegisterJsonTypes(app.ctx)

	for _, info := range info.Imports {
		env := &envctx.Env{Path: info.Path, Aliases: info.Aliases}
		pcache := scache.SetEnv(env)
		for _, typeBytes := range info.Types {
			if err := parser.ProcessTypeSourceBytes(app.ctx, env, typeBytes, pcache, nil); err != nil {
				return kerr.Wrap("UJLXYWCVUC", err)
			}
		}
	}

	pcache, ok := scache.Get(app.env.Path)
	if !ok {
		return kerr.New("SRPHQPBBRX", "%s not found in sys ctx", app.env.Path)
	}
	types := map[string][]byte{}
	for _, name := range pcache.Files.Keys() {
		if b, ok := pcache.Files.Get(name); ok {
			types[name] = b
		}
	}

	editorNode, err := editor.UnmarshalNode(app.ctx, []byte(info.Package))
	if err != nil {
		return kerr.Wrap("KXIKEWOKJI", err)
	}

	// We dial the websocket connection to the server
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/_socket", window.Location().Hostname, window.Location().Port))
	if err != nil {
		return kerr.Wrap("XBMAKPJICG", err)
	}

	// socket allows us to specify the message type - binary or string.
	s := &socket{ws, connection.MESSAGE_TYPE}

	app.conn = connection.New(app.ctx, s, app.fail, false)
	go handle(app.conn.Receive)

	app.spinner = body.GetElementsByClassName("mdl-spinner")[0].(*dom.HTMLDivElement)

	nav := body.GetElementsByClassName("mdl-navigation")[0].(*dom.BasicHTMLElement)
	content := body.GetElementsByClassName("page-content")[0].(*dom.HTMLDivElement)

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

	return nil
}

func handle(f func() error) {
	if err := f(); err != nil {
		app.fail <- err
	}
}

func getInfo(body dom.Element) (info shared.Info, err error) {
	infoJson, err := url.QueryUnescape(body.GetAttribute("info"))
	if err != nil {
		return shared.Info{}, kerr.Wrap("CENGCYKHHP", err)
	}
	err = json.UnmarshalPlain([]byte(infoJson), &info)
	if err != nil {
		return shared.Info{}, kerr.Wrap("AAFXLQRUEW", err)
	}
	return info, nil
}
