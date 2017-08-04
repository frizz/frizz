package utils

import (
	"go/ast"
	"go/parser"

	"github.com/pkg/errors"
)

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
