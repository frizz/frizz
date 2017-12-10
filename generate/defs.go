package generate

import (
	"go/token"

	"frizz.io/generate/jast"

	"go/types"

	"frizz.io/generate/scanner"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

type progDef struct {
	fset    *token.FileSet
	env     vos.Env
	path    string
	pcache  *patsy.Cache
	dir     string
	types   map[string]*scanner.TypeDef // Type name: type def
	stubs   map[string]*stub            // File name: stub
	imports map[string]bool
	jast    *jast.Cache
	info    types.Info
}
