package client // import "kego.io/editor/client"

import (
	"net/url"

	"encoding/json"

	"fmt"

	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
	"kego.io/editor/client/items"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared"
	"kego.io/editor/shared/connection"
	"kego.io/js/console"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system"
)

type appData struct {
	path    string
	aliases map[string]string
	fail    chan error
	conn    *connection.Conn
}

var app appData
var window dom.Window
var doc dom.Document
var body dom.Element

func Start(path string) error {

	window = dom.GetWindow()
	doc = window.Document()
	body = doc.GetElementByID("body")

	// We parse the json info attribute from the body tag
	info, err := getInfo(body)
	if err != nil {
		return kerr.New("MGLVIQIDDY", err, "getInfo")
	}

	app.path = path
	app.aliases = info.Aliases
	app.fail = make(chan error)

	packageNode := &system.Node{}
	if err := ke.UnmarshalNode([]byte(info.Package), packageNode, app.path, app.aliases); err != nil {
		return kerr.New("KXIKEWOKJI", err, "UnmarshalNode")
	}

	// We dial the websocket connection to the server
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/_socket", window.Location().Hostname, window.Location().Port))
	if err != nil {
		return kerr.New("XBMAKPJICG", err, "websocket.Dial")
	}

	// socket allows us to specify the message type - binary or string.
	s := &socket{ws, connection.MESSAGE_TYPE}

	app.conn = connection.New(s, app.fail, app.path, app.aliases)
	go handle(app.conn.Receive)

	nav := body.GetElementsByClassName("mdl-navigation")[0].(*dom.BasicHTMLElement)
	content := body.GetElementsByClassName("page-content")[0].(*dom.HTMLDivElement)
	// We create a new root tree element
	t := tree.New(nav, content, app.conn, items.Root(), app.fail, app.path, app.aliases)
	items.AddPackage(packageNode, t.Root, info.Sources)

	window.AddEventListener("keydown", true, func(e dom.Event) {
		t.KeyboardEvent(e.(*dom.KeyboardEvent))
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
		return shared.Info{}, kerr.New("CENGCYKHHP", err, "url.QueryUnescape (info)")
	}
	err = json.Unmarshal([]byte(infoJson), &info)
	if err != nil {
		return shared.Info{}, kerr.New("AAFXLQRUEW", err, "json.Unmarshal (info)")
	}
	return info, nil
}
