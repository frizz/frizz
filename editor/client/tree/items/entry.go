package items

import (
	"fmt"

	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/kerr"
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

func addEntry(name string, index int, node *editor.Node, parentBranch *tree.Branch, parentEditor editor.Editor) error {

	ed := node.Editor()
	if !parentEditor.AddChildTreeEntry(ed) {
		return nil
	}

	e := &entry{name: name, index: index, Node: node}
	e.Initialise(parentBranch)
	parentBranch.Append(e.branch)

	if err := ed.Initialize(e.branch, editor.Page, e.branch.Tree.Path, e.branch.Tree.Aliases); err != nil {
		return kerr.New("PMLOGADEVK", err, "Initialize")
	}

	if err := addEntryChildren(node, e.branch, ed); err != nil {
		return kerr.New("UPRWSRECVR", err, "addEntryChildren")
	}

	e.branch.Close()

	return nil
}

func addEntryChildren(parentNode *editor.Node, parentBranch *tree.Branch, parentEditor editor.Editor) error {
	if parentNode == nil {
		return nil
	}
	switch parentNode.Type.Native.Value() {
	case "array":
		for i, child := range parentNode.Array {
			if err := addEntry("", i, child, parentBranch, parentEditor); err != nil {
				return kerr.New("IOXSWBQDXH", err, "addEntry (array)")
			}
		}
	case "map":
		for name, child := range parentNode.Map {
			if err := addEntry(name, -1, child, parentBranch, parentEditor); err != nil {
				return kerr.New("YVTQCADGJF", err, "addEntry (map)")
			}
		}
	case "object":
		for name, child := range parentNode.Map {
			if child.Missing {
				continue
			}
			if err := addEntry(name, -1, child, parentBranch, parentEditor); err != nil {
				return kerr.New("SIBWLRIXRG", err, "addEntry (object)")
			}
		}
	}
	return nil
}

var _ tree.Editable = (*entry)(nil)
