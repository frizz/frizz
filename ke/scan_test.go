package ke_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"strconv"

	"context"

	"encoding/json"

	"github.com/davelondon/kerr"
	"github.com/davelondon/kerr/ksrc"
	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/process/packages"
)

var pkgs = map[string]*pkgDef{}
var all = map[string]*errDef{}

func TestAll(t *testing.T) {
	dir, err := packages.GetDirFromPackage(context.Background(), "kego.io")
	require.NoError(t, err)

	// walk each file in the working directory
	walker := func(path string, file os.FileInfo, err error) error {
		// skip anything starting with a period, or not ending in .go.
		if strings.HasPrefix(path, ".") || !strings.HasSuffix(path, ".go") {
			return nil
		}
		return walkFile(path, t)
	}
	filepath.Walk(dir, walker)

	untested := map[string][]string{}
	for id, def := range all {
		if !def.thrown && (def.tested || def.skipped) {
			// Well, if the test isn't thrown the test should fail. Let's test
			// for it here anyway.
			assert.Fail(t, "Error tested but not thrown", id)
		}
		if def.thrown && def.new && !def.tested && !def.skipped {
			arr, ok := untested[def.pkg]
			if !ok {
				arr = []string{}
			}
			untested[def.pkg] = append(arr, id)
		}
	}

	//for pkg, def := range pkgs {
	//	if !def.tested && !def.notest {
	//assert.Fail(t, fmt.Sprintf("%s has no tests.", pkg))
	//	}
	//}

	//if len(untested) > 0 {
	//	for pkg, tests := range untested {
	//		assert.Fail(t, fmt.Sprintf("Errors thrown in %s but not tested: %v", pkg, tests))
	//	}
	//}
}

func walkFile(path string, t *testing.T) error {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return kerr.Wrap("GJULIOKHIR", err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return kerr.Wrap("HMRIYHYYYH", err)
	}

	pkg, err := packages.GetPackageFromDir(context.Background(), filepath.Dir(path))
	if err != nil {
		return kerr.Wrap("FIPPWAKAGK", err)
	}

	// ignore anything that starts with kego.io/demo/
	if strings.HasPrefix(pkg, "kego.io/demo/") {
		return nil
	}

	def := getPkgDef(pkg)
	if strings.HasSuffix(path, "_test.go") {
		def.tested = true
	}

	kerrName := ""
	assertName := ""
	requireName := ""
	for _, is := range file.Imports {
		importPath, _ := strconv.Unquote(is.Path.Value)
		switch importPath {
		case "github.com/davelondon/kerr":
			if is.Name != nil {
				kerrName, _ = strconv.Unquote(is.Name.Name)
			} else {
				kerrName = "kerr"
			}
		case "github.com/davelondon/ktest/assert":
			if is.Name != nil {
				assertName, _ = strconv.Unquote(is.Name.Name)
			} else {
				assertName = "assert"
			}
		case "github.com/davelondon/ktest/require":
			if is.Name != nil {
				requireName, _ = strconv.Unquote(is.Name.Name)
			} else {
				requireName = "require"
			}
		}
	}

	// visitor implements ast.Visitor
	v := &visitor{
		Bytes:       b,
		t:           t,
		kerrName:    kerrName,
		assertName:  assertName,
		requireName: requireName,
		filePath:    path,
		fset:        fset,
		pkg:         pkg,
	}
	ast.Walk(v, file)

	return nil
}

type visitor struct {
	Bytes       []byte
	t           *testing.T
	kerrName    string
	assertName  string
	requireName string
	filePath    string
	fset        *token.FileSet
	pkg         string
}

type pkgDef struct {
	tested bool
	notest bool
}
type errDef struct {
	id      string
	new     bool
	thrown  bool
	tested  bool
	skipped bool
	pkg     string
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	switch ty := node.(type) {
	case *ast.File:
		for _, cg := range ty.Comments {
			for _, c := range cg.List {
				if strings.HasPrefix(c.Text, "// ke: ") {
					val := struct{ Package struct{ Notest bool } }{}
					err := json.Unmarshal([]byte(c.Text[7:]), &val)
					require.NoError(v.t, err)
					if val.Package.Notest {
						def := getPkgDef(v.pkg)
						def.notest = true
					}
				}
			}
		}
	case *ast.CallExpr:
		name := ""
		pkg := ""
		if f, ok := ty.Fun.(*ast.SelectorExpr); ok {
			name = f.Sel.Name
			if i, ok := f.X.(*ast.Ident); ok {
				pkg = i.Name
			}
			if pkg == "" {
				return v
			}
			if pkg == v.kerrName && (name == "New" || name == "Wrap") {
				def := getErrData(v.t, ty.Args, 0, v.filePath, v.fset.Position(f.Pos()))
				if def == nil {
					return v
				}
				if def.thrown {
					assert.Fail(v.t, fmt.Sprint("Duplicate kerr id: ", def.id))
				}
				def.pkg = v.pkg
				def.new = name == "New"
				def.thrown = true
			} else if (pkg == v.assertName || pkg == v.requireName) && (name == "IsError" || name == "HasError") {
				def := getErrData(v.t, ty.Args, 2, v.filePath, v.fset.Position(f.Pos()))
				if def == nil {
					return v
				}
				def.tested = true
			} else if (pkg == v.assertName || pkg == v.requireName) && name == "SkipError" {
				def := getErrData(v.t, ty.Args, 0, v.filePath, v.fset.Position(f.Pos()))
				if def == nil {
					return v
				}
				def.skipped = true
			}

		}
	}
	return v
}

func getErrData(t *testing.T, args []ast.Expr, arg int, file string, pos token.Position) *errDef {
	require.True(t, len(args) > arg, "Not enough args (%s:%d)", file, pos.Line)
	b, ok := args[arg].(*ast.BasicLit)
	if !ok {
		return nil
	}
	require.Equal(t, b.Kind, token.STRING, "kind should be token.STRING (%s:%d)", file, pos.Line)
	id, err := strconv.Unquote(b.Value)
	require.NoError(t, err, "Error unquoting arg (%s:%d)", file, pos.Line)
	assert.True(t, ksrc.IsId(id), "Invalid kerr ID %s (%s:%d)", id, file, pos.Line)
	def, ok := all[id]
	if ok {
		return def
	}
	def = &errDef{}
	def.id = id
	all[id] = def
	return def
}

func getPkgDef(path string) *pkgDef {
	def, ok := pkgs[path]
	if ok {
		return def
	}
	def = &pkgDef{}
	pkgs[path] = def
	return def
}
