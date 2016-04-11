package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BranchView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	c        chan struct{}
	children vecty.List
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	if model == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	v := &BranchView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	v.Mount()
	return v
}

func (v *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		v.Body = old.Body
		v.model = old.model
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (b *BranchView) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (v *BranchView) Mount() {
	if v.c != nil {
		panic("mounting a mounted branch")
	}
	v.c = v.app.Branches.Watch(v.model)
	go func() {
		for range v.c {
			v.ReconcileBody()
		}
	}()
}

func (v *BranchView) Unmount() {
	if v.c == nil {
		return
	}
	v.app.Branches.Delete(v.c)
	close(v.c)
	v.c = nil
	v.Body.Unmount()
}

func (v *BranchView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	v.children = vecty.List{}
	if v.model.Open {
		for _, c := range v.model.Children {
			v.children = append(v.children, NewBranchView(v.ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(v.ctx, v.model),
		elem.Div(
			prop.Class("children"),
			v.children,
		),
	)
}
