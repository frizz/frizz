package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type BranchView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	notifs   chan flux.Notif
	children vecty.List
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	v := &BranchView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		model: model,
	}
	v.Mount()
	return v
}

func (v *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BranchView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BranchView) Mount() {

	v.notifs = v.app.Branches.Watch(v.model,
		stores.BranchOpen,
		stores.BranchOpenPostLoad,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchSelect,
	)

	go func() {
		for notif := range v.notifs {
			switch notif {
			case stores.BranchOpen:
				loaded := LoadBranch(v.ctx, v.app, v.model)
				v.app.Dispatch(&actions.BranchOpenPostLoad{Branch: v.model, Loaded: loaded})
			case stores.BranchSelect:
				loaded := LoadBranch(v.ctx, v.app, v.model)
				if v.model.LastOp == models.BranchOpClickToggle && loaded {
					v.app.Dispatch(&actions.BranchOpen{Branch: v.model})
				}
				v.app.Dispatch(&actions.BranchSelectPostLoad{Branch: v.model, Loaded: loaded})
			case stores.BranchOpenPostLoad,
				stores.BranchClose,
				stores.BranchLoaded:
				v.ReconcileBody()
			}
		}
	}()
}

func (v *BranchView) Unmount() {
	if v.notifs != nil {
		v.app.Branches.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *BranchView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	v.children = vecty.List{}
	if v.model.Open {
		for _, c := range v.model.Children {
			v.children = append(v.children, NewBranchView(v.ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(v.ctx, v.model),
		elem.Div(
			prop.Class("children"),
			v.children,
		),
	)
}
