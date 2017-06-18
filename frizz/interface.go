package frizz

import (
	"go/ast"
	"go/parser"

	"github.com/pkg/errors"
)

func (r *Root) UnpackInterface(stack Stack, in interface{}) (interface{}, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("%s: unpacking into interface, value should be a map", stack)
	}
	ti, ok := m["_type"]
	if !ok {
		return nil, errors.Errorf("%s: unpacking into interface, _type field missing", stack)
	}
	ts, ok := ti.(string)
	if !ok {
		return nil, errors.Errorf("%s: unpacking into interface, _type should be a string", stack)
	}

	v := m["_value"]

	if im, ok := m["_import"]; ok {
		imp, ok := im.(map[string]interface{})
		if !ok {
			return nil, errors.Errorf("%s: unpacking into interface, _import should be a map", stack)
		}
		for alias, path := range imp {
			path, ok := path.(string)
			if !ok {
				return nil, errors.Errorf("%s: unpacking into interface, _import values should be strings", stack)
			}
			r.Imports[alias] = path
		}
	}

	if r.Path == "" {
		return nil, errors.Errorf("%s: unpacking into interface, local path is not set", stack)
	}

	expr, err := parser.ParseExpr(ts)
	if err != nil {
		return nil, errors.Wrapf(err, "%s: parsing type string", stack)
	}
	switch expr := expr.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			if v == nil {
				return nil, errors.Errorf("%s: unpacking native type into interface, _value field missing", stack)
			}
			return UnpackNative(stack, expr.Name, v)
		default:
			p, ok := r.Packers[r.Path]
			if !ok {
				return nil, errors.Errorf("%s: unpacking into interface, packer for %s not registered", stack, r.Path)
			}
			if v != nil {
				return p.Unpack(r, stack, v, expr.Name)
			}
			return p.Unpack(r, stack, in, expr.Name)
		}
	case *ast.SelectorExpr:
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			return nil, errors.Errorf("%s: unpacking into interface, SelectorExpr.X should be *ast.Ident", stack)
		}
		path, ok := r.Imports[x.Name]
		if !ok {
			return nil, errors.Errorf("%s: unpacking into interface, can't find %s in imports", stack, x.Name)
		}
		p, ok := r.Packers[path]
		if !ok {
			return nil, errors.Errorf("%s: unpacking into interface, packer for %s not registered", stack, path)
		}
		if v != nil {
			return p.Unpack(r, stack, v, expr.Sel.Name)
		}
		return p.Unpack(r, stack, in, expr.Sel.Name)
	}
	return nil, errors.Errorf("%s: unsupported type %s", stack, ts)
}
