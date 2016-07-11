package models

func (b *BranchModel) FirstChild() *BranchModel {
	if len(b.Children) == 0 {
		return nil
	}
	return b.Children[0]
}

func (b *BranchModel) lastChild() *BranchModel {
	if len(b.Children) == 0 {
		return nil
	}
	return b.Children[len(b.Children)-1]
}

// lastVisible returns the last visible descendant in list order. If the branch is closed or has no
// children, we return the branch itself.
func (b *BranchModel) LastVisible() *BranchModel {

	for b.Open && len(b.Children) > 0 {
		// if the node is open, test it's last child
		b = b.lastChild()
	}
	// return the first node we find that's closed
	return b
}

// prevVisible returns the previous visible branch in list order
func (b *BranchModel) PrevVisible() *BranchModel {
	if b.isFirstSibling() {
		// at the start of a list of children, the previous in list order is always the parent
		return b.Parent
	}
	// in the middle of a list of children, the previous in list order is the lastVisible of the
	// previous sibling.
	return b.prevSibling().LastVisible()
}

// nextVisible returns the next visible branch in list order.
func (b *BranchModel) NextVisible(includeChildren bool) *BranchModel {
	if b.Parent == nil {
		// if we're testing the root, always return it's first child, or nil if it's empty.
		if len(b.Children) > 0 {
			if !b.Open {
				return nil
			}
			return b.Children[0]
		}
		return nil
	}
	if includeChildren && b.Open && len(b.Children) > 0 {
		// If the branch is open and has children, we return it's first child.
		return b.Children[0]
	}
	current := b
	for current.isLastSibling() {
		// if the node is the last of the siblings, test it's parent
		if current.Parent == nil {
			// if we get to the root node, we're testing the last visible node, so we return nil
			return nil
		}
		current = current.Parent
	}
	// return the next sibling of the first ancestor that has one
	return current.nextSibling()
}

func (b *BranchModel) nextSibling() *BranchModel {
	if b.isLastSibling() {
		return nil
	}
	if b.Parent == nil {
		return nil
	}
	return b.Parent.Children[b.Index+1]
}

func (b *BranchModel) prevSibling() *BranchModel {
	if b.isFirstSibling() {
		return nil
	}
	if b.Parent == nil {
		return nil
	}
	return b.Parent.Children[b.Index-1]
}

func (b *BranchModel) isFirstSibling() bool {
	return b.Index == 0
}

func (b *BranchModel) isLastSibling() bool {
	if b.Parent == nil {
		// Only ever one root node.
		return true
	}
	return b.Index >= len(b.Parent.Children)-1
}

// isDescendantOf tells us if this branch is a direct descendant of the specified branch
func (b *BranchModel) IsDescendantOf(ancestor *BranchModel) bool {
	current := b.Parent
	for current != nil {
		if current == ancestor {
			return true
		}
		current = current.Parent
	}
	return false
}

// isAncestorOf tells us if this branch is a direct ancestor of the specified branch
func (b *BranchModel) IsAncestorOf(descendant *BranchModel) bool {
	current := descendant.Parent
	for current != nil {
		if current == b {
			return true
		}
		current = current.Parent
	}
	return false
}
