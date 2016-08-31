package views // import "kego.io/editor/client/views"

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type PanelView struct {
	*View

	branch *models.BranchModel
	node   *node.Node
}

func NewPanelView(ctx context.Context) *PanelView {
	v := &PanelView{}
	v.View = New(ctx, v)
	v.Watch(nil,
		stores.BranchSelected,
		stores.ViewChanged,
		stores.BranchInitialStateLoaded,
	)
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *PanelView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.branch = v.App.Branches.Selected()
	v.node = v.App.Nodes.Selected()
	v.ReconcileBody()
	if notif.Type == stores.BranchSelected {
		v.Node().Get("parentNode").Set("scrollTop", "0")
	}
}

func (v *PanelView) Render() vecty.Component {
	var editor vecty.Component
	if v.node != nil {
		if v.node.Type.IsNativeMap() {
			editor = NewMapView(v.Ctx, v.node)
		} else if v.node.Type.IsNativeArray() {
			editor = NewArrayView(v.Ctx, v.node)
		} else {
			editor = NewStructView(v.Ctx, v.node)
		}
	} else if v.branch != nil {
		editor = NewPanelNavView(v.Ctx, v.branch)
	}
	return elem.Div(
		prop.Class("content content-panel"),
		editor,
	)
}
