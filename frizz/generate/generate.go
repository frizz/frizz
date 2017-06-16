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

	"strconv"

	"frizz.io/frizz/generate/jast"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func Save(path string) error {
	// notest
	env := vos.Os()

	dir, err := patsy.Dir(env, path)
	if err != nil {
		return errors.Wrapf(err, "can't find dir for package %s", path)
	}

	buf := &bytes.Buffer{}
	if err := Generate(buf, env, path, dir); err != nil {
		return errors.WithMessage(err, "error generating source")
	}

	if err := ioutil.WriteFile(filepath.Join(dir, "types.frizz.go"), buf.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "error saving source")
	}
	return nil
}

type progDef struct {
	fset   *token.FileSet
	path   string
	pcache *patsy.Cache
	dir    string
}

type fileDef struct {
	*progDef
	imports map[string]string
	jast    *jast.Cache
}

type typeDef struct {
	*fileDef
	packer bool
	spec   ast.Expr
	name   string
}

type typeInfo struct {
	pointer bool
	path    string
	name    string
}

func (f *fileDef) getTypeInfo(e ast.Expr) typeInfo {
	switch e := e.(type) {
	case *ast.Ident:
		return typeInfo{path: f.path, name: e.Name}
	case *ast.SelectorExpr:
		if x, ok := e.X.(*ast.Ident); ok && x.Obj == nil {
			if path, ok := f.imports[x.Name]; ok {
				return typeInfo{path: path, name: e.Sel.Name}
			}
		}
	case *ast.StarExpr:
		x := f.getTypeInfo(e.X)
		x.pointer = true
		return x
	}
	return typeInfo{}
}

