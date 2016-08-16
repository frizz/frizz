package views

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type ViewMenuView struct {
	*View
}

func NewViewMenuView(ctx context.Context) *ViewMenuView {
	v := &ViewMenuView{}
	v.View = New(ctx, v)
	v.Watch(nil, stores.ViewChanged)
	return v
}

func (v *ViewMenuView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ViewMenuView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ViewMenuView) Render() vecty.Component {
	return elem.ListItem(
		prop.Class("dropdown"),
		elem.Anchor(
			prop.Href("#"),
			prop.Class("dropdown-toggle"),
			vecty.Data("toggle", "dropdown"),
			vecty.Text("View"),
			elem.Span(
				prop.Class("caret"),
			),
		),
		elem.UnorderedList(
			prop.Class("dropdown-menu"),
			elem.ListItem(
				vecty.ClassMap{
					"disabled": v.App.Branches.View() == models.Data,
				},
				elem.Anchor(
					event.Click(func(ev *vecty.Event) {
						v.App.Dispatch(&actions.ChangeView{View: models.Data})
					}).PreventDefault(),
					prop.Href("#"),
					vecty.Text("Data"),
				),
			),
			elem.ListItem(
				vecty.ClassMap{
					"disabled": v.App.Branches.View() == models.Types,
				},
				elem.Anchor(
					event.Click(func(ev *vecty.Event) {
						v.App.Dispatch(&actions.ChangeView{View: models.Types})
					}).PreventDefault(),
					prop.Href("#"),
					vecty.Text("Types"),
				),
			),
			elem.ListItem(
				vecty.ClassMap{
					"disabled": v.App.Branches.View() == models.Package,
				},
				elem.Anchor(
					event.Click(func(ev *vecty.Event) {
						v.App.Dispatch(&actions.ChangeView{View: models.Package})
					}).PreventDefault(),
					prop.Href("#"),
					vecty.Text("Package"),
				),
			),
		),
	)
}
