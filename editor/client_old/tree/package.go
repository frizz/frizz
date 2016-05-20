package tree

import (
	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
	"kego.io/editor/client_old"
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

func (parent *Root) AddPackage(node *editor.Node, data map[string]string, types map[string][]byte) error {

	ed := node.Editor()

	p := &pkg{Node: node}
	p.Branch = NewBranch(p, parent)

	env := envctx.FromContext(p.tree.ctx)

	p.setLabel(env.Path)
	parent.Append(p)

	if err := ed.Initialize(p.tree.ctx, p, editor.Page, p.tree.Fail); err != nil {
		return kerr.Wrap("NMIESKDFVN", err)
	}

	p.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(node, p, ed); err != nil {
		return kerr.Wrap("RBISWQVLFN", err)
	}

	dataHolder := p.addHolder("data")
	dataHolder.open()
	p.data = dataHolder
	p.data.addSources(data)

	typesHolder := p.addHolder("types")
	p.types = typesHolder
	p.types.addTypes(types)

	p.open()

	return nil
}
