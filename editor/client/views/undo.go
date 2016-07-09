package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
)

type UndoView struct {
	*View
	Undo bool
}

func NewUndoView(ctx context.Context, undo bool) *UndoView {
	v := &UndoView{}
	v.View = New(ctx, v)
	v.Undo = undo
	v.Watch(nil, stores.ActionsChanged)
	return v
}

func (v *UndoView) Reconcile(old vecty.Component) {
	if old, ok := old.(*UndoView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *UndoView) click(e *vecty.Event) {

	var action actions.Undoable
	if v.Undo {
		action = v.App.Actions.Undo()
	} else {
		action = v.App.Actions.Redo()
	}

	if action == nil {
		return
	}

	v.App.Dispatch(action)
}

func (v *UndoView) Render() vecty.Component {

	var label string
	var disabled bool
	if v.Undo {
		label = "Undo"
		if v.App.Actions.Index() <= 0 {
			disabled = true
		}

	} else {
		label = "Redo"
		if v.App.Actions.Index() >= v.App.Actions.Count() {
			disabled = true
		}
	}

	return elem.ListItem(
		vecty.ClassMap{
			"disabled": disabled,
		},
		elem.Anchor(
			event.Click(v.click).PreventDefault().StopPropagation(),
			prop.Href("#"),
			vecty.Text(label),
		),
	)

}
