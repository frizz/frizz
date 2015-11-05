package editor // import "kego.io/editor"

import (
	"honnef.co/go/js/dom"
	"kego.io/system/node"
)

type Editable interface {
	GetEditor(n *node.Node) Editor
}

type Dirtyable interface {
	MarkDirty(bool)
}

type Editor interface {
	Initialize(panel *dom.HTMLDivElement, d Dirtyable, path string, aliases map[string]string) error
	Show()
	Hide()
	IsInitialized() bool
}

type Common struct {
	Path        string
	Aliases     map[string]string
	Panel       *dom.HTMLDivElement
	initialized bool
	dirtyable   Dirtyable
}

func (c *Common) Initialize(panel *dom.HTMLDivElement, dirtyable Dirtyable, path string, aliases map[string]string) error {

	c.Panel = panel
	c.Path = path
	c.Aliases = aliases
	c.dirtyable = dirtyable
	c.initialized = true
	return nil

}

func (e *Common) IsInitialized() bool {
	return e.initialized
}

func (e *Common) Show() {
	e.Panel.Style().Set("display", "block")
}

func (e *Common) Hide() {
	e.Panel.Style().Set("display", "none")
}
