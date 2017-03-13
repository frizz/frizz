package images

// ke: {"package": {"notest": true}}

import (
	"context"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/prop"
	"github.com/dave/vecty/style"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/editors"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

var _ editable.Editable = (*Icon)(nil)

func (s *Icon) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Block
}

func (s *Icon) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewIconEditorView(ctx, node, format)
}

type IconEditorView struct {
	*views.View

	model  *models.EditorModel
	icon   *Icon
	editor *editors.StringEditorView
}

func NewIconEditorView(ctx context.Context, node *node.Node, format editable.Format) *IconEditorView {
	v := &IconEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.icon = v.model.Node.Value.(*Icon)
	v.Watch(v.model.Node,
		stores.NodeValueChanged,
		stores.NodeDescendantChanged,
		stores.NodeFocus,
	)
	return v
}

func (v *IconEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.icon = v.model.Node.Value.(*Icon)
	vecty.Rerender(v)
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *IconEditorView) Focus() {
	v.editor.Focus()
}

func (v *IconEditorView) Render() *vecty.HTML {
	v.editor = editors.NewStringEditorView(v.Ctx, v.model.Node.Map["url"], editable.Inline)
	url := ""
	if v.icon.Url != nil {
		url = v.icon.Url.Value()
	}
	return elem.Div(
		prop.Class("container-fluid"),
		elem.Div(
			prop.Class("row"),
			elem.Div(
				prop.Class("col-sm-10"),
				vecty.Style("padding-left", "0"),
				vecty.Style("padding-right", "0"),
				v.editor,
			),
			elem.Div(
				prop.Class("col-sm-2"),
				elem.Image(
					prop.Class("img-responsive"),
					style.MaxHeight("200px"),
					prop.Src(url),
				),
			),
		),
	)
}
