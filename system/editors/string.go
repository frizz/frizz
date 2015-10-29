package editors // import "kego.io/system/editors"

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/mdl"
	"kego.io/system"
	"kego.io/system/node"
)

func init() {
	// Will leave this disabled for now... I think it's not really
	// needed. "Exists" checkbox shouldn't be in the editor UI.

	//editor.Register(system.NewReference("kego.io/system", "string"), editor.Factory(func(n *node.Node) editor.Editor {
	//return &StringEditor{String: n.Value.(system.String), node: n, Common: &editor.Common{}}
	//}))
}

var _ editor.Editor = (*StringEditor)(nil)

type StringEditor struct {
	system.String
	*editor.Common
	node     *node.Node
	textbox  *mdl.Textbox
	checkbox *mdl.Checkbox
}

func (e *StringEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.Panel = panel
	e.Path = path
	e.Aliases = aliases

	e.checkbox = mdl.NewCheckbox(e.Exists, "Exists")
	e.Panel.AppendChild(e.checkbox)
	e.checkbox.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.textbox.SetDisabled(!e.checkbox.Input.Checked)
		e.node.Null = !e.checkbox.Input.Checked
	})

	e.textbox = mdl.NewTextbox(e.Value, e.node.Key)
	e.Panel.AppendChild(e.textbox)
	e.textbox.AddEventListener("change", true, func(ev dom.Event) {
		e.node.ValueString = e.textbox.Input.Value
	})

	e.Initialized = true
	return nil
}
