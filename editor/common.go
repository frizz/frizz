package editor

import (
	"github.com/tjgq/broadcast"
	"honnef.co/go/js/dom"
)

type Editor struct {
	*dom.HTMLDivElement
	Path    string
	Aliases map[string]string
	Editors []EditorInterface
	layout  Layout
	holder  Holder
	changes *broadcast.Broadcaster
}

func (c *Editor) Send(v interface{}) {
	c.changes.Send(v)
}

func (c *Editor) Listen() *broadcast.Listener {
	return c.changes.Listen()
}

func (c *Editor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	c.HTMLDivElement = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	if layout == Page {
		c.Style().Set("display", "none")
		c.Class().SetString("mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-grid")
	}
	if layout == Inline {
		c.Style().Set("display", "inline-block")
	}

	c.changes = broadcast.New(0)

	holder.ListenForEditorChanges(c.changes.Listen().Ch)

	c.Path = path
	c.Aliases = aliases
	c.holder = holder
	c.layout = layout
	return nil

}

func (c *Editor) Show() {
	if c.layout == Inline {
		c.Style().Set("display", "inline-block")
	} else {
		c.Style().Set("display", "block")
	}
}

func (c *Editor) Hide() {
	c.Style().Set("display", "none")
}

func (c *Editor) AddChildTreeEntry(child EditorInterface) bool {
	return true
}

func (c *Editor) Dirty() bool {
	for _, e := range c.Editors {
		if e.Dirty() {
			return true
		}
	}
	return false
}

func (c *Editor) Select() {
	c.holder.Select(false)
}

func (c *Editor) Focus() {}
