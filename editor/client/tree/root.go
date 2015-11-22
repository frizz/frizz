package tree

import "honnef.co/go/js/dom"

type Root struct {
	*Branch
}

var _ BranchInterface = (*Root)(nil)

func NewRoot(t *Tree, nav *dom.BasicHTMLElement) *Root {

	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node root")

	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")

	nav.AppendChild(element)
	element.AppendChild(inner)

	r := newRoot(t, element, inner)

	t.Root = r

	return r

}

func newRoot(tree *Tree, element *dom.HTMLDivElement, inner *dom.HTMLDivElement) *Root {
	r := &Root{}
	r.Branch = &Branch{
		tree:    tree,
		open:    true,
		element: element,
		inner:   inner,
		self:    r,
	}
	return r
}
