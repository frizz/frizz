package editor // import "kego.io/editor"

import "honnef.co/go/js/dom"

type Editable interface {
	GetEditor(n *Node) Editor
}

type Holder interface {
	MarkDirty(bool)
}

type Layout int

const (
	Inline Layout = iota
	Block
	Page
)

type Editor interface {
	Initialize(panel *dom.HTMLDivElement, holder Holder, layout Layout, path string, aliases map[string]string) error
	Show()
	Hide()
	Layout() Layout
	AddChildTreeEntry(child Editor) bool
}
