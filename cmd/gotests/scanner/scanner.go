package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/kerr/kerrsource"
	"kego.io/process/packages"
)

// ke: {"package": {"notest": true}}

var baseDir string

type Source struct {
	Wraps            []ErrDef
	ExcludedBlocks   []PosDef
	ExcludedPackages map[string]bool
	ExcludedFuncs    []FuncDef
	Skipped          map[string]bool
	All              map[string]ErrDef
	JsTestPackages   map[string]bool
}
type ErrDef struct {
	Id   string
	File string
	Line int
}
type PosDef struct {
	File string
	Line int
}
type FuncDef struct {
	File      string
	LineStart int
	LineEnd   int
}

var source = &Source{
	Skipped:          map[string]bool{},
	ExcludedPackages: map[string]bool{},
	All:              map[string]ErrDef{},
	JsTestPackages:   map[string]bool{},
}

func Get(dir string) (*Source, error) {
	baseDir = dir

	// walk each file in the working directory
	walker := func(filename string, file os.FileInfo, err error) error {
		if strings.HasPrefix(getParentDir(filename), ".") {
			return filepath.SkipDir
		}
		if strings.HasPrefix(file.Name(), ".") {
			return nil
		}
		if file.IsDir() {
			return nil
		}
		if strings.HasSuffix(filename, ".go") {
			return scanFile(filename)
		}
		return nil
	}
	if err := filepath.Walk(baseDir, walker); err != nil {
		return nil, err
	}

	return source, nil
}

func getParentDir(filename string) string {
	dir, _ := filepath.Split(filename)
	dirs := strings.Split(dir, string(os.PathSeparator))
	return dirs[len(dirs)-2]
}

func scanFile(filename string) error {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	pkg, err := packages.GetPackageFromDir(context.Background(), filepath.Dir(filename))
	if err != nil {
		return err
	}

	kerrName := ""
	assertName := ""
	for _, is := range file.Imports {
		importPath, _ := strconv.Unquote(is.Path.Value)
		switch importPath {
		case "kego.io/kerr":
			if is.Name != nil {
				kerrName, _ = strconv.Unquote(is.Name.Name)
			} else {
				kerrName = "kerr"
			}
		case "kego.io/kerr/assert":
			if is.Name != nil {
				assertName, _ = strconv.Unquote(is.Name.Name)
			} else {
				assertName = "assert"
			}
		}
	}

	relfilename, err := filepath.Rel(filepath.Join(baseDir, ".."), filename)
	if err != nil {
		return err
	}

	for _, cg := range file.Comments {
		for _, c := range cg.List {
			if strings.HasPrefix(c.Text, "// ke: ") {
				val := struct {
					Block struct {
						Notest bool
					}
					Package struct {
						Jstest bool
						Notest bool
					}
				}{}
				err := json.UnmarshalPlain([]byte(c.Text[7:]), &val)
				if err != nil {
					return err
				}
				if val.Block.Notest {
					source.ExcludedBlocks = append(source.ExcludedBlocks, PosDef{
						File: relfilename,
						Line: fset.Position(c.Pos()).Line,
					})
				}
				if val.Package.Notest {
					source.ExcludedPackages[pkg] = true
				}
				if val.Package.Jstest {
					source.JsTestPackages[pkg] = true
				}
			}
		}
	}

	// visitor implements ast.Visitor
	v := &visitor{
		Bytes:      b,
		kerrName:   kerrName,
		assertName: assertName,
		file:       relfilename,
		pkg:        pkg,
		fset:       fset,
	}
	ast.Walk(v, file)

	return v.err
}

