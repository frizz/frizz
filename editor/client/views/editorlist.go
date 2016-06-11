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

type EditorListView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model  *models.EditorModel
	filter *system.Reference
}

func NewEditorListView(ctx context.Context, model *models.EditorModel, filter *system.Reference) *EditorListView {
	v := &EditorListView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = model
	v.filter = filter
	v.Mount()
	return v
}

func (v *EditorListView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorListView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *EditorListView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *EditorListView) Mount() {
	v.notifs = v.app.Watch(v.model,
		stores.EditorArrayOrderChanged,
		stores.EditorChildAdded,
		stores.EditorChildDeleted,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *EditorListView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *EditorListView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *EditorListView) render() vecty.Component {

	if v.model == nil || v.model.Node.Missing || v.model.Node.Null {
		return elem.Div(vecty.Text("editor (nil)"))
	}

	children := vecty.List{}

	add := func(n *node.Node) {
		e := models.GetEditable(v.ctx, n)
		f := e.Format(n.Rule)
		if f == editable.Block || f == editable.Inline {
			children = append(children, e.EditorView(v.ctx, n))
		}
	}

	for _, n := range v.model.Node.Map {
		if n.Missing || n.Null {
			continue
		}
		if v.filter != nil && *n.Origin != *v.filter {
			continue
		}
		add(n)
	}

	for _, n := range v.model.Node.Array {
		add(n)
	}

	if len(children) == 0 {
		return elem.Div()
	}
	return elem.Div(
		elem.Form(
			children,
		),
	)
}
