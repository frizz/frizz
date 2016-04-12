package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
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

func NewSummaryView(ctx context.Context, n *node.Node) *SummaryView {

	v := &SummaryView{
		ctx:      ctx,
		app:      stores.FromContext(ctx),
		selected: n,
	}

	return v
}

func (v *SummaryView) Reconcile(old vecty.Component) {
	if old, ok := old.(*SummaryView); ok {
		v.Body = old.Body
		v.selected = old.selected
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *SummaryView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *SummaryView) render() vecty.Component {
	if v.selected == nil {
		return elem.Div()
	}
	return elem.Div(
		vecty.Text("Summary table"),
	)
}