func (p *progDef) newFileDef(f *ast.File) (*fileDef, error) {
	imports := map[string]string{}
	for _, is := range f.Imports {
		path, err := strconv.Unquote(is.Path.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if is.Name != nil && is.Name.Name != "" && is.Name.Name != "_" {
			imports[is.Name.Name] = path
		}
		name, err := p.pcache.Name(path, p.dir)
		if err != nil {
			return nil, err
		}
		imports[name] = path
	}
	return &fileDef{
		progDef: p,
		imports: imports,
		jast:    jast.New(imports),
	}, nil
}

func Generate(writer io.Writer, env vos.Env, path string, dir string) error {

	f := NewFilePath(path)

	prog := &progDef{
		fset:   token.NewFileSet(),
		path:   path,
		pcache: patsy.NewCache(env),
		dir:    dir,
	}

	pkgs, err := parser.ParseDir(prog.fset, prog.dir, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "error parsing Go code in package %s", path)
	}
	// ignore any test package
	var filtered []*ast.Package
	for name, pkg := range pkgs {
		if strings.HasSuffix(name, "_test") {
			// notest
			continue
		}
		filtered = append(filtered, pkg)
	}
	if len(filtered) == 0 {
		// notest
		return errors.Errorf("can't find Go code in package %s", path)
	}
	if len(filtered) > 1 {
		// notest
		return errors.Errorf("found more than one package in %s", path)
	}
	pkg := filtered[0]

	var all []*typeDef
	var unpackers = map[string]bool{}
	//	var repackers = map[string]bool{}
	for _, f := range pkg.Files {
		var file *fileDef // lazy init because not all files need action
		var outer error
		ast.Inspect(f, func(n ast.Node) bool {
			if err != nil {
				// as soon as error is detected, stop processing
				return false
			}
			switch n := n.(type) {
			case nil:
				return true
			case *ast.FuncDecl:
				if n.Name.Name != "Unpack" { // && n.Name.Name != "Repack" {
					return true
				}
				if n.Recv == nil {
					// function not method
					return true
				}
				if len(n.Recv.List) != 1 {
					return true // impossible?
				}
				s, ok := n.Recv.List[0].Type.(*ast.StarExpr)
				if !ok {
					return true
				}
				id, ok := s.X.(*ast.Ident)
				if !ok {
					return true
				}
				receiverType := id.Name

				if file == nil {
					file, err = prog.newFileDef(f)
					if err != nil {
						outer = err
						return false
					}
				}
				// check the signature
				// Unpack: (*frizz.Root, frizz.Stack, interface{}) error
				switch n.Name.Name {
				case "Unpack":
					if len(n.Type.Params.List) != 3 {
						return true
					}
					param0 := file.getTypeInfo(n.Type.Params.List[0].Type)
					param1 := file.getTypeInfo(n.Type.Params.List[1].Type)
					if param0 != (typeInfo{true, "frizz.io/frizz", "Root"}) {
						return true
					}
					if param1 != (typeInfo{false, "frizz.io/frizz", "Stack"}) {
						return true
					}
					ift, ok := n.Type.Params.List[2].Type.(*ast.InterfaceType)
					if !ok {
						return true
					}
					if len(ift.Methods.List) != 0 {
						return true
					}

					// check return is error
					if len(n.Type.Results.List) != 1 {
						return true
					}
					ide, ok := n.Type.Results.List[0].Type.(*ast.Ident)
					if !ok {
						return true
					}
					if ide.Name != "error" {
						return true
					}
					// unpack method is the correct signature.
					unpackers[receiverType] = true
				}

			case *ast.GenDecl:
				var found bool
				if n.Doc == nil {
					return true
				}
				for _, c := range n.Doc.List {
					if c.Text == "// frizz" {
						found = true
						break
					}
				}
				if !found {
					return true
				}
				if n.Tok != token.TYPE {
					outer = errors.Errorf("unsupported token tagged with frizz comment at %s", prog.fset.Position(n.TokPos))
					return false
				}
				if len(n.Specs) != 1 {
					outer = errors.Errorf("must be a single spec in type definition at %s", prog.fset.Position(n.TokPos))
					return false
				}
				ts, ok := n.Specs[0].(*ast.TypeSpec)
				if !ok {
					outer = errors.Errorf("must be type spec at %s", prog.fset.Position(n.TokPos))
					return false
				}
				if file == nil {
					file, err = prog.newFileDef(f)
					if err != nil {
						outer = err
						return false
					}
				}
				all = append(all, &typeDef{
					fileDef: file,
					name:    ts.Name.Name,
					spec:    ts.Type,
				})
			}

			return true
		})
		if outer != nil {
			return errors.WithMessage(outer, "error parsing Go source")
		}
		for _, td := range all {
			_, hasUnpacker := unpackers[td.name]
			if hasUnpacker {
				td.packer = true
			}
		}
	}
	/*
		var Unpackers = struct{
			<name> func(*frizz.Root, frizz.Stack, interface{})(<name>, error)
			<...>
		}{
			<name>: unpack_<name>
			<...>
		}
		func unpack_<name><unpacker...>
	*/
	f.Var().Id("Unpackers").Op("=").StructFunc(func(g *Group) {
		for _, t := range all {
			g.Id(t.name).Func().Params(
				Op("*").Qual("frizz.io/frizz", "Root"),
				Qual("frizz.io/frizz", "Stack"),
				Interface(),
			).Params(Id(t.name), Error())
		}
	}).Values(DictFunc(func(d Dict) {
		for _, t := range all {
			d[Id(t.name)] = Id("unpack_" + t.name)
		}
	}))
	for _, t := range all {
		f.Add(t.unpacker(t.spec, t.name, "unpack_"+t.name, t.packer))
	}

	/*
		func init() {
			frizz.DefaultRegistry.Set(
				"<path>",
				"<name>",
				func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
					return unpack_<name>(root, stack, in)
				},
			)
			<...>
		}
	*/
	f.Func().Id("init").Params().BlockFunc(func(g *Group) {
		for _, t := range all {
			g.Qual("frizz.io/frizz", "DefaultRegistry").Dot("Set").Call(
				Lit(t.path),
				Lit(t.name),
				Func().Params(
					Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
					Id("stack").Qual("frizz.io/frizz", "Stack"),
					Id("in").Interface(),
				).Params(Interface(), Error()).Block(
					Return(Id("unpack_"+t.name).Call(Id("root"), Id("stack"), Id("in"))),
				),
			)
		}
	})

	if err := f.Render(writer); err != nil {
		return errors.Wrap(err, "error saving generated file")
	}
	return nil
}

func (f *fileDef) unpacker(spec ast.Expr, name, function string, packer bool) *Statement {
	/**
	func (root *frizz.Root, stack frizz.Stack, in interface{}) (value <named or spec>, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if function != "" {
			s.Id(function)
		}
	}).Params(
		Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
		Id("stack").Qual("frizz.io/frizz", "Stack"),
		Id("in").Interface(),
	).Params(
		Id("value").Do(func(s *Statement) {
			if name != "" {
				s.Id(name)
			} else {
				s.Add(f.jast.Expr(spec))
			}
		}),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		if packer {
			/*
				out := new(<name>)
				if err := out.Unpack(root, stack, in), err != nil {
					return value, err
				}
				return *out, nil
			*/
			g.Id("out").Op(":=").New(Id(name))
			g.If(
				Err().Op(":=").Id("out").Dot("Unpack").Call(Id("root"), Id("stack"), Id("in")),
				Err().Op("!=").Nil(),
			).Block(
				Return(Id("value"), Err()),
			)
			g.Return(Op("*").Id("out"), Nil())
			return
		}
		switch spec := spec.(type) {
		default:
			panic(fmt.Sprintf("unsupported type %T", spec))
		case *ast.InterfaceType:
			f.interfaceUnpacker(g, spec, name)
		case *ast.StarExpr:
			f.pointerUnpacker(g, spec, name)
		case *ast.MapType:
			f.mapUnpacker(g, spec, name)
		case *ast.ArrayType:
			f.sliceUnpacker(g, spec, name)
		case *ast.StructType:
			f.structUnpacker(g, spec, name)
		case *ast.Ident:
			switch spec.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				f.nativeUnpacker(g, spec, name)
			default:
				f.localUnpacker(g, spec, name)
			}
		case *ast.SelectorExpr:
			f.selectorUnpacker(g, spec, name)
		}
	})
}

