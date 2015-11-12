package items

import (
	"fmt"

	"kego.io/editor"
	"kego.io/editor/client/tree"
)

// Entry items are nodes. Each branch inside a source branch are entry.
type entry struct {
	*editor.Node
	branch *tree.Branch

	name  string
	index int
}

var _ tree.Item = (*entry)(nil)

func (e *entry) Initialise(parent *tree.Branch) {
	label := ""
	if e.index > -1 {
		label = fmt.Sprint("[", e.index, "]")
	} else {
		label = e.name
	}
	e.branch = tree.NewBranch(e, parent)
	e.branch.SetLabel(label)
}

func addEntry(name string, index int, node *editor.Node, parentBranch *tree.Branch, parentEditor editor.Editor) *entry {

	ed := node.Editor()
	if !parentEditor.AddChildTreeEntry(ed) {
		return nil
	}

	e := &entry{name: name, index: index, Node: node}
	e.Initialise(parentBranch)
	parentBranch.Append(e.branch)

	addEntryChildren(node, e.branch, ed)

	e.branch.Close()

	return e
}

func addEntryChildren(parentNode *editor.Node, parentBranch *tree.Branch, parentEditor editor.Editor) {
	if parentNode == nil {
		return
	}
	switch parentNode.Type.Native.Value() {
	case "array":
		for i, child := range parentNode.Array {
			addEntry("", i, &editor.Node{child}, parentBranch, parentEditor)
		}
	case "map":
		for name, child := range parentNode.Map {
			addEntry(name, -1, &editor.Node{child}, parentBranch, parentEditor)
		}
	case "object":
		for name, child := range parentNode.Map {
			if child.Missing {
				continue
			}
			addEntry(name, -1, &editor.Node{child}, parentBranch, parentEditor)
		}
	}
}

var _ tree.Editable = (*entry)(nil)
