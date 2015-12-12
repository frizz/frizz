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
	t := get("table").(*dom.HTMLTableElement)
	t.Class().SetString("mdl-data-table mdl-js-data-table mdl-shadow--2dp")

	thead := get("thead").(*dom.HTMLTableSectionElement)
	t.AppendChild(thead)

	tr := get("tr").(*dom.HTMLTableRowElement)
	thead.AppendChild(tr)

	tbody := get("tbody").(*dom.HTMLTableSectionElement)
	t.AppendChild(tbody)

	return &TableStruct{t, tr, tbody}
}

func (t *TableStruct) Head(columns ...string) {
	for _, column := range columns {
		th := get("th").(*dom.HTMLTableCellElement)
		th.Class().Add("mdl-data-table__cell--non-numeric")
		th.SetTextContent(column)
		t.head.AppendChild(th)
	}
}

type TableRowStruct struct {
	*dom.HTMLTableRowElement
	table *TableStruct
}

func (r *TableRowStruct) Click(action func(dom.Event)) *TableRowStruct {
	r.AddEventListener("click", true, action)
	r.Style().Set("cursor", "pointer")
	return r
}

func (t *TableStruct) Row() *TableRowStruct {
	tr := get("tr").(*dom.HTMLTableRowElement)
	t.body.AppendChild(tr)
	return &TableRowStruct{tr, t}
}

type TableCellStruct struct {
	*dom.HTMLTableCellElement
	row *TableRowStruct
}

func (r *TableRowStruct) Cell() *TableCellStruct {
	td := get("td").(*dom.HTMLTableCellElement)
	td.Class().Add("mdl-data-table__cell--non-numeric")
	r.AppendChild(td)

	return &TableCellStruct{td, r}
}

func (c *TableCellStruct) Text(text string) *TableCellStruct {
	c.SetTextContent(text)
	return c
}

func (t *TableStruct) Upgrade() {
	js.Global.Get("componentHandler").Call("upgradeElement", t)
}
