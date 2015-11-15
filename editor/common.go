package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/broadcast"
)

type Common struct {
	*dom.HTMLDivElement
	*broadcast.Broadcaster
	Path    string
	Aliases map[string]string
	Editors []Editor
	layout  Layout
	holder  Holder
}

func (c *Common) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	c.HTMLDivElement = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	if layout == Page {
		c.Style().Set("display", "none")
		c.Class().SetString("mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-grid")
	}
	if layout == Inline {
		c.Style().Set("display", "inline-block")
	}

	c.Broadcaster = broadcast.New(0)

	holder.Listen(c.Listen().Ch)

	c.Path = path
	c.Aliases = aliases
	c.holder = holder
	c.layout = layout
	return nil

}

func (c *Common) Show() {
	if c.layout == Inline {
		c.Style().Set("display", "inline-block")
	} else {
		c.Style().Set("display", "block")
	}
}

func (c *Common) Hide() {
	c.Style().Set("display", "none")
}

func (c *Common) AddChildTreeEntry(child Editor) bool {
	return true
}

// Notify sends a notification that the editor data has been changed.
func (c *Common) Notify(editor Editor) {
	c.Send(editor)
}

func (c *Common) Dirty() bool {
	for _, e := range c.Editors {
		if e.Dirty() {
			return true
		}
	}
	return false
}

func (c *Common) Select() {
	c.holder.Select(false)
}

func (c *Common) Focus() {}
