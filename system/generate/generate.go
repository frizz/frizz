package generate

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"path/filepath"

	"strings"

	"fmt"

	"unicode"

	"unicode/utf8"

	. "github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/fatih/camelcase"
	"github.com/pkg/errors"
)

func Save(path string) error {
	buf := &bytes.Buffer{}
	if err := Generate(buf, path); err != nil {
		return errors.WithMessage(err, "error generating source")
	}
	// if Generate succeeded then patsy.Dir will not fail
	dir, _ := patsy.Dir(vos.Os(), path)
	if err := ioutil.WriteFile(filepath.Join(dir, "generated.go"), buf.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "error saving source")
	}
	return nil
}

func Generate(writer io.Writer, path string) error {

	f := NewFilePath(path)

	dir, err := patsy.Dir(vos.Os(), path)
	if err != nil {
		return errors.Wrapf(err, "can't find dir for package %s", path)
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "error parsing Go code in package %s", path)
	}
	// ignore any test package
	var filtered []*ast.Package
	for name, pkg := range pkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		filtered = append(filtered, pkg)
	}
	if len(filtered) == 0 {
		return errors.Errorf("can't find Go code in package %s", path)
	}
	if len(filtered) > 1 {
		return errors.Errorf("found more than one package in %s", path)
	}
	pkg := filtered[0]
	var outer error
	ast.Inspect(pkg, func(n ast.Node) bool {
		if err != nil {
			// as soon as error is detected, stop processing
			return false
		}
		if n == nil {
			return true
		}
		gd, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}
		var found bool
		if gd.Doc == nil {
			return true
		}
		for _, c := range gd.Doc.List {
			if c.Text == "// frizz" {
				found = true
				break
			}
		}
		if !found {
			return true
		}
		if gd.Tok != token.TYPE {
			outer = errors.Errorf("unsupported token tagged with frizz comment at %s", fset.Position(gd.TokPos))
			return false
		}
		if len(gd.Specs) != 1 {
			outer = errors.Errorf("must be a single spec in type definition at %s", fset.Position(gd.TokPos))
			return false
		}
		ts, ok := gd.Specs[0].(*ast.TypeSpec)
		if !ok {
			outer = errors.Errorf("must be type spec at %s", fset.Position(gd.TokPos))
			return false
		}
		name := ts.Name.Name
		spec := ts.Type

		f.Add(unpacker(spec, Id("unpack"+name), Id(name)))

		return true
	})
	if outer != nil {
		return errors.WithMessage(outer, "error parsing Go source")
	}
	if err := f.Render(writer); err != nil {
		return errors.Wrap(err, "error saving generated file")
	}
	return nil
}

