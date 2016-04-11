package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type BreadcrumbsView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	selected *node.Node
}

func NewBreadcrumbsView(ctx context.Context) *BreadcrumbsView {

	e := &BreadcrumbsView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	return e
}

func (e *BreadcrumbsView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BreadcrumbsView); ok {
		e.Body = old.Body
		e.selected = old.selected
	}
	e.RenderFunc = e.render
	e.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (e *BreadcrumbsView) Apply(element *vecty.Element) {
	element.AddChild(e)
}

func (e *BreadcrumbsView) render() vecty.Component {
	return elem.Div(
		vecty.Text("Breadcrmbs"),
	)
}
