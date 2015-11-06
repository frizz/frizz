package branch

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
)

// Holders (e.g. 'data' and 'types') hold groups of source nodes
type holder struct {
	*Common
	pkg   *pkg
	label *dom.HTMLSpanElement
	name  string
}

func (h *holder) Initialise(parent tree.Branch) {
	h.Common = &Common{this: h, tree: parent.Tree()}
	h.Common.Initialise(parent)
	h.Common.SetLabel(h.name)
}

func (parent *pkg) addHolder(name string) *holder {
	h := &holder{name: name, pkg: parent}
	h.Initialise(parent)

	parent.Append(h)

	return h
}
