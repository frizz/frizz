package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type MapView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewMapView(ctx context.Context, node *node.Node) *MapView {
	v := &MapView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *MapView) Reconcile(old vecty.Component) {
	if old, ok := old.(*MapView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *MapView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *MapView) Mount() {
	v.notifs = v.app.Nodes.Watch(v.model.Node,
		stores.NodeInitialised,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *MapView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *MapView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *MapView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Map (nil)"))
	}

	return elem.Div(
		elem.Div(
			elem.Button(
				prop.Type(prop.TypeButton),
				prop.Class("btn btn-primary"),
				vecty.Text("Add"),
				event.Click(func(ev *vecty.Event) {
					v.app.Dispatch(&actions.AddCollectionItem{
						Parent: v.model.Node,
					})
				}).PreventDefault(),
			),
		),
		NewMapTableView(v.ctx, v.model.Node),
	)

}
