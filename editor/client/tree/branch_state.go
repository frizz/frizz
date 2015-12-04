package tree

import "kego.io/editor/client/icons"

// Open opens the branch. For asynchronously loaded branches it initialises the load.
func (b *Branch) open() {

	if b.parent == nil {
		return
	}

	done, ok := b.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	// we wrap the success action in a closure, because it must be executed syncronously or in a
	// goroutine depending on the return from ensureContentLoaded
	success := func(fromAsync bool) {

		if b.inner != nil {
			b.inner.Style().Set("display", "block")
		}

		if fromAsync && !b.canOpen() {
			// if after loading the content, we can't we can't open the branch, this means we just
			// there were no children. In this case we want to select the branch.
			b.Select(false)
		} else {
			b.opened = true
			b.afterStateChange()
		}
	}

	if done == nil {
		// if the done chanel is nil, the operation was synchronous, so we should call success
		// synchronously
		success(false)
	} else {
		go func() {
			// block and wait until the response arrives
			<-done
			success(true)
		}()
	}
	return
}

// Close closes a branch and hides the children
func (b *Branch) close() {
	b.closeWithoutUpdate()
	b.afterStateChange()
	return
}

func (b *Branch) closeWithoutUpdate() {
	if b.parent == nil {
		return
	}
	if !b.opened {
		return
	}
	if b.inner != nil {
		b.inner.Style().Set("display", "none")
	}
	b.opened = false
	if async, ok := b.self.(AsyncInterface); ok {
		async.Cancel()
	}
	for _, child := range b.children {
		if child.opened {
			child.closeWithoutUpdate()
		}
	}
}

// toggle inverts the open/closed state of a branch
func (b *Branch) toggle() {
	if b.opened {
		b.close()
	} else {
		b.open()
	}
}

// afterStateChange is fired every time a branch is opened or closed.
func (b *Branch) afterStateChange() {
	t := b.tree
	if t.Selected != nil {
		// if the selected branch is now hidden, we should un-select it.
		t.Selected.UnselectIfHidden()
	}
	b.updateOpenerIcon()
	b.setDirtyIconState()
	if b.parent != nil {
		b.parent.setDirtyIconState()
	}
}

// canOpen gives us an indication of whether the branch can be opened. We use this to work out if we
// should display the plus icon, or the empty icon. If a branch is async but the children aren't
// loaded, we don't know if it has children or not. In that case we assume it can be opened.
func (b *Branch) canOpen() bool {
	if b.isAsyncAndNotLoaded() {
		return true
	}
	if len(b.children) == 0 {
		return false
	}
	return true
}

// isVisible tells us if the branch is currently visible. For a branch to be visible, all it's
// ancestors must be open.
func (b *Branch) isVisible() bool {
	current := b.parent
	for current != nil {
		if !current.opened {
			return false
		}
		current = current.parent
	}
	return true
}

// updateOpenerIcon updates the opener button with the correct icon
func (b *Branch) updateOpenerIcon() {
	if b.opener != nil {
		if b.isAsyncAndNotLoaded() {
			b.opener.SetInnerHTML(icons.Unknown)
		} else if b.children == nil || len(b.children) == 0 {
			b.opener.SetInnerHTML(icons.Empty)
		} else if b.opened {
			b.opener.SetInnerHTML(icons.Minus)
		} else {
			b.opener.SetInnerHTML(icons.Plus)
		}
	}
}
