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
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/json"
	"kego.io/system"
	"kego.io/system/node"
)

type AddPopupView struct {
	*View

	model      *models.AddPopupModel
	nameInput  *vecty.Element
	typeSelect *vecty.Element
}

func NewAddPopupView(ctx context.Context) *AddPopupView {
	v := &AddPopupView{}
	v.View = New(ctx, v)
	v.Mount()
	return v
}

func (v *AddPopupView) Reconcile(old vecty.Component) {
	if old, ok := old.(*AddPopupView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *AddPopupView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *AddPopupView) Mount() {
	v.Notifs = v.App.Watch(nil,
		stores.AddPopupChange,
	)

	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()
}

func (v *AddPopupView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.model = v.App.Misc.AddPopup()
	v.ReconcileBody()
	if v.model.Visible {
		js.Global.Call("$", "#add-modal").Call("modal", "show")
		if v.model.Parent.JsonType == json.J_MAP {
			js.Global.Call("$", "#add-modal-name").Call("focus")
		} else {
			js.Global.Call("$", "#add-modal-type").Call("focus")
		}
	} else {
		js.Global.Call("$", "#add-modal").Call("modal", "hide")
	}
}

func (v *AddPopupView) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *AddPopupView) Render() vecty.Component {
	if v.model == nil || !v.model.Visible {
		return v.modal()
	}

	v.nameInput = nil
	v.typeSelect = nil

	var nameControl, typeControl vecty.Markup
	if v.model.Parent.JsonType == json.J_MAP {
		v.nameInput = elem.Input(
			prop.Class("form-control"),
			prop.ID("add-modal-name"),
		)
		nameControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-name"),
				vecty.Text("name"),
			),
			v.nameInput,
			elem.Paragraph(
				prop.Class("help-block"),
				vecty.Text("Enter a name for the new item"),
			),
		)
	}
	if len(v.model.Types) > 1 {

		var options vecty.List
		for _, t := range v.model.Types {
			displayName, err := t.Id.ValueContext(v.Ctx)
			if err != nil {
				// we shouldn't be able to get here
				v.App.Fail <- kerr.Wrap("NTSWPIEAHC", err)
			}
			options = append(options, elem.Option(
				prop.ID(t.Id.String()),
				vecty.Text(displayName),
			))
		}

		v.typeSelect = elem.Select(
			prop.Class("form-control"),
			prop.ID("add-modal-type"),
			elem.Option(
				prop.ID(""),
				vecty.Text(""),
			),
			options,
		)

		typeControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-type"),
				vecty.Text("type"),
			),
			v.typeSelect,
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
				event.Submit(func(ev *vecty.Event) {
					v.save()
				}).PreventDefault(),
				nameControl,
				typeControl,
			),
		),
	)
}

func (v *AddPopupView) modal(markup ...vecty.Markup) *vecty.Element {

	return elem.Div(
		prop.ID("add-modal"),
		prop.Class("modal"),
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
							v.App.Dispatch(&actions.CloseAddPopup{})
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
							v.App.Dispatch(&actions.CloseAddPopup{})
						}).PreventDefault(),
					),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("btn btn-primary"),
						elem.Span(
							vecty.Text("Save"),
						),
						event.Click(func(ev *vecty.Event) {
							v.save()
						}).PreventDefault(),
					),
				),
			),
		),
	)
}

func (v *AddPopupView) save() {

	var t *system.Type
	if len(v.model.Types) == 1 {
		t = v.model.Types[0]
	} else {
		options := v.typeSelect.Node().Get("options")
		selectedIndex := v.typeSelect.Node().Get("selectedIndex").Int()
		value := options.Index(selectedIndex).Get("id").String()
		if value == "" {
			return
		}
		r, err := system.NewReferenceFromString(v.Ctx, value)
		if err != nil {
			v.App.Fail <- kerr.Wrap("SEMCIELKRN", err)
			return
		}
		ty, ok := system.GetTypeFromCache(v.Ctx, r.Package, r.Name)
		if !ok {
			v.App.Fail <- kerr.New("RWHSCOFNQM", "Type %s not found in cache", r.Value())
			return
		}
		t = ty
	}

	if v.model.Parent.Type.IsNativeMap() {
		name := v.nameInput.Node().Get("value").String()
		if name == "" {
			// TODO: show an error
			return
		}
		if _, duplicate := v.model.Parent.Map[name]; duplicate {
			// TODO: show an error
			return
		}
		v.App.Dispatch(&actions.InitializeNode{
			Node:   node.NewNode(),
			Parent: v.model.Parent,
			Key:    name,
			Type:   t,
		})
	} else if v.model.Parent.Type.IsNativeArray() {
		v.App.Dispatch(&actions.InitializeNode{
			Node:   node.NewNode(),
			Parent: v.model.Parent,
			Type:   t,
		})
	} else {
		v.App.Dispatch(&actions.InitializeNode{
			Parent: v.model.Parent,
			Node:   v.model.Node,
			Type:   t,
		})
	}

	v.App.Dispatch(&actions.CloseAddPopup{})

}
