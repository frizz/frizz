package tree

func (b *Branch) firstChild() *Branch {
	if len(b.children) == 0 {
		return nil
	}
	return b.children[0]
}

func (b *Branch) lastChild() *Branch {
	if len(b.children) == 0 {
		return nil
	}
	return b.children[len(b.children)-1]
}

// lastVisible returns the last visible descendant in list order. If the branch is closed or has no
// children, we return the branch itself.
func (b *Branch) lastVisible() *Branch {

	for b.opened && len(b.children) > 0 {
		// if the node is open, test it's last child
		b = b.lastChild()
	}
	// return the first node we find that's closed
	return b
}

// prevVisible returns the previous visible branch in list order
func (b *Branch) prevVisible() *Branch {
	if b.isFirstSibling() {
		// at the start of a list of children, the previous in list order is always the parent
		return b.parent
	}
	// in the middle of a list of children, the previous in list order is the lastVisible of the
	// previous sibling.
	return b.prevSibling().lastVisible()
}

// nextVisible returns the next visible branch in list order.
func (b *Branch) nextVisible(includeChildren bool) *Branch {
	if b.parent == nil {
		// if we're testing the root, always return it's first child, or nil if it's empty.
		if len(b.children) > 0 {
			return b.children[0]
		}
		return nil
	}
	if includeChildren && b.opened && len(b.children) > 0 {
		// If the branch is open and has children, we return it's first child.
		return b.children[0]
	}
	current := b
	for current.isLastSibling() {
		// if the node is the last of the siblings, test it's parent
		if current.parent == nil {
			// if we get to the root node, we're testing the last visible node, so we return nil
			return nil
		}
		current = current.parent
	}
	// return the next sibling of the first ancestor that has one
	return current.nextSibling()
}

func (b *Branch) nextSibling() *Branch {
	if b.isLastSibling() {
		return nil
	}
	if b.parent == nil {
		return nil
	}
	return b.parent.children[b.index+1]
}

func (b *Branch) prevSibling() *Branch {
	if b.isFirstSibling() {
		return nil
	}
	if b.parent == nil {
		return nil
	}
	return b.parent.children[b.index-1]
}

func (b *Branch) isFirstSibling() bool {
	return b.index == 0
}

func (b *Branch) isLastSibling() bool {
	if b.parent == nil {
		// Only ever one root node.
		return true
	}
	return b.index >= len(b.parent.children)-1
}

// isDescendantOf tells us if this branch is a direct descendant of the specified branch
func (b *Branch) isDescendantOf(ancestor *Branch) bool {
	current := b.parent
	for current != nil {
		if current == ancestor {
			return true
		}
		current = current.parent
	}
	return false
}

// isAncestorOf tells us if this branch is a direct ancestor of the specified branch
func (b *Branch) isAncestorOf(descendant *Branch) bool {
	current := descendant.parent
	for current != nil {
		if current == b {
			return true
		}
		current = current.parent
	}
	return false
}
