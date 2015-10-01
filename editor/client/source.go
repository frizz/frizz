package client

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared/messages"
	"kego.io/helper"
	"kego.io/json"
	"kego.io/kerr"
)

type source struct {
	name   string
	loaded bool
	label  *dom.HTMLDivElement
	node   *tree.Node
}

func (s *source) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent(s.name)
	s.label = label
	div.AppendChild(label)
}

func (s *source) LoadContent() chan bool {

	successChannel := make(chan bool)

	if s.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it. This has to be done with a goroutine.
		go func() { successChannel <- true }()
		return successChannel
	}

	responseChannel := app.conn.Request(messages.NewSourceRequest(s.name))

	go func() {
		if err := s.awaitSourceResponse(responseChannel, successChannel); err != nil {
			app.fail <- kerr.New("DLTCGMSREX", err, "awaitSourceResponse")
		}
	}()

	return successChannel

}

func (s *source) awaitSourceResponse(responseChannel chan messages.Message, successChannel chan bool) error {

	m := <-responseChannel

	gr, ok := m.(*messages.SourceResponse)
	if !ok {
		return kerr.New("MVPKNNVHOX", nil, "%T is not a *messages.SourceResponse", m)
	}

	n := &helper.Node{}
	if err := json.UnmarshalPlainContext([]byte(gr.Data.Value), n, app.path, app.aliases); err != nil {
		return kerr.New("ACODETSACJ", err, "UnmarshalPlainContext")
	}

	child := &entry{index: -1, name: gr.Name.Value, node: n}
	if err := addEntryChildren(child, s.node); err != nil {
		return kerr.New("PFVRMGEJMF", err, "addEntryChildren (root)")
	}

	s.loaded = true
	successChannel <- true
	return nil
}

func (s *source) ContentLoaded() bool {
	return s.loaded
}

var _ tree.Item = (*source)(nil)
var _ tree.AsyncItem = (*source)(nil)

func addSource(name string, parent *tree.Node) {
	s := &source{name: name}
	n := tree.NewNode(s)
	s.node = n
	parent.AppendNodes(n)
}
