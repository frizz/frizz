package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
	"kego.io/system"
)

type pkg struct {
	*item
	node  *system.Node
	label *dom.HTMLDivElement
	data  *holder
	types *holder
}

var _ tree.Item = (*pkg)(nil)
var _ tree.HasNode = (*pkg)(nil)

func (p *pkg) Node() *system.Node {
	return p.node
}

func (p *pkg) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(p.tree.Path)
	p.label = label
	div.AppendChild(label)
}

func AddPackage(node *system.Node, parentBranch *tree.Branch, sourcesData []string, sourcesTypes []string) *pkg {
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
