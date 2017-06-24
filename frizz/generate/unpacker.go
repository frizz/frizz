package generate

import (
	"fmt"
	"go/ast"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func (f *fileDef) unpacker(spec ast.Expr, name string, method bool, custom bool) *Statement {
	/**
	func <if method>(p packer) Unpack<name></if>(root *frizz.Root, stack frizz.Stack, in interface{}) (value <named or spec>, null bool, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if method {
			s.Params(Id("p").Id("packer")).Id("Unpack" + name)
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
		Id("null").Bool(),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		/*
			if in == nil {
				return value, true, nil
			}
		*/
		g.If(Id("in").Op("==").Nil()).Block(
			Return(Id("value"), True(), Nil()),
		)
		if custom {
			f.customUnpacker(g, name)
			return
		}
		if name != "" {
			f.aliasUnpacker(g, spec, name)
			return
		}
		switch spec := spec.(type) {
		default:
			panic(fmt.Sprintf("unsupported type %T", spec))
		case *ast.InterfaceType:
			f.interfaceUnpacker(g, spec)
		case *ast.StarExpr:
			f.pointerUnpacker(g, spec)
		case *ast.MapType:
			f.mapUnpacker(g, spec)
		case *ast.ArrayType:
			f.sliceUnpacker(g, spec)
		case *ast.StructType:
			f.structUnpacker(g, spec)
		case *ast.Ident:
			switch spec.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				f.nativeUnpacker(g, spec)
			default:
				f.localUnpacker(g, spec)
			}
		case *ast.SelectorExpr:
			f.selectorUnpacker(g, spec)
		}
	})
}

func (f *fileDef) aliasUnpacker(g *Group, spec ast.Expr, name string) {
	/*
		out, null, err := <unpacker>(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return <name>(out), false, nil
	*/
	g.Comment("aliasUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Add(f.unpacker(spec, "", false, false)).Call(
		Id("root"),
		Id("stack"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id(name).Parens(Id("out")), False(), Nil())
}

func (f *fileDef) customUnpacker(g *Group, name string) {
	/*
		out := new(<name>)
		null, err = out.Unpack(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return *out, false, nil
	*/
	g.Comment("customUnpacker")
	g.Id("out").Op(":=").New(Id(name))
	g.List(Id("null"), Err()).Op("=").Id("out").Dot("Unpack").Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Op("*").Id("out"), False(), Nil())
}

func (f *fileDef) interfaceUnpacker(g *Group, spec *ast.InterfaceType) {
	/*
		out, null, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(<spec>)
		if !ok {
			return value, false, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	*/
	g.Comment("interfaceUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Id("root").Dot("UnpackInterface").Call(Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.List(Id("iface"), Id("ok")).Op(":=").Id("out").Assert(f.jast.Expr(spec))
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("unpacking into interface, type %T does not implement interface"),
				Id("out"),
			),
		),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id("iface"), False(), Nil())
}

func (f *fileDef) selectorUnpacker(g *Group, spec *ast.SelectorExpr) {
	/*
		out, null, err := <spec.X>.Packer.Unpack<spec.Sel.Name>(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, false, nil
	*/
	g.Comment("selectorUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Qual(f.pathFromSelector(spec), "Packer").Dot("Unpack"+spec.Sel.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id("out"), False(), Nil())
}

func (f *fileDef) localUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, null, err := p.Unpack<spec name>(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, false, nil
	*/
	g.Comment("localUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Id("p").Dot("Unpack"+spec.Name).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id("out"), False(), Nil())
}

func (f *fileDef) pointerUnpacker(g *Group, spec *ast.StarExpr) {
	/*
		out, null, err := <unpacker>(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return &out, false, nil
	*/
	g.Comment("pointerUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Add(f.unpacker(spec.X, "", false, false)).Call(Id("root"), Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Op("&").Id("out"), False(), Nil())
}

func (f *fileDef) mapUnpacker(g *Group, spec *ast.MapType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into map, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out = make(<spec>, len(m))
		for k, v := range m {
			stack := stack.Append(frizz.MapItem(k))
			u, null, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[k] = u
			}
		}
		return out, false, nil
	*/
	g.Comment("mapUnpacker")
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			False(),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into map, value should be a map"),
			),
		),
	)
	g.If(Len(Id("m")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Var().Id("out").Op("=").Make(f.jast.Expr(spec), Len(Id("m")))
	g.For(List(Id("k"), Id("v")).Op(":=").Range().Id("m")).Block(
		Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "MapItem").Parens(Id("k"))),
		List(Id("u"), Id("null"), Err()).Op(":=").Add(f.unpacker(spec.Value, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), False(), Err()),
		),
		If(Op("!").Id("null")).Block(
			Id("out").Index(Id("k")).Op("=").Id("u"),
		),
	)
	g.Return(Id("out"), False(), Nil())
}

