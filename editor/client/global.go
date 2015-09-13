package client

import (
	"reflect"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
	"kego.io/editor/tree"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system"
)

type global struct {
	name    string
	loaded  bool
	path    string
	aliases map[string]string
	label   *dom.HTMLDivElement
}

func (g *global) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(g.name)
	g.label = label
	div.AppendChild(label)
}

func (g *global) Open(node *tree.Node, success func(), fail func(error)) {
	if g.loaded {
		success()
		return
	}
	data, err := xhr.Send("GET", g.name, nil)
	if err != nil {
		fail(kerr.New("KHBCONMBSF", err, "xrh.Send"))
		return
	}
	var i interface{}
	if err = ke.Unmarshal(data, &i, g.path, g.aliases); err != nil {
		fail(kerr.New("GXMCMOPRFK", err, "ke.Unmarshal"))
		return
	}
	o, ok := i.(system.Object)
	if !ok {
		fail(kerr.New("HEHNRKNHIC", nil, "%T is not a system.Object", i))
		return
	}
	b := o.GetBase()
	t, ok := b.Type.GetType()
	if !ok {
		fail(kerr.New("FODKTVUAAS", nil, "Type not found"))
		return
	}
	rule := system.NewMinimalRuleHolder(t)
	child := &element{data: o, rule: rule, value: reflect.ValueOf(o), index: -1, name: b.Id.Name}
	if err := addChildren(child, node); err != nil {
		fail(kerr.New("HBLEFLWPOG", err, "addNodes (root)"))
		return
	}
	g.loaded = true
	success()
	return
}

func (g *global) AsyncOpen() bool {
	return true
}

func (g *global) ContentLoaded() bool {
	return g.loaded
}

var _ = tree.Item(&global{})

func addGlobal(name string, parent *tree.Node, path string, aliases map[string]string) {
	n := tree.NewNode(&global{name: name, path: path, aliases: aliases})
	parent.AppendNodes(n)
}
