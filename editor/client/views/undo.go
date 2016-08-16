package views

import (
	"fmt"

	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
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
		if v.App.Actions.Index() <= 0 {
			disabled = true
			label = "Undo"
		} else {
			label = fmt.Sprintf("Undo (%d)", v.App.Actions.Index())
		}

	} else {
		if v.App.Actions.Index() >= v.App.Actions.Count() {
			disabled = true
			label = "Redo"
		} else {
			label = fmt.Sprintf("Redo (%d)", v.App.Actions.Count()-v.App.Actions.Index())
		}
	}

	return elem.ListItem(
		vecty.ClassMap{
			"disabled": disabled,
		},
		elem.Anchor(
			event.Click(v.click).PreventDefault(),
			prop.Href("#"),
			vecty.Text(label),
		),
	)

}
