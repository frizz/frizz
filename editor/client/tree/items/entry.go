package items

import (
	"fmt"

	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Entry items are nodes. Each branch inside a source branch are entry.
type entry struct {
	*editor.Node
	branch *tree.Branch

	name   string
	index  int
	editor editor.Editor
}

var _ tree.Item = (*entry)(nil)

func (e *entry) Initialise(parent *tree.Branch) {
	label := ""
	if e.index > -1 {
		label = fmt.Sprint("[", e.index, "]")
	} else {
		label = e.name
	}
	e.branch.SetLabel(label)
}

func addEntry(name string, index int, node *node.Node, parent *tree.Branch) *entry {

	editNode := editor.Node{node}
	edit := editNode.Editor()
	if edit.Layout() != editor.Page {
		return nil
	}

	e := &entry{name: name, index: index, Node: &editor.Node{node}}
	e.branch = tree.NewBranch(e, parent)
	e.Initialise(parent)
	parent.Append(e.branch)

	addEntryChildren(node, e.branch)

	e.branch.Close()

	return e
}

func addEntryChildren(node *node.Node, parent *tree.Branch) {
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

var _ tree.Editable = (*entry)(nil)
