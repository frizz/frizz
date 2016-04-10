package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/kerr"
	"kego.io/system/node"
)

type TypeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	types map[string]*node.Node
}

func NewTypeStore(ctx context.Context) *TypeStore {
	s := &TypeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
		types: map[string]*node.Node{},
	}
	s.Init(s)
	return s
}

func (s *TypeStore) Names() []string {
	var names []string
	for n, _ := range s.types {
		names = append(names, n)
	}
	return names
}

func (s *TypeStore) Get(name string) (*node.Node, bool) {
	n, ok := s.types[name]
	return n, ok
}

func (s *TypeStore) All() map[string]*node.Node {
	return s.types
}

func (s *TypeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		types := action.Info.Imports[action.Info.Path].Types
		if len(types) == 0 {
			break
		}
		for name, bytes := range types {
			typ, err := node.Unmarshal(s.ctx, bytes)
			if err != nil {
				s.app.Fail <- kerr.Wrap("QDOVEIKABS", err)
			}
			s.types[name] = typ
		}
		s.NotifyAll()
	}
	return true
}
