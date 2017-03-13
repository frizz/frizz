package views

import (
	"context"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/prop"
)

type Header struct {
	*View
}

func NewHeader(ctx context.Context) *Header {
	v := &Header{}
	v.View = New(ctx, v)
	return v
}

func (v *Header) Render() *vecty.HTML {
	return elem.Navigation(
		prop.Class("navbar navbar-inverse navbar-fixed-top"),
		elem.Div(
			prop.Class("container-fluid"),
			elem.Div(
				prop.Class("navbar-header"),
				elem.Button(
					prop.Type("button"),
					prop.Class("navbar-toggle collapsed"),
					vecty.Data("toggle", "collapse"),
					vecty.Data("target", "#navbar"),
					elem.Span(
						prop.Class("sr-only"),
						vecty.Text("Toggle navigation"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
					elem.Span(
						prop.Class("icon-bar"),
					),
				),
				/* elem.Anchor(
					prop.Class("navbar-brand"),
					prop.Href("#"),
					vecty.Text("Ke"),
				), */
			),
			elem.Div(
				prop.ID("navbar"),
				prop.Class("navbar-collapse collapse"),
				elem.UnorderedList(
					prop.Class("nav navbar-nav"),
					NewViewMenuView(v.Ctx),
					NewEditMenuView(v.Ctx),
					NewSaveView(v.Ctx),
				),
				elem.Form(
					prop.Class("navbar-form navbar-right"),
					elem.Input(
						prop.Type("text"),
						prop.Class("form-control"),
						prop.Placeholder("Search..."),
					),
				),
			),
		),
	)
}
