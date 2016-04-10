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

func (b *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		b.Body = old.Body
		b.model = old.model
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	if model == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	b := &BranchView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	b.Mount()
	return b
}

func (b *BranchView) Mount() {
	if b.c != nil {
		panic("mounting a mounted branch")
	}
	b.c = b.app.Branches.WatchAll(b.model)
	go func() {
		for range b.c {
			b.ReconcileBody()
		}
	}()
}

func (b *BranchView) Unmount() {
	if b.c == nil {
		return
	}
	b.app.Branches.DeleteAll(b.c)
	close(b.c)
	b.c = nil
	b.Body.Unmount()
}

// Apply implements the vecty.Markup interface.
func (b *BranchView) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *BranchView) render() vecty.Component {
	if b.model == nil {
		return elem.Div()
	}

	b.children = vecty.List{}
	if b.model.Open {
		for _, c := range b.model.Children {
			b.children = append(b.children, NewBranchView(b.ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(b.ctx, b.model),
		elem.Div(
			prop.Class("children"),
			b.children,
		),
	)
}
