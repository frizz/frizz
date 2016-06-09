package stores

import (
	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node

	addModalVisible bool
	addModalParent  *node.Node
	addModelNode    *node.Node
	addModalTypes   []*system.Type
}

func (b *NodeStore) Selected() *node.Node {
	return b.selected
}

func (b *NodeStore) AddModalVisible() bool {
	return b.addModalVisible
}

func (b *NodeStore) AddModalParent() *node.Node {
	return b.addModalParent
}

func (b *NodeStore) AddModalNode() *node.Node {
	return b.addModelNode
}

func (b *NodeStore) AddModalTypes() []*system.Type {
	return b.addModalTypes
}

type nodeNotif string

func (b nodeNotif) IsNotif() {}

const (
	NodeInitialised nodeNotif = "NodeInitialised"
	AddModalChange  nodeNotif = "AddModalChange"
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

		s.addModalVisible = true
		s.addModalParent = action.Node.Parent
		s.addModelNode = action.Node
		s.addModalTypes = types
		s.Notify(nil, AddModalChange)

	case *actions.AddCollectionItemClick:

		rw, err := action.Parent.Rule.ItemsRule()
		if err != nil {
			s.app.Fail <- kerr.Wrap("EWYOMNAQMU", err)
		}

		s.addModalVisible = true
		s.addModalParent = action.Parent
		s.addModelNode = nil
		s.addModalTypes = rw.PermittedTypes()
		s.Notify(nil, AddModalChange)

		return true
	case *actions.AddModalCloseClick:

		s.addModalVisible = false
		s.addModalParent = nil
		s.addModelNode = nil
		s.addModalTypes = nil
		s.Notify(nil, AddModalChange)

		return true

	}
	return true
}
