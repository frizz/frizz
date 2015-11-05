package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system/node"
)

type source struct {
	*item
	node   *node.Node
	name   string
	loaded bool
	label  *dom.HTMLDivElement
	conn   *connection.Conn
	pkg    *pkg
	data   *holder
}

var _ tree.Item = (*source)(nil)
var _ tree.Async = (*source)(nil)
var _ tree.Noder = (*source)(nil)

func (s *source) Node() *node.Node {
	return s.node
}

func (s *source) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(s.name)
	s.label = label
	div.AppendChild(label)
}

func (s *source) LoadContent() chan bool {

	successChannel := make(chan bool, 1)

	if s.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it. This has to be done with a goroutine.
		successChannel <- true
		return successChannel
	}

	responseChannel := s.branch.Tree.Conn.Request(messages.NewSourceRequest(s.name))

	go func() {
		if err := s.awaitSourceResponse(responseChannel, successChannel); err != nil {
			s.branch.Tree.Fail <- kerr.New("DLTCGMSREX", err, "awaitSourceResponse")
		}
	}()

	return successChannel

}

func (s *source) awaitSourceResponse(responseChannel chan messages.MessageInterface, successChannel chan bool) error {

	m := <-responseChannel

	gr, ok := m.(*messages.SourceResponse)
	if !ok {
		return kerr.New("MVPKNNVHOX", nil, "%T is not a *messages.SourceResponse", m)
	}

	n := &node.Node{}
	if err := ke.UnmarshalNode([]byte(gr.Data.Value()), n, s.tree.Path, s.tree.Aliases); err != nil {
		return kerr.New("ACODETSACJ", err, "UnmarshalNode")
	}

	s.node = n

	addNodeChildren(n, s.branch)

	s.loaded = true
	successChannel <- true
	return nil
}

func (s *source) ContentLoaded() bool {
	return s.loaded
}

func (d *holder) addSource(name string, parent *tree.Branch) *source {
	newSource := &source{item: &item{tree: parent.Tree}, name: name, data: d, pkg: d.pkg}
	newBranch := parent.Tree.Branch(newSource)
	newSource.branch = newBranch

	parent.Append(newBranch)

	return newSource
}

func (d *holder) addSources(sources []string) {
	for _, name := range sources {
		d.addSource(name, d.branch)
	}
}
