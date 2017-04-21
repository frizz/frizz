package stores

import (
	"sort"

	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/models"
	"frizz.io/flux"
	"frizz.io/system/node"
	"github.com/dave/kerr"
)

type TypeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	types map[string]*models.TypeModel
}

type typeNotif string

func (b typeNotif) IsNotif() {}

const (
	TypeChanged typeNotif = "TypeChanged"
)

func NewTypeStore(ctx context.Context) *TypeStore {
	s := &TypeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
		types: map[string]*models.TypeModel{},
	}
	s.Init(s)
	return s
}

func (s *TypeStore) Names() []string {
	var names []string
	for n, _ := range s.types {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func (s *TypeStore) Get(name string) *models.TypeModel {
	ti, ok := s.types[name]
	if !ok {
		return nil
	}
	return ti
}

func (s *TypeStore) All() map[string]*models.TypeModel {
	return s.types
}

func (s *TypeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		types := action.Info.Imports[action.Info.Path].Types
		if len(types) == 0 {
			break
		}
		for name, ti := range types {
			typ, err := node.Unmarshal(s.ctx, ti.Bytes)
			if err != nil {
				s.app.Fail <- kerr.Wrap("QDOVEIKABS", err)
			}
			s.types[name] = &models.TypeModel{
				Node: typ,
				File: ti.File,
			}

		}
		payload.Notify(nil, TypeChanged)
	}
	return true
}
