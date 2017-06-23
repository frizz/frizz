package generate

import (
	"fmt"
	"go/ast"

	. "github.com/dave/jennifer/jen"
)

func (f *fileDef) repacker(spec ast.Expr, name string, method bool, custom bool) *Statement {
	/*
		func <if method>(p packer) Repack<name></if>(root *frizz.Root, stack frizz.Stack, in <name or spec>) (interface{}, bool, error) {
			<...>
		}
	*/
	return Func().Do(func(s *Statement) {
		if method {
			s.Params(Id("p").Id("packer")).Id("Repack" + name)
		}
	}).Params(
		Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
		Id("stack").Qual("frizz.io/frizz", "Stack"),
		Id("in").Do(func(s *Statement) {
			if name != "" {
				s.Id(name)
			} else {
				s.Add(f.jast.Expr(spec))
			}
		}),
	).Params(
		Interface(),
		Bool(),
		Error(),
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
		return <repacker>(root, stack, (<spec>)(in))
	*/
	g.Return(
		f.repacker(spec, "", false, false).Call(
			Id("root"),
			Id("stack"),
			// Add parens around type
			Parens(f.jast.Expr(spec)).Parens(Id("in")),
		),
	)
}

func (f *fileDef) customRepacker(g *Group, spec ast.Expr, name string) {
	/*
		out, dict, err := in.Repack(root, stack)
		if err != nil {
			return nil, false, err
		}
		return out, dict, nil
	*/
	g.Comment("customRepacker")
	g.List(Id("out"), Id("dict"), Err()).Op(":=").Id("in").Dot("Repack").Call(Id("root"), Id("stack"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Nil())
}

func (f *fileDef) interfaceRepacker(g *Group, spec *ast.InterfaceType) {
	/*
		return frizz.RepackInterface(root, stack, in)
	*/
	g.Comment("interfaceRepacker")
	g.Return(
		Qual("frizz.io/frizz", "RepackInterface").Call(
			Id("root"),
			Id("stack"),
			Id("in"),
		),
	)
}

func (f *fileDef) pointerRepacker(g *Group, spec *ast.StarExpr) {
	/*
		out, dict, err := <unpacker>(root, stack, *in)
		if err != nil {
			return nil, false, err
		}
		return out, dict, nil
	*/
	g.Comment("pointerRepacker")
	g.List(Id("out"), Id("dict"), Err()).Op(":=").Add(f.repacker(spec.X, "", false, false)).Call(
		Id("root"),
		Id("stack"),
		Op("*").Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Nil())
}

func (f *fileDef) mapRepacker(g *Group, spec *ast.MapType) {
	/*
		out := make(map[string]interface{}, <len>)
		for k, item := range in {
			v, _, err := <repacker>(root, stack, item)
			if err != nil {
				return nil, false, err
			}
			out[k] = v
		}
		return out, true, nil
	*/
	g.Comment("mapRepacker")
	g.Id("out").Op(":=").Make(Map(String()).Interface(), Len(Id("in")))
	g.For(List(Id("k"), Id("item")).Op(":=").Range().Id("in")).Block(
		List(Id("v"), Id("_"), Err()).Op(":=").Add(f.repacker(spec.Value, "", false, false)).Call(
			Id("root"),
			Id("stack"),
			Id("item"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), Err()),
		),
		Id("out").Index(Id("k")).Op("=").Id("v"),
	)
	g.Return(Id("out"), True(), Nil())
}

func (f *fileDef) sliceRepacker(g *Group, spec *ast.ArrayType) {
	/*
		out := make([]interface{}, <len>)
		for i, item := range in {
			v, _, err := <repacker>(root, stack, item)
			if err != nil {
				return nil, false, err
			}
			out[i] = v
		}
		return out, false, nil
	*/
	g.Comment("sliceRepacker")
	g.Id("out").Op(":=").Make(Index().Interface(), Len(Id("in")))
	g.For(List(Id("i"), Id("item")).Op(":=").Range().Id("in")).Block(
		List(Id("v"), Id("_"), Err()).Op(":=").Add(f.repacker(spec.Elt, "", false, false)).Call(
			Id("root"),
			Id("stack"),
			Id("item"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), False(), Err()),
		),
		Id("out").Index(Id("i")).Op("=").Id("v"),
	)
	g.Return(Id("out"), False(), Nil())
}

func (f *fileDef) selectorRepacker(g *Group, spec *ast.SelectorExpr) {
	/*
		out, dict, err := <spec.pkg>.Packer.Repack<spec.name>(root, stack, in)
		if err != nil {
			return nil, false, err
		}
		return out, dict, nil
	*/
	g.Comment("selectorRepacker")
	g.List(Id("out"), Id("dict"), Err()).Op(":=").Qual(f.pathFromSelector(spec), "Packer").Dot("Repack"+spec.Sel.Name).Call(
		Id("root"),
		Id("stack"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Nil())
}

func (f *fileDef) localRepacker(g *Group, spec *ast.Ident) {
	/*
		out, dict, err := p.Repack<spec name>(root, stack, in)
		if err != nil {
			return nil, false, err
		}
		return out, dict, nil
	*/
	g.Comment("localRepacker")
	g.List(Id("out"), Id("dict"), Err()).Op(":=").Id("p").Dot("Repack"+spec.Name).Call(
		Id("root"),
		Id("stack"),
		Id("in"),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Nil(), False(), Err()),
	)
	g.Return(Id("out"), Id("dict"), Nil())
}

func (f *fileDef) nativeRepacker(g *Group, spec *ast.Ident) {
	/*
		return system.Repack<type>(in), false, nil
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
	g.Return(Qual("frizz.io/frizz", repackFunction).Call(Id("in")), False(), Nil())
}

func (f *fileDef) structRepacker(g *Group, spec *ast.StructType) {
	/*
		out := make(map[string]interface{}, <fields len>)
		<fields...>
		if v, _, err := <repacker>(root, stack, in.<field name>); err != nil {
			return nil, false, err
		} else {
			out["<field name>"] = v
		}
		</fields>
		return out, false, nil
	*/
	g.Comment("structRepacker")
	// len +1 to leave room for _type
	g.Id("out").Op(":=").Make(Map(String()).Interface(), Lit(len(spec.Fields.List)+1))
	for _, field := range spec.Fields.List {
		g.If(
			List(Id("v"), Id("_"), Err()).Op(":=").Add(f.repacker(field.Type, "", false, false)).Call(Id("root"), Id("stack"), Id("in").Dot(fieldName(field))),
			Err().Op("!=").Nil(),
		).Block(
			Return(Nil(), False(), Err()),
		).Else().Block(
			Id("out").Index(Lit(fieldName(field))).Op("=").Id("v"),
		)
	}
	g.Return(Id("out"), False(), Nil())
}
