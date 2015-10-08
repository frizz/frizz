package generator

/*
import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"testing"
)

// TODO: this is WIP - currently stuck getting comments to line up.
// TODO: See http://stackoverflow.com/questions/31628613 for my question.
func TestAstGen(t *testing.T) {

	t.Skip()

	source := `package a

	// B comment
	type B struct {
		// C comment
		C string
		// D comment
		D string
	}`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", []byte(source), parser.ParseComments)
	if err != nil {
		t.Error(err)
	}

	ast.Print(fset, file)

	v := &visitor{
		b:    []byte(source),
		file: file,
	}
	ast.Walk(v, file)

	var output []byte
	buf := bytes.NewBuffer(output)
	if err := printer.Fprint(buf, fset, file); err != nil {
		t.Error(err)
	}

	expected := `package a

	// B comment
	type B struct {
		// C comment
		C string
		// D comment
		D int
		// E comment
		E float64
	}
	`

	if buf.String() != expected {
		t.Error(fmt.Sprintf("Test failed. Expected:\n%s\nGot:\n%s", expected, buf.String()))
	}

	/*
		actual output = `package a

		// B comment
		type B struct {
			// C comment
			// D comment
			// E comment
			C	string
			D	int
			E	float64
		}
		`


}

type visitor struct {
	b    []byte
	file *ast.File
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	snippet := ""
	if node.Pos() > 0 && node.End() > node.Pos() && len(v.b) >= int(node.End())-2 {
		snippet = string(v.b[node.Pos()-1 : node.End()-1])
		snippet = strings.Replace(strings.Replace(snippet, "\n", "", -1), "\t", " ", -1)
	}
	fmt.Printf("%T: %d %d %s\n", node, node.Pos(), node.End(), snippet)
	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			break
		}
		ts := n.Specs[0].(*ast.TypeSpec)
		if ts.Name.Name == "B" {
			fields := ts.Type.(*ast.StructType).Fields
			addStructField(fields, v.file, "int", "E", "E comment")
			addStructField(fields, v.file, "float64", "F", "F comment")
		}
	}

	return v
}

func addStructField(fields *ast.FieldList, file *ast.File, typ string, name string, comment string) {
	previousField := fields.List[len(fields.List)-1]
	commentPos := previousField.End()
	c := &ast.Comment{Text: fmt.Sprint("// ", comment), Slash: commentPos}
	cg := &ast.CommentGroup{List: []*ast.Comment{c}}
	file.Comments = append(file.Comments, cg)

	structPos := c.End()

	id := ast.NewIdent(name)
	id.NamePos = token.Pos(structPos)

	f := &ast.Field{
		Doc:   cg,
		Names: []*ast.Ident{id},
		Type:  ast.NewIdent(typ),
	}
	fields.List = append(fields.List, f)

}
*/
