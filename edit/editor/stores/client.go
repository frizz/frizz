package stores

import (
	"context"

	"frizz.io/edit/common"
	"frizz.io/edit/editor/actions"
	"frizz.io/flux"
)

type ClientStore struct {
	*flux.Store
	ctx context.Context
	app *App

	client *common.Client
}

type clientNotif string

func (b clientNotif) IsNotif() {}

const (
	ClientChanged clientNotif = "ClientChanged"
)

func NewClientStore(ctx context.Context, app *App) *ClientStore {
	s := &ClientStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   app,
	}
	s.Init(s)
	return s
}

func (s *ClientStore) Client() *common.Client {
	return s.client
}

func (s *ClientStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		s.client = action.Client
		payload.Notify(nil, ClientChanged)
	}
	return true
}
