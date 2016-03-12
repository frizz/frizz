package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type BranchStore struct {
	*flux.Store
	ctx context.Context
	app *App

	branches map[*node.Node]*models.BranchModel
}

func (b *BranchStore) All() map[*node.Node]*models.BranchModel {
	return b.branches
}

func NewBranchStore(ctx context.Context) *BranchStore {
	return &BranchStore{
		Store:    &flux.Store{},
		ctx:      ctx,
		app:      FromContext(ctx),
		branches: map[*node.Node]*models.BranchModel{},
	}
}

func (s *BranchStore) Get(node *node.Node) *models.BranchModel {
	return s.branches[node]
}

func (s *BranchStore) Handle(payload *flux.Payload) bool {

	switch action := payload.Action.(type) {
	case *actions.InitialState:
		payload.WaitFor(s.app.Nodes)
		root := s.app.Nodes.Get()
		nodes := root.Flatten(false)
		for _, n := range nodes {
			s.branches[n] = &models.BranchModel{
				Root: n == root,
				Open: true,
			}
		}
		s.Notify()
	case *actions.ToggleNode:
		b, ok := s.branches[action.Node]
		if !ok {
			// branch not found - ignore
			break
		}
		if !b.CanOpen() {
			// branch can't open - ignore
			break
		}
		b.Open = !b.Open
		s.Notify()
	}

	return true
}
