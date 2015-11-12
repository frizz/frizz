package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type TableStruct struct {
	*dom.HTMLTableElement
	head *dom.HTMLTableRowElement
	body *dom.HTMLTableSectionElement
}

func Table() *TableStruct {
	t := dom.GetWindow().Document().CreateElement("table").(*dom.HTMLTableElement)
	t.Class().SetString("mdl-data-table mdl-js-data-table mdl-shadow--2dp")

	thead := dom.GetWindow().Document().CreateElement("thead").(*dom.HTMLTableSectionElement)
	t.AppendChild(thead)

	tr := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
	thead.AppendChild(tr)

	tbody := dom.GetWindow().Document().CreateElement("tbody").(*dom.HTMLTableSectionElement)
	t.AppendChild(tbody)

	return &TableStruct{t, tr, tbody}
}

func (t *TableStruct) Head(columns ...string) {
	for _, column := range columns {
		th := dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
		th.Class().Add("mdl-data-table__cell--non-numeric")
		th.SetTextContent(column)
		t.head.AppendChild(th)
	}
}

type row struct {
	*dom.HTMLTableRowElement
	table *TableStruct
}

func (t *TableStruct) Row() *row {
	tr := dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
	t.body.AppendChild(tr)
	return &row{tr, t}
}

type cell struct {
	*dom.HTMLTableCellElement
	row *row
}

func (r *row) Cell() *cell {
	td := dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
	td.Class().Add("mdl-data-table__cell--non-numeric")
	r.AppendChild(td)

	return &cell{td, r}
}

func (c *cell) Text(text string) *cell {
	c.SetTextContent(text)
	return c
}

func (t *TableStruct) Upgrade() {
	js.Global.Get("componentHandler").Call("upgradeElement", t)
}
