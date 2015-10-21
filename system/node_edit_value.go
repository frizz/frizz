package system

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/json"
)

type NodeValueEditor struct {
	*Node
	*editorCommon
	input *dom.HTMLInputElement
}

var _ Editor = (*NodeValueEditor)(nil)

func (e *NodeValueEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	switch e.JsonType {
	case json.J_STRING:
		tb := mdl.NewTextbox(e.ValueString, e.Node.Key)
		e.panel.AppendChild(tb)
	case json.J_NUMBER:
		tb := mdl.NewTextbox(strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
		e.panel.AppendChild(tb)
	case json.J_BOOL:
		cb := mdl.NewCheckbox(e.ValueBool, e.Node.Key)
		e.panel.AppendChild(cb)
	}

	e.initialized = true
	return nil
}

func (e *NodeValueEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
