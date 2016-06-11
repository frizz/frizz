package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type MapTableView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewMapTableView(ctx context.Context, model *models.EditorModel) *MapTableView {
	v := &MapTableView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		model: model,
	}
	v.Mount()
	return v
}

func (v *MapTableView) Reconcile(old vecty.Component) {
	if old, ok := old.(*MapTableView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *MapTableView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *MapTableView) Mount() {
	v.notifs = v.app.Editors.Watch(nil,
		stores.EditorChildDeleted,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *MapTableView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *MapTableView) Unmount() {
	if v.notifs != nil {
		v.app.Editors.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *MapTableView) render() vecty.Component {

	if v.model.Node == nil || len(v.model.Node.Map) == 0 {
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(vecty.Text("name")),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.model.Node.Map {
		rows = append(rows, NewMapRowView(v.ctx, c))
	}

	return elem.Div(
		elem.Table(
			prop.Class("table table-hover"),
			elem.TableHead(
				head,
			),
			elem.TableBody(
				rows,
			),
		),
	)
}
