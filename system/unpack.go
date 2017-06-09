package system

import (
	"context"
	"go/ast"
	"go/parser"

	"github.com/pkg/errors"
)

func UnpackInterface(ctx context.Context, in interface{}) (interface{}, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.New("unpacking into interface, value should be a map")
	}
	ti, ok := m["_type"]
	if !ok {
		return nil, errors.New("unpacking into interface, _type field missing")
	}
	ts, ok := ti.(string)
	if !ok {
		return nil, errors.New("unpacking into interface, _type should be a string")
	}

	v := m["_value"]

	uci := ctx.Value(UnpackContextKey)
	if uci == nil {
		return nil, errors.New("unpacking into interface, unpack context not found")
	}
	uc, ok := uci.(*UnpackContext)
	if !ok {
		// notest
		return nil, errors.New("unpacking into interface, unpack context should be *UnpackContext")
	}

	rci := ctx.Value(RegistryContextKey)
	if rci == nil {
		// notest
		return nil, errors.New("unpacking into interface, registry context not found")
	}
	rc, ok := rci.(*RegistryContext)
	if !ok {
		// notest
		return nil, errors.New("unpacking into interface, registry context should by *RegistryContext")
	}

	if im, ok := m["_import"]; ok {
		imp, ok := im.(map[string]interface{})
		if !ok {
			return nil, errors.New("unpacking into interface, _import should be a map")
		}
		for alias, path := range imp {
			path, ok := path.(string)
			if !ok {
				return nil, errors.New("unpacking into interface, _import values should be strings")
			}
			uc.Set(alias, path)
		}
	}

	if uc.Path == "" {
		return nil, errors.New("unpacking into interface, local path is not set")
	}

	expr, err := parser.ParseExpr(ts)
	if err != nil {
		return nil, errors.Wrap(err, "parsing type string")
	}
	switch expr := expr.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			if v == nil {
				return nil, errors.New("unpacking native type into interface, _value field missing")
			}
			return Convert(expr.Name, v)
		default:
			rt, ok := rc.Get(RegistryTypeKey{Path: uc.Path, Name: expr.Name})
			if !ok {
				return nil, errors.Errorf("unpacking into interface, can't find %s.%s in type registry", uc.Path, expr.Name)
			}
			if v != nil {
				return rt.Unpacker(ctx, v)
			}
			return rt.Unpacker(ctx, in)
		}
	case *ast.SelectorExpr:
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			return nil, errors.New("unpacking into interface, SelectorExpr.X should be *ast.Ident")
		}
		path, ok := uc.Get(x.Name)
		if !ok {
			return nil, errors.Errorf("unpacking into interface, can't find %s in imports", x.Name)
		}
		rt, ok := rc.Get(RegistryTypeKey{Path: path, Name: expr.Sel.Name})
		if !ok {
			return nil, errors.Errorf("unpacking into interface, can't find %s.%s in registry", path, expr.Sel.Name)
		}
		if v != nil {
			return rt.Unpacker(ctx, v)
		}
		return rt.Unpacker(ctx, in)
	case *ast.MapType:
		// TODO
		//if key, ok := expr.Key.(*ast.Ident); !ok || key.Name != "string" {
		//	return nil, errors.New("unpacking map, key must be string")
		//}
	case *ast.ArrayType:
		// TODO
	}
	return nil, nil
}
