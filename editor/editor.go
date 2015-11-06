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
}

type Common struct {
	Path      string
	Aliases   map[string]string
	Panel     *dom.HTMLDivElement
	dirtyable Dirtyable
}

func (c *Common) Initialize(content *dom.HTMLDivElement, dirtyable Dirtyable, path string, aliases map[string]string) error {

	c.Panel = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	c.Panel.Style().Set("display", "none")
	c.Panel.Class().SetString("mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-grid")
	content.AppendChild(c.Panel)

	c.Path = path
	c.Aliases = aliases
	c.dirtyable = dirtyable
	return nil

}

func (e *Common) Show() {
	e.Panel.Style().Set("display", "block")
}

func (e *Common) Hide() {
	e.Panel.Style().Set("display", "none")
}
