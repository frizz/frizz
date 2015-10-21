package system

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

func (s String) GetEditor(node *Node) Editor {
	return &StringEditor{String: s, node: node, editorCommon: &editorCommon{}}
}

var _ HasEditor = (*String)(nil)

var _ Editor = (*StringEditor)(nil)

type StringEditor struct {
	String
	*editorCommon
	node     *Node
	textbox  *mdl.Textbox
	checkbox *mdl.Checkbox
}

func (e *StringEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	e.checkbox = mdl.NewCheckbox(e.Exists, "Exists")
	e.panel.AppendChild(e.checkbox)
	e.checkbox.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.textbox.SetDisabled(!e.checkbox.Input.Checked)
		e.Update()
	})

	e.textbox = mdl.NewTextbox(e.Value, e.node.Key)
	e.panel.AppendChild(e.textbox)
	e.textbox.AddEventListener("change", true, func(ev dom.Event) {
		e.Update()
	})

	e.initialized = true
	return nil
}

func (e *StringEditor) Update() {
	e.node.ValueString = e.textbox.Input.Value
	e.node.Null = !e.checkbox.Input.Checked
}
