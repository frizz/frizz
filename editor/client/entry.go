package client

import (
	"fmt"

	"honnef.co/go/js/dom"

	"kego.io/editor/client/tree"
	"kego.io/kerr"
	"kego.io/system"
)

type entry struct {
	name  string
	index int
	node  *system.Node
	label *dom.HTMLDivElement
}

var _ tree.Item = (*entry)(nil)

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
		value = fmt.Sprintf("%v", e.node.ValueBool)
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

func addEntry(en *entry, parent *tree.Branch) error {
	n := tree.NewBranch(en)
	parent.AppendBranches(n)
	n.Open()

	if en.node == nil {
		return nil
	}
	return addEntryChildren(en, n)
}

func addEntryChildren(en *entry, n *tree.Branch) error {
	switch en.node.Type.Native.Value {
	case "array":
		for i, childNode := range en.node.Array {
			child := &entry{index: i, node: childNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("MJTEOBEYVQ", err, "addEntry (array)")
			}
		}
	case "map":
		for name, childNode := range en.node.Map {
			child := &entry{name: name, index: -1, node: childNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("RWNXXYDUCS", err, "addEntry (map)")
			}
		}
	case "object":
		for name, childNode := range en.node.Fields {
			if childNode.Missing {
				continue
			}
			child := &entry{name: name, index: -1, node: childNode}
			if err := addEntry(child, n); err != nil {
				return kerr.New("VJOGBGNGIB", err, "addEntry (fields)")
			}
		}
	}
	return nil
}
