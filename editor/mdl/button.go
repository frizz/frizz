package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

/*
<button class="mdl-button mdl-js-button mdl-button--raised">
Button
</button>
*/

type button struct {
	*dom.HTMLButtonElement
}

func Button() *button {
	b := dom.GetWindow().Document().CreateElement("button").(*dom.HTMLButtonElement)
	b.Class().SetString("mdl-button mdl-js-button mdl-button--raised")
	js.Global.Get("componentHandler").Call("upgradeElement", b)
	return &button{b}
}

func (b *button) Text(text string) *button {
	b.SetTextContent(text)
	return b
}

func (b *button) Click(action func(dom.Event)) *button {
	b.AddEventListener("click", true, action)
	return b
}
