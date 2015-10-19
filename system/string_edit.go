package system

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

func (s String) GetEditor(node *Node) Editor {
	return &StringEditor{String: s, node: node}
}

var _ HasEditor = (*String)(nil)

var _ Editor = (*StringEditor)(nil)

type StringEditor struct {
	String
	*editorCommon
	node    *Node
	textbox *dom.HTMLInputElement
}

func (e *StringEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	cb := mdl.NewCheckbox(e.Exists, "Exists")
	e.panel.AppendChild(cb)

	tb := mdl.NewTextbox(e.Value, e.node.Key)
	e.panel.AppendChild(tb)

	cb.Input.AddEventListener("change", true, func(e dom.Event) {
		tb.SetDisabled(!cb.Input.Checked)
	})

	e.initialized = true
	return nil
}

func (e *StringEditor) Update() {
	if e.Exists {
		e.node.ValueString = e.textbox.Value
	}
}
