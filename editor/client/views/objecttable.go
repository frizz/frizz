package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
)

type ObjectTableView struct {
	*View

	model  *models.EditorModel
	origin *system.Reference
}

func NewObjectTableView(ctx context.Context, model *models.EditorModel, origin *system.Reference) *ObjectTableView {
	v := &ObjectTableView{}
	v.View = New(ctx, v)
	v.model = model
	v.origin = origin
	v.Mount()
	return v
}

func (v *ObjectTableView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectTableView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ObjectTableView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ObjectTableView) Mount() {
	v.Notifs = v.App.Watch(v.model,
		stores.EditorChildAdded,
		stores.EditorChildDeleted,
	)

	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ObjectTableView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ObjectTableView) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *ObjectTableView) Render() vecty.Component {

	if v.model.Node == nil || len(v.model.Node.Map) == 0 {
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(vecty.Text("name")),
		elem.TableHeader(vecty.Text("holds")),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.model.Node.Map {
		if *c.Origin != *v.origin {
			continue
		}
		rows = append(rows, NewObjectRowView(v.Ctx, c))
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
