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

	addPop *models.AddPopModel
}

func (s *NodeStore) Selected() *node.Node {
	return s.selected
}

func (s *NodeStore) AddPop() *models.AddPopModel {
	return s.addPop
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
			s.app.Dispatch(&actions.NodeInitialized{Node: action.Node})
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

		types := rw.PermittedTypes()

		if len(types) == 1 && action.Parent.Type.IsNativeArray() {
			// if only one type is compatible and adding to an array, don't show the popup, just
			// add it.
			n := node.NewNode()
			n.Parent = action.Parent
			n.Index = len(n.Parent.Array)
			n.Parent.Array = append(n.Parent.Array, n)
			if err := n.InitialiseWithConcreteType(s.ctx, types[0]); err != nil {
				s.app.Fail <- kerr.Wrap("JRFSDHXIRY", err)
			}
			s.app.Dispatch(&actions.NodeInitialized{Node: n})
			return true
		}

		s.addPop = &models.AddPopModel{
			Visible: true,
			Parent:  action.Parent,
			Types:   types,
		}
		s.Notify(nil, AddPopChange)

		return true
	case *actions.AddPopClose:

		s.addPop = &models.AddPopModel{
			Visible: false,
		}
		s.Notify(nil, AddPopChange)

		return true
	case *actions.AddPopNameChange:
		s.addPop.Name = action.Value
		return true
	case *actions.AddPopTypeChange:
		if action.Value == "" {
			s.addPop.Type = nil
		}
		r, err := system.NewReferenceFromString(s.ctx, action.Value)
		if err != nil {
			s.app.Fail <- kerr.Wrap("DQUOTUHOJR", err)
		}
		t, ok := system.GetTypeFromCache(s.ctx, r.Package, r.Name)
		if !ok {
			s.app.Fail <- kerr.New("HRHFMOPYGO", "Type %s not found in cache", r.Value())
		}
		s.addPop.Type = t
		return true
	case *actions.AddPopSaveClick:
		var t *system.Type
		if len(s.addPop.Types) == 1 {
			t = s.addPop.Types[0]
		} else if s.addPop.Type != nil {
			t = s.addPop.Type
		} else {
			return true
		}
		var n *node.Node

		if s.addPop.Node != nil {
			n = s.addPop.Node
		} else if s.addPop.Parent.Type.IsNativeMap() {
			n = node.NewNode()
			n.Parent = s.addPop.Parent
			n.Key = s.addPop.Name
			s.addPop.Parent.Map[s.addPop.Name] = n
		} else if s.addPop.Parent.Type.IsNativeArray() {
			n = node.NewNode()
			n.Parent = s.addPop.Parent
			n.Index = len(s.addPop.Parent.Array)
			s.addPop.Parent.Array = append(s.addPop.Parent.Array, n)
		}

		if err := n.InitialiseWithConcreteType(s.ctx, t); err != nil {
			s.app.Fail <- kerr.Wrap("NGCOVJKEXG", err)
		}

		s.app.Dispatch(&actions.NodeInitialized{Node: n})

		s.addPop = &models.AddPopModel{
			Visible: false,
		}
		s.Notify(nil, AddPopChange)

		return true
	}
	return true
}
