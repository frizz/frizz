package tree

type KeyboardActions interface {
	KeyboardSelectPrev()
	KeyboardSelectNext()
	KeyboardSelectParent()
	KeyboardClose() bool
	KeyboardOpen() bool
}

// KeyboardSelectNext attempts to select the previous visible branch in list order
func (b *Branch) KeyboardSelectPrev() {
	if p := b.prevVisible(); p != nil {
		p.Select(true)
	}
}

// KeyboardSelectNext attempts to select the next visible branch in list order
func (b *Branch) KeyboardSelectNext() {
	if n := b.nextVisible(true); n != nil {
		n.Select(true)
	}
}

// KeyboardSelectParent attempts to select the parent
func (b *Branch) KeyboardSelectParent() {
	if p := b.parent; p != nil {
		p.Select(true)
	}
}

// KeyboardClose attempts to close the branch, and returns true if the branch was closed. If it was
// already closed or not openable, we return false.
func (b *Branch) KeyboardClose() bool {
	if !b.opened || !b.canOpen() {
		return false
	}
	b.close()
	return true
}

// KeyboardOpen attempts to open the branch, and returns true if the branch was opened. If it was
// already open or not openable, we return false.
func (b *Branch) KeyboardOpen() bool {
	if b.opened || !b.canOpen() {
		return false
	}
	b.open()
	return true
}
