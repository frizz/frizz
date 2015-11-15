package editor // import "kego.io/editor"

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/broadcast"
)

type Editable interface {
	GetEditor(n *Node) Editor
}

type Holder interface {
	Select(fromKeyboard bool)
	Listen(changes <-chan interface{})
}

type Layout int

const (
	Inline Layout = iota
	Block
	Page
)

type Editor interface {
	dom.Node
	Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error
	Show()
	Hide()
	Layout() Layout
	AddChildTreeEntry(child Editor) bool
	Dirty() bool
	Select()
	Focus()
	Value() interface{}
	Send(v interface{})
	Listen() *broadcast.Listener
}
