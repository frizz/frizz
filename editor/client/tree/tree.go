package tree // import "kego.io/editor/client/tree"

// ke: {"package": {"jstest": true}}

import (
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/shared/connection"
)

type Tree struct {
	Root     *Root
	Conn     *connection.Conn
	Fail     chan error
	Selected BranchInterface
	Editor   editor.EditorInterface
	Content  *dom.HTMLDivElement
	ctx      context.Context
}

func New(ctx context.Context, content *dom.HTMLDivElement, conn *connection.Conn, fail chan error) *Tree {
	return &Tree{Content: content, Conn: conn, Fail: fail, ctx: ctx}
}

func (t *Tree) KeyboardEvent(e *dom.KeyboardEvent) {
	switch e.KeyCode {
	case 38: // up
		e.PreventDefault()
		if t.Selected == nil {
			t.Root.KeyboardSelectLast()
			return
		}
		t.Selected.KeyboardSelectPrev()
	case 40: // down
		e.PreventDefault()
		if t.Selected == nil {
			t.Root.KeyboardSelectFirst()
			return
		}
		t.Selected.KeyboardSelectNext()
	case 37: // left
		e.PreventDefault()
		if t.Selected == nil {
			return
		}
		if success := t.Selected.KeyboardClose(); !success {
			t.Selected.KeyboardSelectParent()
		}
	case 39: // right
		e.PreventDefault()
		if t.Selected == nil {
			return
		}
		t.Selected.KeyboardOpen()
	}
}
