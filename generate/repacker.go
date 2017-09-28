package generate

import (
	"fmt"
	"go/ast"

	. "github.com/dave/jennifer/jen"
)

func (f *fileDef) repacker(spec ast.Expr, name string, method bool, custom bool) *Statement {
	/*
		func <if method>(p packageType) Repack<name></if>(context global.DataContext, in <name or spec>) (value interface{}, dict bool, null bool, err error) {
			<...>
		}
	*/
	return Func().Do(func(s *Statement) {
		if method {
			s.Params(Id("p").Id("packageType")).Id("Repack" + name)
		}
	}).Params(
		Id("context").Qual("frizz.io/global", "DataContext"),
		Id("in").Do(func(s *Statement) {
			if name != "" {
				s.Id(name)
			} else {
				s.Add(f.jast.Expr(spec))
			}
		}),
	).Params(
		Id("value").Interface(),
		Id("dict").Bool(),
		Id("null").Bool(),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		if custom {
			f.customRepacker(g, spec, name)
			return
		}
		if name != "" {
			f.aliasRepacker(g, spec, name)
			return
		}
		switch spec := spec.(type) {
		default:
			panic(fmt.Sprintf("unsupported type %T", spec))
		case *ast.InterfaceType:
			f.interfaceRepacker(g, spec)
		case *ast.StarExpr:
			f.pointerRepacker(g, spec)
		case *ast.MapType:
			f.mapRepacker(g, spec)
		case *ast.ArrayType:
			f.sliceRepacker(g, spec)
		case *ast.StructType:
			f.structRepacker(g, spec)
		case *ast.Ident:
			switch spec.Name {
			case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
				f.nativeRepacker(g, spec)
			default:
				f.localRepacker(g, spec)
			}
		case *ast.SelectorExpr:
			f.selectorRepacker(g, spec)
		}
	})
}

func (f *fileDef) aliasRepacker(g *Group, spec ast.Expr, name string) {
	/*
		return <repacker>(context, (<spec>)(in))
	*/
	g.Return(
		f.repacker(spec, "", false, false).Call(
			Id("context"),
			// Add parens around type
			Parens(f.jast.Expr(spec)).Parens(Id("in")),
		),
	)
}

func (f *fileDef) customRepacker(g *Group, spec ast.Expr, name string) {
	/*
		out, dict, null, err := in.Repack(context)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	*/
	g.Comment("customRepacker")
	g.List(Id("out"), Id("dict"), Id("null"), Err()).Op(":=").Id("in").Dot("Repack").Call(
		Id("context"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Id("null"), Nil())
}

func (f *fileDef) interfaceRepacker(g *Group, spec *ast.InterfaceType) {
	/*
		return pack.RepackInterface(context, false, in)
	*/
	g.Comment("interfaceRepacker")
	g.Return(
		Qual("frizz.io/pack", "RepackInterface").Call(
			Id("context"),
			False(),
			Id("in"),
		),
	)
}

func (f *fileDef) pointerRepacker(g *Group, spec *ast.StarExpr) {
	/*
		if in = nil {
			return nil, false, true, nil
		}
		out, dict, null, err := <unpacker>(context, *in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, false, nil
	*/
	g.Comment("pointerRepacker")
	g.If(Id("in").Op("==").Nil()).Block(
		Return(Nil(), False(), True(), Nil()),
	)
	g.List(Id("out"), Id("dict"), Id("null"), Err()).Op(":=").Add(f.repacker(spec.X, "", false, false)).Call(
		Id("context"),
		Op("*").Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), False(), Nil())
}

func (f *fileDef) mapRepacker(g *Group, spec *ast.MapType) {
	/*
		out := make(map[string]interface{}, <len>)
		for k, item := range in {
			v, _, _, err := <repacker>(context, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	*/
	g.Comment("mapRepacker")
	g.Id("out").Op(":=").Make(Map(String()).Interface(), Len(Id("in")))
	g.For(List(Id("k"), Id("item")).Op(":=").Range().Id("in")).Block(
		List(Id("v"), Id("_"), Id("_"), Err()).Op(":=").Add(f.repacker(spec.Value, "", false, false)).Call(
			Id("context"),
			Id("item"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), False(), Err()),
		),
		Id("out").Index(Id("k")).Op("=").Id("v"),
	)
	g.Return(Id("out"), True(), Len(Id("in")).Op("==").Lit(0), Nil())
}

func (f *fileDef) sliceRepacker(g *Group, spec *ast.ArrayType) {
	/*
		out := make([]interface{}, <len>)
		empty := true
		for i, item := range in {
			v, _, null, err := <repacker>(context, item)
			if err != nil {
				return nil, false, false, err
			}
			if !null {
				empty = false
			}
			out[i] = v
		}
		return out, false, empty, nil
	*/
	g.Comment("sliceRepacker")
	g.Id("out").Op(":=").Make(Index().Interface(), Len(Id("in")))
	g.Id("empty").Op(":=").True()
	g.For(List(Id("i"), Id("item")).Op(":=").Range().Id("in")).Block(
		List(Id("v"), Id("_"), Id("null"), Err()).Op(":=").Add(f.repacker(spec.Elt, "", false, false)).Call(
			Id("context"),
			Id("item"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), False(), Err()),
		),
		If(Op("!").Id("null")).Block(
			Id("empty").Op("=").False(),
		),
		Id("out").Index(Id("i")).Op("=").Id("v"),
	)
	g.Return(Id("out"), False(), Id("empty"), Nil())
}

