package items

import "kego.io/editor/client/tree"

// Holders (e.g. 'data' and 'types') hold groups of source nodes
type holder struct {
	branch *tree.Branch

	name string
	pkg  *pkg
}

var _ tree.Item = (*holder)(nil)

func (h *holder) Initialise() {
	h.branch.SetLabel(h.name)
}

func (parent *pkg) addHolder(name string) *holder {
	h := &holder{name: name, pkg: parent}
	h.branch = tree.NewBranch(h, parent.branch)
	h.Initialise()

	parent.branch.Append(h.branch)

	return h
}
