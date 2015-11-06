package branch

import (
	"fmt"

	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Entry items are nodes. Each branch inside a source branch are entry.
type entry struct {
	*Common
	name   string
	index  int
	node   *node.Node
	editor editor.Editor
}

var _ tree.Editable = (*source)(nil)

func (e *entry) Editor() editor.Editor {

	if ed, ok := e.node.Value.(editor.Editable); ok {
		return ed.GetEditor(e.node)
	}

	if factory := editor.Get(*e.node.Type.Id); factory != nil {
		return factory(e.node)
	}

	return editor.Default(e.node)
}

func (e *entry) Initialise(parent tree.Branch) {

	label := ""
	if e.index > -1 {
		label = fmt.Sprint("[", e.index, "]")
	} else {
		label = e.name
	}

	e.Common = &Common{this: e, tree: parent.Tree()}
	e.Common.Initialise(parent)
	e.Common.SetLabel(label)

}

func addEntry(name string, index int, node *node.Node, parent tree.Branch) *entry {

	e := &entry{name: name, index: index, node: node}
	e.Initialise(parent)
	parent.Append(e)

	addEntryChildren(node, e)

	e.Close()

	return e
}

func addEntryChildren(node *node.Node, parent tree.Branch) {
	if node == nil {
		return
	}
	switch node.Type.Native.Value() {
	case "array":
		for i, childNode := range node.Array {
			addEntry("", i, childNode, parent)
		}
	case "map":
		for name, childNode := range node.Map {
			addEntry(name, -1, childNode, parent)
		}
	case "object":
		for name, childNode := range node.Fields {
			if childNode.Missing {
				continue
			}
			addEntry(name, -1, childNode, parent)
		}
	}
}
