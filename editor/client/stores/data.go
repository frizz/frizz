package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type DataStore struct {
	*flux.Store
	ctx context.Context
	app *App

	data map[string]*models.DataModel
}

type dataKey string

const DataChanged dataKey = "DataChanged"

func NewDataStore(ctx context.Context) *DataStore {
	s := &DataStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
		data:  map[string]*models.DataModel{},
	}
	s.Init(s)
	return s
}

func (s *DataStore) Names() []string {
	var names []string
	for n, _ := range s.data {
		names = append(names, n)
	}
	return names
}

func (s *DataStore) Get(name string) (node *node.Node, file string, ok bool) {
	info, ok := s.data[name]
	if !ok {
		return nil, "", false
	}
	return info.Node, info.File, true
}

func (s *DataStore) All() map[string]*models.DataModel {
	return s.data
}

func (s *DataStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		if len(action.Info.Data) == 0 {
			break
		}
		for name, filename := range action.Info.Data {
			s.data[name] = &models.DataModel{Name: name, File: filename}
		}
		s.Notify(DataChanged)
	}
	return true
}
