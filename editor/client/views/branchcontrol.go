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
	*View

	model    *models.BranchModel
	children vecty.List
}

func NewBranchControlView(ctx context.Context, model *models.BranchModel) *BranchControlView {
	v := &BranchControlView{}
	v.View = New(ctx, v)
	v.model = model
	v.Watch(v.reaction, v.model,
		stores.BranchSelectControl,
		stores.BranchUnselectControl,
	)
	return v
}

// ke: {"func": {"notest": true}}
func (v *BranchControlView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchControlView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
	if v.model != nil && v.App.Branches.Selected() == v.model {
		v.focus()
	}
}

func (v *BranchControlView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if v.model != nil && v.App.Branches.Selected() == v.model {
		v.focus()
	}
}

func (v *BranchControlView) focus() {
	v.Node().Call("scrollIntoViewIfNeeded")
}

func (v *BranchControlView) toggleClick(*vecty.Event) {
	go func() {
		if !v.model.CanOpen() {
			v.App.Dispatch(&actions.BranchSelecting{Branch: v.model, Op: models.BranchOpClickToggle})
			return
		}
		if v.model.Open {
			v.App.Dispatch(&actions.BranchClose{Branch: v.model})
			return
		}
		v.App.Dispatch(&actions.BranchOpening{Branch: v.model})
	}()
}

func (v *BranchControlView) labelClick(*vecty.Event) {
	v.App.Dispatch(&actions.BranchSelecting{Branch: v.model, Op: models.BranchOpClickLabel})
}

func (v *BranchControlView) Render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	selected := v.App.Branches.Selected() == v.model

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
				vecty.Text(v.model.Contents.Label(v.Ctx)),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
	)

}
