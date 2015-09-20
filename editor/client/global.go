package client

import (
	"reflect"

	"honnef.co/go/js/dom"
	"kego.io/editor/shared/messages"
	"kego.io/editor/tree"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system"
)

type global struct {
	name   string
	loaded bool
	label  *dom.HTMLDivElement
	node   *tree.Node
}

func (g *global) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(g.name)
	g.label = label
	div.AppendChild(label)
}

func (g *global) LoadContent() chan bool {

	successChannel := make(chan bool)

	if g.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it. This has to be done with a goroutine.
		go func() { successChannel <- true }()
		return successChannel
	}

	responseChannel := app.conn.Request(messages.NewGlobalRequest(g.name), app.fail)

	go func() {
		if err := g.awaitGlobalResponse(responseChannel, successChannel); err != nil {
			app.fail <- kerr.New("ATCPPWKQOF", err, "awaitGlobalResponse")
		}
	}()

	return successChannel

}

func (g *global) awaitGlobalResponse(responseChannel chan messages.Message, successChannel chan bool) error {

	m := <-responseChannel

	gr, ok := m.(*messages.GlobalResponse)
	if !ok {
		return kerr.New("YCFRNPEYGI", nil, "%T is not a *messages.GlobalResponse", m)
	}

	var i interface{}
	if err := ke.Unmarshal([]byte(gr.Data.Value), &i, app.path, app.aliases); err != nil {
		return kerr.New("GXMCMOPRFK", err, "ke.Unmarshal")
	}

	o, ok := i.(system.Object)
	if !ok {
		return kerr.New("HEHNRKNHIC", nil, "%T is not a system.Object", i)
	}

	b := o.Object()
	t, ok := b.Type.GetType()
	if !ok {
		return kerr.New("FODKTVUAAS", nil, "Type not found")
	}

	rule := system.NewMinimalRuleHolder(t)
	child := &element{data: o, rule: rule, value: reflect.ValueOf(o), index: -1, name: b.Id.Name}
	if err := addChildren(child, g.node); err != nil {
		return kerr.New("HBLEFLWPOG", err, "addNodes (root)")
	}

	g.loaded = true
	successChannel <- true
	return nil
}

func (g *global) ContentLoaded() bool {
	return g.loaded
}

var _ = tree.Item(&global{})
var _ = tree.AsyncItem(&global{})

func addGlobal(name string, parent *tree.Node) {
	g := &global{name: name}
	n := tree.NewNode(g)
	g.node = n
	parent.AppendNodes(n)
}
