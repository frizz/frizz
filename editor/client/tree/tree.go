package tree // import "kego.io/editor/client/tree"

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/shared/connection"
	"kego.io/system"
)

type Tree struct {
	Root     *Branch
	Conn     *connection.Conn
	Fail     chan error
	Path     string
	Aliases  map[string]string
	selected *Branch
	editor   system.Editor
	content  *dom.HTMLDivElement
}

func (t *Tree) KeyboardEvent(e *dom.KeyboardEvent) {
	switch e.KeyCode {
	case 38: // up
		e.PreventDefault()
		if t.selected == nil {
			b := t.Root.lastVisible()
			if b != nil {
				b.Select(true)
			}
			return
		}
		b := t.selected.prev
		if b != nil && !b.root {
			b.Select(true)
		}
	case 40: // down
		e.PreventDefault()
		if t.selected == nil {
			b := t.Root.children[0]
			if b != nil {
				b.Select(true)
			}
			return
		}
		b := t.selected.next
		if b != nil {
			b.Select(true)
		}
	case 37: // left
		e.PreventDefault()
		if t.selected == nil {
			return
		}
		if t.selected.open && t.selected.canOpen() {
			t.selected.Close()
			return
		}
		b := t.selected.parent
		if b != nil && !b.root {
			b.Select(true)
		}
	case 39: // right
		e.PreventDefault()
		if t.selected == nil {
			return
		}
		if t.selected.open || !t.selected.canOpen() {
			return
		}
		t.selected.Open()
	}
}

func New(parent *dom.BasicHTMLElement, content *dom.HTMLDivElement, conn *connection.Conn, root Item, fail chan error, path string, aliases map[string]string) *Tree {

	tree := &Tree{Conn: conn, Fail: fail, Path: path, Aliases: aliases, content: content}
	tree.Root = &Branch{Tree: tree, root: true, open: true, item: root}

	// We must tolerate passing in a nil dom element in order to run tests in pure go
	if parent == nil {
		tree.Root.Initialise()
		return tree
	}

	nodeDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	nodeDiv.SetAttribute("class", "node root")
	parent.AppendChild(nodeDiv)

	innerDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	innerDiv.SetAttribute("class", "children")
	nodeDiv.AppendChild(innerDiv)

	tree.Root.element = nodeDiv
	tree.Root.inner = innerDiv

	tree.Root.Initialise()
	return tree
}

func (t *Tree) Branch(item Item) *Branch {
	return &Branch{item: item, Tree: t}
}
