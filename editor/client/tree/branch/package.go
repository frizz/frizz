package branch

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Pkg is the top-level item. It holds the entry items for th package object, and
// the "data" and "types" holder items.
type pkg struct {
	*Common
	node  *node.Node
	label *dom.HTMLSpanElement
	data  *holder
	types *holder
}

var _ tree.Editable = (*source)(nil)

func (p *pkg) Editor() editor.Editor {
	return editor.Default(p.node)
}

func (p *pkg) Initialise(parent tree.Branch) {
	p.Common = &Common{this: p, tree: parent.Tree()}
	p.Common.Initialise(parent)
	p.Common.SetLabel(p.tree.Path)
}
