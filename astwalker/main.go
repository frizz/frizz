package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"fmt"
	"io/ioutil"
)

func main() {
	if false {
		walkFile("/Users/dave/dev/proj/go/src/kego.io/process/edit.go")
	} else {
		walker := func(path string, file os.FileInfo, err error) error {
			if strings.HasPrefix(path, ".") || !strings.HasSuffix(path, ".go") {
				return nil
			}
			return walkFile(path)
		}
		filepath.Walk(".", walker)
	}
}

func walkFile(path string) error {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	v := &FuncVisitor{b, false}
	ast.Walk(v, file)

	if v.Found {
		fmt.Println(path)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := printer.Fprint(f, fset, file); err != nil {
			return err
		}
	}

	return nil
}

type FuncVisitor struct {
	Bytes []byte
	Found bool
}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {

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
		}
		if name == "New" && pkg == "kerr" {
			t.Args = append(t.Args[0:2], t.Args[3:]...)
			v.Found = true
		}
	}

	return v
}
