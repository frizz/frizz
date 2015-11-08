package editor

import "honnef.co/go/js/dom"

type Common struct {
	*dom.HTMLDivElement
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

	c.Path = path
	c.Aliases = aliases
	c.holder = holder
	c.layout = layout
	return nil

}

func (e *Common) Show() {
	if e.layout == Inline {
		e.Style().Set("display", "inline-block")
	} else {
		e.Style().Set("display", "block")
	}
}

func (e *Common) Hide() {
	e.Style().Set("display", "none")
}

func (e *Common) AddChildTreeEntry(child Editor) bool {
	return true
}
