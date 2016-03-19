package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/system/node"
)

type SelectedStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node
}

func NewSelectedStore(ctx context.Context) *SelectedStore {
	return &SelectedStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (s *SelectedStore) Get() *node.Node {
	return s.selected
}

func (s *SelectedStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.SelectNode:
		s.selected = action.Node
		s.Notify()
	}
	return true
}
