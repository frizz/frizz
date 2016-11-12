package views

import (
	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type StructView struct {
	*View

	model *models.EditorModel
}

func NewStructView(ctx context.Context, node *node.Node) *StructView {
	v := &StructView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	return v
}

func (v *StructView) Render() *vecty.HTML {
	if v.model == nil {
		return elem.Div(vecty.Text("Struct (nil)"))
	}

	if v.model.Node.Type.Basic {
		// This view is only for types that embed system:object
		panic(kerr.New("QHDQMXTNIH", "Basic type %s not supported by StructView", v.model.Node.Type.Id.String()))
	}

	out := vecty.List{}

	// Always show the editor for system:object first
	objectEditor, ok := clientctx.FromContext(v.Ctx).Get("kego.io/system:object")
	if !ok {
		panic(kerr.New("BJRMXESSUV", "Can't find editor for system:object"))
	}
	out = append(out, objectEditor.EditorView(v.Ctx, v.model.Node, editable.Block))

	out = append(out, NewStructFragmentView(v.Ctx, v.model.Node, v.model.Node.Type.Id))

	return elem.Div(
		out,
	)

}
