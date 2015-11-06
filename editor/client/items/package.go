package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Pkg is the top-level item. It holds the entry items for th package object, and
// the "data" and "types" holder items.
type pkg struct {
	*item
	node  *node.Node
	label *dom.HTMLSpanElement
	data  *holder
	types *holder
}

var _ tree.Item = (*pkg)(nil)
var _ tree.Editable = (*source)(nil)

func (p *pkg) Editor() editor.Editor {
	return editor.Default(p.node)
}

func (p *pkg) Initialise(label *dom.HTMLSpanElement) {
	label.SetTextContent(p.tree.Path)
	p.label = label
}

func AddPackage(node *node.Node, parentBranch *tree.Branch, sourcesData []string, sourcesTypes []string) *pkg {
	newPackage := &pkg{item: &item{tree: parentBranch.Tree}, node: node}
	newBranch := parentBranch.Tree.Branch(newPackage)
	newPackage.branch = newBranch

	parentBranch.Append(newBranch)

	addNodeChildren(node, newBranch)

	data := newPackage.addHolder("data")
	data.branch.Open()
	newPackage.data = data
	newPackage.data.addSources(sourcesData)

	types := newPackage.addHolder("types")
	newPackage.types = types
	newPackage.types.addSources(sourcesTypes)

	newBranch.Open()

	return newPackage
}
