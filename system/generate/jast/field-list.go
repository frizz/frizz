package jast

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

// FieldList adds fields to a group - works for both Params() and Struct()
func (c *Cache) FieldList(g *jen.Group, fl *ast.FieldList) {
	if fl == nil {
		return
	}
	for _, v := range fl.List {
		g.Do(func(s *jen.Statement) {
			if v.Names != nil {
				s.ListFunc(func(g *jen.Group) {
					for _, n := range v.Names {
						g.Id(n.Name)
					}
				})
			}
			s.Add(c.Expr(v.Type))
		})
	}
}
