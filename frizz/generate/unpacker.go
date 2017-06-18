package generate

import (
	"fmt"
	"go/ast"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func (f *fileDef) unpacker(spec ast.Expr, name string, method bool, custom bool) *Statement {
	/**
	func <if method>(Packer) Unpack<name></if>(root *frizz.Root, stack frizz.Stack, in interface{}) (value <named or spec>, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if method {
			s.Params(Id("p").Id("Packer")).Id("Unpack" + name)
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
		if custom {
			/*
				out := new(<name>)
				if err := out.Unpack(root, stack, in); err != nil {
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

func (f *fileDef) interfaceUnpacker(g *Group, spec *ast.InterfaceType, name string) {
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
		if name != "" {
			s.Id(name)
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

func (f *fileDef) selectorUnpacker(g *Group, spec *ast.SelectorExpr, name string) {
	/*
		out, err := <spec.X>.Packer{}.Unpack<name>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(out), nil
	*/
	g.Comment("selectorUnpacker")
	g.List(Id("out"), Err()).Op(":=").Do(func(s *Statement) {
		x, ok := spec.X.(*ast.Ident)
		if !ok {
			panic("spec.X must be *ast.Ident")
		}
		if x.Obj != nil {
			panic("x.Obj must be nil")
		}
		path, ok := f.imports[x.Name]
		if !ok {
			panic(fmt.Sprintf("%s not found in imports", x.Name))
		}
		s.Qual(path, "Packer")
	}).Values().Dot("Unpack"+spec.Sel.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if name == "" {
			s.Id("out")
		} else {
			s.Id(name).Parens(Id("out"))
		}
	}), Nil())
}

func (f *fileDef) localUnpacker(g *Group, spec *ast.Ident, name string) {
	/*
		out, err := p.Unpack<name>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(out), nil
	*/
	g.Comment("localUnpacker")
	g.List(Id("out"), Err()).Op(":=").Id("p").Dot("Unpack"+spec.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if name == "" {
			s.Id("out")
		} else {
			s.Id(name).Parens(Id("out"))
		}
	}), Nil())
}

func (f *fileDef) pointerUnpacker(g *Group, spec *ast.StarExpr, name string) {
	/*
		out, err := <unpacker>(root, stack, in)
		if err != nil {
			return value, err
		}
		return <alias?>(&out), nil
	*/
	g.Comment("pointerUnpacker")
	g.List(Id("out"), Err()).Op(":=").Add(f.unpacker(spec.X, "", false, false)).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), Err()),
	)
	g.Return(Do(func(s *Statement) {
		if name == "" {
			s.Op("&").Id("out")
		} else {
			s.Id(name).Parens(Op("&").Id("out"))
		}
	}), Nil())
}

func (f *fileDef) mapUnpacker(g *Group, spec *ast.MapType, name string) {
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
		if name != "" {
			s.Id(name)
		} else {
			s.Add(f.jast.Expr(spec))
		}
	}), Len(Id("m")))
	g.For(List(Id("k"), Id("v")).Op(":=").Range().Id("m")).Block(
		Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "MapItem").Parens(Id("k"))),
		List(Id("u"), Err()).Op(":=").Add(f.unpacker(spec.Value, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), Err()),
		),
		Id("out").Index(Id("k")).Op("=").Id("u"),
	)
	g.Return(Id("out"), Nil())
}

func (f *fileDef) sliceUnpacker(g *Group, spec *ast.ArrayType, name string) {
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
	if name != "" {
		def = Id(name)
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
		List(Id("u"), Err()).Op(":=").Add(f.unpacker(spec.Elt, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
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

func (f *fileDef) structUnpacker(g *Group, spec *ast.StructType, name string) {
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
		if name != "" {
			s.Id(name)
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
			List(Id("u"), Err()).Op(":=").Add(f.unpacker(field.Type, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
			If(Err().Op("!=").Nil()).Block(
				Return(Id("value"), Err()),
			),
			Id("out").Dot(name).Op("=").Id("u"),
		)

	}
	g.Return(Id("out"), Nil())
}

func (f *fileDef) nativeUnpacker(g *Group, spec *ast.Ident, name string) {
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
		if name == "" {
			s.Id("out")
		} else {
			s.Id(name).Parens(Id("out"))
		}
	}), Nil())
}
