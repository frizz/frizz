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

	/*
		<div class="mdl-textfield mdl-js-textfield">
			<input class="mdl-textfield__input" type="text" id="sample1" />
			<label class="mdl-textfield__label" for="sample1">Text...</label>
		  </div>
	*/

	s.textbox = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	s.textbox.Type = "text"
	s.textbox.Value = s.Value

	s.panel.AppendChild(s.textbox)
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