func (f *fileDef) selectorRepacker(g *Group, spec *ast.SelectorExpr) {
	pkg := f.pathFromSelector(spec)
	if pkg == "encoding/json" && spec.Sel.Name == "Number" {
		g.Comment("selectorRepacker (json.Number)")
		/*
			if in == "" {
				return nil, false, true, nil
			}
			return in, false, false, nil
		*/
		g.If(Id("in").Op("==").Lit("")).Block(
			Return(Id("value"), False(), True(), Nil()),
		)
		g.Return(Id("in"), False(), False(), Nil())
		return
	}

	/*
		p := context.Package().Get("<pkg>")
		if p != nil {
			out, dict, null, err := p.Repack(context, in, "<spec.name>")
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}
	*/
	g.Comment("selectorRepacker")
	g.Id("p").Op(":=").Id("context").Dot("Package").Call().Dot("Get").Call(Lit(pkg))
	g.If(Id("p").Op("!=").Nil()).Block(
		List(Id("out"), Id("dict"), Id("null"), Err()).Op(":=").Id("p").Dot("Repack").Call(
			Id("context"),
			Id("in"),
			Lit(spec.Sel.Name),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), False(), Err()),
		),
		Return(Id("out"), Id("dict"), Id("null"), Nil()),
	)

	/*
		var vi interface{} = &in
		if vu, ok := vi.(json.Marshaler); ok {
			b, err := vu.MarshalJSON()
			if err != nil {
				return value, false, false, err
			}
			if err := json.Unmarshal(b, &value); err != nil {
				return value, false, false, err
			}
			return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
		}
		return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "<pkg>", "<spec.Sel.Name>")
	*/
	g.Var().Id("vi").Interface().Op("=").Op("&").Id("in")
	g.If(List(Id("vu"), Id("ok")).Op(":=").Id("vi").Assert(Qual("encoding/json", "Marshaler")), Id("ok")).Block(
		List(Id("b"), Err()).Op(":=").Id("vu").Dot("MarshalJSON").Call(),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), False(), Err()),
		),
		If(Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("value")), Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), False(), Err()),
		),
		Return(
			Id("value"),
			Id("b").Index(Lit(0)).Op("==").LitRune('{'),
			Len(Id("b")).Op("==").Lit(4).Op("&&").String().Parens(Id("b")).Op("==").Lit("null"),
			Nil(),
		),
	)
	g.Return(
		Id("value"),
		False(),
		False(),
		Qual("github.com/pkg/errors", "Errorf").Call(
			Lit("%s: can't repack %s.%s"),
			Id("context").Dot("Location").Call(),
			Lit(pkg),
			Lit(spec.Sel.Name),
		),
	)
}

func (f *fileDef) localRepacker(g *Group, spec *ast.Ident) {
	/*
		out, dict, null, err := p.Repack<spec name>(context, in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	*/
	g.Comment("localRepacker")
	g.List(Id("out"), Id("dict"), Id("null"), Err()).Op(":=").Id("p").Dot("Repack"+spec.Name).Call(
		Id("context"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Id("null"), Nil())
}

func (f *fileDef) nativeRepacker(g *Group, spec *ast.Ident) {
	/*
		return pack.Repack<type>(in), false, nil
	*/
	g.Comment("nativeRepacker")
	var repackFunction string
	switch spec.Name {
	case "bool":
		repackFunction = "RepackBool"
	case "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64":
		repackFunction = "RepackNumber"
	case "rune", "string":
		repackFunction = "RepackString"
	}
	g.Return(Qual("frizz.io/pack", repackFunction).Call(Id("in")))
}

func (f *fileDef) structRepacker(g *Group, spec *ast.StructType) {
	/*
		out := make(map[string]interface{}, <fields len>)
		empty := true
		<fields...>
		if v, _, null, err := <repacker>(context, in.<field name>); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["<field name>"] = v
			}
		}
		</fields>
		return out, false, empty, nil
	*/
	g.Comment("structRepacker")
	// len +1 to leave room for _type
	g.Id("out").Op(":=").Make(Map(String()).Interface(), Lit(len(spec.Fields.List)+1))
	g.Id("empty").Op(":=").True()
	for _, field := range spec.Fields.List {
		g.If(
			List(Id("v"), Id("_"), Id("null"), Err()).Op(":=").Add(
				f.repacker(field.Type, "", false, false),
			).Call(
				Id("context"),
				Id("in").Dot(fieldName(field)),
			),
			Err().Op("!=").Nil(),
		).Block(
			Return(Nil(), False(), False(), Err()),
		).Else().Block(
			If(Op("!").Id("null")).Block(
				Id("empty").Op("=").False(),
				Id("out").Index(Lit(fieldName(field))).Op("=").Id("v"),
			),
		)
	}
	g.Return(Id("out"), False(), Id("empty"), Nil())
}