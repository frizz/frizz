// package jast is a collection of utilities for converting AST to jennifer.
// Note that jast is experimental and the API will change.
package jast

import (
	"go/ast"
	"go/types"
)

// New creates a new jast cache.
func New(imports map[string]string, uses map[*ast.Ident]types.Object) *Cache {
	if (imports == nil && uses == nil) || (imports != nil && uses != nil) {
		panic("must specify either imports or uses but not both")
	}
	return &Cache{
		imports: imports,
		uses:    uses,
	}
}

type Cache struct {
	imports map[string]string
	uses    map[*ast.Ident]types.Object
}
