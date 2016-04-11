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
	c   chan struct{}

	branch *models.BranchModel
}

func NewPanelView(ctx context.Context) *PanelView {
	v := &PanelView{
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
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
		v.branch = old.branch
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *PanelView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *PanelView) Mount() {
	if v.c != nil {
		panic("mounting a mounted panel")
	}
	v.c = v.app.Branches.WatchSingle(stores.BranchSelectedChanged, v.branch)
	go func() {
		for range v.c {
			v.branch = v.app.Branches.Selected()
			v.ReconcileBody()
		}
	}()
}

func (v *PanelView) Unmount() {
	if v.c == nil {
		return
	}
	v.app.Branches.DeleteSingle(stores.BranchSelectedChanged, v.c)
	close(v.c)
	v.c = nil
	v.Body.Unmount()
}

func (v *PanelView) render() vecty.Component {

	var n *node.Node
	if v.branch != nil {
		ni, ok := v.branch.Contents.(models.NodeContentsInterface)
		if ok {
			n = ni.GetNode()
		}
	}

	return elem.Div(
		prop.Class("content panel"),
		vecty.If(
			n != nil,
			NewBreadcrumbsView(v.ctx),
			NewEditorView(v.ctx, n),
			NewSummaryView(v.ctx),
		),
	)
}
