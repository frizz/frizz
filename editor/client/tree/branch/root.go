package branch

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

type Root struct {
	*Common
}

func NewRoot(tree *tree.Tree, nav *dom.BasicHTMLElement) *Root {

	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node root")

	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")

	nav.AppendChild(element)
	element.AppendChild(inner)

	r := &Root{}
	r.Common = &Common{this: r, tree: tree, root: true, open: true, element: element, inner: inner}
	return r

}

func (parent *Root) AddPackage(node *node.Node, sourcesData []string, sourcesTypes []string) *pkg {
	p := &pkg{node: node}
	p.Initialise(parent)

	parent.Append(p)

	addEntryChildren(node, p)

	data := p.addHolder("data")
	data.Open()
	p.data = data
	p.data.addSources(sourcesData)

	types := p.addHolder("types")
	p.types = types
	p.types.addSources(sourcesTypes)

	p.Open()

	return p
}
