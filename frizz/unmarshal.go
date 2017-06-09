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

func Package(path string) (map[string]interface{}, error) {
	u := New(path)
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
			Unpacker: u,
			Imports:  make(map[string]string),
		}
		s := Stack{RootItem(name)}
		i, err := r.UnpackInterface(s, v)
		if err != nil {
			return nil, errors.Wrapf(err, "unpacking %s", fname)
		}
		out[name] = i
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

func (u *Unpacker) Unmarshal(in []byte) (interface{}, error) {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer(in))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}
	return u.Unpack(v)
}

func (u *Unpacker) Unpack(in interface{}) (interface{}, error) {
	r := &Root{
		Unpacker: u,
		Imports:  make(map[string]string),
	}
	s := Stack{RootItem("root")}
	return r.UnpackInterface(s, in)
}
