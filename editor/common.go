package editor

import "honnef.co/go/js/dom"

type Common struct {
	Path    string
	Aliases map[string]string
	Panel   *dom.HTMLDivElement
	holder  Holder
}

func (c *Common) Initialize(content *dom.HTMLDivElement, holder Holder, layout Layout, path string, aliases map[string]string) error {

	if layout == Page {
		c.Panel = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		c.Panel.Style().Set("display", "none")
		c.Panel.Class().SetString("mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-grid")
		content.AppendChild(c.Panel)
	} else {
		c.Panel = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		content.AppendChild(c.Panel)
	}

	c.Path = path
	c.Aliases = aliases
	c.holder = holder
	return nil

}

func (e *Common) Show() {
	e.Panel.Style().Set("display", "block")
}

func (e *Common) Hide() {
	e.Panel.Style().Set("display", "none")
}

func (e *Common) AddChildTreeEntry(child Editor) bool {
	return true
}
