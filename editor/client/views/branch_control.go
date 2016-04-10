package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BranchControlView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	c        chan struct{}
	children vecty.List
}

func (b *BranchControlView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchControlView); ok {
		b.Body = old.Body
		b.model = old.model
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func NewBranchControlView(ctx context.Context, model *models.BranchModel) *BranchControlView {
	if model == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	b := &BranchControlView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	b.Mount()
	return b
}

func (b *BranchControlView) Mount() {
	if b.c != nil {
		panic("mounting a mounted BranchControl")
	}
	b.c = b.app.Branches.WatchSingle(stores.BranchSelectedChanged, b.model)
	go func() {
		for range b.c {
			b.ReconcileBody()
		}
	}()
}

func (b *BranchControlView) Unmount() {
	if b.c == nil {
		return
	}
	b.app.Branches.DeleteSingle(stores.BranchSelectedChanged, b.c)
	close(b.c)
	b.c = nil
	b.Body.Unmount()
}

// Apply implements the vecty.Markup interface.
func (b *BranchControlView) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *BranchControlView) toggleClick(*vecty.Event) {
	go func() {
		LoadBranch(b.ctx, b.app, b.model, true)
		if b.model.CanOpen() {
			b.app.Dispatch(&actions.ToggleBranch{Branch: b.model})
		} else {
			b.app.Dispatch(&actions.SelectBranch{Branch: b.model})
		}
	}()
}

func (b *BranchControlView) labelClick(*vecty.Event) {
	go func() {
		LoadBranch(b.ctx, b.app, b.model, true)
		b.app.Dispatch(&actions.SelectBranch{Branch: b.model})
	}()
}

func (b *BranchControlView) render() vecty.Component {
	if b.model == nil {
		return elem.Div()
	}

	selected := b.app.Branches.Selected() == b.model

	icon := b.model.Icon()

	return elem.Div(
		elem.Anchor(
			vecty.ClassMap{
				"toggle":   true,
				"selected": selected,
				"plus":     icon == "plus",
				"minus":    icon == "minus",
				"unknown":  icon == "unknown",
				"empty":    icon == "empty",
			},
			event.Click(b.toggleClick),
		),
		elem.Div(
			vecty.ClassMap{
				"node-content": true,
				"selected":     selected,
			},
			elem.Span(
				prop.Class("node-label"),
				event.Click(b.labelClick),
				vecty.Text(b.model.Contents.Label()),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
	)

}
