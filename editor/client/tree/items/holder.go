package items

import "kego.io/editor/client/tree"

// Holders (e.g. 'data' and 'types') hold groups of source nodes
type holder struct {
	branch *tree.Branch

	name string
	pkg  *pkg
}

var _ tree.Item = (*holder)(nil)

func (h *holder) Initialise(parent *tree.Branch) {
	h.branch = tree.NewBranch(h, parent)
	h.branch.SetLabel(h.name)
}

func (parent *pkg) addHolder(name string) *holder {
	h := &holder{name: name, pkg: parent}
	h.Initialise(parent.branch)
	parent.branch.Append(h.branch)
	return h
}
