package stores

import (
	"context"

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
	if node == nil {
		return nil
	}
	e, ok := s.editors[node]
	if !ok {
		return nil
	}
	if e.Deleted {
		return nil
	}
	return e
}

func (s *EditorStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.Add:
		payload.Wait(s.app.Branches)
		switch action.Direction() {
		case actions.New, actions.Redo:
			mutateAddEditor(s, action.Node, action.Parent)
		case actions.Undo:
			mutateDeleteEditor(s, action.Node, action.Parent)
		}

	case *actions.Delete:
		payload.Wait(s.app.Branches)
		switch action.Direction() {
		case actions.New, actions.Redo:
			mutateDeleteEditor(s, action.Node, action.Parent)
		case actions.Undo:
			mutateAddEditor(s, action.Node, action.Parent)
		}

	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types, s.app.Data)
		s.AddEditorsRecursively(s.app.Package.Node())
		for _, ti := range s.app.Types.All() {
			s.AddEditorsRecursively(ti.Node)
		}
	case *actions.LoadFileSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			break
		}
		n := ni.GetNode()
		s.AddEditorsRecursively(n)
	}
	return true
}

func (s *EditorStore) AddEditorsRecursively(n *node.Node) {
	var e *models.EditorModel
	var ok bool
	if e, ok = s.editors[n]; !ok {
		e = models.NewEditor(n)
		s.editors[n] = e
	}
	e.Deleted = false
	for _, c := range n.Map {
		s.AddEditorsRecursively(c)
	}
	for _, c := range n.Array {
		s.AddEditorsRecursively(c)
	}
}

func mutateAddEditor(s *EditorStore, n *node.Node, p *node.Node) {
	s.AddEditorsRecursively(n)
}

func mutateDeleteEditor(s *EditorStore, n *node.Node, p *node.Node) {
	if e, ok := s.editors[n]; ok && p != nil && p.Type.IsNativeCollection() {
		e.Deleted = true
	}
}
