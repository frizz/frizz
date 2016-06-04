package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type ObjectView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model  *models.EditorModel
	origin *system.Reference
}

func NewObjectView(ctx context.Context, node *node.Node, origin *system.Reference) *ObjectView {
	v := &ObjectView{
		ctx:    ctx,
		app:    stores.FromContext(ctx),
		origin: origin,
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *ObjectView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ObjectView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ObjectView) Mount() {
	v.notifs = v.app.Nodes.Watch(v.model.Node,
		stores.NodeInitialised,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ObjectView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ObjectView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ObjectView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("default editor (nil)"))
	}

	children := vecty.List{}
	for _, n := range v.model.Node.Map {
		if n.Missing || n.Null {
			continue
		}
		if *n.Origin != *v.origin {
			continue
		}
		e := models.GetEditable(v.ctx, n)
		if e.Format() == editable.Block || e.Format() == editable.Inline {
			children = append(children, e.EditorView(v.ctx, n))
		}
	}
	if len(children) == 0 {
		return elem.Div(vecty.Text("default editor"))
	}
	return elem.Div(
		elem.Form(
			children,
		),
	)
}
