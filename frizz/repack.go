package frizz

import (
	"reflect"

	"github.com/pkg/errors"
)

func RepackInterface(root *Root, stack Stack, in interface{}) (interface{}, bool, error) {
	t := reflect.TypeOf(in)
	p, ok := root.Packers[t.PkgPath()]
	if !ok {
		return nil, false, errors.Errorf("Can't find packer for %s", t.PkgPath())
	}
	v, dict, err := p.Repack(root, stack, in, t.Name())
	if err != nil {
		return nil, false, err
	}

	typeName := t.Name()
	if t.PkgPath() != root.Path {
		// find package alias
		alias := ""
		for a, p := range root.Imports {
			if p == t.PkgPath() {
				alias = a
				break
			}
		}
		if alias == "" {
			return nil, false, errors.Errorf("Can't find alias for %s", t.PkgPath())
		}
		typeName = alias + "." + t.Name()
	}

	if v, ok := v.(map[string]interface{}); ok && !dict {
		// if value is a map, and dict == false (e.g. value is a struct), add
		// type to struct and return
		v["_type"] = typeName
		return v, false, nil
	}

	out := map[string]interface{}{
		"_type":  typeName,
		"_value": v,
	}
	return out, false, nil
}
