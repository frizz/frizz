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
		models.AddEditorsRecursively(s.editors, s.app.Package.Node())
		for _, n := range s.app.Types.All() {
			models.AddEditorsRecursively(s.editors, n)
		}
		s.Notify(nil, EditorInitialStateLoaded)
	case *actions.LoadSourceSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			return true
		}
		n := ni.GetNode()
		e := models.AddEditorsRecursively(s.editors, n)
		s.Notify(e, EditorLoaded)
	}
	return true
}
