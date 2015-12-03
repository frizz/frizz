package tree

// Holders (e.g. 'data' and 'types') hold groups of source nodes
type holder struct {
	*Branch

	name string
	pkg  *pkg
}

var _ BranchInterface = (*holder)(nil)

func (parent *pkg) addHolder(name string) *holder {
	h := &holder{name: name, pkg: parent}
	h.Branch = NewBranch(h, parent)
	h.setLabel(name)
	parent.Append(h)
	return h
}
