package system

import "honnef.co/go/js/dom"

type Editor interface {
	Initialize()
	Update()
}

func (s *String) GetEditor(node *Node, panel *dom.HTMLDivElement) Editor {
	return &StringEditor{}
}

type StringEditor struct {
	String
	node        *Node
	panel       *dom.HTMLDivElement
	textbox     *dom.HTMLInputElement
	initialized bool
}

func (s *StringEditor) Initialize() {
	if s.initialized {
		return
	}
	s.textbox = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	s.textbox.Type = "text"
	s.textbox.Value = s.Value
	s.panel.AppendChild(s.textbox)
	s.initialized = true
}

func (s *StringEditor) Update() {
	if s.Exists {
		s.node.ValueString = s.textbox.Value
	}
}