func unpacker(spec ast.Expr, namedFunc *Statement, namedType *Statement) *Statement {
	/**
	func <named>(in interface{}) (value <namedType>, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if namedFunc != nil {
			// Optional - unpacker is used to generate both anonymous and named
			// functions
			s.Add(namedFunc)
		}
	}).Params(
		Id("in").Interface(),
	).Params(
		Id("value").Do(func(s *Statement) {
			if namedType != nil {
				s.Add(namedType)
			} else {
				s.Add(expr(spec))
			}
		}),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		switch s := spec.(type) {
		case *ast.StructType:
			structUnpacker(g, namedType, s)
		case *ast.Ident:
			switch s.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				nativeUnpacker(g, s)
			}
		}
	})
}

/*
case *BadExpr:
case *Ident:
case *Ellipsis:
case *BasicLit:
case *FuncLit:
case *CompositeLit:
case *ParenExpr:
case *SelectorExpr:
case *IndexExpr:
case *SliceExpr:
case *TypeAssertExpr:
case *CallExpr:
case *StarExpr:
case *UnaryExpr:
case *BinaryExpr:
case *KeyValueExpr:

case *ArrayType:
case *StructType:
case *FuncType:
case *InterfaceType:
case *MapType:
case *ChanType:
*/

// fieldList adds fields to a group - works for both Params() and Struct()
func fieldList(g *Group, fl *ast.FieldList) {
	if fl == nil {
		return
	}
	for _, v := range fl.List {
		g.Do(func(s *Statement) {
			if v.Names != nil {
				s.ListFunc(func(g *Group) {
					for _, n := range v.Names {
						g.Id(n.Name)
					}
				})
			}
			s.Add(expr(v.Type))
		})
	}
}

func expr(spec ast.Expr) *Statement {
	switch spec := spec.(type) {
	case nil:
		return Null()
	case *ast.StructType:
		/*
			struct {
				<fields...>
					<name(s)> <type>
				</fields>
			}
		*/
		return StructFunc(func(g *Group) {
			fieldList(g, spec.Fields)
		})
	case *ast.Ident:
		return Id(spec.Name)
	case *ast.StarExpr:
		return Op("*").Add(expr(spec.X))
	case *ast.ArrayType:
		return Index(expr(spec.Len)).Add(expr(spec.Elt))
	case *ast.FuncType:
		return Func().ParamsFunc(func(g *Group) {
			fieldList(g, spec.Params)
		}).ParamsFunc(func(g *Group) {
			fieldList(g, spec.Results)
		})
	case *ast.InterfaceType:
		return InterfaceFunc(func(g *Group) {
			fieldList(g, spec.Methods)
		})
	case *ast.MapType:
		return Map(expr(spec.Key)).Add(expr(spec.Value))
	case *ast.ChanType:
		return Do(func(s *Statement) {
			if spec.Dir == ast.SEND {
				s.Op("<-")
			}
		}).Chan().Do(func(s *Statement) {
			if spec.Dir == ast.RECV {
				s.Op("<-")
			}
		}).Add(expr(spec.Value))
	case *ast.ParenExpr:
		return Parens(expr(spec.X))
	case *ast.Ellipsis:
		return Op("...").Add(expr(spec.Elt))
	case *ast.BasicLit:
		return Op(spec.Value)
	case *ast.SelectorExpr:
		return expr(spec.X).Dot(spec.Sel.Name)
	case *ast.IndexExpr:
		return expr(spec.X).Index(expr(spec.Index))
	case *ast.SliceExpr:
		return expr(spec.X).IndexFunc(func(g *Group) {
			if spec.Low != nil {
				g.Add(expr(spec.Low))
			} else {
				g.Empty()
			}
			if spec.High != nil {
				g.Add(expr(spec.High))
			} else {
				g.Empty()
			}
			if spec.Max != nil {
				g.Add(expr(spec.Max))
			} else {
				g.Null()
			}
		})
	case *ast.TypeAssertExpr:
		return expr(spec.X).Assert(expr(spec.Type))
	case *ast.CallExpr:
		return expr(spec.Fun).CallFunc(func(g *Group) {
			for _, arg := range spec.Args {
				g.Add(expr(arg))
			}
		})
	case *ast.UnaryExpr:
		return Op(spec.Op.String()).Add(expr(spec.X))
	case *ast.BinaryExpr:
		return expr(spec.X).Op(spec.Op.String()).Add(expr(spec.Y))
	case *ast.KeyValueExpr:
		// kludge: would usually use Dict here but it's not a *Statement so
		// this will work.
		return expr(spec.Key).Op(":").Add(expr(spec.Value))
	case *ast.CompositeLit:
		return ValuesFunc(func(g *Group) {
			for _, element := range spec.Elts {
				g.Add(expr(element))
			}
		})
		//case *ast.BadExpr:
		//case *ast.FuncLit:
	}
	return Null()
}

func structUnpacker(g *Group, namedType *Statement, spec *ast.StructType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("error unpacking into struct, value should be a map")
		}
		var out <type>
		<fields...>
		if v, ok := m["<jsonName>"]; ok {
			u, err := <unpacker>(v)
			if err != nil {
				return value, err
			}
			out.<fieldName> = u
		}
		</fields>
		return out, nil
	*/
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("error unpacking into struct, value should be a map"),
			),
		),
	)
	g.Var().Id("out").Do(func(s *Statement) {
		if namedType != nil {
			s.Add(namedType)
		} else {
			s.Add(expr(spec))
		}
	})
	for _, f := range spec.Fields.List {
		jsonName := goToJson(f.Names[0].Name)
		g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(jsonName)), Id("ok")).Block(
			List(Id("u"), Err()).Op(":=").Add(unpacker(f.Type, nil, nil)).Call(Id("v")),
			If(Err().Op("!=").Nil()).Block(
				Return(Id("value"), Err()),
			),
			Id("out").Dot(f.Names[0].Name).Op("=").Id("u"),
		)
	}
	g.Return(Id("out"), Nil())
}

func nativeUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, err := system.Convert_<spec.Name>(in)
		if err != nil {
			return nil, err
		}
		return out, nil
	*/
	g.List(Id("out"), Err()).Op(":=").Qual("frizz.io/system", fmt.Sprintf("Convert_%s", spec.Name)).Call(Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Id("out"), Nil())
}

func goToJson(name string) string {

	words := camelcase.Split(name)

	if len(words) == 0 {
		return ""
	}

	firstRune, _ := utf8.DecodeRuneInString(words[0])
	public := unicode.IsUpper(firstRune)

	var prefix string
	if !public {
		prefix = "_"
	}
	return prefix + strings.ToLower(strings.Join(words, "-"))
}
