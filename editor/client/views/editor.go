package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

func RegisterDefaultEditor(ctx context.Context) {
	editors := clientctx.FromContext(ctx)
	editors.Set("", new(DefaultEditor))
}

type DefaultEditor struct{}

func (s *DefaultEditor) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Branch
}

func (s *DefaultEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewEditorView(ctx, node, format)
}

type EditorView struct {
	*View

	model *models.EditorModel
}

func NewEditorView(ctx context.Context, node *node.Node, format editable.Format) *EditorView {
	v := &EditorView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.Watch(v.reaction, v.model,
		stores.EditorValueChanged,
	)
	return v
}

func (v *EditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *EditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *EditorView) Render() vecty.Component {
	return elem.Div(vecty.Text("default editor"))
}
