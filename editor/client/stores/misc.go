package stores

import (
	"context"

	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
)

type MiscStore struct {
	*flux.Store
	ctx context.Context
	app *App

	addPopup *models.AddPopupModel
	info     bool
}

func (s *MiscStore) AddPopup() *models.AddPopupModel {
	return s.addPopup
}

func (s *MiscStore) Info() bool {
	return s.info
}

type miscNotif string

func (b miscNotif) IsNotif() {}

const (
	AddPopupChange  miscNotif = "AddPopupChange"
	InfoStateChange miscNotif = "InfoStateChange"
)

func NewMiscStore(ctx context.Context) *MiscStore {
	s := &MiscStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *MiscStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.ToggleInfoState:
		s.info = !s.info
		payload.Notify(nil, InfoStateChange)
	case *actions.OpenAddPopup:
		s.addPopup = &models.AddPopupModel{
			Visible: true,
			Parent:  action.Parent,
			Node:    action.Node,
			Types:   action.Types,
		}
		payload.Notify(nil, AddPopupChange)
	case *actions.CloseAddPopup:
		s.addPopup = &models.AddPopupModel{
			Visible: false,
		}
		payload.Notify(nil, AddPopupChange)
	}
	return true
}
