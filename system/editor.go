package system

import "honnef.co/go/js/dom"

type Editor interface {
	Initialized() bool
	Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error
	Update()
	Show()
	Hide()
}

type HasEditor interface {
	GetEditor(node *Node) Editor
}

type editorCommon struct {
	path        string
	aliases     map[string]string
	panel       *dom.HTMLDivElement
	initialized bool
}

func (e *editorCommon) Initialized() bool {
	return e.initialized
}

func (e *editorCommon) Show() {
	e.panel.Style().Set("display", "block")
}

func (e *editorCommon) Hide() {
	e.panel.Style().Set("display", "none")
}
