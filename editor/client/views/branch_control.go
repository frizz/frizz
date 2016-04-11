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

func NewBranchControlView(ctx context.Context, model *models.BranchModel) *BranchControlView {
	if model == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	v := &BranchControlView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	v.Mount()
	return v
}

func (v *BranchControlView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchControlView); ok {
		v.Body = old.Body
		v.model = old.model
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
	if v.model != nil && v.app.Branches.Selected() == v.model {
		v.focus()
	}
}

// Apply implements the vecty.Markup interface.
func (v *BranchControlView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BranchControlView) Mount() {
	if v.c != nil {
		panic("mounting a mounted BranchControl")
	}
	v.c = v.app.Branches.WatchSingle(stores.BranchSelectedChanged, v.model)
	go func() {
		for range v.c {
			v.ReconcileBody()
			if v.model != nil && v.app.Branches.Selected() == v.model {
				v.focus()
			}
		}
	}()
}

func (v *BranchControlView) Unmount() {
	if v.c == nil {
		return
	}
	v.app.Branches.DeleteSingle(stores.BranchSelectedChanged, v.c)
	close(v.c)
	v.c = nil
	v.Body.Unmount()
}

func (v *BranchControlView) focus() {
	v.Node().Call("scrollIntoViewIfNeeded")
}

func (v *BranchControlView) toggleClick(*vecty.Event) {
	go func() {
		LoadBranch(v.ctx, v.app, v.model)
		if v.model.CanOpen() {
			v.app.Dispatch(&actions.ToggleBranch{Branch: v.model})
		} else {
			v.app.Dispatch(&actions.SelectBranch{Branch: v.model})
		}
	}()
}

func (v *BranchControlView) labelClick(*vecty.Event) {
	go func() {
		loaded := LoadBranch(v.ctx, v.app, v.model)
		v.app.Dispatch(&actions.SelectBranch{Branch: v.model})
		if loaded {
			v.app.Dispatch(&actions.OpenBranch{Branch: v.model})
		}
	}()
}

func (v *BranchControlView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	selected := v.app.Branches.Selected() == v.model

	icon := v.model.Icon()

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
			event.Click(v.toggleClick),
		),
		elem.Div(
			vecty.ClassMap{
				"node-content": true,
				"selected":     selected,
			},
			elem.Span(
				prop.Class("node-label"),
				event.Click(v.labelClick),
				vecty.Text(v.model.Contents.Label()),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
	)

}
