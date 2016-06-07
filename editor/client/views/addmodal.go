package views

import (
	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/json"
	"kego.io/system/node"
)

type AddModalView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	parent *node.Node
	node   *node.Node
}

func NewAddModalView(ctx context.Context, parent *node.Node, node *node.Node) *AddModalView {
	v := &AddModalView{
		ctx:    ctx,
		app:    stores.FromContext(ctx),
		parent: parent,
		node:   node,
	}
	v.Mount()
	return v
}

func (v *AddModalView) Reconcile(old vecty.Component) {
	if old, ok := old.(*AddModalView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *AddModalView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *AddModalView) Mount() {
	v.notifs = v.app.Nodes.Watch(nil, stores.AddModalChange)

	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *AddModalView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.node = v.app.Nodes.AddModalNode()
	v.parent = v.app.Nodes.AddModalParent()
	v.ReconcileBody()
	if v.app.Nodes.AddModalVisible() {
		js.Global.Call("$", "#add-modal").Call("modal", "show")
	} else {
		js.Global.Call("$", "#add-modal").Call("modal", "hide")
	}
}

func (v *AddModalView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *AddModalView) render() vecty.Component {
	if v.parent == nil {
		return v.modal()
	}

	/*


		card := mdl.Card("Choose a type", "Add")
		options := map[string]string{}

		for _, t := range types {
			displayName, err := t.Id.ValueContext(e.ctx)
			if err != nil {
				// we shouldn't be able to get here
				e.fail <- kerr.Wrap("IPLHSXDWQK", err)
			}
			options[t.Id.String()] = displayName
		}

		dropdown := mdl.Select(options)
		card.Content.AppendChild(dropdown)

		go func() {
			result := <-card.Result
			if result {
				r, err := system.NewReferenceFromString(e.ctx, dropdown.Value)
				if err != nil {
					e.fail <- kerr.Wrap("KHJGQXORPD", err)
				}
				t, ok := system.GetTypeFromCache(e.ctx, r.Package, r.Name)
				if !ok {
					e.fail <- kerr.New("WEADSXTPYC", "Type %s not found in cache", r.Value())
				}
				e.InitialiseChildWithConcreteType(node, t)
			}
			card.Hide()
		}()
	*/

	var nameControl, typeControl vecty.Markup
	if v.parent.JsonType == json.J_MAP {
		nameControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-name"),
				vecty.Text("name"),
			),
			elem.Input(
				prop.Class("form-control"),
				prop.ID("add-modal-name"),
			),
			elem.Paragraph(
				prop.Class("help-block"),
				vecty.Text("Enter a name for the new item"),
			),
		)
	}
	if len(v.app.Nodes.AddModalTypes()) > 1 {

		var options vecty.List
		for _, t := range v.app.Nodes.AddModalTypes() {
			displayName, err := t.Id.ValueContext(v.ctx)
			if err != nil {
				// we shouldn't be able to get here
				v.app.Fail <- kerr.Wrap("NTSWPIEAHC", err)
			}
			options = append(options, elem.Option(
				prop.ID(t.Id.String()),
				vecty.Text(displayName),
			))
		}

		typeControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-type"),
				vecty.Text("type"),
			),
			elem.Select(
				prop.Class("form-control"),
				prop.ID("add-modal-type"),
				options,
			),
			elem.Paragraph(
				prop.Class("help-block"),
				vecty.Text("Select a type for the new object"),
			),
		)
	}

	return v.modal(
		elem.Div(
			prop.Class("modal-body"),
			elem.Form(
				nameControl,
				typeControl,
			),
		),
	)
}

func (v *AddModalView) modal(markup ...vecty.Markup) *vecty.Element {

	return elem.Div(
		prop.ID("add-modal"),
		prop.Class("modal fade"),
		vecty.Data("backdrop", "static"),
		vecty.Data("keyboard", "false"),
		elem.Div(
			prop.Class("modal-dialog"),
			elem.Div(
				prop.Class("modal-content"),
				elem.Div(
					prop.Class("modal-header"),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("close"),
						elem.Span(
							vecty.Text("×"),
						),
						event.Click(func(ev *vecty.Event) {
							v.app.Dispatch(&actions.AddModalCloseClick{})
						}).PreventDefault(),
					),
					elem.Header4(
						prop.Class("modal-title"),
						vecty.Text("Add item"),
					),
				),
				vecty.List(markup),
				elem.Div(
					prop.Class("modal-footer"),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("btn btn-default"),
						elem.Span(
							vecty.Text("Close"),
						),
						event.Click(func(ev *vecty.Event) {
							v.app.Dispatch(&actions.AddModalCloseClick{})
						}).PreventDefault(),
					),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("btn btn-primary"),
						elem.Span(
							vecty.Text("Save"),
						),
						event.Click(func(ev *vecty.Event) {
							v.app.Dispatch(nil)
						}).PreventDefault(),
					),
				),
			),
		),
	)
}
