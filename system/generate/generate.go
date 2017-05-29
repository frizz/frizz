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
		if err := addUnpacker(fset, f, name, spec); err != nil {
			outer = errors.WithMessage(err, fmt.Sprintf("error printing unpacker for type at %s", fset.Position(gd.TokPos)))
			return false
		}

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

func addUnpacker(fset *token.FileSet, f *File, name string, spec ast.Expr) error {
	var outer error
	/*
		func unpack<name>(in interface{}) (interface{}, error) {
			<...>
		}
	*/
	f.Func().Id("unpack"+name).Params(
		Id("in").Interface(),
	).Params(
		Interface(),
		Error(),
	).BlockFunc(func(g *Group) {
		switch s := spec.(type) {
		// TODO
		case *ast.StructType:
			/*
				m, ok := in.(map[string]interface{})
				if !ok {
					return nil, errors.Errorf("error unpacking into %s, value should be a map", name)
				}
				out := new(<name>)
				<...>
				return out, nil
			*/
			g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
			g.If(Op("!").Id("ok")).Block(
				Return(
					Nil(),
					Qual("github.com/pkg/errors", "Errorf").Call(
						Lit("error unpacking into %s, value should be a map"),
						Lit(name),
					),
				),
			)
			g.Id("out").Op(":=").New(Id(name))
			for _, f := range s.Fields.List {
				fieldName := f.Names[0].Name
				jsonName := goToJson(f.Names[0].Name)
				switch ft := f.Type.(type) {
				default:
					ast.Print(fset, f)
				case *ast.Ident:
					//var native string
					switch ft.Name {
					case "bool":
						// native = "bool"
					case "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64":
						// native = "float64"
					case "rune", "string":
						// native = "string"
					case "complex64", "complex128", "error", "uintptr":
						outer = errors.Errorf("type %s not supported", ft.Name)
						return
					default:
						ast.Print(fset, f)
					}
					/*
						if v, ok := m["<jsonName>"]; ok {
							c, err := system.Convert_<ft.Name>(v)
							if err != nil {
								return nil, errors.WithMessage(err, "error unpacking")
							}
							out.<fieldName> = c
						}
					*/
					g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(jsonName)), Id("ok")).Block(
						List(Id("c"), Err()).Op(":=").Qual("frizz.io/system", fmt.Sprintf("Convert_%s", ft.Name)).Call(Id("v")),
						If(Err().Op("!=").Nil()).Block(
							Return(
								Nil(),
								Qual("github.com/pkg/errors", "WithMessage").Call(
									Err(),
									Lit("error unpacking"),
								),
							),
						),
						Id("out").Dot(fieldName).Op("=").Id("c"),
					)
				}
			}
			g.Return(Id("out"), Nil())
		default:
			ast.Print(fset, s)
		}
	})
	if outer != nil {
		return errors.WithMessage(outer, "error generating source")
	}
	return nil
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
