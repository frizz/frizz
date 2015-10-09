package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Checkbox struct {
	*dom.HTMLLabelElement
	Input *dom.HTMLInputElement
	Span  *dom.HTMLSpanElement
}

func NewCheckbox(value bool, label string) *Checkbox {

	id := randomId()

	c := &Checkbox{}
	c.HTMLLabelElement = dom.GetWindow().Document().CreateElement("label").(*dom.HTMLLabelElement)
	c.Class().SetString("mdl-switch mdl-js-switch mdl-js-ripple-effect")
	c.For = id

	c.Input = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	c.Input.SetID(id)
	c.Input.Class().Add("mdl-switch__input")
	c.Input.Type = "checkbox"
	c.Input.Checked = value

	c.AppendChild(c.Input)

	c.Span = dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	c.Span.Class().Add("mdl-switch__label")
	c.Span.SetTextContent(label)
	c.AppendChild(c.Span)

	js.Global.Get("componentHandler").Call("upgradeElement", c)

	return c
}
