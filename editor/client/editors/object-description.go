package editors

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/system/node"
)

type ObjectDescriptionView struct {
	*views.View

	model *models.NodeModel
}

func NewObjectDescriptionView(ctx context.Context, node *node.Node) *ObjectDescriptionView {
	v := &ObjectDescriptionView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Nodes.Get(node)
	v.Watch(v.model.Node,
		stores.NodeValueChanged,
		stores.NodeDeleted,
		stores.NodeInitialised,
	)
	return v
}

func (v *ObjectDescriptionView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectDescriptionView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ObjectDescriptionView) Render() vecty.Component {
	n := v.model.Node
	if n.Null || n.Missing || n.ValueString == "" {
		return elem.Div()
	}
	return elem.Paragraph(
		prop.Class("lead"),
		vecty.Text(n.ValueString),
	)
}
