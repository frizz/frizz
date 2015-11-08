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

func (c *Common) MarkDirty(dirty bool) {
	// note we can't use c.Dirty() here because we need to call the overridden function on types
	// with Common embedded.
	c.holder.MarkDirty(c, dirty)
}

func (c *Common) Dirty() bool {
	for _, e := range c.Editors {
		if e.Dirty() {
			return true
		}
	}
	return false
}
