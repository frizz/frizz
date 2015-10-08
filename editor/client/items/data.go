package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
)

type data struct {
	*item
	pkg   *pkg
	label *dom.HTMLDivElement
}

var _ tree.Item = (*data)(nil)

func (d *data) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent("data")
	d.label = label
	div.AppendChild(label)
}

func (p *pkg) addData() *data {
	newData := &data{item: &item{tree: p.tree}, pkg: p}
	newBranch := newData.tree.Branch(newData)
	newData.branch = newBranch

	p.branch.Append(newBranch)

	p.data = newData

	newBranch.Open()

	return newData
}
