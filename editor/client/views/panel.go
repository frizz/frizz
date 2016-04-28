package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type PanelView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App
	c   chan struct{}

	branch *models.BranchModel
}

func NewPanelView(ctx context.Context) *PanelView {
	v := &PanelView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.Mount()
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *PanelView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *PanelView) Mount() {
	v.c = v.app.Branches.Watch(stores.BranchSelectPostLoad)
	go func() {
		for range v.c {
			v.branch = v.app.Branches.Selected()
			v.ReconcileBody()
		}
	}()
}

func (v *PanelView) Unmount() {
	if v.c != nil {
		v.app.Branches.Delete(v.c)
		v.c = nil
	}
	v.Body.Unmount()
}

func (v *PanelView) render() vecty.Component {

	var n *node.Node
	label := ""
	if v.branch != nil {
		label = v.branch.Contents.Label()
		ni, ok := v.branch.Contents.(models.NodeContentsInterface)
		if ok {
			n = ni.GetNode()
		}
	}

	return elem.Div(
		prop.Class("content panel"),
		vecty.Text(label),
		NewBreadcrumbsView(v.ctx, v.branch),
		NewEditorView(v.ctx, n),
		NewSummaryView(v.ctx, n),
	)
}
