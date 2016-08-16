package stores

import (
	"context"

	"kego.io/editor/client/actions"
	"kego.io/flux"
)

type EnvStore struct {
	*flux.Store
	ctx context.Context
	app *App

	path string
	hash []byte
}

func NewEnvStore(ctx context.Context) *EnvStore {
	s := &EnvStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *EnvStore) Path() string {
	return s.path
}

func (s *EnvStore) Hash() []byte {
	return s.hash
}

func (s *EnvStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		s.hash = action.Info.Hash
		s.path = action.Info.Path
	}
	return true
}
