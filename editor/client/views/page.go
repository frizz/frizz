package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type PageView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Environment *envctx.Env
}

func NewPage(ctx context.Context, env *envctx.Env) *PageView {
	v := &PageView{
		ctx:         ctx,
		app:         stores.FromContext(ctx),
		Environment: env,
	}
	v.addKeyboardEvents()
	return v
}

func (v *PageView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PageView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *PageView) Apply(element *vecty.Element) {
	element.AddChild(v)
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
				v.app.Dispatch(&actions.AddCollectionItem{Parent: v.app.Nodes.Selected()})
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
	selected := v.app.Branches.Selected()
	switch code {
	case 38: // up
		if selected == nil {
			if b := v.app.Branches.Root().LastVisible(); b != nil {
				v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
			}
			return
		}
		if b := selected.PrevVisible(); b != nil {
			v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
	case 40: // down
		if selected == nil {
			b := v.app.Branches.Root()
			v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
		if b := selected.NextVisible(true); b != nil {
			v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
			return
		}
	case 37: // left
		if selected == nil {
			return
		}
		if selected.CanOpen() && selected.Open {
			// if the branch is open, left arrow should close it.
			v.app.Dispatcher.Dispatch(&actions.BranchClose{Branch: selected})
			return
		} else {
			if b := selected.Parent; b != nil {
				v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
				return
			}
		}
	case 39: // right
		if selected == nil {
			return
		}
		if selected.CanOpen() && !selected.Open {
			// if the branch is closed, right arrow should open it
			loadBranch(v.ctx, v.app, selected, &flux.Waiter{})
			v.app.Dispatcher.Dispatch(&actions.BranchOpen{Branch: selected})
			return
		} else {
			if b := selected.FirstChild(); b != nil {
				v.app.Dispatcher.Dispatch(&actions.BranchSelect{Branch: b, Op: models.BranchOpKeyboard})
				return
			}
		}
	}
}

func (v *PageView) render() vecty.Component {
	return elem.Div(
		prop.ID("wrapper"),
		NewHeader(v.ctx, v.Environment),
		elem.Div(
			prop.Class("wrapper"),
			elem.Div(
				prop.ID("tree"),
				prop.Class("split split-horizontal"),
				NewTreeView(v.ctx),
			),
			elem.Div(
				prop.ID("main"),
				prop.Class("split split-horizontal"),
				NewPanelView(v.ctx),
			),
		),
		NewAddPopView(v.ctx),
	)
}
