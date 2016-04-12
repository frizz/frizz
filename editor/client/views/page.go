package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
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
		v.Environment = old.Environment
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
		case "INPUT", "TEXTAREA":
			if k.KeyCode == 27 {
				// escape
				document.ActiveElement().Blur()
			}
			return
		default:
			if k.KeyCode >= 37 && k.KeyCode <= 40 {
				// up, down, left, right
				k.PreventDefault()
				go func() {
					v.keyPress(k.KeyCode)
				}()
			}
		}
	})
}

func (v *PageView) keyPress(code int) {
	selected := v.app.Branches.Selected()
	switch code {
	case 38: // up
		if selected == nil {
			if b := v.app.Branches.Root().LastVisible(); b != nil {
				LoadBranchDebounced(v.ctx, v.app, b)
				v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
			}
			return
		}
		if b := selected.PrevVisible(); b != nil {
			LoadBranchDebounced(v.ctx, v.app, b)
			v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
			return
		}
	case 40: // down
		if selected == nil {
			b := v.app.Branches.Root()
			LoadBranchDebounced(v.ctx, v.app, b)
			v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
			return
		}
		if b := selected.NextVisible(true); b != nil {
			LoadBranchDebounced(v.ctx, v.app, b)
			v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
			return
		}
	case 37: // left
		if selected == nil {
			return
		}
		if selected.CanOpen() && selected.Open {
			// if the branch is open, left arrow should close it.
			v.app.Dispatcher.Dispatch(&actions.CloseBranch{Branch: selected})
			return
		} else {
			if b := selected.Parent; b != nil {
				LoadBranchDebounced(v.ctx, v.app, b)
				v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
				return
			}
		}
	case 39: // right
		if selected == nil {
			return
		}
		if selected.CanOpen() && !selected.Open {
			// if the branch is closed, right arrow should open it
			LoadBranch(v.ctx, v.app, selected)
			v.app.Dispatcher.Dispatch(&actions.OpenBranch{Branch: selected})
			return
		} else {
			if b := selected.FirstChild(); b != nil {
				LoadBranchDebounced(v.ctx, v.app, b)
				v.app.Dispatcher.Dispatch(&actions.SelectBranch{Branch: b})
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
	)
}
