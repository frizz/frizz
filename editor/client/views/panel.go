package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type PanelView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Selected *node.Node
}

func NewPanelView(ctx context.Context) *PanelView {
	p := &PanelView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	go func() {
		sc := p.app.Selected.Changed()
		for {
			select {
			case <-sc:
				p.Selected = p.app.Selected.Get()
				p.ReconcileBody()
			}
		}
	}()

	return p
}

// Apply implements the vecty.Markup interface.
func (p *PanelView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
		v.Selected = old.Selected
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

func (p *PanelView) render() vecty.Component {

	return elem.Div(
		prop.Class("content"),
		vecty.If(
			p.Selected == nil,
			vecty.Text("None selected"),
		),
		vecty.If(
			p.Selected != nil,
			NewBreadcrumbsView(p.ctx),
			NewEditorView(p.ctx, p.Selected),
			NewSummaryView(p.ctx),
		),
	)
}
