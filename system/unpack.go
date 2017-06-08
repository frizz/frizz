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
	ti, ok := m["__type"]
	if !ok {
		return nil, errors.New("unpacking into interface, __type field missing")
	}
	ts, ok := ti.(string)
	if !ok {
		return nil, errors.New("unpacking into interface, __type should be a string")
	}

	v := m["__value"]

	uc := ctx.Value(UnpackContextKey).(*UnpackContext)
	if uc == nil {
		return nil, errors.New("unpacking into interface, unpack context not found")
	}

	rc := ctx.Value(RegistryContextKey).(*RegistryContext)
	if rc == nil {
		return nil, errors.New("unpacking into interface, registry context not found")
	}

	if uc.Path == "" {
		return nil, errors.New("unpacking into interface, local path is not set")
	}

	expr, err := parser.ParseExpr(ts)
	if err != nil {
		return nil, errors.New("parsing type string")
	}
	switch expr := expr.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			if v == nil {
				return nil, errors.New("unpacking into interface, __value field missing")
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
	case *ast.ArrayType:
		// TODO
	}
	return nil, nil
}
