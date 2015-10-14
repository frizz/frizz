package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type table struct {
	columns []*column
}
type column struct {
	heading string
	cells   []string
}

func Table() *table {
	return &table{}
}
func (t *table) Column(heading string) *column {
	c := &column{heading, []string{}}
	t.columns = append(t.columns, c)
	return c
}
func (c *column) Cell(value string) {
	c.cells = append(c.cells, value)
}
func (t *table) Build() *dom.HTMLTableElement {

	table := dom.GetWindow().Document().CreateElement("table").(*dom.HTMLTableElement)
	table.Class().SetString("mdl-data-table mdl-js-data-table mdl-shadow--2dp")

	thead := dom.GetWindow().Document().CreateElement("thead").(*dom.HTMLTableSectionElement)

	tr := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
	thead.AppendChild(tr)

	for _, c := range t.columns {
		th := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
		th.Class().Add("mdl-data-table__cell--non-numeric")
		th.SetTextContent(c.heading)
		tr.AppendChild(th)
	}

	tbody := dom.GetWindow().Document().CreateElement("tbody").(*dom.HTMLTableSectionElement)

	for i, _ := range t.columns[0].cells {
		tr := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
		tbody.AppendChild(tr)

		for j, _ := range t.columns {
			td := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
			td.Class().Add("mdl-data-table__cell--non-numeric")
			td.SetTextContent(t.columns[j].cells[i])
			tr.AppendChild(td)
		}
	}
	table.AppendChild(thead)
	table.AppendChild(tbody)
	js.Global.Get("componentHandler").Call("upgradeElement", table)
	return table
}
