package images

// notest

import (
	"context"

	"frizz.io/editor/client/editable"
	"frizz.io/editor/client/editors"
	"frizz.io/editor/client/models"
	"frizz.io/editor/client/stores"
	"frizz.io/editor/client/views"
	"frizz.io/flux"
	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/prop"
	"github.com/dave/vecty/style"
)

var _ editable.Editable = (*Photo)(nil)

func (s *Photo) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Block
}

func (s *Photo) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewIconEditorView(ctx, node, format)
}

type IconEditorView struct {
	*views.View

	model  *models.EditorModel
	photo  *Photo
	editor *editors.StringEditorView
}

func NewIconEditorView(ctx context.Context, node *node.Node, format editable.Format) *IconEditorView {
	v := &IconEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.photo = v.model.Node.Value.(*Photo)
	v.Watch(v.model.Node,
		stores.NodeValueChanged,
		stores.NodeDescendantChanged,
		stores.NodeFocus,
	)
	return v
}

func (v *IconEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.photo = v.model.Node.Value.(*Photo)
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
	if v.photo.Url != nil {
		url = v.photo.Url.Value()
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
