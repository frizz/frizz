package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
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
	return elem.Body(
		vecty.Markup(
			prop.ID("wrapper"),
		),
		vecty.Text("Hello world"),
	)
}
