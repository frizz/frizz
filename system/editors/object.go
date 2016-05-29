package editors

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type ObjectEditor struct{}

func (s *ObjectEditor) Format() editable.Format {
	return editable.Branch
}

func (s *ObjectEditor) EditorView(ctx context.Context, node *node.Node) vecty.Component {
	return NewObjectEditorView(ctx, node)
}

type ObjectEditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewObjectEditorView(ctx context.Context, node *node.Node) *ObjectEditorView {
	v := &ObjectEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *ObjectEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectEditorView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ObjectEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ObjectEditorView) Mount() {
	v.notifs = v.app.Editors.Watch(v.model,
		stores.EditorLoaded,
		stores.EditorInitialStateLoaded,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ObjectEditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ObjectEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Editors.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ObjectEditorView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("default editor (nil)"))
	}

	children := vecty.List{}
	for label, n := range v.model.Node.Map {
		if n.Missing || n.Null {
			continue
		}
		e := models.GetEditable(v.ctx, n)
		if e.Format() == editable.Block || e.Format() == editable.Inline {
			children = append(children, elem.Div(vecty.Text(label)))
			children = append(children, elem.Div(e.EditorView(v.ctx, n)))
		}
	}
	if len(children) > 0 {
		return elem.Div(children)
	}

	typeName := ""
	if v.model.Node.Type != nil {
		typeName = v.model.Node.Type.Id.String()
	}

	return elem.Div(vecty.Text("default editor (" + typeName + ")"))
}
