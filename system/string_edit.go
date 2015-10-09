package system

import "honnef.co/go/js/dom"

type HasEditor interface {
	GetEditor(node *Node) Editor
}

var _ HasEditor = String{}

func (s String) GetEditor(node *Node) Editor {
	return &StringEditor{String: s, node: node}
}

type Editor interface {
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

func (s *StringEditor) Initialize(content *dom.HTMLDivElement) {
	if s.initialized {
		return
	}
	s.panel = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	// panel should be hidden by default
	s.panel.Style().Set("display", "none")

	s.textbox = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	s.textbox.Type = "text"
	s.textbox.Value = s.Value

	s.panel.AppendChild(s.textbox)

	content.AppendChild(s.panel)

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
