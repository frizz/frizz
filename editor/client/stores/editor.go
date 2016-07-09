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
	return s.editors[node]
}

func (s *EditorStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.Add:
		payload.Wait(s.app.Branches)
		if action.Forward() {
			if e := mutateAddEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e, EditorChildAdded)
			}
		} else {
			if e := mutateDeleteEditor(s, action.Node, action.Parent); e != nil {
				s.app.Notify(e, EditorChildDeleted)
			}
		}

	case *actions.Delete:
		payload.Wait(s.app.Branches)
		if action.Forward() {
			if e := mutateDeleteEditor(s, action.Node, action.Node.Parent); e != nil {
				s.app.Notify(e, EditorChildDeleted)
			}
		} else {
			if e := mutateAddEditor(s, action.Node, action.Node.Parent); e != nil {
				s.app.Notify(e, EditorChildAdded)
			}
		}

	case *actions.Reorder:
		payload.Wait(s.app.Branches)
		s.app.Notify(action.Model, EditorArrayOrderChanged)

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
		s.app.Notify(e, EditorLoaded)

	case *actions.BranchSelected:
		payload.Wait(s.app.Nodes)
		if e := s.Get(s.app.Nodes.Selected()); e != nil {
			s.app.Notify(e, EditorSelected)
		}
		s.app.Notify(nil, EditorSelected)

	case *actions.EditorFocus:
		s.app.Notify(action.Editor, EditorFocus)
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

func mutateAddEditor(s *EditorStore, n *node.Node, p *node.Node) *models.EditorModel {
	s.AddEditorsRecursively(n)
	return s.Get(p)
}

func mutateDeleteEditor(s *EditorStore, n *node.Node, p *node.Node) *models.EditorModel {
	if p.Type.IsNativeCollection() && s.editors[n] != nil {
		delete(s.editors, n)
	}
	return s.Get(p)
}
