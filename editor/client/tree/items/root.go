package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

type Root struct {
	branch *tree.Branch
}

var _ tree.Item = (*Root)(nil)

func NewRoot(t *tree.Tree, nav *dom.BasicHTMLElement) *Root {

	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node root")

	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")

	nav.AppendChild(element)
	element.AppendChild(inner)

	r := &Root{}
	r.branch = tree.NewRootBranch(t, element, inner)

	t.Root = r.branch

	return r

}

func (parent *Root) AddPackage(node *node.Node, sourcesData []string, sourcesTypes []string) *pkg {
	p := &pkg{node: node}
	p.branch = tree.NewBranch(p, parent.branch)
	p.Initialise()

	parent.branch.Append(p.branch)

	addEntryChildren(node, p.branch)

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
