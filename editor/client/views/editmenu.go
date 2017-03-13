package views

import (
	"context"

	"fmt"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
	"kego.io/editor/client/stores"
)

type EditMenuView struct {
	*View
}

func NewEditMenuView(ctx context.Context) *EditMenuView {
	v := &EditMenuView{}
	v.View = New(ctx, v)
	v.Watch(nil,
		stores.ActionsChanged,
		stores.InfoStateChange,
	)
	return v
}

func (v *EditMenuView) Render() *vecty.HTML {

	var undoLabel, redoLabel string
	var undoEnabled, redoEnabled bool

	if v.App.Actions.Index() <= 0 {
		undoLabel = "Undo"
	} else {
		undoEnabled = true
		undoLabel = fmt.Sprintf("Undo %s", v.App.Actions.UndoPeek().Description())
	}

	if v.App.Actions.Index() >= v.App.Actions.Count() {
		redoLabel = "Redo"
	} else {
		redoEnabled = true
		redoLabel = fmt.Sprintf("Redo %s", v.App.Actions.RedoPeek().Description())
	}

	return elem.ListItem(
		prop.Class("dropdown"),
		elem.Anchor(
			prop.Href("#"),
			prop.Class("dropdown-toggle"),
			vecty.Data("toggle", "dropdown"),
			vecty.Text("Edit"),
			elem.Span(
				prop.Class("caret"),
			),
		),
		elem.UnorderedList(
			prop.Class("dropdown-menu"),
			elem.ListItem(
				vecty.ClassMap{
					"disabled": !undoEnabled,
				},
				elem.Anchor(
					event.Click(func(e *vecty.Event) {
						action := v.App.Actions.Undo()
						if action == nil {
							return
						}
						v.App.Dispatch(action)
					}).PreventDefault().StopPropagation(),
					prop.Href("#"),
					vecty.Text(undoLabel),
				),
			),
			elem.ListItem(
				vecty.ClassMap{
					"disabled": !redoEnabled,
				},
				elem.Anchor(
					event.Click(func(e *vecty.Event) {
						action := v.App.Actions.Redo()
						if action == nil {
							return
						}
						v.App.Dispatch(action)
					}).PreventDefault().StopPropagation(),
					prop.Href("#"),
					vecty.Text(redoLabel),
				),
			),
		),
	)
}
