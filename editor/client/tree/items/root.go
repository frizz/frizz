package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
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
