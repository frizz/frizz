package generate

import (
	"fmt"
	"go/ast"
)

func fieldName(field *ast.Field) string {
	if len(field.Names) > 0 {
		// named field
		return field.Names[0].Name
	}
	// anonymous field
	typ := field.Type
	for {
		// remove any number of stars
		star, ok := typ.(*ast.StarExpr)
		if !ok {
			break
		}
		typ = star.X
	}
	switch typ := typ.(type) {
	case *ast.Ident:
		return typ.Name
	case *ast.SelectorExpr:
		return typ.Sel.Name
	default:
		panic(fmt.Sprintf("can't find name for anonymous field %#v", typ))
	}
}
