package builder

/*
import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"

	"fmt"

	"path/filepath"

	"io/ioutil"

	"github.com/davelondon/kerr"
	"kego.io/system"
)

// TODO: this is WIP - currently stuck getting comments to line up.
// TODO: See http://stackoverflow.com/questions/31628613 for my question.
func UpdateStruct(path string, outputDir string, t *system.Type) error {

	baseName := t.Id.Name
	if strings.HasPrefix(t.Id.Name, "@") {
		//baseName = t.Id.Name[1:]
		return nil // temporary disable
	}
	fileName := fmt.Sprint(baseName, ".go")
	filePath := filepath.Join(outputDir, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(filePath, []byte(fmt.Sprint("package ", getPackageName(path), "\n")), 0777); err != nil {
			return kerr.New("XHCKADDXEJ", err, "ioutil.WriteFile")
		}
	}

	if err := walkFile(filePath, t); err != nil {
		return kerr.New("MPKDIUPXSQ", err, "walkFile")
	}

	return nil
}

func walkFile(filePath string, t *system.Type) error {

	newSource := `package a

	import "kego.io/system"

	type B struct {
		// This is a changed field comment
		C system.String
		// This is a the second field comment
		D system.Number
	}`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	vGet := &funcVisitor{
		bytes:  []byte(newSource),
		t:      t,
		done:   false,
		err:    nil,
		action: A_GET,
		found:  nil,
		file:   file,
	}
	ast.Walk(vGet, file)

	if vGet.found == nil {
		return fmt.Errorf("Not found")
	}

	newTypeSpec := vGet.found

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fset = token.NewFileSet()
	file, err = parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Print(fset, file)

	v := &funcVisitor{
		bytes:  b,
		t:      t,
		done:   false,
		err:    nil,
		action: A_UPDATE,
		found:  newTypeSpec,
		file:   file,
	}
	ast.Walk(v, file)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := printer.Fprint(f, fset, file); err != nil {
		return err
	}

	return nil
}

type funcVisitor struct {
	bytes  []byte
	t      *system.Type
	done   bool
	err    error
	action actions
	found  *ast.GenDecl
	file   *ast.File
}

type actions string

var A_GET actions = "get"
var A_UPDATE actions = "update"

func (v *funcVisitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil || v.done || v.err != nil {
		return v
	}

	//fmt.Printf("%T: %s\n", node, v.Bytes[node.Pos()-1:node.End()])
	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			break
		}
		//fmt.Printf("%#v\n", n.Specs[0])
		ts := n.Specs[0].(*ast.TypeSpec)
		//fmt.Printf("%#v\n", ts.Name.Name)
		if ts.Name.Name == system.GoName(v.t.Id.Name) {
			if v.action == A_GET {
				v.found = n
				v.done = true
				return v
			}

			if v.action == A_UPDATE {
				fields := ts.Type.(*ast.StructType).Fields

				addField(fields, v.file, "system", "Int", "G", "G comment")
				addField(fields, v.file, "system", "Number", "H", "H comment")
			}
			//if err := updateStruct(ts, v.t); err != nil {
			//	v.err = kerr.New("KTBBWSDARN", err, "updateStruct")
			//	return v
			//}
		}
	}

	return v
}

func addField(fields *ast.FieldList, file *ast.File, pkg string, typ string, name string, comment string) {
	c := &ast.Comment{Text: fmt.Sprint("\n// ", comment)}
	cg := &ast.CommentGroup{List: []*ast.Comment{c}}
	f := &ast.Field{
		Doc:   cg,
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type: &ast.SelectorExpr{
			X:   ast.NewIdent(pkg),
			Sel: ast.NewIdent(typ),
		},
	}
	fields.List = append(fields.List, f)
	file.Comments = append(file.Comments, cg)
}

func updateStruct(ts *ast.TypeSpec, t *system.Type) error {

	fmt.Println("updateStructs")
	fset := token.NewFileSet()
	exp, err := parser.ParseExpr(`type B struct {
			C system.String
		}`)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	ast.Print(fset, exp)

	return nil
}
*/
