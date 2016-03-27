package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type PanelView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Selected *models.BranchModel
}

func NewPanelView(ctx context.Context) *PanelView {
	p := &PanelView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	/*
		go func() {
			sc := p.app.Branches.Changed()
			for {
				select {
				case <-sc:
					p.Selected = p.app.Branches.Selected()
					p.ReconcileBody()
				}
			}
		}()
	*/
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

	var n *node.Node
	if p.Selected != nil {
		ni, ok := p.Selected.Contents.(models.NodeContentsInterface)
		if ok {
			n = ni.GetNode()
		}
	}

	return elem.Div(
		prop.Class("content panel"),
		vecty.If(
			n != nil,
			NewBreadcrumbsView(p.ctx),
			NewEditorView(p.ctx, n),
			NewSummaryView(p.ctx),
		),
	)
}
