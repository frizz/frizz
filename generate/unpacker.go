package generate

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func (p *packageDef) unpacker(spec ast.Expr, name string, method bool, custom bool) *Statement {
	/**
	func <if method>(p packageType) Unpack<name></if>(context global.DataContext, in interface{}) (value <named or spec>, null bool, err error) {
		<...>
	}
	*/
	return Func().Do(func(s *Statement) {
		if method {
			s.Params(Id("p").Id("packageType")).Id("Unpack" + name)
		}
	}).Params(
		Id("context").Qual("frizz.io/global", "DataContext"),
		Id("in").Interface(),
	).Params(
		Id("value").Do(func(s *Statement) {
			if name != "" {
				s.Id(name)
			} else {
				s.Add(p.jast.Expr(spec))
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
			p.customUnpacker(g, name)
			return
		}
		if name != "" {
			p.aliasUnpacker(g, spec, p.Path, name)
			return
		}
		switch spec := spec.(type) {
		default:
			panic(fmt.Sprintf("unsupported type %T", spec))
		case *ast.InterfaceType:
			p.interfaceUnpacker(g, spec)
		case *ast.StarExpr:
			p.pointerUnpacker(g, spec)
		case *ast.MapType:
			p.mapUnpacker(g, spec)
		case *ast.ArrayType:
			p.sliceUnpacker(g, spec)
		case *ast.StructType:
			p.structUnpacker(g, spec)
		case *ast.Ident:
			switch spec.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				p.nativeUnpacker(g, spec)
			default:
				p.localUnpacker(g, spec)
			}
		case *ast.SelectorExpr:
			p.selectorUnpacker(g, spec)
		}
	})
}

func (p *packageDef) aliasUnpacker(g *Group, spec ast.Expr, pkg, name string) {
	/*
		out, null, err := <unpacker>(context, root, stack, in)
		if err != nil || null {
			return value, null, err
		}
		return <name>(out), false, nil
	*/
	g.Comment("aliasUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Add(p.unpacker(spec, "", false, false)).Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
		Return(Id("value"), Id("null"), Err()),
	)
	g.Return(Qual(pkg, name).Parens(Id("out")), False(), Nil())
}

func (p *packageDef) customUnpacker(g *Group, name string) {
	/*
		out := new(<name>)
		null, err = out.Unpack(context, in)
		if err != nil || null {
			return value, null, err
		}
		return *out, false, nil
	*/
	g.Comment("customUnpacker")
	g.Id("out").Op(":=").New(Id(name))
	g.List(Id("null"), Err()).Op("=").Id("out").Dot("Unpack").Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
		Return(Id("value"), Id("null"), Err()),
	)
	g.Return(Op("*").Id("out"), False(), Nil())
}

func (p *packageDef) interfaceUnpacker(g *Group, spec *ast.InterfaceType) {
	/*
		out, null, err := pack.UnpackInterface(context, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(<spec>)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into interface, type %T does not implement interface", context.Location(), out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	*/
	g.Comment("interfaceUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Qual("frizz.io/pack", "UnpackInterface").Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("value"), False(), Err()),
	)
	g.List(Id("iface"), Id("ok")).Op(":=").Id("out").Assert(p.jast.Expr(spec))
	g.If(Op("!").Id("ok")).Block(
		Return(
			Id("value"),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: unpacking into interface, type %T does not implement interface"),
				Id("context").Dot("Location").Call(),
				Id("out"),
			),
		),
	)
	g.If(Id("null")).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Return(Id("iface"), False(), Nil())
}

