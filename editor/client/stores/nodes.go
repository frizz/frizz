package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/kerr"
)

type NodeStore struct {
	*flux.Store
	ctx  context.Context
	app  *App
	root *editor.Node
}

func NewNodeStore(ctx context.Context) *NodeStore {
	return &NodeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (n *NodeStore) Root() *editor.Node {
	return n.root
}

func (n *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		pkg, err := editor.UnmarshalNode(n.ctx, []byte(action.Info.Package))
		if err != nil {
			n.app.Fail <- kerr.Wrap("KXIKEWOKJI", err)
		}
		n.root = pkg
		n.Notify()
	}
	return true
}
