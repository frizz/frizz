package tree // import "kego.io/editor/client/tree"

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/shared/connection"
)

type Tree struct {
	Root     *Branch
	Conn     *connection.Conn
	Fail     chan error
	Path     string
	Aliases  map[string]string
	Selected *Branch
	Editor   editor.Editor
	Content  *dom.HTMLDivElement
}

func (t *Tree) KeyboardEvent(e *dom.KeyboardEvent) {
	switch e.KeyCode {
	case 38: // up
		e.PreventDefault()
		if t.Selected == nil {
			b := t.Root.LastVisible()
			if b != nil {
				b.Select(true)
			}
			return
		}
		b := t.Selected.PrevVisible()
		if b != nil {
			b.Select(true)
		}
	case 40: // down
		e.PreventDefault()
		if t.Selected == nil {
			b := t.Root.children[0]
			if b != nil {
				b.Select(true)
			}
			return
		}
		b := t.Selected.NextVisible(true)
		if b != nil {
			b.Select(true)
		}
	case 37: // left
		e.PreventDefault()
		if t.Selected == nil {
			return
		}
		if t.Selected.open && t.Selected.CanOpen() {
			t.Selected.Close()
			return
		}
		b := t.Selected.Parent
		if b != nil {
			b.Select(true)
		}
	case 39: // right
		e.PreventDefault()
		if t.Selected == nil {
			return
		}
		if t.Selected.open || !t.Selected.CanOpen() {
			return
		}
		t.Selected.Open()
	}
}

func New(content *dom.HTMLDivElement, conn *connection.Conn, fail chan error, path string, aliases map[string]string) *Tree {
	return &Tree{Content: content, Conn: conn, Fail: fail, Path: path, Aliases: aliases}
}
