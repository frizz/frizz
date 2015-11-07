package items

import (
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Pkg is the top-level item. It holds the entry items for th package object, and
// the "data" and "types" holder items.
type pkg struct {
	branch *tree.Branch

	node  *node.Node
	data  *holder
	types *holder
}

var _ tree.Item = (*pkg)(nil)

func (p *pkg) Initialise() {
	p.branch.SetLabel(p.branch.Tree.Path)
}

var _ tree.Editable = (*pkg)(nil)

func (p *pkg) Editor() editor.Editor {
	return editor.Default(p.node)
}
