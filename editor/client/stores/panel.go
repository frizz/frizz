package stores

import (
	"context"

	"frizz.io/editor/client/models"
	"frizz.io/flux"
)

type PanelStore struct {
	*flux.Store
	ctx context.Context
	app *App

	panels map[*models.BranchModel]*models.PanelModel
}

type panelNotif string

func (b panelNotif) IsNotif() {}

const (
	PanelChanged panelNotif = "PanelChanged"
)

func NewPanelStore(ctx context.Context) *PanelStore {
	s := &PanelStore{
		Store:  &flux.Store{},
		ctx:    ctx,
		app:    FromContext(ctx),
		panels: map[*models.BranchModel]*models.PanelModel{},
	}
	s.Init(s)
	return s
}

func (s *PanelStore) Selected() *models.PanelModel {
	if s.app.Branches.Selected() == nil {
		return nil
	}
	return s.panels[s.app.Branches.Selected()]
}

func (s *PanelStore) Handle(payload *flux.Payload) bool {
	/*
		switch action := payload.Action.(type) {
		case *actions.SelectNode:
			if s.selected == action.Node {
				break
			}
			s.selected = action.Node
			payload.Notify()
		}
	*/
	return true
}