func (p *packageDef) selectorUnpacker(g *Group, spec *ast.SelectorExpr) {

	ob := p.info.Uses[spec.Sel]
	pkg := ob.Pkg().Path()
	name := ob.Name()

	if pkg == "encoding/json" && name == "Number" {
		// special case for json.Number
		/*
			out, ok := in.(json.Number)
			if !ok {
				return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", context.Location(), in)
			}
			if out == "" {
				return value, true, nil
			}
			return out, false, nil
		*/
		g.Comment("selectorUnpacker (json.Number)")
		g.List(Id("out"), Id("ok")).Op(":=").Id("in").Assert(Qual("encoding/json", "Number"))
		g.If(Op("!").Id("ok")).Block(
			Return(
				Id("value"),
				False(),
				Qual("github.com/pkg/errors", "Errorf").Call(
					Lit("%s: unpacking into json.Number, found %T"),
					Id("context").Dot("Location").Call(),
					Id("in"),
				),
			),
		)
		g.If(Id("out").Op("==").Lit("")).Block(
			Return(Id("value"), True(), Nil()),
		)
		g.Return(Id("out"), False(), Nil())
		return
	}

	if fpkg, ok := p.scan.Packages[pkg]; ok {
		if _, ok := fpkg.types[name]; ok {
			// current type is a frizz type
			/*
				out, null, err := <pkg>.Unpack<name>(context, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			*/
			g.Comment("selectorUnpacker (frizz type)")
			g.List(Id("out"), Id("null"), Err()).Op(":=").Qual(pkg, "Package").Dot("Unpack"+name).Call(Id("context"), Id("in"))
			g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
				Return(Id("value"), Id("null"), Err()),
			)
			g.Return(Id("out"), False(), Nil())
			return
		}
	}

	if types.Implements(types.NewPointer(ob.Type()), p.scan.unmarshaler) {
		// type implements json.Unmarshaler
		/*
			b, err := json.Marshal(in)
			if err != nil {
				return value, false, err
			}
			if err := value.UnmarshalJSON(b); err != nil {
				return value, false, err
			}
			return value, false, nil
		*/
		g.Comment("selectorUnpacker (json.Unmarshaler)")
		g.List(Id("b"), Err()).Op(":=").Qual("encoding/json", "Marshal").Call(Id("in"))
		g.If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), False(), Err()),
		)
		g.If(
			Err().Op(":=").Id("value").Dot("UnmarshalJSON").Call(Id("b")),
			Err().Op("!=").Nil(),
		).Block(
			Return(Id("value"), False(), Err()),
		)
		g.Return(Id("value"), False(), Nil())
		return
	}

	if b, ok := ob.Type().Underlying().(*types.Basic); ok {
		// underlying type is basic, so can use aliasUnpacker?
		g.Comment("selectorUnpacker (basic type)")
		p.aliasUnpacker(g, &ast.Ident{Name: b.String()}, ob.Pkg().Path(), ob.Name())
		return
	}

	panic(fmt.Sprintf("Can't unpack %s.%s", pkg, name))
}

func (p *packageDef) localUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, null, err := p.Unpack<spec name>(context, in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	*/
	g.Comment("localUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Id("p").Dot("Unpack"+spec.Name).Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
		Return(Id("value"), Id("null"), Err()),
	)
	g.Return(Id("out"), False(), Nil())
}

func (p *packageDef) pointerUnpacker(g *Group, spec *ast.StarExpr) {
	/*
		out, null, err := <unpacker>(context, in)
		if err != nil || null {
			return value, null, err
		}
		return &out, false, nil
	*/
	g.Comment("pointerUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Add(p.unpacker(spec.X, "", false, false)).Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
		Return(Id("value"), Id("null"), Err()),
	)
	g.Return(Op("&").Id("out"), False(), Nil())
}

