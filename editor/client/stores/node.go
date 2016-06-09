package stores

import (
	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node

	addPop *models.AddPopModel
}

func (b *NodeStore) Selected() *node.Node {
	return b.selected
}

func (b *NodeStore) AddPop() *models.AddPopModel {
	return b.addPop
}

type nodeNotif string

func (b nodeNotif) IsNotif() {}

const (
	NodeInitialised nodeNotif = "NodeInitialised"
	AddPopChange    nodeNotif = "AddPopChange"
)

func NewNodeStore(ctx context.Context) *NodeStore {
	s := &NodeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.BranchSelect:
		s.selected = nil
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		}
	case *actions.BranchSelectPostLoad:
		s.selected = nil
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		}
	case *actions.AddNodeClick:

		types := action.Node.Rule.PermittedTypes()

		if len(types) == 1 {
			// if only one type is compatible, don't show the popup, just add it.
			if err := action.Node.InitialiseWithConcreteType(s.ctx, types[0]); err != nil {
				s.app.Fail <- kerr.Wrap("NEWJOIFGBH", err)
			}
			return true
		}

		s.addPop = &models.AddPopModel{
			Visible: true,
			Parent:  action.Node.Parent,
			Node:    action.Node,
			Types:   types,
		}
		s.Notify(nil, AddPopChange)

	case *actions.AddCollectionItemClick:

		rw, err := action.Parent.Rule.ItemsRule()
		if err != nil {
			s.app.Fail <- kerr.Wrap("EWYOMNAQMU", err)
		}

		s.addPop = &models.AddPopModel{
			Visible: true,
			Parent:  action.Parent,
			Types:   rw.PermittedTypes(),
		}
		s.Notify(nil, AddPopChange)

		return true
	case *actions.AddPopCloseClick:

		s.addPop = &models.AddPopModel{
			Visible: false,
		}
		s.Notify(nil, AddPopChange)

		return true

	}
	return true
}
