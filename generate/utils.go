package generate

import (
	"fmt"
	"go/ast"
	"go/types"
)

func (p *progDef) pathFromSelector(spec *ast.SelectorExpr) string {
	x, ok := spec.X.(*ast.Ident)
	if !ok {
		panic("spec.X must be *ast.Ident")
	}
	if x.Obj != nil {
		panic("x.Obj must be nil")
	}
	u, ok := p.info.Uses[x]
	if !ok {
		panic("x not found in uses")
	}
	pn, ok := u.(*types.PkgName)
	if !ok {
		panic("u not *ast.PkgName")
	}
	return pn.Imported().Path()
}

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
