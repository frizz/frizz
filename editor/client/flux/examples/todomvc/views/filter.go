package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
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

func (b *FilterButton) Reconcile(old vecty.Component) {
	if old, ok := old.(*FilterButton); ok {
		b.Body = old.Body
		b.Label = old.Label
		b.Filter = old.Filter
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
