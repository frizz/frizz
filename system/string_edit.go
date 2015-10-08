package system

import "honnef.co/go/js/dom"

func (s String) InitializeEditControl(node *Node, panel dom.HTMLDivElement) {
	textbox := dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
	textbox.Type = "text"
	textbox.Value = s.Value
	panel.AppendChild(textbox)
}

func (s String) UpdateNode(node *Node) {
	if s.Exists {
		node.ValueString = s.Value
	}

}