type visitor struct {
	Bytes      []byte
	t          *testing.T
	kerrName   string
	assertName string
	file       string
	pkg        string
	fset       *token.FileSet
	err        error
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	switch ob := node.(type) {
	/*
		case *ast.File:

			for _, cg := range ty.Comments {
				for _, c := range cg.List {
						if strings.HasPrefix(c.Text, "// ke: ") {
							val := struct{ Notest bool }{}
							err := json.UnmarshalPlain([]byte(c.Text[7:]), &val)
							assert.NoError(v.t, err)
							if val.Notest {
								def := getPkgDef(v.pkg)
								def.notest = true
							}
						}
				}
			}
	*/
	case *ast.FuncDecl:
		if ob.Doc == nil {
			return v
		}
		for _, c := range ob.Doc.List {
			if strings.HasPrefix(c.Text, "// ke: ") {
				val := struct {
					Func struct {
						Notest bool
					}
				}{}
				err := json.UnmarshalPlain([]byte(c.Text[7:]), &val)
				if err != nil {
					v.err = err
					return nil
				}
				if val.Func.Notest {
					source.ExcludedFuncs = append(source.ExcludedFuncs, FuncDef{
						File:      v.file,
						LineStart: v.fset.Position(ob.Pos()).Line,
						LineEnd:   v.fset.Position(ob.End()).Line,
					})
				}
			}
		}
	case *ast.CallExpr:
		name := ""
		pkg := ""
		if f, ok := ob.Fun.(*ast.SelectorExpr); ok {
			name = f.Sel.Name
			if i, ok := f.X.(*ast.Ident); ok {
				pkg = i.Name
			}
			if pkg == "" {
				return v
			}
			if pkg == v.kerrName && name == "Wrap" {
				id, err := getErrorId(v.t, ob.Args, 0, v.file)
				if err != nil {
					v.err = err
					return nil
				}
				source.Wraps = append(source.Wraps, ErrDef{
					Id:   id,
					File: v.file,
					Line: v.fset.Position(ob.Pos()).Line,
				})
			} else if pkg == v.kerrName && name == "New" {
				id, err := getErrorId(v.t, ob.Args, 0, v.file)
				if err != nil {
					v.err = err
					return nil
				}
				source.All[id] = ErrDef{
					Id:   id,
					File: v.file,
					Line: v.fset.Position(ob.Pos()).Line,
				}
			} else if pkg == v.assertName && (name == "SkipError") {
				id, err := getErrorId(v.t, ob.Args, 0, v.file)
				if err != nil {
					v.err = err
					return nil
				}
				source.Skipped[id] = true
			} /*else if pkg == v.assertName && (name == "IsError" || name == "HasError") {
				def := getErrData(v.t, ty.Args, 2, v.file)
				def.tested = true
			} else if pkg == v.assertName && (name == "SkipError") {
				def := getErrData(v.t, ty.Args, 0, v.file)
				def.skipped = true
			} else if pkg == v.assertName && (name == "IsErrorMulti" || name == "HasErrorMulti") {
				assert.True(v.t, len(ty.Args) > 2, "Not enough args (%s)", v.file)
				for arg := 2; arg < len(ty.Args); arg++ {
					def := getErrData(v.t, ty.Args, arg, v.file)
					def.tested = true
				}
			}*/

		}
	}
	return v
}

func getErrorId(t *testing.T, args []ast.Expr, arg int, file string) (string, error) {
	if len(args) <= arg {
		return "", kerr.New("FLKWAYFFYL", "Not enough args (%s)", file)
	}
	b, ok := args[arg].(*ast.BasicLit)
	if !ok {
		return "", kerr.New("YFKEYRIUWJ", "Arg should be *ast.BasicLit (%s)", file)
	}
	if b.Kind != token.STRING {
		return "", kerr.New("YWPXGGRLEL", "kind should be token.STRING (%s)", file)
	}
	id, err := strconv.Unquote(b.Value)
	if err != nil {
		return "", kerr.Wrap("NTUDBTYOQY", err)
	}
	if !kerrsource.IsId(id) {
		return "", kerr.New("YLTSLGODKR", "Invalid kerr ID %s (%s)", id, file)
	}
	return id, nil
}

/*
func getErrData(t *testing.T, args []ast.Expr, arg int, file string) *errDef {
	assert.True(t, len(args) > arg, "Not enough args (%s)", file)
	b, ok := args[arg].(*ast.BasicLit)
	assert.True(t, ok, "Arg should be *ast.BasicLit (%s)", file)
	assert.Equal(t, b.Kind, token.STRING, "kind should be token.STRING (%s)", file)
	id, err := strconv.Unquote(b.Value)
	assert.NoError(t, err, "Error unquoting arg (%s)", file)
	assert.True(t, kerrsource.IsId(id), "Invalid kerr ID %s (%s)", id, file)
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
*/
