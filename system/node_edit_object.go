package system

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"kego.io/kerr"
)

type NodeObjectEditor struct {
	*Node
	path        string
	aliases     map[string]string
	panel       *dom.HTMLDivElement
	input       *dom.HTMLInputElement
	initialized bool
}

var _ Editor = (*NodeObjectEditor)(nil)

func (e *NodeObjectEditor) Initialized() bool {
	return e.initialized
}

func (e *NodeObjectEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	/*
		<table class="mdl-data-table mdl-js-data-table mdl-data-table--selectable mdl-shadow--2dp">
		  <thead>
		    <tr>
		      <th class="mdl-data-table__cell--non-numeric">Material</th>
		      <th>Quantity</th>
		      <th>Unit price</th>
		    </tr>
		  </thead>
		  <tbody>
		    <tr>
		      <td class="mdl-data-table__cell--non-numeric">Acrylic (Transparent)</td>
		      <td>25</td>
		      <td>$2.90</td>
		    </tr>
		    <tr>
		      <td class="mdl-data-table__cell--non-numeric">Plywood (Birch)</td>
		      <td>50</td>
		      <td>$1.25</td>
		    </tr>
		    <tr>
		      <td class="mdl-data-table__cell--non-numeric">Laminate (Gold on Blue)</td>
		      <td>10</td>
		      <td>$2.35</td>
		    </tr>
		  </tbody>
		</table>
	*/

	table := dom.GetWindow().Document().CreateElement("table").(*dom.HTMLTableElement)
	table.Class().SetString("mdl-data-table mdl-js-data-table mdl-shadow--2dp")

	thead := dom.GetWindow().Document().CreateElement("thead").(*dom.HTMLTableSectionElement)

	theadRow := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
	thead.AppendChild(theadRow)

	headName := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
	headName.Class().Add("mdl-data-table__cell--non-numeric")
	headName.SetTextContent("name")
	theadRow.AppendChild(headName)

	headOrigin := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
	headOrigin.Class().Add("mdl-data-table__cell--non-numeric")
	headOrigin.SetTextContent("origin")
	theadRow.AppendChild(headOrigin)

	headType := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
	headType.Class().Add("mdl-data-table__cell--non-numeric")
	headType.SetTextContent("holds")
	theadRow.AppendChild(headType)

	headValue := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
	headValue.Class().Add("mdl-data-table__cell--non-numeric")
	headValue.SetTextContent("value")
	theadRow.AppendChild(headValue)

	tbody := dom.GetWindow().Document().CreateElement("tbody").(*dom.HTMLTableSectionElement)

	for name, field := range e.Fields {

		tr := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
		tbody.AppendChild(tr)

		nameCell := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
		nameCell.Class().Add("mdl-data-table__cell--non-numeric")
		nameCell.SetTextContent(name)
		tr.AppendChild(nameCell)

		originCell := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
		originCell.Class().Add("mdl-data-table__cell--non-numeric")
		val, err := field.Origin.ValueContext(e.path, e.aliases)
		if err != nil {
			return kerr.New("ACQLJXWYQX", err, "ValueContext")
		}
		originCell.SetTextContent(val)
		tr.AppendChild(originCell)

		typeCell := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
		typeCell.Class().Add("mdl-data-table__cell--non-numeric")
		val, err = field.Rule.ParentType.Id.ValueContext(e.path, e.aliases)
		if err != nil {
			return kerr.New("XDKOSFJVQV", err, "ValueContext")
		}
		typeCell.SetTextContent(val)
		tr.AppendChild(typeCell)

		valueCell := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
		valueCell.Class().Add("mdl-data-table__cell--non-numeric")
		if field.Missing || field.Null {
			valueCell.SetTextContent("")
		} else {
			val, err = field.Type.Id.ValueContext(e.path, e.aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			valueCell.SetTextContent(val)
		}
		tr.AppendChild(valueCell)

	}

	table.AppendChild(thead)
	table.AppendChild(tbody)

	e.panel.AppendChild(table)

	js.Global.Get("componentHandler").Call("upgradeElement", table)

	e.initialized = true
	return nil
}

func (e *NodeObjectEditor) Show() {
	e.panel.Style().Set("display", "block")
}
func (e *NodeObjectEditor) Hide() {
	e.panel.Style().Set("display", "none")
}

func (e *NodeObjectEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
