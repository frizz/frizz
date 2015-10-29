package editor // import "kego.io/editor"

import (
	"honnef.co/go/js/dom"
	"kego.io/system/node"
)

type Editable interface {
	GetEditor(n *node.Node) Editor
}

type Editor interface {
	Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error
	Show()
	Hide()
	IsInitialized() bool
}

type Common struct {
	Path        string
	Aliases     map[string]string
	Panel       *dom.HTMLDivElement
	Initialized bool
}

func (e *Common) IsInitialized() bool {
	return e.Initialized
}

func (e *Common) Show() {
	e.Panel.Style().Set("display", "block")
}

func (e *Common) Hide() {
	e.Panel.Style().Set("display", "none")
}
