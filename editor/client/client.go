package client // import "kego.io/editor/client"

import (
	"net/url"

	"encoding/json"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
	"kego.io/editor/shared"
	"kego.io/editor/tree"
	"kego.io/kerr"
)

var window dom.Window
var doc dom.Document
var body dom.Element

func Start(path string) error {

	window = dom.GetWindow()
	doc = window.Document()
	body = doc.GetElementByID("body")
	window.AddEventListener("hashchange", false, navigate)

	// We ping the server every 500ms. When we close the browser
	// the server will exit after a 2sec timeout.
	window.SetInterval(keepAlive, 500)

	// We parse the json info attribute from the body tag
	info, err := getInfo(body)
	if err != nil {
		return kerr.New("MGLVIQIDDY", err, "getInfo")
	}

	// We create a new root tree element
	root := tree.New(body)

	// We add a new node to the tree for each global
	for _, g := range info.Globals {
		addGlobal(g, root, info.Path, info.Aliases)
	}

	/*
		err = root.Each(func(n *tree.Node) error {
			if err := n.WriteDom(body, doc); err != nil {
				return kerr.New("EWSIOWILSD", err, "WriteDom")
			}
			return nil
		})
		if err != nil {
			return kerr.New("MYWBOWUMIY", err, "root.Each")
		}
	*/

	/*for _, e := range doc.GetElementsByClassName("global") {
		href := e.GetAttribute("href") // must read this now to get it in the closure
		e.AddEventListener("click", true, func(ev dom.Event) {
			ev.PreventDefault()
			go func() {
				if err := loadFromServer(href, path, aliases, root); err != nil {
					console.Error(err.Error())
				}
			}()
		})
	}*/

	return nil
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

// TODO: do this instead with a websocket.
func keepAlive() {
	go func() { xhr.Send("GET", "/_ping", nil) }()
}

func navigate(e dom.Event) {
	id := dom.GetWindow().Location().Hash[1:]
	element := doc.GetElementByID(id)
	div := doc.CreateElement("div")
	div.SetInnerHTML(id)
	if element.NextSibling() != nil {
		body.InsertBefore(div, element.NextSibling())
	} else {
		body.AppendChild(div)
	}
}