func (f *fileDef) interfaceUnpacker(g *Group, spec *ast.InterfaceType, alias string) {
	/*
		out, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, err
		}
		iface, ok := out.(<name or spec>)
		if !ok {
			return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		return iface, nil
	*/
	g.Comment("interfaceUnpacker")
	g.List(Id("out"), Err()).Op(":=").Id("root").Dot("UnpackInterface").Call(Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.List(Id("iface"), Id("ok")).Op(":=").Id("out").Assert(Do(func(s *Statement) {
		if alias != "" {
			s.Id(alias)
		} else {
			s.Add(f.jast.Expr(spec))
		}
	}))
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("unpacking into interface, type %T does not implement interface"),
				Id("out"),
			),
		),
	)
	g.Return(Id("iface"), Nil())
}

func (f *fileDef) selectorUnpacker(g *Group, spec *ast.SelectorExpr, alias string) {
	/*
		out, err := <spec.X>.Unpackers.<spec.Sel>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(out), nil
	*/
	g.Comment("selectorUnpacker")
	g.List(Id("out"), Err()).Op(":=").Do(func(s *Statement) {
		x, ok := spec.X.(*ast.Ident)
		if !ok {
			panic("spec.X nust be *ast.Ident")
		}
		if x.Obj != nil {
			panic("x.Obj must be nil")
		}
		path, ok := f.imports[x.Name]
		if !ok {
			panic(fmt.Sprintf("%s not found in imports", x.Name))
		}
		s.Qual(path, "Unpackers")
	}).Dot(spec.Sel.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if alias == "" {
			s.Id("out")
		} else {
			s.Id(alias).Parens(Id("out"))
		}
	}), Nil())
}

func (f *fileDef) localUnpacker(g *Group, spec *ast.Ident, alias string) {
	/*
		out, err := unpack_<spec.Name>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(out), nil
	*/
	g.Comment("localUnpacker")
	g.List(Id("out"), Err()).Op(":=").Id("unpack_"+spec.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if alias == "" {
			s.Id("out")
		} else {
			s.Id(alias).Parens(Id("out"))
		}
	}), Nil())
}

func (f *fileDef) pointerUnpacker(g *Group, spec *ast.StarExpr, alias string) {
	/*
		out, err := <unpacker>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(&out), nil
	*/
	g.Comment("pointerUnpacker")
	g.List(Id("out"), Err()).Op(":=").Add(f.unpacker(spec.X, "", "", false)).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if alias == "" {
			s.Op("&").Id("out")
		} else {
			s.Id(alias).Parens(Op("&").Id("out"))
		}
	}), Nil())
}

func (f *fileDef) mapUnpacker(g *Group, spec *ast.MapType, alias string) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into map, value should be a map")
		}
		var out = make(<name or spec>, len(m))
		for k, v := range m {
			stack := stack.Append(frizz.MapItem(k))
			u, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, err
			}
			out[k] = u
		}
		return out, nil
	*/
	g.Comment("mapUnpacker")
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into map, value should be a map"),
			),
		),
	)
	g.Var().Id("out").Op("=").Make(Do(func(s *Statement) {
		if alias != "" {
			s.Id(alias)
		} else {
			s.Add(f.jast.Expr(spec))
		}
	}), Len(Id("m")))
	g.For(List(Id("k"), Id("v")).Op(":=").Range().Id("m")).Block(
		Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "MapItem").Parens(Id("k"))),
		List(Id("u"), Err()).Op(":=").Add(f.unpacker(spec.Value, "", "", false)).Call(Id("root"), Id("stack"), Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), Err()),
		),
		Id("out").Index(Id("k")).Op("=").Id("u"),
	)
	g.Return(Id("out"), Nil())
}