func (f *fileDef) sliceUnpacker(g *Group, spec *ast.ArrayType) {
	/*
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.New("unpacking into slice, value should be an array")
		}
		if len(a) == 0 {
			return value, true, nil
		}
		<if slice type...>
			var out = make(<spec>, len(a))
		<if array type...>
			var out <name or spec>
			if len(a) > <len expr> {
				return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), <len expr>)
			}
		<endif>
		for i, v := range a {
			stack := stack.Append(frizz.ArrayItem(i))
			u, null, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[i] = u
			}
		}
		return out<"[:]" if slice type>, false, nil
	*/
	g.Comment("sliceUnpacker")
	g.List(Id("a"), Id("ok")).Op(":=").Id("in").Assert(Index().Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			False(),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into slice, value should be an array"),
			),
		),
	)
	g.If(Len(Id("a")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	if spec.Len == nil {
		// return type is a slice type, so we use make to get a slice
		// of the correct size for the data
		g.Var().Id("out").Op("=").Make(f.jast.Expr(spec), Len(Id("a")))
	} else {
		// return type is an array type, so we create an array of the
		// same type
		g.Var().Id("out").Add(f.jast.Expr(spec))
	}

	if spec.Len != nil {
		// check to make sure the data fits
		g.If(Len(Id("a")).Op(">").Add(f.jast.Expr(spec.Len))).Block(
			Return(
				Id("value"),
				False(),
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
		List(Id("u"), Id("null"), Err()).Op(":=").Add(f.unpacker(spec.Elt, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), False(), Err()),
		),
		If(Op("!").Id("null")).Block(
			Id("out").Index(Id("i")).Op("=").Id("u"),
		),
	)
	g.Return(
		Id("out").Do(func(s *Statement) {
			if spec.Len == nil {
				// return type is a slice type, so we should take a slice of the array
				s.Index(Empty(), Empty())
			}
		}),
		False(),
		Nil(),
	)
}

func (f *fileDef) structUnpacker(g *Group, spec *ast.StructType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into struct, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out <spec>
		<fields...>
		if v, ok := m["<field name>"]; ok {
			stack := stack.Append(frizz.FieldItem("<field name>"))
			u, null, err := <unpacker>(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.<field name> = u
			}
		}
		</fields>
		return out, false, nil
	*/
	g.Comment("structUnpacker")
	g.List(Id("m"), Id("ok")).Op(":=").Id("in").Assert(Map(String()).Interface())
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			False(),
			Qual("github.com/pkg/errors", "New").Call(
				Lit("unpacking into struct, value should be a map"),
			),
		),
	)
	g.If(Len(Id("m")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Var().Id("out").Add(f.jast.Expr(spec))
	for _, field := range spec.Fields.List {
		g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(fieldName(field))), Id("ok")).Block(
			Id("stack").Op(":=").Id("stack").Dot("Append").Call(Qual("frizz.io/frizz", "FieldItem").Parens(Lit(fieldName(field)))),
			List(Id("u"), Id("null"), Err()).Op(":=").Add(f.unpacker(field.Type, "", false, false)).Call(Id("root"), Id("stack"), Id("v")),
			If(Err().Op("!=").Nil()).Block(
				Return(Id("value"), False(), Err()),
			),
			If(Op("!").Id("null")).Block(
				Id("out").Dot(fieldName(field)).Op("=").Id("u"),
			),
		)
	}
	g.Return(Id("out"), False(), Nil())
}

func (f *fileDef) nativeUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, null, err := frizz.Unpack<strings.Title(spec.Name)>(stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, nil
	*/
	g.Comment("nativeUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Qual("frizz.io/frizz", fmt.Sprintf("Unpack%s", strings.Title(spec.Name))).Call(Id("stack"), Id("in"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id("out"), False(), Nil())
}
