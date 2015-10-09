package system

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/json"
)

func (n *Node) GetEditor() Editor {
	return &NodeEditor{Node: n}
}

var _ Editor = (*NodeEditor)(nil)

type NodeEditor struct {
	*Node
	panel       *dom.HTMLDivElement
	input       *dom.HTMLInputElement
	initialized bool
}

func (e *NodeEditor) Initialize(content *dom.HTMLDivElement) {
	if e.initialized {
		return
	}
	e.panel = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	// panel should be hidden by default
	e.panel.Style().Set("display", "none")

	typeLabel := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	typeLabel.SetTextContent(e.Type.Id.Value())
	e.panel.AppendChild(typeLabel)

	switch e.JsonType {
	case json.J_STRING:
		e.input = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
		e.input.Type = "text"
		e.input.Value = e.ValueString
		e.panel.AppendChild(e.input)
	case json.J_NUMBER:
		e.input = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
		e.input.Type = "text"
		e.input.Value = strconv.FormatFloat(e.ValueNumber, 'f', -1, 64)
		e.panel.AppendChild(e.input)
	case json.J_BOOL:
		e.input = dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
		e.input.Type = "checkbox"
		e.input.Checked = e.ValueBool
		e.panel.AppendChild(e.input)
	}

	content.AppendChild(e.panel)

	e.initialized = true
}

func (e *NodeEditor) Show() {
	e.panel.Style().Set("display", "block")
}
func (e *NodeEditor) Hide() {
	e.panel.Style().Set("display", "none")
}

func (e *NodeEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
