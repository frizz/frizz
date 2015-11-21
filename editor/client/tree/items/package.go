package items

import (
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/kerr"
)

// Pkg is the top-level item. It holds the entry items for th package object, and
// the "data" and "types" holder items.
type pkg struct {
	*editor.Node
	branch *tree.Branch

	data  *holder
	types *holder
}

var _ tree.Item = (*pkg)(nil)

func (p *pkg) Initialise(parent *tree.Branch) {
	p.branch = tree.NewBranch(p, parent)
	p.branch.SetLabel(p.branch.Tree.Path)
}

func (parent *Root) AddPackage(node *editor.Node, sourcesData []string, sourcesTypes []string) error {

	ed := node.Editor()

	p := &pkg{Node: node}
	p.Initialise(parent.branch)

	parent.branch.Append(p.branch)

	if err := ed.Initialize(p.branch, editor.Page, p.branch.Tree.Path, p.branch.Tree.Aliases); err != nil {
		return kerr.New("NMIESKDFVN", err, "Initialize")
	}

	p.branch.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(node, p.branch, ed); err != nil {
		return kerr.New("RBISWQVLFN", err, "addEntryChildren")
	}

	data := p.addHolder("data")
	data.branch.Open()
	p.data = data
	p.data.addSources(sourcesData)

	types := p.addHolder("types")
	p.types = types
	p.types.addSources(sourcesTypes)

	p.branch.Open()

	return nil
}

var _ tree.Editable = (*pkg)(nil)
