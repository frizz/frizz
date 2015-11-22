package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type TextboxStruct struct {
	*dom.HTMLDivElement
	Input *dom.HTMLInputElement
	Label *dom.HTMLLabelElement
}

func Textbox(value string, label string) *TextboxStruct {

	id := randomId()

	t := &TextboxStruct{}
	t.HTMLDivElement = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	t.Class().SetString("mdl-textfield mdl-js-textfield mdl-textfield--floating-label")
	t.Style().Set("width", "100%")

	t.Input = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	t.Input.Class().Add("mdl-textfield__input")
	t.Input.SetID(id)
	t.Input.Type = "text"
	t.Input.Value = value

	t.AppendChild(t.Input)

	t.Label = dom.GetWindow().Document().CreateElement("label").(*dom.HTMLLabelElement)
	t.Label.Class().Add("mdl-textfield__label")
	t.Label.For = id
	t.Label.SetTextContent(label)
	t.AppendChild(t.Label)

	js.Global.Get("componentHandler").Call("upgradeElement", t)

	return t
}

func (t *TextboxStruct) SetDisabled(state bool) {
	//console.Log("setting disabled", state)
	if state {
		t.Input.Disabled = true
		t.Class().Add("is-disabled")
	} else {
		t.Input.Disabled = false
		t.Class().Remove("is-disabled")
	}
	js.Global.Get("componentHandler").Call("upgradeElement", t)
}
