package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
)

type pkg struct {
	path    string
	aliases map[string]string
	label   *dom.HTMLDivElement
	branch  *tree.Branch
}

var _ tree.Item = (*source)(nil)

func (p *pkg) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(p.path)
	p.label = label
	div.AppendChild(label)
}

func (p *pkg) AddSources(sources []string) {
	for _, name := range sources {
		p.addSource(name, p.branch)
	}
}

func AddPackage(path string, aliases map[string]string, parent *tree.Branch) *pkg {
	p := &pkg{path: path, aliases: aliases}
	b := parent.Tree.Branch(p)
	p.branch = b
	parent.Append(b)
	b.Open()
	return p
}
