package stores

import (
	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/flux"
	"frizz.io/system/node"
	"github.com/dave/kerr"
)

type PackageStore struct {
	*flux.Store
	ctx context.Context
	app *App

	node     *node.Node
	filename string
}

type packageNotif string

func (b packageNotif) IsNotif() {}

const (
	PackageChanged packageNotif = "PackageChanged"
)

func NewPackageStore(ctx context.Context) *PackageStore {
	s := &PackageStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *PackageStore) Node() *node.Node {
	return s.node
}

func (s *PackageStore) Filename() string {
	return s.filename
}

func (s *PackageStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		pkgNode, err := node.Unmarshal(s.ctx, action.Info.Package)
		if err != nil {
			s.app.Fail <- kerr.Wrap("KXIKEWOKJI", err)
		}
		s.node = pkgNode
		s.filename = action.Info.PackageFilename
		payload.Notify(nil, PackageChanged)
	}
	return true
}
