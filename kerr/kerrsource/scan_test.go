package kerrsource_test

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

	"golang.org/x/net/context"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/process/packages"
)

var all = map[string]bool{}

func TestAll(t *testing.T) {
	dir, err := packages.GetDirFromPackage(context.Background(), "kego.io")
	assert.NoError(t, err)

	// walk each file in the working directory
	walker := func(path string, file os.FileInfo, err error) error {
		// skip anything starting with a period, or not ending in .go.
		if strings.HasPrefix(path, ".") || !strings.HasSuffix(path, ".go") {
			return nil
		}
		return walkFile(path, t)
	}
	filepath.Walk(dir, walker)

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

	// visitor implements ast.Visitor
	v := &visitor{Bytes: b, t: t}
	ast.Walk(v, file)

	return nil
}

type visitor struct {
	Bytes []byte
	t     *testing.T
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	switch ty := node.(type) {
	case *ast.CallExpr:
		name := ""
		pkg := ""
		if f, ok := ty.Fun.(*ast.SelectorExpr); ok {
			name = f.Sel.Name
			if i, ok := f.X.(*ast.Ident); ok {
				pkg = i.Name
			}

			if pkg == "kerr" && (name == "New" || name == "Wrap") {
				if len(ty.Args) == 0 {
					assert.Fail(v.t, "Zero args")
				}

				if b, ok := ty.Args[0].(*ast.BasicLit); ok && b.Kind == token.STRING {
					id, err := strconv.Unquote(b.Value)
					assert.NoError(v.t, err)
					if _, found := all[id]; found {
						assert.Fail(v.t, fmt.Sprint("Duplicate kerr id: ", id))
					}
					all[id] = true
				}
			}
		}

	}

	return v
}