func (f *fileDef) sliceUnpacker(g *Group, spec *ast.ArrayType, alias string) {
	/*
		a, ok := in.([]interface{})
		if !ok {
			return value, errors.New("unpacking into slice, value should be an array")
		}
		<if slice type...>
			var out = make(<name or spec>, len(a))
		<if array type...>
			var out <name or spec>
			if len(a) > <len expr> {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), <len expr>)
			}
		<endif>
		for i, v := range a {
			stack := stack.Append(frizz.ArrayItem(i))
			u, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, err
			}
			out[i] = u
		}
		return out<"[:]" if slice type>, nil
	*/
	g.Comment("sliceUnpacker")
	g.List(Id("a"), Id("ok")).Op(":=").Id("in").Assert(Index().Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into slice, value should be an array"),
			),
		),
	)

	var def *Statement
	if alias != "" {
		def = Id(alias)
	} else {
		def = f.jast.Expr(spec)
	}

	if spec.Len == nil {
		// return type is a slice type, so we use make to get a slice
		// of the correct size for the data
		g.Var().Id("out").Op("=").Make(def, Len(Id("a")))
	} else {
		// return type is an array type, so we create an array of the
		// same type
		g.Var().Id("out").Add(def)
	}

	if spec.Len != nil {
		// check to make sure the data fits
		g.If(Len(Id("a")).Op(">").Add(f.jast.Expr(spec.Len))).Block(
			Return(
				Id("value"),
				Qual("github.com/pkg/errors", "Errorf").Call(
					Lit("data length %d does not fit in array of length %d"),
					Len(Id("a")),
					f.jast.Expr(spec.Len),
				),
			),
		)
	}
	g.For(List(Id("i"), Id("v")).Op(":=").Range().Id("a")).Block(
		Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "ArrayItem").Parens(Id("i"))),
		List(Id("u"), Err()).Op(":=").Add(f.unpacker(spec.Elt, "", "", false)).Call(Id("root"), Id("stack"), Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), Err()),
		),
		Id("out").Index(Id("i")).Op("=").Id("u"),
	)
	g.Return(Id("out").Do(func(s *Statement) {
		if spec.Len == nil {
			// return type is a slice type, so we should take a slice of the array
			s.Index(Empty(), Empty())
		}
	}), Nil())
}

func (f *fileDef) structUnpacker(g *Group, spec *ast.StructType, alias string) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out <alias or type>
		<fields...>
		if v, ok := m["<jsonName>"]; ok {
			stack := stack.Append(frizz.FieldItem("<jsonName>"))
			u, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, err
			}
			out.<fieldName> = u
		}
		</fields>
		return out, nil
	*/
	g.Comment("structUnpacker")
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into struct, value should be a map"),
			),
		),
	)
	g.Var().Id("out").Do(func(s *Statement) {
		if alias != "" {
			s.Id(alias)
		} else {
			s.Add(f.jast.Expr(spec))
		}
	})
	for _, field := range spec.Fields.List {
		var name string
		if len(field.Names) == 0 {
			// anonymous field
			typ := field.Type
			for {
				// remove any number of stars
				star, ok := typ.(*ast.StarExpr)
				if !ok {
					break
				}
				typ = star.X
			}
			switch typ := typ.(type) {
			case *ast.Ident:
				name = typ.Name
			case *ast.SelectorExpr:
				name = typ.Sel.Name
			default:
				panic(fmt.Sprintf("can't find name for anonymous field %#v", typ))
			}
		} else {
			// named field
			name = field.Names[0].Name
		}
		g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(name)), Id("ok")).Block(
			Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "FieldItem").Parens(Lit(name))),
			List(Id("u"), Err()).Op(":=").Add(f.unpacker(field.Type, "", "", false)).Call(Id("root"), Id("stack"), Id("v")),
			If(Err().Op("!=").Nil()).Block(
				Return(Id("value"), Err()),
			),
			Id("out").Dot(name).Op("=").Id("u"),
		)

	}
	g.Return(Id("out"), Nil())
}

func (f *fileDef) nativeUnpacker(g *Group, spec *ast.Ident, alias string) {
	/*
		out, err := frizz.Unpack<strings.Title(spec.Name)>(stack, in)
		if err != nil {
			return nil, err
		}
		return <alias?>(out), nil
	*/
	g.Comment("nativeUnpacker")
	g.List(Id("out"), Err()).Op(":=").Qual("frizz.io/frizz", fmt.Sprintf("Unpack%s", strings.Title(spec.Name))).Call(Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if alias == "" {
			s.Id("out")
		} else {
			s.Id(alias).Parens(Id("out"))
		}
	}), Nil())
}
