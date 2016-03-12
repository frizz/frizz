package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	root *node.Node
}

func NewNodeStore(ctx context.Context) *NodeStore {
	return &NodeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (n *NodeStore) Get() *node.Node {
	return n.root
}

func (n *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		/*
			pkg := node.NewNode()
			c1 := node.NewNode()
			c2 := node.NewNode()
			c2.SetValueString(n.ctx, "qux")
			pkg.SetMapValue("foo", c1)
			c1.SetMapValue("bar", c2)
		*/
		pkg := node.NewNode()
		err := ke.UnmarshalUntyped(n.ctx, []byte(action.Info.Package), pkg)
		if err != nil {
			n.app.Fail <- kerr.Wrap("KXIKEWOKJI", err)
		}
		n.root = pkg
		n.Notify()
	}
	return true
}
