package client

import (
	"fmt"

	"honnef.co/go/js/dom"

	"kego.io/editor/client/tree"
	"kego.io/helper"
	"kego.io/kerr"
)

type entry struct {
	name  string
	index int
	node  *helper.Node
	label *dom.HTMLDivElement
}

func (e *entry) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)

	name := ""
	if e.index > -1 {
		name = fmt.Sprint("[", e.index, "]")
	} else {
		name = e.name
	}

	value := ""
	if e.node.Type.Native.Value == "bool" {
		value = fmt.Sprintf("%s", e.node.ValueBool)
	} else if e.node.Type.Native.Value == "number" {
		value = fmt.Sprintf("%v", e.node.ValueNumber)
	} else if e.node.Type.Native.Value == "string" {
		value = shortenString(e.node.ValueString)
	}

	if value == "" {
		label.SetTextContent(name)
	} else {
		label.SetTextContent(fmt.Sprint(name, " = ", value))
	}

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

var _ tree.Item = (*entry)(nil)

func addEntry(en *entry, parent *tree.Node) error {
	n := tree.NewNode(en)
	parent.AppendNodes(n)
	n.Open()

	if en.node == nil {
		return nil
	}
	return addEntryChildren(en, n)
}

func addEntryChildren(en *entry, n *tree.Node) error {
	switch en.node.Type.Native.Value {
	case "array":
		for i, childHelperNode := range en.node.Array {
			child := &entry{index: i, node: childHelperNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("MJTEOBEYVQ", err, "addEntry (array)")
			}
		}
	case "map":
		for name, childHelperNode := range en.node.Map {
			child := &entry{name: name, index: -1, node: childHelperNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("RWNXXYDUCS", err, "addEntry (map)")
			}
		}
	case "object":
		for name, childHelperNode := range en.node.Fields {
			if childHelperNode.Missing {
				continue
			}
			child := &entry{name: name, index: -1, node: childHelperNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("VJOGBGNGIB", err, "addEntry (fields)")
			}
		}
	}
	return nil
}
