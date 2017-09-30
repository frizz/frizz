package generate

import (
	"go/ast"
	"go/token"
	"strconv"

	"frizz.io/generate/jast"

	"frizz.io/generate/scanner"
	"github.com/dave/patsy"
	"github.com/pkg/errors"
)

type progDef struct {
	fset    *token.FileSet
	scanner *scanner.Scanner
	path    string
	pcache  *patsy.Cache
	dir     string
	types   map[string]*typeDef // Type name: type def
	stubs   map[string]*stub    // File name: stub
	imports map[string]bool
}

type fileDef struct {
	*progDef
	imports map[string]string
	jast    *jast.Cache
}

type typeDef struct {
	*fileDef
	packable     bool
	spec         ast.Expr
	name         string
	typeFilename string
}

type typeInfo struct {
	pointer bool
	path    string
	name    string
}

func (f *fileDef) getTypeInfo(e ast.Expr) typeInfo {
	switch e := e.(type) {
	case *ast.Ident:
		return typeInfo{path: f.path, name: e.Name}
	case *ast.SelectorExpr:
		if x, ok := e.X.(*ast.Ident); ok && x.Obj == nil {
			if path, ok := f.imports[x.Name]; ok {
				return typeInfo{path: path, name: e.Sel.Name}
			}
		}
	case *ast.StarExpr:
		x := f.getTypeInfo(e.X)
		x.pointer = true
		return x
	}
	return typeInfo{}
}

func (p *progDef) newFileDef(f *ast.File) (*fileDef, error) {
	imports := map[string]string{}
	for _, is := range f.Imports {
		path, err := strconv.Unquote(is.Path.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if is.Name != nil && is.Name.Name != "" && is.Name.Name != "_" {
			imports[is.Name.Name] = path
		}
		name, err := p.pcache.Name(path, p.dir)
		if err != nil {
			return nil, err
		}
		imports[name] = path
	}
	return &fileDef{
		progDef: p,
		imports: imports,
		jast:    jast.New(imports),
	}, nil
}
