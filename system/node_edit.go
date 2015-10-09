package system

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
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

func (e *NodeEditor) Initialized() bool {
	return e.initialized
}

func (e *NodeEditor) Initialize(panel *dom.HTMLDivElement) {

	e.panel = panel

	/*
		if e.Key != "" {
			nameLabel := dom.GetWindow().Document().CreateElement("h1").(*dom.HTMLHeadingElement)
			nameLabel.SetTextContent(e.Key)
			e.panel.AppendChild(nameLabel)
		}

		typeLabel := dom.GetWindow().Document().CreateElement("h2").(*dom.HTMLHeadingElement)
		typeLabel.SetTextContent(e.Type.Id.Value())
		e.panel.AppendChild(typeLabel)
	*/
	switch e.JsonType {
	case json.J_STRING:
		tb := mdl.NewTextbox("textbox1", e.ValueString, e.Node.Key)
		e.panel.AppendChild(tb)
	case json.J_NUMBER:
		tb := mdl.NewTextbox("textbox1", strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
		e.panel.AppendChild(tb)
	case json.J_BOOL:
		cb := mdl.NewCheckbox("checkbox1", e.ValueBool, e.Node.Key)
		e.panel.AppendChild(cb)
	}

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
