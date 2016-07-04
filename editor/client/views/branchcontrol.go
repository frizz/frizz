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
	"kego.io/flux"
)

type BranchControlView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	notifs   chan flux.NotifPayload
	children vecty.List
}

func NewBranchControlView(ctx context.Context, model *models.BranchModel) *BranchControlView {
	app := stores.FromContext(ctx)
	v := &BranchControlView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	v.RenderFunc = v.render
	v.Mount()
	return v
}

// ke: {"func": {"notest": true}}
func (v *BranchControlView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchControlView); ok {
		v.Body = old.Body
	}
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
	v.notifs = v.app.Watch(v.model,
		stores.BranchSelectControl,
		stores.BranchUnselectControl,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *BranchControlView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if v.model != nil && v.app.Branches.Selected() == v.model {
		v.focus()
	}
}

func (v *BranchControlView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *BranchControlView) focus() {
	v.Node().Call("scrollIntoViewIfNeeded")
}

func (v *BranchControlView) toggleClick(*vecty.Event) {
	go func() {
		if !v.model.CanOpen() {
			v.app.Dispatch(&actions.BranchSelecting{Branch: v.model, Op: models.BranchOpClickToggle})
			return
		}
		if v.model.Open {
			v.app.Dispatch(&actions.BranchClose{Branch: v.model})
			return
		}
		v.app.Dispatch(&actions.BranchOpening{Branch: v.model})
	}()
}

func (v *BranchControlView) labelClick(*vecty.Event) {
	v.app.Dispatch(&actions.BranchSelecting{Branch: v.model, Op: models.BranchOpClickLabel})
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
				vecty.Text(v.model.Contents.Label(v.ctx)),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
	)

}
