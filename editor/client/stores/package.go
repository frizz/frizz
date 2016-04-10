package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/kerr"
	"kego.io/system/node"
)

type PackageStore struct {
	*flux.Store
	ctx context.Context
	app *App

	packageNode *node.Node
}

func NewPackageStore(ctx context.Context) *PackageStore {
	s := &PackageStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *PackageStore) Get() *node.Node {
	return s.packageNode
}

func (s *PackageStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		pkg, err := node.Unmarshal(s.ctx, action.Info.Package)
		if err != nil {
			s.app.Fail <- kerr.Wrap("KXIKEWOKJI", err)
		}
		s.packageNode = pkg
		s.NotifyAll()
	}
	return true
}
