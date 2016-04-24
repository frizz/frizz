package tree

import (
	"net/rpc"

	"kego.io/context/envctx"
	"kego.io/editor"
	"kego.io/editor/shared"
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

func (s *source) MakeCall(fail chan error) *rpc.Call {
	request := &shared.DataRequest{
		File:    s.file,
		Name:    s.name,
		Package: envctx.FromContext(s.tree.ctx).Path,
	}
	data := shared.DataResponse{}
	done := make(chan *rpc.Call, 1)
	return s.tree.Conn.Go("Server.Data", request, &data, done, fail)
}

func (s *source) ProcessResponse(response interface{}) error {

	gr, ok := response.(*shared.DataResponse)
	if !ok {
		return kerr.New("TIBHFJGNNP", "%T is not a *shared.DataResponse", response)
	}

	n, err := editor.UnmarshalNode(s.tree.ctx, gr.Data)
	if err != nil {
		return kerr.Wrap("RVVWUVBKTN", err)
	}
	s.Node = n

	ed := s.Node.Editor()

	if err := ed.Initialize(s.tree.ctx, s, editor.Page, s.tree.Fail); err != nil {
		return kerr.Wrap("WXAHRIGKHI", err)
	}

	s.ListenForEditorChanges(ed.Listen().Ch)

	if err := addEntryChildren(s.Node, s, ed); err != nil {
		return kerr.Wrap("IQMWLPORUF", err)
	}

	return nil
}
