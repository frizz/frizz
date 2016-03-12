package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor"
	"kego.io/editor/client/flux"
	"kego.io/system/node"
)

type EditorStore struct {
	*flux.Store
	ctx context.Context
	app *App

	editors map[*node.Node]editor.EditorInterface
}

func NewEditorStore(ctx context.Context) *EditorStore {
	return &EditorStore{
		Store:   &flux.Store{},
		ctx:     ctx,
		app:     FromContext(ctx),
		editors: map[*node.Node]editor.EditorInterface{},
	}
}

func (n *EditorStore) Get(node *node.Node) editor.EditorInterface {
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
