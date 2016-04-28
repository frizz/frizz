package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type TreeView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App
	c   chan struct{}

	Root     *models.BranchModel
	Selected *models.BranchModel
}

func NewTreeView(ctx context.Context) *TreeView {
	v := &TreeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.Mount()
	return v
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *TreeView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *TreeView) Mount() {
	v.c = v.app.Branches.Watch(stores.BranchInitialStateLoaded, v.app.Branches.Root())
	go func() {
		for range v.c {
			v.Root = v.app.Branches.Root()
			v.ReconcileBody()
		}
	}()
}

func (v *TreeView) Unmount() {
	if v.c != nil {
		v.app.Branches.Delete(v.c)
		v.c = nil
	}
	v.Body.Unmount()
}

func (v *TreeView) render() vecty.Component {
	if v.Root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content tree"),
		NewBranchView(v.ctx, v.Root),
	)
}
