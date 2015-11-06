package branch

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system/node"
)

type source struct {
	*Common
	node   *node.Node
	name   string
	loaded bool
	label  *dom.HTMLSpanElement
	conn   *connection.Conn
	pkg    *pkg
	data   *holder
	editor editor.Editor
}

var _ tree.Async = (*source)(nil)
var _ tree.Editable = (*source)(nil)

func (s *source) Editor() editor.Editor {
	return editor.Default(s.node)
}

func (s *source) Initialise(parent tree.Branch) {
	s.Common = &Common{this: s, tree: parent.Tree()}
	s.Common.Initialise(parent)
	s.Common.SetLabel(s.name)
}

func (s *source) LoadContent() chan bool {

	successChannel := make(chan bool, 1)

	if s.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it. This has to be done with a goroutine.
		successChannel <- true
		return successChannel
	}

	responseChannel := s.tree.Conn.Request(messages.NewSourceRequest(s.name))

	go func() {
		if err := s.awaitSourceResponse(responseChannel, successChannel); err != nil {
			s.tree.Fail <- kerr.New("DLTCGMSREX", err, "awaitSourceResponse")
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

	addEntryChildren(n, s)

	s.loaded = true
	successChannel <- true
	return nil
}

func (s *source) ContentLoaded() bool {
	return s.loaded
}

func (parent *holder) addSource(name string) *source {
	s := &source{name: name, data: parent, pkg: parent.pkg}
	s.Initialise(parent)

	parent.Append(s)

	return s
}

func (parent *holder) addSources(sources []string) {
	for _, name := range sources {
		parent.addSource(name)
	}
}
