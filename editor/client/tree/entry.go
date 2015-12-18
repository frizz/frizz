package tree

import (
	"fmt"

	"kego.io/editor"
	"kego.io/kerr"
)

// Entry items are nodes. Each branch inside a source branch are entry.
type entry struct {
	*Branch
	*editor.Node
	name  string
	index int
}

var _ BranchInterface = (*entry)(nil)
var _ Editable = (*entry)(nil)

func addEntry(name string, index int, node *editor.Node, parentBranch BranchInterface, parentEditor editor.EditorInterface) (*entry, error) {

	ed := node.Editor()
	if !parentEditor.AddChildTreeEntry(ed) {
		return nil, nil
	}

	e := &entry{name: name, index: index, Node: node}
	e.Branch = NewBranch(e, parentBranch)

	if index > -1 {
		e.setLabel(fmt.Sprint("[", index, "]"))
	} else {
		e.setLabel(name)
	}

	parentBranch.Append(e)

	if err := ed.Initialize(e, editor.Page, e.tree.Fail, e.tree.Path, e.tree.Aliases); err != nil {
		return nil, kerr.New("PMLOGADEVK", err, "Initialize")
	}

	e.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(node, e, ed); err != nil {
		return nil, kerr.New("UPRWSRECVR", err, "addEntryChildren")
	}

	e.close()

	return e, nil
}

func addEntryChildren(parentNode *editor.Node, parentBranch BranchInterface, parentEditor editor.EditorInterface) error {
	if parentNode == nil {
		return nil
	}
	switch parentNode.Type.Native.Value() {
	case "array":
		for i, child := range parentNode.Array {
			if _, err := addEntry("", i, child, parentBranch, parentEditor); err != nil {
				return kerr.New("IOXSWBQDXH", err, "addEntry (array)")
			}
		}
	case "map":
		for name, child := range parentNode.Map {
			if _, err := addEntry(name, -1, child, parentBranch, parentEditor); err != nil {
				return kerr.New("YVTQCADGJF", err, "addEntry (map)")
			}
		}
	case "object":
		for name, child := range parentNode.Map {
			if child.Missing {
				continue
			}
			if _, err := addEntry(name, -1, child, parentBranch, parentEditor); err != nil {
				return kerr.New("SIBWLRIXRG", err, "addEntry (object)")
			}
		}
	}
	return nil
}
