package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
)

type SelectedStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *models.BranchModel
}

func NewSelectedStore(ctx context.Context) *SelectedStore {
	return &SelectedStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (s *SelectedStore) Get() *models.BranchModel {
	return s.selected
}

func (s *SelectedStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.SelectBranch:
		s.selected = action.Branch
		s.Notify()
	}
	return true
}
