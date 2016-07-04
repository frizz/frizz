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
	v.Mount()
	return v
}

func (v *EditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *EditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *EditorView) Mount() {
	v.Notifs = v.App.Watch(v.model,
		stores.EditorValueChanged,
	)
	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()
}

func (v *EditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *EditorView) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *EditorView) Render() vecty.Component {
	return elem.Div(vecty.Text("default editor"))
}
