// kerrfmt walks through the given source file, finds the kerr.New function, if it has an empty
// string as the first parameter, it replaces with a kerr id.
package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"math/rand"
	"os"
	"time"

	"kego.io/kerr/util"

	"io/ioutil"
	"strconv"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()
	//file := flag.Arg(0)
	file := "/Users/dave/dev/src/kego.io/kerr/cmd/kerrfmt/test.go"
	if err := walkFile(file); err != nil {
		panic(err.Error())
	}

}

func walkFile(filename string) error {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// visitor implements ast.Visitor
	v := &visitor{b, false}
	ast.Walk(v, file)

	// if we made a replacement, re-write the modified file
	if v.Found {
		temp := filename + "_temp"
		f, err := os.Create(temp)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := printer.Fprint(f, fset, file); err != nil {
			os.Remove(temp)
			return err
		}
		if err := os.Rename(temp, filename); err != nil {
			return err
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
		}
		if pkg != "kerr" {
			return v
		}
		switch name {
		case "New":
			if b, ok := t.Args[0].(*ast.BasicLit); ok && b.Kind == token.STRING && b.Value == "\"\"" {
				b.Value = strconv.Quote(util.GetRandomId())
			}
			v.Found = true
		case "Wrap":

		}
	}

	return v
}
