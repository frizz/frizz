package tree

import (
	"kego.io/context/envctx"
	"kego.io/editor"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
)

type source struct {
	*Branch
	*Async
	*editor.Node

	name   string
	file   string
	pkg    *pkg
	holder *holder

	editor editor.EditorInterface
}

var _ BranchInterface = (*source)(nil)
var _ Editable = (*source)(nil)
var _ AsyncInterface = (*source)(nil)

func (parent *holder) addSource(name string, file string) *source {
	s := &source{name: name, file: file, holder: parent, pkg: parent.pkg}
	s.Branch = NewBranch(s, parent)
	s.Async = NewAsync(s)
	s.setLabel(name)
	parent.Append(s)
	return s
}

func (parent *holder) addSources(sources map[string]string) {
	for name, file := range sources {
		parent.addSource(name, file)
	}
}

func (s *source) ContentRequest() messages.MessageInterface {
	env := envctx.FromContext(s.tree.ctx)
	return messages.NewDataRequest(env.Path, s.name, s.file)
}

func (s *source) ProcessResponse(response messages.MessageInterface) error {

	gr, ok := response.(*messages.DataResponse)
	if !ok {
		return kerr.New("MVPKNNVHOX", nil, "%T is not a *messages.SourceResponse", response)
	}

	s.Node = editor.NewEditorNode()
	if err := ke.UnmarshalUntyped(s.tree.ctx, []byte(gr.Data.Value()), s.Node); err != nil {
		return kerr.New("ACODETSACJ", err, "UnmarshalNode")
	}

	ed := s.Node.Editor()

	if err := ed.Initialize(s.tree.ctx, s, editor.Page, s.tree.Fail); err != nil {
		return kerr.New("UOPUXTANHO", err, "Initialize")
	}

	s.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(s.Node, s, ed); err != nil {
		return kerr.New("MLUGRXOWHC", err, "addEntryChildren")
	}

	return nil
}
