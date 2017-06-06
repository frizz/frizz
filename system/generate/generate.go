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

	"strconv"

	"frizz.io/system/generate/jast"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/fatih/camelcase"
	"github.com/pkg/errors"
)

func Save(path string) error {
	env := vos.Os()

	dir, err := patsy.Dir(env, path)
	if err != nil {
		return errors.Wrapf(err, "can't find dir for package %s", path)
	}

	buf := &bytes.Buffer{}
	if err := Generate(buf, env, path, dir); err != nil {
		return errors.WithMessage(err, "error generating source")
	}

	if err := ioutil.WriteFile(filepath.Join(dir, "generated.frizz.go"), buf.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "error saving source")
	}
	return nil
}

type progDef struct {
	fset *token.FileSet
}

type fileDef struct {
	*progDef
	imports map[string]string
	jast    *jast.Cache
}

type typeDef struct {
	*fileDef
	spec ast.Expr
	name string
}

func Generate(writer io.Writer, env vos.Env, path string, dir string) error {

	f := NewFilePath(path)

	prog := &progDef{
		fset: token.NewFileSet(),
	}

	dir, err := patsy.Dir(vos.Os(), path)
	if err != nil {
		return errors.Wrapf(err, "can't find dir for package %s", path)
	}
	pkgs, err := parser.ParseDir(prog.fset, dir, nil, parser.ParseComments)
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

	pcache := patsy.NewCache(env)
	var all []typeDef
	for _, f := range pkg.Files {
		var file *fileDef // lazy init because not all files need action
		var outer error
		ast.Inspect(f, func(n ast.Node) bool {
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
				outer = errors.Errorf("unsupported token tagged with frizz comment at %s", prog.fset.Position(gd.TokPos))
				return false
			}
			if len(gd.Specs) != 1 {
				outer = errors.Errorf("must be a single spec in type definition at %s", prog.fset.Position(gd.TokPos))
				return false
			}
			ts, ok := gd.Specs[0].(*ast.TypeSpec)
			if !ok {
				outer = errors.Errorf("must be type spec at %s", prog.fset.Position(gd.TokPos))
				return false
			}
			if file == nil {
				imports := map[string]string{}
				for _, is := range f.Imports {
					p, err := strconv.Unquote(is.Path.Value)
					if err != nil {
						outer = err
						return false
					}
					if is.Name != nil && is.Name.Name != "" && is.Name.Name != "_" {
						imports[is.Name.Name] = p
					}
					name, err := pcache.Name(p, dir)
					if err != nil {
						outer = err
						return false
					}
					imports[name] = p
				}
				file = &fileDef{
					progDef: prog,
					imports: imports,
					jast:    jast.New(imports),
				}
			}
			all = append(all, typeDef{
				fileDef: file,
				name:    ts.Name.Name,
				spec:    ts.Type,
			})
			return true
		})
		if outer != nil {
			return errors.WithMessage(outer, "error parsing Go source")
		}
	}
	/*
		var Unpackers = struct{
			<name> func(interface{})(<name>, error)
			<...>
		}{
			<name>: unpacker_<name>
			<...>
		}
		func unpacker_<name><unpacker...>
	*/
	f.Var().Id("Unpackers").Op("=").StructFunc(func(g *Group) {
		for _, t := range all {
			g.Id(t.name).Func().Params(Interface()).Params(Id(t.name), Error())
		}
	}).Values(DictFunc(func(d Dict) {
		for _, t := range all {
			d[Id(t.name)] = Id("unpacker_" + t.name)
		}
	}))
	for _, t := range all {
		f.Add(t.unpacker("unpacker_" + t.name))
	}
	if err := f.Render(writer); err != nil {
		return errors.Wrap(err, "error saving generated file")
	}
	return nil
}

func (t *typeDef) unpacker(function string) *Statement {
	/**
	func (in interface{}) (value <named or spec>, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if function != "" {
			s.Id(function)
		}
	}).Params(
		Id("in").Interface(),
	).Params(
		Id("value").Do(func(s *Statement) {
			if t.name != "" {
				s.Id(t.name)
			} else {
				s.Add(t.jast.Expr(t.spec))
			}
		}),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		switch spec := t.spec.(type) {
		default:
			panic(fmt.Sprintf("unsupported type %T", t.spec))
		case *ast.StarExpr:
			t.pointerUnpacker(g, spec)
		case *ast.MapType:
			t.mapUnpacker(g, spec)
		case *ast.ArrayType:
			t.sliceUnpacker(g, spec)
		case *ast.StructType:
			t.structUnpacker(g, spec)
		case *ast.Ident:
			switch spec.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				t.nativeUnpacker(g, spec)
			default:
				t.localUnpacker(g, spec)
			}
		case *ast.SelectorExpr:
			t.selectorUnpacker(g, spec)
		}
	})
}

func (t *typeDef) selectorUnpacker(g *Group, spec *ast.SelectorExpr) {
	/*
		out, err := <spec.X>.Unpackers.<spec.Sel>(in)
		if err != nil {
			return value, err
		}
		return out, nil
	*/
	x, ok := spec.X.(*ast.Ident)
	if !ok {
		panic("spec.X nust be *ast.Ident")
	}
	if x.Obj != nil {
		panic("x.Obj must be nil")
	}
	path, ok := t.imports[x.Name]
	if !ok {
		panic(fmt.Sprintf("%s not found in imports", x.Name))
	}
	g.List(Id("out"), Err()).Op(":=").Qual(path, "Unpackers").Dot(spec.Sel.Name).Call(Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Id("out"), Nil())
}

