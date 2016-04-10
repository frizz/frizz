package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
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

func (n *EditorStore) Get(node *node.Node) *models.EditorModel {
	return n.editors[node]
}

func (n *EditorStore) Handle(payload *flux.Payload) bool {
	/*
		switch action := payload.Action.(type) {
		case *actions.InitialState:
		}
	*/
	return true
}
