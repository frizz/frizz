package tree

import (
	"kego.io/editor"
	"kego.io/kerr"
)

func (b *Branch) showEditPanel(fromKeyboard bool) {

	if b.parent == nil {
		return
	}

	if b.editor != nil && b.tree.Editor != nil && b.editor == b.tree.Editor {
		// if this editor is already visible, we don't need to do anything
		return
	}

	if b.tree.Editor != nil {
		b.tree.Editor.Hide()
		b.tree.Editor = nil
	}

	done, ok := b.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	// we wrap the success action in a closure, because it must be executed syncronously or in a
	// goroutine depending on the return from ensureContentLoaded
	success := func() {

		if !fromKeyboard && b.canOpen() && !b.opened {
			// if we clicked on an item, and it's not open, we should open it
			b.open()
		}

		if b.editor == nil {
			ed, ok := b.self.(Editable)
			if !ok {
				return
			}
			b.editor = ed.Editor()
			if err := b.editor.Initialize(b.self, editor.Page, b.tree.Path, b.tree.Aliases); err != nil {
				b.tree.Fail <- kerr.New("KKOBKWJDBI", err, "Initialize")
				return
			}
			b.ListenForEditorChanges(b.editor.Listen().Ch)
			b.tree.Content.AppendChild(b.editor)
		}
		if b.editor == nil {
			return
		}
		if b.tree.Editor != nil {
			b.tree.Editor.Hide()
		}
		b.editor.Show()
		b.tree.Editor = b.editor

	}

	if done == nil {
		// if the done chanel is nil, the operation was synchronous, so we should call success
		// synchronously
		success()
	} else {
		go func() {
			// block and wait until the response arrives
			<-done
			success()
		}()
	}
	return
}

func (b *Branch) markEditorDirtyState(e editor.EditorInterface, state bool, self bool) {

	if b.dirtySelf == nil {
		b.dirtySelf = map[editor.EditorInterface]bool{}
	}

	if b.dirtyChild == nil {
		b.dirtyChild = map[editor.EditorInterface]bool{}
	}

	m := b.dirtyChild
	if self {
		m = b.dirtySelf
	}

	if state {
		m[e] = true
	} else {
		delete(m, e)
	}

	b.setDirtyIconState()
}

func (b *Branch) setDirtyIconState() {
	if b.badge == nil {
		return
	}

	if b.dirtyIconState() {
		b.badge.Style().Set("display", "inline")
		return
	}

	b.badge.Style().Set("display", "none")

}

func (b *Branch) dirtyIconState() bool {

	if b.dirtySelf != nil && len(b.dirtySelf) > 0 {
		// if an editor in this branch is dirty, always display the icon
		return true
	}

	if b.dirtyChild == nil || len(b.dirtyChild) == 0 {
		// if there at no dirty editors in descendant branches, don't display the icon
		return false
	}

	if b.opened {
		// if descendants are dirty, only show the icon if this branch is closed
		return false
	}

	return true
}

func (b *Branch) notify(editor editor.EditorInterface) {
	current := b
	for current != nil {
		current.markEditorDirtyState(editor, editor.Dirty(), current == b)
		current = current.parent
	}
}

func (b *Branch) ListenForEditorChanges(changes <-chan interface{}) {
	go func() {
		for e := range changes {
			b.notify(e.(editor.EditorInterface))
		}
	}()
}
