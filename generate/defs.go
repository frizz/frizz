package generate

import (
	"go/ast"
	"go/token"

	"frizz.io/generate/jast"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"golang.org/x/tools/go/loader"
)

type scanDef struct {
	fset     *token.FileSet
	env      vos.Env
	input    []struct{ path, dir string }
	pcache   *patsy.Cache
	loaded   *loader.Program
	Packages map[string]*packageDef // Package path: package def
}

type packageDef struct {
	scan    *scanDef
	Path    string
	Dir     string
	types   map[string]*typeDef // Type name: type def
	stubs   map[string]*stub    // File name: stub
	imports map[string]bool     // Package path: true
	jast    *jast.Cache
	info    *loader.PackageInfo
}

type typeDef struct {
	packable   bool
	spec       ast.Expr
	name       string
	annotation string
}
