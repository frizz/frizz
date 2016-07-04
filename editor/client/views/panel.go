package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/editable"
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
	v.Mount()
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *PanelView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *PanelView) Mount() {
	v.Notifs = v.App.Watch(nil,
		stores.BranchSelected,
	)
	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()

}

func (v *PanelView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.branch = v.App.Branches.Selected()
	v.node = v.App.Nodes.Selected()
	v.ReconcileBody()
	if notif.Type == stores.BranchSelected {
		v.Node().Get("parentNode").Set("scrollTop", "0")
	}
}

func (v *PanelView) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *PanelView) Render() vecty.Component {
	var editor, breadcrumbs vecty.Component
	if v.branch != nil {
		breadcrumbs = NewBreadcrumbsView(v.Ctx, v.branch)
	}
	if v.node != nil {
		if ed, ok := v.node.Value.(editable.Editable); ok {
			editor = ed.EditorView(v.Ctx, v.node, editable.Branch)
		} else if v.node.Type.IsNativeMap() {
			editor = NewMapView(v.Ctx, v.node)
		} else if v.node.Type.IsNativeArray() {
			editor = NewArrayView(v.Ctx, v.node)
		} else {
			editor = NewObjectView(v.Ctx, v.node)
		}
	}
	return elem.Div(
		prop.Class("content panel"),
		breadcrumbs,
		editor,
	)
}
