// package jast is a collection of utilities for converting AST to jennifer.
// Note that jast is experimental and the API will change.
package jast

func New(imports map[string]string) *Cache {
	return &Cache{
		imports: imports,
	}
}

type Cache struct {
	imports map[string]string
}
