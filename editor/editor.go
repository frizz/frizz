package editor // import "kego.io/editor"

import (
	"github.com/tjgq/broadcast"
	"honnef.co/go/js/dom"
)

type Editable interface {
	GetEditor(n *Node) EditorInterface
}

type Holder interface {
	Select(fromKeyboard bool)
	ListenForEditorChanges(changes <-chan interface{})
}

type Layout int

const (
	Inline Layout = iota
	Block
	Page
)

type EditorInterface interface {
	dom.Node
	Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error
	Show()
	Hide()
	Layout() Layout
	AddChildTreeEntry(child EditorInterface) bool
	Dirty() bool
	Select()
	Focus()
	Value() interface{}
	Send(v interface{})
	Listen() *broadcast.Listener
}
