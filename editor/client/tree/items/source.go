package items

import (
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system/node"
)

type source struct {
	branch *tree.Branch

	name   string
	pkg    *pkg
	holder *holder

	loaded bool
	node   *node.Node
	editor editor.Editor
}

var _ tree.Item = (*source)(nil)

func (s *source) Initialise() {
	s.branch.SetLabel(s.name)
}

func (parent *holder) addSource(name string) *source {
	s := &source{name: name, holder: parent, pkg: parent.pkg}
	s.branch = tree.NewBranch(s, parent.branch)
	s.Initialise()

	parent.branch.Append(s.branch)

	return s
}

func (parent *holder) addSources(sources []string) {
	for _, name := range sources {
		parent.addSource(name)
	}
}

var _ tree.Editable = (*source)(nil)

func (s *source) Editor() editor.Editor {
	return editor.Default(s.node)
}

var _ tree.Async = (*source)(nil)

func (s *source) ContentLoaded() bool {
	return s.loaded
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
	if err := ke.UnmarshalNode([]byte(gr.Data.Value()), n, s.branch.Tree.Path, s.branch.Tree.Aliases); err != nil {
		return kerr.New("ACODETSACJ", err, "UnmarshalNode")
	}

	s.node = n

	addEntryChildren(n, s.branch)

	s.loaded = true
	successChannel <- true
	return nil
}
