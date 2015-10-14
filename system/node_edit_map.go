package system

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
)

type NodeMapEditor struct {
	*Node
	path        string
	aliases     map[string]string
	panel       *dom.HTMLDivElement
	input       *dom.HTMLInputElement
	initialized bool
}

var _ Editor = (*NodeMapEditor)(nil)

func (e *NodeMapEditor) Initialized() bool {
	return e.initialized
}

func (e *NodeMapEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	table := mdl.Table()

	items := e.Rule.Rule.(CollectionRule).GetItemsRule()
	itemsHolds := items.(ObjectInterface).GetObject().Type
	if items.GetRule().Interface {
		itemsHolds = itemsHolds.TypeInterfaceReference()
	}
	hold, err := itemsHolds.ValueContext(e.path, e.aliases)
	if err != nil {
		return kerr.New("XDKOSFJVQV", err, "ValueContext")
	}

	names := table.Column("name")
	holds := table.Column("holds")
	values := table.Column("value")

	for name, item := range e.Map {

		names.Cell(name)
		holds.Cell(hold)

		if item.Missing || item.Null {
			values.Cell("")
		} else {
			val, err := item.Type.Id.ValueContext(e.path, e.aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			values.Cell(val)
		}

	}
	e.panel.AppendChild(table.Build())

	e.initialized = true
	return nil
}

func (e *NodeMapEditor) Show() {
	e.panel.Style().Set("display", "block")
}
func (e *NodeMapEditor) Hide() {
	e.panel.Style().Set("display", "none")
}

func (e *NodeMapEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
