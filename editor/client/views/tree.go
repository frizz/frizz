package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type TreeView struct {
	*View

	Root     *models.BranchModel
	Selected *models.BranchModel
}

func NewTreeView(ctx context.Context) *TreeView {
	v := &TreeView{}
	v.View = New(ctx, v)
	v.Mount()
	return v
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *TreeView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *TreeView) Mount() {
	v.Notifs = v.App.Watch(v.App.Branches.Root(),
		stores.BranchInitialStateLoaded,
	)
	go func() {
		for notif := range v.Notifs {
			v.Root = v.App.Branches.Root()
			v.ReconcileBody()
			close(notif.Done)
		}
	}()
}

func (v *TreeView) Render() vecty.Component {
	if v.Root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content tree"),
		NewBranchView(v.Ctx, v.Root),
	)
}
