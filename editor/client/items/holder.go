package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
)

// Holders (e.g. 'data' and 'types') hold groups of source nodes
type holder struct {
	*item
	pkg   *pkg
	label *dom.HTMLSpanElement
	name  string
}

var _ tree.Item = (*holder)(nil)

func (d *holder) Initialise(label *dom.HTMLSpanElement) {
	label.SetTextContent(d.name)
	d.label = label
}

func (p *pkg) addHolder(name string) *holder {
	newHolder := &holder{item: &item{tree: p.tree}, name: name, pkg: p}
	newBranch := newHolder.tree.Branch(newHolder)
	newHolder.branch = newBranch

	p.branch.Append(newBranch)

	return newHolder
}
