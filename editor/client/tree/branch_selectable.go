package tree

import "time"

type Selectable interface {
	// Select highlights the branch and opens the edit panel. This may happen asynchronously if the
	// contents are not loaded yet. If there are contents to be loaded, and the select action comes
	// from keyboard input, we add a 50ms delay before initiating the call.
	Select(fromKeyboard bool)
	// Unselect clears the selected state and hides the children
	Unselect()
	// UnselectIfHidden clears the selected state and hides the children but only if the branch is
	// not visible
	UnselectIfHidden()
}

// Select highlights the branch and opens the edit panel. This may happen asynchronously if the
// contents are not loaded yet. If there are contents to be loaded, and the select action comes from
// keyboard input, we add a 50ms delay before initiating the call.
func (b *Branch) Select(fromKeyboard bool) {

	if b.parent == nil {
		// we never select the root node
		return
	}

	if !b.isVisible() {
		// if the branch we're selecting isn't visible, we should open all it's closed ancestors.
		ancestor := b.parent
		for ancestor != nil {
			if !ancestor.opened && ancestor.canOpen() {
				ancestor.open()
			}
			ancestor = ancestor.parent
		}
	}

	if b.tree.Selected != nil {
		b.tree.Selected.Unselect()
	}
	b.content.Class().Add("selected")
	b.tree.Selected = b
	b.selected = true

	if fromKeyboard && b.isAsyncAndNotLoaded() {
		// wait 50ms before showing the edit panel so we don't generate lots of content load
		// requests if we scroll quickly. We only do this for keyboard events.
		go func() {
			time.Sleep(time.Millisecond * 50)
			if b.selected {
				b.showEditPanel(fromKeyboard)
			}
		}()
	} else {
		b.showEditPanel(fromKeyboard)
	}

	return
}

// Unselect clears the selected state and hides the children
func (b *Branch) Unselect() {
	// un-select
	b.content.Class().Remove("selected")
	b.tree.Selected = nil
	b.selected = false
}

// UnselectIfHidden clears the selected state and hides the children but only if the branch is not
// visible
func (b *Branch) UnselectIfHidden() {
	if !b.isVisible() {
		b.Unselect()
	}
}
