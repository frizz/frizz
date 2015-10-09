package system

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type HasEditor interface {
	GetEditor(node *Node) Editor
}

var _ HasEditor = String{}

func (s String) GetEditor(node *Node) Editor {
	return &StringEditor{String: s, node: node}
}

type Editor interface {
	Initialized() bool
	Initialize(*dom.HTMLDivElement, string, map[string]string) error
	Update()
	Show()
	Hide()
}

var _ Editor = (*StringEditor)(nil)

type StringEditor struct {
	String
	node        *Node
	panel       *dom.HTMLDivElement
	textbox     *dom.HTMLInputElement
	initialized bool
	path        string
	aliases     map[string]string
}

func (e *StringEditor) Initialized() bool {
	return e.initialized
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

func (e *StringEditor) Show() {
	e.panel.Style().Set("display", "block")
}
func (e *StringEditor) Hide() {
	e.panel.Style().Set("display", "none")
}

func (e *StringEditor) Update() {
	if e.Exists {
		e.node.ValueString = e.textbox.Value
	}
}
