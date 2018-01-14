package views

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"honnef.co/go/js/dom"
)

type PageView struct {
	vecty.Core
	//*View
}

func NewPage() *PageView {
	p := &PageView{}
	return p
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	input := elem.Input(
		vecty.Markup(
			prop.Type(prop.TypeText),
			prop.Value("frizz.io/tests/data"),
		),
	)
	auth := dom.GetWindow().Document().GetElementByID("wrapper").GetAttribute("auth")
	return elem.Body(
		vecty.Markup(
			vecty.Attribute("auth", auth),
			prop.ID("wrapper"),
		),
		input,
		elem.Button(
			vecty.Markup(
				event.Click(func(event *vecty.Event) {
					js.Global.Call("bootstrap", input.Node().Get("value"))

				}).PreventDefault(),
			),
			vecty.Text("Bootstrap"),
		),
		elem.Span(
			vecty.Markup(
				prop.ID("log"),
			),
		),
	)
}