func (p *packageDef) mapUnpacker(g *Group, spec *ast.MapType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into map, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out = make(<spec>, len(m))
		for k, v := range m {
			context := context.New(context.Location().Child(global.MapItem(k)))
			u, null, err := <unpacker>(context, v)
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
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("unpacking into map, value should be a map, found: %#v"),
				Id("context").Dot("Location").Call(),
				Id("in"),
			),
		),
	)
	g.If(Len(Id("m")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Var().Id("out").Op("=").Make(p.jast.Expr(spec), Len(Id("m")))
	g.For(List(Id("k"), Id("v")).Op(":=").Range().Id("m")).Block(
		Id("context").Op(":=").Id("context").Dot("New").Call(
			Id("context").Dot("Location").Call().Dot("Child").Call(
				Qual("frizz.io/global", "MapItem").Parens(Id("k")),
			),
		),
		List(Id("u"), Id("null"), Err()).Op(":=").Add(p.unpacker(spec.Value, "", false, false)).Call(
			Id("context"),
			Id("v"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("value"), False(), Err()),
		),
		If(Op("!").Id("null")).Block(
			Id("out").Index(Id("k")).Op("=").Id("u"),
		),
	)
	g.Return(Id("out"), False(), Nil())
}

func (p *packageDef) sliceUnpacker(g *Group, spec *ast.ArrayType) {
	/*
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array, found: %#v", context.Location(), in)
		}
		if len(a) == 0 {
			return value, true, nil
		}
		<if slice type...>
			var out = make(<spec>, len(a))
		<if array type...>
			var out <name or spec>
			if len(a) > <len expr> {
				return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), <len expr>)
			}
		<endif>
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := <unpacker>(context, v)
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
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: unpacking into slice, value should be an array, found: %#v"),
				Id("context").Dot("Location").Call(),
				Id("in"),
			),
		),
	)
	g.If(Len(Id("a")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	if spec.Len == nil {
		// return type is a slice type, so we use make to get a slice
		// of the correct size for the data
		g.Var().Id("out").Op("=").Make(p.jast.Expr(spec), Len(Id("a")))
	} else {
		// return type is an array type, so we create an array of the
		// same type
		g.Var().Id("out").Add(p.jast.Expr(spec))
	}

	if spec.Len != nil {
		// check to make sure the data fits
		g.If(Len(Id("a")).Op(">").Add(p.jast.Expr(spec.Len))).Block(
			Return(
				Id("value"),
				False(),
				Qual("github.com/pkg/errors", "Errorf").Call(
					Lit("%s: data length %d does not fit in array of length %d"),
					Id("context").Dot("Location").Call(),
					Len(Id("a")),
					p.jast.Expr(spec.Len),
				),
			),
		)
	}
	g.For(List(Id("i"), Id("v")).Op(":=").Range().Id("a")).Block(
		Id("context").Op(":=").Id("context").Dot("New").Call(
			Id("context").Dot("Location").Call().Dot("Child").Call(
				Qual("frizz.io/global", "ArrayItem").Parens(Id("i")),
			),
		),
		List(Id("u"), Id("null"), Err()).Op(":=").Add(p.unpacker(spec.Elt, "", false, false)).Call(
			Id("context"),
			Id("v"),
		),
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

func (p *packageDef) structUnpacker(g *Group, spec *ast.StructType) {
	/*
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out <spec>
		<fields...>
		if v, ok := m["<field name>"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("<field name>")))
			u, null, err := <unpacker>(context, v)
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
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: unpacking into struct, value should be a map, found: %#v"),
				Id("context").Dot("Location").Call(),
				Id("in"),
			),
		),
	)
	g.If(Len(Id("m")).Op("==").Lit(0)).Block(
		Return(Id("value"), True(), Nil()),
	)
	g.Var().Id("out").Add(p.jast.Expr(spec))
	for _, field := range spec.Fields.List {
		g.If(List(Id("v"), Id("ok")).Op(":=").Id("m").Index(Lit(fieldName(field))), Id("ok")).Block(
			Id("context").Op(":=").Id("context").Dot("New").Call(
				Id("context").Dot("Location").Call().Dot("Child").Call(
					Qual("frizz.io/global", "FieldItem").Parens(Lit(fieldName(field))),
				),
			),
			List(Id("u"), Id("null"), Err()).Op(":=").Add(p.unpacker(field.Type, "", false, false)).Call(
				Id("context"),
				Id("v"),
			),
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

func (p *packageDef) nativeUnpacker(g *Group, spec *ast.Ident) {
	/*
		out, null, err := pack.Unpack<strings.Title(spec.Name)>(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, nil
	*/
	g.Comment("nativeUnpacker")
	g.List(Id("out"), Id("null"), Err()).Op(":=").Qual("frizz.io/pack", fmt.Sprintf("Unpack%s", strings.Title(spec.Name))).Call(Id("context").Dot("Location").Call(), Id("in"))
	g.If(Err().Op("!=").Nil().Op("||").Id("null")).Block(
		Return(Id("value"), Id("null"), Err()),
	)
	g.Return(Id("out"), False(), Nil())
}
