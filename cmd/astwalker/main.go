// astwalker walks through all source files, finds the kerr.New function, and
// removes the third parameter. This was used to refactor the code base and
// learn about the ast package.
package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/kerr"

	"fmt"
	"io/ioutil"
)

// notest

func main() {

	// disabled
	return

	if true {
		// for testing
		walkFile("/Users/dave/dev/src/frizz.io/cmd/astwalker/test.go")
	} else {
		// walk each file in the working directory
		walker := func(path string, file os.FileInfo, err error) error {
			// skip anything starting with a period, or not ending in .go.
			if strings.HasPrefix(path, ".") || !strings.HasSuffix(path, ".go") {
				return nil
			}
			fmt.Println(path)
			return walkFile(path)
		}
		filepath.Walk("/Users/dave/dev/src/frizz.io/", walker)
	}
}

func walkFile(path string) error {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return kerr.Wrap("FGAHVRNUPV", err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return kerr.Wrap("ALRIEYJBJE", err)
	}

	// visitor implements ast.Visitor
	v := &visitor{b, false}
	ast.Walk(v, file)

	// if we made a replacement, re-write the modified file
	if v.Found {
		fmt.Println(path)
		f, err := os.Create(path)
		if err != nil {
			return kerr.Wrap("QFPDQRTIRS", err)
		}
		defer f.Close()
		if err := printer.Fprint(f, fset, file); err != nil {
			return kerr.Wrap("VEPVDWFWEF", err)
		}
	}

	return nil
}

type visitor struct {
	Bytes []byte
	Found bool
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	switch t := node.(type) {
	case *ast.CallExpr:
		name := ""
		pkg := ""
		if f, ok := t.Fun.(*ast.SelectorExpr); ok {
			name = f.Sel.Name
			if i, ok := f.X.(*ast.Ident); ok {
				pkg = i.Name
			}

			if pkg == "kerr" && name == "New" {
				// for call expressions kerr.New, remove the third parameter
				//t.Args = append(t.Args[0:2], t.Args[3:]...)

				if i, ok := t.Args[1].(*ast.Ident); ok && i.Name == "nil" && i.Obj == nil {
					// If the 2nd parameter is nil, rename to New1 and remove the 2nd parameter.
					f.Sel.Name = "New1"
					t.Args = append(t.Args[0:1], t.Args[2:]...)
				} else {
					// If not, rename to Wrap and only use the first 2 parameters
					f.Sel.Name = "Wrap"
					t.Args = t.Args[0:2]
				}

				v.Found = true
			}
		}

	}

	return v
}
