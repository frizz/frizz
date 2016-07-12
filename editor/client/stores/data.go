package stores

import (
	"sort"

	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
)

type DataStore struct {
	*flux.Store
	ctx context.Context
	app *App

	data map[string]*models.DataModel
}

type dataNotif string

func (b dataNotif) IsNotif() {}

const (
	DataChanged dataNotif = "DataChanged"
)

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
	sort.Strings(names)
	return names
}

func (s *DataStore) Get(name string) *models.DataModel {
	info, ok := s.data[name]
	if !ok {
		return nil
	}
	return info
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
		payload.Notify(nil, DataChanged)
	}
	return true
}
