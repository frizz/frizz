package views // import "kego.io/editor/client/views"

import (
	"context"

	"reflect"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/context/sysctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type PanelView struct {
	*View

	branch *models.BranchModel
	node   *node.Node
}

func NewPanelView(ctx context.Context) *PanelView {
	v := &PanelView{}
	v.View = New(ctx, v)
	v.Watch(nil,
		stores.BranchSelected,
		stores.ViewChanged,
		stores.BranchInitialStateLoaded,
	)
	return v
}

func (v *PanelView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PanelView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *PanelView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.branch = v.App.Branches.Selected()
	v.node = v.App.Nodes.Selected()
	v.ReconcileBody()
	if notif.Type == stores.BranchSelected {
		v.Node().Get("parentNode").Set("scrollTop", "0")
	}
}

func (v *PanelView) Render() vecty.Component {
	var editor vecty.Component
	if v.node != nil {
		if v.node.Type.IsNativeMap() {
			editor = NewMapView(v.Ctx, v.node)
		} else if v.node.Type.IsNativeArray() {
			editor = NewArrayView(v.Ctx, v.node)
		} else {
			editor = NewStructView(v.Ctx, v.node)
		}
	} else if v.branch != nil {
		switch v.branch.Contents.(type) {
		case *models.DataContents:
			editor = NewPanelNavView(v.Ctx, v.branch).Contents(
				elem.UnorderedList(
					prop.Class("nav navbar-nav navbar-right"),
					elem.ListItem(
						elem.Anchor(
							vecty.Text("Add"),
							prop.Href("#"),
							event.Click(func(ev *vecty.Event) {
								addNewFile(v.Ctx, v.App, true)
							}).PreventDefault(),
						),
					),
				),
			)
		case *models.TypesContents:
			editor = NewPanelNavView(v.Ctx, v.branch).Contents(
				elem.UnorderedList(
					prop.Class("nav navbar-nav navbar-right"),
					elem.ListItem(
						elem.Anchor(
							vecty.Text("Add"),
							prop.Href("#"),
							event.Click(func(ev *vecty.Event) {
								addNewFile(v.Ctx, v.App, false)
							}).PreventDefault(),
						),
					),
				),
			)
		default:
			editor = NewPanelNavView(v.Ctx, v.branch)
		}
	}
	return elem.Div(
		prop.Class("content content-panel"),
		editor,
	)
}

func addNewFile(ctx context.Context, app *stores.App, all bool) {

	var types []*system.Type
	if all {
		rt := reflect.TypeOf((*system.ObjectInterface)(nil)).Elem()
		types = system.GetAllTypesThatImplementReflectInterface(ctx, rt)
	} else {
		syscache := sysctx.FromContext(ctx)
		t, ok := syscache.GetType("kego.io/system", "type")
		if !ok {
			panic(kerr.New("NNFSJEXNKF", "Can't find system:type in sys ctx").Error())
		}
		types = []*system.Type{t.(*system.Type)}
	}

	app.Dispatch(&actions.OpenAddPopup{
		Types: types,
	})

}
