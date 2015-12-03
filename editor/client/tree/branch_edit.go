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

func (b *Branch) markEditorDirtyState(e editor.EditorInterface, state bool) {
	if state {
		if b.dirty == nil {
			b.dirty = map[editor.EditorInterface]bool{}
		}
		b.dirty[e] = true
	} else {
		if b.dirty != nil {
			delete(b.dirty, e)
		}
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

	if b.dirty == nil || len(b.dirty) == 0 {
		return false
	}

	if !b.opened {
		// if descendants are dirty, only show the icon if this branch is closed
		return true
	}

	return false
}

func (b *Branch) notify(editor editor.EditorInterface) {
	current := b
	for current != nil {
		current.markEditorDirtyState(editor, editor.Dirty())
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
