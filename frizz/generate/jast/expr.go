package jast

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

func (c *Cache) Expr(spec ast.Expr) *jen.Statement {
	switch spec := spec.(type) {
	case nil:
		return jen.Null()
	case *ast.StructType:
		return jen.StructFunc(func(g *jen.Group) {
			c.FieldList(g, spec.Fields)
		})
	case *ast.Ident:
		return jen.Id(spec.Name)
	case *ast.StarExpr:
		return jen.Op("*").Add(c.Expr(spec.X))
	case *ast.ArrayType:
		return jen.Index(c.Expr(spec.Len)).Add(c.Expr(spec.Elt))
	case *ast.FuncType:
		return jen.Func().ParamsFunc(func(g *jen.Group) {
			c.FieldList(g, spec.Params)
		}).ParamsFunc(func(g *jen.Group) {
			c.FieldList(g, spec.Results)
		})
	case *ast.InterfaceType:
		return jen.InterfaceFunc(func(g *jen.Group) {
			c.FieldList(g, spec.Methods)
		})
	case *ast.MapType:
		return jen.Map(c.Expr(spec.Key)).Add(c.Expr(spec.Value))
	case *ast.ChanType:
		return jen.Do(func(s *jen.Statement) {
			if spec.Dir == ast.SEND {
				s.Op("<-")
			}
		}).Chan().Do(func(s *jen.Statement) {
			if spec.Dir == ast.RECV {
				s.Op("<-")
			}
		}).Add(c.Expr(spec.Value))
	case *ast.ParenExpr:
		return jen.Parens(c.Expr(spec.X))
	case *ast.Ellipsis:
		return jen.Op("...").Add(c.Expr(spec.Elt))
	case *ast.BasicLit:
		return jen.Op(spec.Value)
	case *ast.SelectorExpr:
		if x, ok := spec.X.(*ast.Ident); ok && x.Obj == nil {
			if path, ok := c.imports[x.Name]; ok {
				return jen.Qual(path, spec.Sel.Name)
			}
		}
		return c.Expr(spec.X).Dot(spec.Sel.Name)
	case *ast.IndexExpr:
		return c.Expr(spec.X).Index(c.Expr(spec.Index))
	case *ast.SliceExpr:
		return c.Expr(spec.X).IndexFunc(func(g *jen.Group) {
			if spec.Low != nil {
				g.Add(c.Expr(spec.Low))
			} else {
				g.Empty()
			}
			if spec.High != nil {
				g.Add(c.Expr(spec.High))
			} else {
				g.Empty()
			}
			if spec.Max != nil {
				g.Add(c.Expr(spec.Max))
			} else {
				g.Null()
			}
		})
	case *ast.TypeAssertExpr:
		return c.Expr(spec.X).Assert(c.Expr(spec.Type))
	case *ast.CallExpr:
		return c.Expr(spec.Fun).CallFunc(func(g *jen.Group) {
			for _, arg := range spec.Args {
				g.Add(c.Expr(arg))
			}
		})
	case *ast.UnaryExpr:
		return jen.Op(spec.Op.String()).Add(c.Expr(spec.X))
	case *ast.BinaryExpr:
		return c.Expr(spec.X).Op(spec.Op.String()).Add(c.Expr(spec.Y))
	case *ast.KeyValueExpr:
		// kludge: would usually use Dict here but it's not a *Statement so
		// this will work.
		return c.Expr(spec.Key).Op(":").Add(c.Expr(spec.Value))
	case *ast.CompositeLit:
		return jen.ValuesFunc(func(g *jen.Group) {
			for _, element := range spec.Elts {
				g.Add(c.Expr(element))
			}
		})
		//case *ast.BadExpr:
		//case *ast.FuncLit:
	}
	return jen.Null()
}
