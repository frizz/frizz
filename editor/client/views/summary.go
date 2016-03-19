package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type SummaryView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	selected *node.Node
}

func NewSummaryView(ctx context.Context) *SummaryView {

	e := &SummaryView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	return e
}

// Apply implements the vecty.Markup interface.
func (e *SummaryView) Apply(element *vecty.Element) {
	element.AddChild(e)
}

func (e *SummaryView) Reconcile(old vecty.Component) {
	if old, ok := old.(*SummaryView); ok {
		e.Body = old.Body
		e.selected = old.selected
	}
	e.RenderFunc = e.render
	e.ReconcileBody()
}

func (e *SummaryView) render() vecty.Component {
	return elem.Div(
		vecty.Text("Summary table"),
	)
}