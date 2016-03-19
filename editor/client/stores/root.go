package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system/node"
)

type RootStore struct {
	*flux.Store
	ctx context.Context
	app *App

	root *node.Node
}

func NewRootStore(ctx context.Context) *RootStore {
	return &RootStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (s *RootStore) Get() *node.Node {
	return s.root
}

func (s *RootStore) Handle(payload *flux.Payload) bool {
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
		err := ke.UnmarshalUntyped(s.ctx, []byte(action.Info.Package), pkg)
		if err != nil {
			s.app.Fail <- kerr.Wrap("KXIKEWOKJI", err)
		}
		s.root = pkg
		s.Notify()
	}
	return true
}
