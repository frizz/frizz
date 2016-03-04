package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/flux/examples/todomvc/actions"
	"kego.io/editor/client/flux/examples/todomvc/model"
	"kego.io/editor/client/flux/examples/todomvc/stores"
)

type FilterButton struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Label  string
	Filter model.FilterState
}

func NewFilterButton(ctx context.Context, label string, filter model.FilterState) *FilterButton {
	return &FilterButton{
		ctx:    ctx,
		app:    stores.FromContext(ctx),
		Label:  label,
		Filter: filter,
	}
}

// Apply implements the vecty.Markup interface.
func (b *FilterButton) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *FilterButton) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*FilterButton); ok {
		b.Body = oldComp.Body
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func (b *FilterButton) onClick(event *vecty.Event) {
	go func() {
		b.app.Dispatcher.Dispatch(&actions.SetFilter{Filter: b.Filter})
	}()
}

func (b *FilterButton) render() vecty.Component {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(b.app.Todos.Filter() == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),
			vecty.Text(b.Label),
		),
	)
}