func (t *typeDef) localUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, err := unpacker_<spec.Name>(in)
		if err != nil {
			return value, err
		}
		return out, nil
	*/
	g.List(Id("out"), Err()).Op(":=").Id("unpacker_" + spec.Name).Call(Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Id("out"), Nil())
}

func (t *typeDef) pointerUnpacker(g *Group, spec *ast.StarExpr) {
	/*
		out, err := <unpacker>(in)
		if err != nil {
			return value, err
		}
		return &out, nil
	*/
	inner := &typeDef{
		fileDef: t.fileDef,
		spec:    spec.X,
		name:    t.name,
	}
	g.List(Id("out"), Err()).Op(":=").Add(inner.unpacker("")).Call(Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Op("&").Id("out"), Nil())
}

func (t *typeDef) mapUnpacker(g *Group, spec *ast.MapType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("error unpacking into map, value should be a map")
		}
		var out = make(<spec>, len(m))
		for k, v := range m {
			u, err := <unpacker>(v)
			if err != nil {
				return value, err
			}
			out[k] = u
		}
		return out, nil
	*/
	inner := &typeDef{
		fileDef: t.fileDef,
		spec:    spec.Value,
		name:    "",
	}
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("error unpacking into map, value should be a map"),
			),
		),
	)
	g.Var().Id("out").Op("=").Make(t.jast.Expr(spec), Len(Id("m")))
	g.For(List(Id("k"), Id("v")).Op(":=").Range().Id("m")).Block(
		List(Id("u"), Err()).Op(":=").Add(inner.unpacker("")).Call(Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), Err()),
		),
		Id("out").Index(Id("k")).Op("=").Id("u"),
	)
	g.Return(Id("out"), Nil())
}

func (t *typeDef) sliceUnpacker(g *Group, spec *ast.ArrayType) {
	/*
		a, ok := in.([]interface{})
		if !ok {
			return value, errors.New("error unpacking into slice, value should be an array")
		}
		<if slice type...>
			var out = make(<spec>, len(a))
		<if array type...>
			var out <spec>
			if len(a) > <len expr> {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), <len expr>)
			}
		<endif>
		for i, v := range a {
			u, err := <unpacker>(v)
			if err != nil {
				return value, err
			}
			out[i] = u
		}
		return out<"[:]" if slice type>, nil
	*/
	inner := &typeDef{
		fileDef: t.fileDef,
		spec:    spec.Elt,
		name:    "",
	}
	g.List(Id("a"), Id("ok")).Op(":=").Id("in").Assert(Index().Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("error unpacking into slice, value should be an array"),
			),
		),
	)
	if spec.Len == nil {
		// return type is a slice type, so we use make to get a slice of the
		// correct size for the data
		g.Var().Id("out").Op("=").Make(t.jast.Expr(spec), Len(Id("a")))
	} else {
		// return type is an array type, so we create an array of the same type
		// and check to make sure the data fits
		g.Var().Id("out").Add(t.jast.Expr(spec))
		g.If(Len(Id("a")).Op(">").Add(t.jast.Expr(spec.Len))).Block(
			Return(
				Id("value"),
				Qual("github.com/pkg/errors", "Errorf").Call(
					Lit("data length %d does not fit in array of length %d"),
					Len(Id("a")),
					t.jast.Expr(spec.Len),
				),
			),
		)
	}
	g.For(List(Id("i"), Id("v")).Op(":=").Range().Id("a")).Block(
		List(Id("u"), Err()).Op(":=").Add(inner.unpacker("")).Call(Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), Err()),
		),
		Id("out").Index(Id("i")).Op("=").Id("u"),
	)
	if spec.Len == nil {
		// return type is a slice type, so we should take a slice of the array
		g.Return(Id("out").Index(Empty(), Empty()), Nil())
	} else {
		// return type is an array type, so we return the bare array
		g.Return(Id("out"), Nil())
	}
}

func (t *typeDef) structUnpacker(g *Group, spec *ast.StructType) {
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
		if t.name != "" {
			s.Id(t.name)
		} else {
			s.Add(t.jast.Expr(spec))
		}
	})
	for _, f := range spec.Fields.List {
		field := &typeDef{
			fileDef: t.fileDef,
			spec:    f.Type,
			name:    "",
		}
		jsonName := goToJson(f.Names[0].Name)
		g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(jsonName)), Id("ok")).Block(
			List(Id("u"), Err()).Op(":=").Add(field.unpacker("")).Call(Id("v")),
			If(Err().Op("!=").Nil()).Block(
				Return(Id("value"), Err()),
			),
			Id("out").Dot(f.Names[0].Name).Op("=").Id("u"),
		)
	}
	g.Return(Id("out"), Nil())
}

func (t *typeDef) nativeUnpacker(g *Group, spec *ast.Ident) {
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
	if t.name == "" {
		g.Return(Id("out"), Nil())
	} else {
		g.Return(Id(t.name).Parens(Id("out")), Nil())
	}
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
