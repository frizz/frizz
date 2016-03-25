package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
)

type PanelStore struct {
	*flux.Store
	ctx context.Context
	app *App

	panels map[*models.BranchModel]*models.PanelModel
}

func NewPanelStore(ctx context.Context) *PanelStore {
	return &PanelStore{
		Store:  &flux.Store{},
		ctx:    ctx,
		app:    FromContext(ctx),
		panels: map[*models.BranchModel]*models.PanelModel{},
	}
}

func (s *PanelStore) Selected() *models.PanelModel {
	if s.app.Selected.Get() == nil {
		return nil
	}
	return s.panels[s.app.Selected.Get()]
}

func (s *PanelStore) Handle(payload *flux.Payload) bool {
	/*
		switch action := payload.Action.(type) {
		case *actions.SelectNode:
			if s.selected == action.Node {
				break
			}
			s.selected = action.Node
			s.Notify()
		}
	*/
	return true
}
