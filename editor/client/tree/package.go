package tree

import (
	"kego.io/context/envctx"
	"kego.io/editor"
	"kego.io/kerr"
)

// Pkg is the top-level item. It holds the entry items for th package object, and
// the "data" and "types" holder items.
type pkg struct {
	*Branch
	*editor.Node

	data  *holder
	types *holder
}

var _ BranchInterface = (*pkg)(nil)
var _ Editable = (*pkg)(nil)

func (parent *Root) AddPackage(node *editor.Node, sourcesData []string, sourcesTypes []string) error {

	ed := node.Editor()

	p := &pkg{Node: node}
	p.Branch = NewBranch(p, parent)

	env := envctx.FromContext(p.tree.ctx)

	p.setLabel(env.Path)
	parent.Append(p)

	if err := ed.Initialize(p.tree.ctx, p, editor.Page, p.tree.Fail); err != nil {
		return kerr.New("NMIESKDFVN", err, "Initialize")
	}

	p.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(node, p, ed); err != nil {
		return kerr.New("RBISWQVLFN", err, "addEntryChildren")
	}

	data := p.addHolder("data")
	data.open()
	p.data = data
	p.data.addSources(sourcesData)

	types := p.addHolder("types")
	p.types = types
	p.types.addSources(sourcesTypes)

	p.open()

	return nil
}
