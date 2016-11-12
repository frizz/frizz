package views // import "kego.io/editor/client/views"

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"honnef.co/go/js/dom"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
)

type PageView struct {
	*View
}

func NewPage(ctx context.Context) *PageView {
	v := &PageView{}
	v.View = New(ctx, v)
	v.addKeyboardEvents()
	return v
}

func (v *PageView) addKeyboardEvents() {
	window := dom.GetWindow()
	document := window.Document().(dom.HTMLDocument)
	window.AddEventListener("keydown", true, func(e dom.Event) {
		k := e.(*dom.KeyboardEvent)
		switch document.ActiveElement().TagName() {
		case "INPUT", "TEXTAREA", "SELECT":
			if k.KeyCode == 27 {
				// escape
				document.ActiveElement().Blur()
			}
			return
		default:
			switch k.KeyCode {
			case 65:
				// "a"
				addCollectionItem(v.Ctx, v.App, v.App.Nodes.Selected())
			case 37, 38, 39, 40:
				// up, down, left, right
				k.PreventDefault()
				go func() {
					v.KeyPress(k.KeyCode)
				}()
			}
		}
	})
}

func (v *PageView) KeyPress(code int) {
	selected := v.App.Branches.Selected()
	switch code {
	case 38: // up
		if selected == nil {
			if b := v.App.Branches.Root().LastVisible(); b != nil {
				v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
			}
			return
		}
		if b := selected.PrevVisible(); b != nil {
			v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
	case 40: // down
		if selected == nil {
			b := v.App.Branches.Root()
			v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
		if b := selected.NextVisible(true); b != nil {
			v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
		if selected.CanOpen() && !selected.Open {
			// if the branch at the end of the list is closed, down arrow
			// should open it
			loadBranch(v.Ctx, v.App, selected, &flux.Waiter{})
			v.App.Dispatcher.Dispatch(&actions.BranchOpening{Branch: selected})
			return
		}
	case 37: // left
		if selected == nil {
			return
		}
		if selected.CanOpen() && selected.Open {
			// if the branch is open, left arrow should close it.
			v.App.Dispatcher.Dispatch(&actions.BranchClose{Branch: selected})
			return
		} else {
			if b := selected.Parent; b != nil {
				v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
				return
			}
		}
	case 39: // right
		if selected == nil {
			return
		}
		if selected.CanOpen() && !selected.Open {
			// if the branch is closed, right arrow should open it
			loadBranch(v.Ctx, v.App, selected, &flux.Waiter{})
			v.App.Dispatcher.Dispatch(&actions.BranchOpening{Branch: selected})
			return
		} else {
			if b := selected.FirstChild(); b != nil {
				v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpKeyboard})
				return
			}
		}
	}
}

func (v *PageView) Render() *vecty.HTML {
	return elem.Body(
		prop.ID("wrapper"),
		NewHeader(v.Ctx),
		elem.Div(
			prop.Class("wrapper"),
			elem.Div(
				prop.ID("tree"),
				prop.Class("split split-horizontal"),
				NewTreeView(v.Ctx),
			),
			elem.Div(
				prop.ID("main"),
				prop.Class("split split-horizontal"),
				NewPanelView(v.Ctx),
			),
		),
		NewAddPopupView(v.Ctx),
	)
}
