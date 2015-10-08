package items

import (
	"fmt"

	"honnef.co/go/js/dom"

	"kego.io/editor/client/tree"
	"kego.io/system"
)

type entry struct {
	name  string
	index int
	node  *system.Node
	label *dom.HTMLDivElement
}

var _ tree.Item = (*entry)(nil)

func (e *entry) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)

	name := ""
	if e.index > -1 {
		name = fmt.Sprint("[", e.index, "]")
	} else {
		name = e.name
	}

	label.SetTextContent(name)
	e.label = label
	div.AppendChild(label)
}

func shortenString(in string) string {
	var runes = 0
	for i, _ := range in {
		runes++
		if runes > 25 {
			return fmt.Sprint(in[:i], "...")
		}
	}
	return in
}

func addEntry(name string, index int, node *system.Node, parent *tree.Branch) *entry {
	e := &entry{name: name, index: index, node: node}
	b := parent.Tree.Branch(e)
	parent.Append(b)
	b.Open()

	if e.node != nil {
		addEntryChildren(e, b)
	}
	return e
}

func addEntryChildren(en *entry, n *tree.Branch) {
	switch en.node.Type.Native.Value {
	case "array":
		for i, childNode := range en.node.Array {
			addEntry("", i, childNode, n)
		}
	case "map":
		for name, childNode := range en.node.Map {
			addEntry(name, -1, childNode, n)
		}
	case "object":
		for name, childNode := range en.node.Fields {
			if childNode.Missing {
				continue
			}
			addEntry(name, -1, childNode, n)
		}
	}
}
