package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BranchControlView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	cUnsel   chan struct{}
	cSel     chan struct{}
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
	v.cSel = v.app.Branches.WatchSingle(stores.BranchPreSelect, v.model)
	v.cUnsel = v.app.Branches.WatchSingle(stores.BranchUnselect, v.model)
	go func() {
		for range flux.WatchMulti(v.cSel, v.cUnsel) {
			v.ReconcileBody()
			if v.model != nil && v.app.Branches.Selected() == v.model {
				v.focus()
			}
		}
	}()
}

func (v *BranchControlView) Unmount() {
	if v.cSel != nil {
		v.app.Branches.Delete(v.cSel)
		v.cSel = nil
	}
	if v.cUnsel != nil {
		v.app.Branches.Delete(v.cUnsel)
		v.cUnsel = nil
	}
	v.Body.Unmount()
}

func (v *BranchControlView) focus() {
	v.Node().Call("scrollIntoViewIfNeeded")
}

func (v *BranchControlView) toggleClick(*vecty.Event) {
	go func() {
		if v.model.CanOpen() {
			if v.model.Open {
				v.app.Dispatch(&actions.BranchClose{Branch: v.model})
			} else {
				v.app.Dispatch(&actions.BranchOpen{Branch: v.model})
			}
		} else {
			v.app.Dispatch(&actions.BranchSelectClick{Branch: v.model})
		}
	}()
}

func (v *BranchControlView) labelClick(*vecty.Event) {
	go func() {
		v.app.Dispatch(&actions.BranchSelectClick{Branch: v.model})
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
