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
	EditorSelected           editorNotif = "EditorSelected"
	EditorLoaded             editorNotif = "EditorLoaded"
	EditorChildAdded         editorNotif = "EditorChildAdded"
	EditorChildDeleted       editorNotif = "EditorChildDeleted"
	EditorInitialStateLoaded editorNotif = "EditorInitialStateLoaded"
	EditorArrayOrderChanged  editorNotif = "EditorArrayOrderChanged"
	EditorValueChanged       editorNotif = "EditorValueChanged"
	EditorFocus              editorNotif = "EditorFocus"
	EditorErrorsChanged      editorNotif = "EditorErrorsChanged"
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
			if e := mutateAddEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e.Node, EditorChildAdded)
			}
		case actions.Undo:
			if e := mutateDeleteEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e.Node, EditorChildDeleted)
			}
		}

	case *actions.Delete:
		payload.Wait(s.app.Branches)
		switch action.Direction() {
		case actions.New, actions.Redo:
			if e := mutateDeleteEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e.Node, EditorChildDeleted)
			}
		case actions.Undo:
			if e := mutateAddEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e.Node, EditorChildAdded)
			}
		}

	case *actions.Reorder:
		payload.Wait(s.app.Branches)
		s.app.Notify(action.Model.Node, EditorArrayOrderChanged)

	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types, s.app.Data)
		s.AddEditorsRecursively(s.app.Package.Node())
		for _, n := range s.app.Types.All() {
			s.AddEditorsRecursively(n)
		}
		s.app.Notify(nil, EditorInitialStateLoaded)
	case *actions.LoadSourceSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			break
		}
		n := ni.GetNode()
		e := s.AddEditorsRecursively(n)
		s.app.Notify(e.Node, EditorLoaded)

	case *actions.BranchSelected:
		payload.Wait(s.app.Nodes)
		if e := s.Get(s.app.Nodes.Selected()); e != nil {
			s.app.Notify(e.Node, EditorSelected)
		}
		s.app.Notify(nil, EditorSelected)

	case *actions.EditorFocus:
		s.app.Notify(action.Editor.Node, EditorFocus)
	}
	return true
}

func (s *EditorStore) AddEditorsRecursively(n *node.Node) *models.EditorModel {
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
	return e
}

func mutateAddEditor(s *EditorStore, n *node.Node, p *node.Node) *models.EditorModel {
	s.AddEditorsRecursively(n)
	return s.Get(p)
}

func mutateDeleteEditor(s *EditorStore, n *node.Node, p *node.Node) *models.EditorModel {
	if e, ok := s.editors[n]; ok && p.Type.IsNativeCollection() {
		e.Deleted = true
	}
	return s.Get(p)
}
