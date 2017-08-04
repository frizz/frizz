package frizz

import (
	"path/filepath"
	"strings"

	"frizz.io/global"
	"frizz.io/pack"
	"frizz.io/utils"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func New(local global.Package, more ...global.Package) global.Context {
	all := map[string]global.Package{}
	local.Add(all)
	for _, p := range more {
		p.Add(all)
	}
	return pack.NewContext(local.Path(), all)
}

func Package(p global.Package) (map[string]interface{}, error) {
	c := New(p)
	out := make(map[string]interface{})
	env := vos.Os()
	dir, err := patsy.Dir(env, p.Path())
	if err != nil {
		return nil, errors.Wrap(err, "getting dir from path")
	}
	files, err := filepath.Glob(filepath.Join(dir, "*.frizz.json"))
	if err != nil {
		return nil, errors.Wrap(err, "getting files")
	}
	for _, fpath := range files {
		_, fname := filepath.Split(fpath)
		name := strings.TrimSuffix(fname, ".frizz.json")

		v, err := utils.DecodeFile(fpath)
		if err != nil {
			return nil, errors.Wrapf(err, "decoding %s", fname)
		}

		r := pack.NewRoot()
		if err := r.Parse(v); err != nil {
			return nil, errors.Wrapf(err, "parsing imports from %s", fname)
		}
		s := global.NewStack(name)
		i, null, err := pack.UnpackInterface(c, r, s, v)
		if err != nil {
			return nil, errors.Wrapf(err, "unpacking %s", fname)
		}
		if !null {
			out[name] = i
		}
	}
	return out, nil
}
