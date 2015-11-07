package items

import (
	"kego.io/editor"
	"kego.io/editor/client/tree"
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

func (parent *Root) AddPackage(node *editor.Node, sourcesData []string, sourcesTypes []string) *pkg {
	p := &pkg{Node: node}
	p.Initialise(parent.branch)

	parent.branch.Append(p.branch)

	addEntryChildren(node, p.branch, node.Editor())

	data := p.addHolder("data")
	data.branch.Open()
	p.data = data
	p.data.addSources(sourcesData)

	types := p.addHolder("types")
	p.types = types
	p.types.addSources(sourcesTypes)

	p.branch.Open()

	return p
}

var _ tree.Editable = (*pkg)(nil)
