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
	Initialize(*dom.HTMLDivElement)
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
}

func (s *StringEditor) Initialized() bool {
	return s.initialized
}

func (s *StringEditor) Initialize(panel *dom.HTMLDivElement) {

	s.panel = panel

	cb := mdl.NewCheckbox("checkbox1", s.Exists, "Exists")
	s.panel.AppendChild(cb)

	tb := mdl.NewTextbox("textbox1", s.Value, s.node.Key)
	s.panel.AppendChild(tb)

	cb.Input.AddEventListener("change", true, func(e dom.Event) {
		tb.SetDisabled(!cb.Input.Checked)
	})

	s.initialized = true
}

func (s *StringEditor) Show() {
	s.panel.Style().Set("display", "block")
}
func (s *StringEditor) Hide() {
	s.panel.Style().Set("display", "none")
}

func (s *StringEditor) Update() {
	if s.Exists {
		s.node.ValueString = s.textbox.Value
	}
}
