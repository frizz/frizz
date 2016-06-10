package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system/node"
)

type EditorStore struct {
	*flux.Store
	ctx context.Context
	app *App

	editors map[*node.Node]*models.EditorModel
}

type editorNotif string

func (b editorNotif) IsNotif() {}

const (
	EditorChanged            editorNotif = "EditorChanged"
	EditorLoaded             editorNotif = "EditorLoaded"
	EditorAdded              editorNotif = "EditorAdded"
	EditorInitialStateLoaded editorNotif = "EditorInitialStateLoaded"
)

func NewEditorStore(ctx context.Context) *EditorStore {
	s := &EditorStore{
		Store:   &flux.Store{},
		ctx:     ctx,
		app:     FromContext(ctx),
		editors: map[*node.Node]*models.EditorModel{},
	}
	s.Init(s)
	return s
}

func (s *EditorStore) Get(node *node.Node) *models.EditorModel {
	return s.editors[node]
}

func (s *EditorStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types, s.app.Data)
		s.AddEditorsRecursively(s.app.Package.Node())
		for _, n := range s.app.Types.All() {
			s.AddEditorsRecursively(n)
		}
		s.Notify(nil, EditorInitialStateLoaded)
	case *actions.LoadSourceSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			break
		}
		n := ni.GetNode()
		e := s.AddEditorsRecursively(n)
		s.Notify(e, EditorLoaded)
	case *actions.InitializeNode:
		payload.Wait(s.app.Branches)
		e := s.AddEditorsRecursively(action.Node)
		s.Notify(e, EditorAdded)
	case *actions.BranchSelected:
		payload.Wait(s.app.Nodes)
		if s.app.Nodes.Selected() == nil {
			break
		}
		e := s.Get(s.app.Nodes.Selected())
		if e == nil {
			break
		}
		s.Notify(e, EditorChanged)
	}
	return true
}

func (s *EditorStore) AddEditorsRecursively(n *node.Node) *models.EditorModel {
	e := models.NewEditor(n)
	s.editors[n] = e
	for _, c := range n.Map {
		s.AddEditorsRecursively(c)
	}
	for _, c := range n.Array {
		s.AddEditorsRecursively(c)
	}
	return e
}
