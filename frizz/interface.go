package frizz

import (
	"go/ast"
	"go/parser"
	"reflect"

	"github.com/pkg/errors"
)

func (r *Root) UnpackInterface(stack Stack, in interface{}) (value interface{}, null bool, err error) {
	if in == nil {
		return nil, true, nil
	}
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, value should be a map", stack)
	}
	ti, ok := m["_type"]
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, _type field missing", stack)
	}
	ts, ok := ti.(string)
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, _type should be a string", stack)
	}

	v := m["_value"]

	if r.Path == "" {
		return nil, false, errors.Errorf("%s: unpacking into interface, local path is not set", stack)
	}

	path, name, err := ParseReference(ts, r.Path, r.Imports)
	if err != nil {
		return nil, false, errors.Wrapf(err, "%s: unpacking into interface", stack)
	} else if path != "" && name != "" {
		// we have a path and a name => unpack with packer
		p, ok := r.Packers[path]
		if !ok {
			return nil, false, errors.Errorf("%s: unpacking into interface, packer for %s not registered", stack, r.Path)
		}
		if v != nil {
			return p.Unpack(r, stack, v, name)
		}
		return p.Unpack(r, stack, in, name)
	} else if name != "" {
		// no path path but we have a name => native type
		if v == nil {
			return nil, false, errors.Errorf("%s: unpacking native type into interface, _value field missing", stack)
		}
		return UnpackNative(stack, name, v)
	} else {
		// no path or name => type not found
		return nil, false, errors.Errorf("%s: unsupported type %s", stack, ts)
	}
}

func (r *Root) RepackInterface(stack Stack, root bool, in interface{}) (value interface{}, dict bool, null bool, err error) {
	if in == nil {
		return nil, false, true, nil
	}
	t := reflect.TypeOf(in)
	if t.PkgPath() == "" {
		switch t.Name() {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			value, dict, null, err = RepackNative(stack, t.Name(), in)
			if err != nil {
				return nil, false, false, err
			}
		default:
			return nil, false, false, errors.Errorf("%s: unsupported type %s", stack, t.Name())
		}
	} else {
		p, ok := r.Packers[t.PkgPath()]
		if !ok {
			return nil, false, false, errors.Errorf("%s: can't find packer for %s", stack, t.PkgPath())
		}
		value, dict, null, err = p.Repack(r, stack, in, t.Name())
		if err != nil {
			return nil, false, false, err
		}
	}

	typeName := t.Name()
	if t.PkgPath() != "" && t.PkgPath() != r.Path {
		alias := r.RegisterPackage(t.PkgPath())
		typeName = alias + "." + t.Name()
	}

	out := map[string]interface{}{}

	vmap, isMap := value.(map[string]interface{})
	isStruct := isMap && !dict

	if isStruct {
		// if value is a struct, use the value as out
		out = vmap
	} else if !null {
		// if value is not a struct, add the value to _value only if not null
		out["_value"] = value
	}

	// always add type information
	out["_type"] = typeName

	// when repacking into a root, add imports to _import
	if root {
		imports := make(map[string]interface{}, len(r.Imports))
		for a, p := range r.Imports {
			imports[a] = p
		}
		out["_import"] = imports
	}

	return out, false, false, nil
}

func ParseReference(value string, local string, imports map[string]string) (path, name string, err error) {
	expr, err := parser.ParseExpr(value)
	if err != nil {
		return "", "", errors.Wrapf(err, "parsing reference %s", value)
	}
	switch expr := expr.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			return "", expr.Name, nil
		default:
			return local, expr.Name, nil
		}
	case *ast.SelectorExpr:
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			return "", "", errors.Errorf("parsing reference %s, SelectorExpr.X should be *ast.Ident", value)
		}
		path, ok := imports[x.Name]
		if !ok {
			return "", "", errors.Errorf("parsing reference %s, can't find %s in imports", value, x.Name)
		}
		return path, expr.Sel.Name, nil
	}
	return "", "", nil
}
