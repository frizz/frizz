package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
)

type MiscStore struct {
	*flux.Store
	ctx context.Context
	app *App

	addPopup *models.AddPopupModel
}

func (s *MiscStore) AddPopup() *models.AddPopupModel {
	return s.addPopup
}

type miscNotif string

func (b miscNotif) IsNotif() {}

const (
	AddPopupChange miscNotif = "AddPopupChange"
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
	case *actions.OpenAddPopup:
		s.addPopup = &models.AddPopupModel{
			Visible: true,
			Parent:  action.Parent,
			Node:    action.Node,
			Types:   action.Types,
		}
		s.app.Notify(nil, AddPopupChange)
	case *actions.CloseAddPopup:
		s.addPopup = &models.AddPopupModel{
			Visible: false,
		}
		s.app.Notify(nil, AddPopupChange)
	}
	return true
}
