package generate

import (
	"go/token"

	"frizz.io/generate/jast"

	"go/types"

	"frizz.io/generate/scanner"
	"github.com/dave/patsy"
)

type progDef struct {
	fset    *token.FileSet
	path    string
	pcache  *patsy.Cache
	dir     string
	types   map[string]*scanner.TypeDef // Type name: type def
	stubs   map[string]*stub            // File name: stub
	imports map[string]bool
	jast    *jast.Cache
	info    types.Info
}
