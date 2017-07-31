package frizz

import (
	"github.com/pkg/errors"
)

func New(imports Importer) *Context {
	packers := map[string]Packer{}
	imports.Add(packers, nil)
	return &Context{
		Path:    imports.Path(),
		Packers: packers,
	}
}

type Context struct {
	Path    string
	Packers map[string]Packer
}

func (u *Context) Register(packers ...Packer) *Context {
	for _, p := range packers {
		u.Packers[p.Path()] = p
	}
	return u
}

type Root struct {
	*Context
	Imports map[string]string
}

func (r *Root) ParseImports(v interface{}) error {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	im, ok := m["_import"]
	if !ok {
		return nil
	}
	imp, ok := im.(map[string]interface{})
	if !ok {
		return errors.Errorf("parsing _import, should be a map, found %T", im)
	}
	for alias, pathi := range imp {
		path, ok := pathi.(string)
		if !ok {
			return errors.Errorf("parsing _import, values should be strings, found %T", pathi)
		}
		r.Imports[alias] = path
	}
	return nil
}

type Packable interface {
	Unpack(root *Root, stack Stack, in interface{}) (null bool, err error)
	Repack(root *Root, stack Stack) (value interface{}, dict bool, null bool, err error)
}

type Packer interface {
	Path() string
	Unpack(root *Root, stack Stack, in interface{}, name string) (value interface{}, null bool, err error)
	Repack(root *Root, stack Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error)
}

type Typer interface {
	Path() string
	Get(name string) string
}

type Importer interface {
	Path() string
	Add(map[string]Packer, map[string]Typer)
}
