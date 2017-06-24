package frizz

import (
	"bytes"
	"encoding/json"

	"path/filepath"

	"strings"

	"os"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func Package(path string, packers ...Packer) (map[string]interface{}, error) {
	u := New(path, packers...)
	out := make(map[string]interface{})
	env := vos.Os()
	dir, err := patsy.Dir(env, path)
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

		v, err := decodeFile(fpath)
		if err != nil {
			return nil, errors.Wrapf(err, "decoding %s", fname)
		}

		r := &Root{
			Context: u,
			Imports: make(map[string]string),
		}
		s := Stack{RootItem(name)}
		i, null, err := r.UnpackInterface(s, v)
		if err != nil {
			return nil, errors.Wrapf(err, "unpacking %s", fname)
		}
		if !null {
			out[name] = i
		}
	}
	return out, nil
}

func decodeFile(fpath string) (interface{}, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var v interface{}
	d := json.NewDecoder(f)
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (u *Context) Unmarshal(in []byte) (interface{}, error) {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer(in))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}
	v, null, err := u.Unpack(v)
	if err != nil {
		return nil, err
	}
	if null {
		return nil, nil
	}
	return v, nil
}

func (u *Context) Unpack(in interface{}) (value interface{}, null bool, err error) {
	r := &Root{
		Context: u,
		Imports: make(map[string]string),
	}
	s := Stack{RootItem("root")}
	return r.UnpackInterface(s, in)
}

func (u *Context) Repack(in interface{}) (value interface{}, dict bool, null bool, err error) {
	r := &Root{
		Context: u,
		Imports: make(map[string]string),
	}
	s := Stack{RootItem("root")}
	return r.RepackInterface(s, true, in)
}
