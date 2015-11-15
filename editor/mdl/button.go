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

type ButtonStruct struct {
	*dom.HTMLButtonElement
}

func Button() *ButtonStruct {
	b := dom.GetWindow().Document().CreateElement("button").(*dom.HTMLButtonElement)
	b.Class().SetString("mdl-button mdl-js-button mdl-button--raised")
	js.Global.Get("componentHandler").Call("upgradeElement", b)
	return &ButtonStruct{b}
}

func (b *ButtonStruct) Text(text string) *ButtonStruct {
	b.SetTextContent(text)
	return b
}

func (b *ButtonStruct) Click(action func(dom.Event)) *ButtonStruct {
	b.AddEventListener("click", true, action)
	return b
}
