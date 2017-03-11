package generate

import (
	"context"

	"encoding/json"

	"fmt"

	"sort"

	"bytes"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/process/generate/builder"
	"kego.io/system"
)

func Structs(ctx context.Context, env *envctx.Env) (source []byte, err error) {

	scache := sysctx.FromContext(ctx)

	pcache, ok := scache.Get(env.Path)
	if !ok {
		return nil, kerr.New("DQVQWTKRSK", "%s not found in sys ctx", env.Path)
	}
	types := pcache.Types
	exports := pcache.Exports

	infoBytes, _ := json.Marshal(InfoStruct{Path: env.Path, Hash: env.Hash})

	f := NewFilePath(env.Path)
	f.PackageComment("info:" + string(infoBytes))
	f.Comment(`ke: {"file": {"notest": true}}`)
	f.Line()

	for _, name := range types.Keys() {
		t, ok := types.Get(name)
		if !ok {
			// ke: {"block": {"notest": true}}
			continue
		}
		typ := t.Type.(*system.Type)

		isRule := typ.Id.IsRule()
		isNativeCollection := typ.IsNativeCollection() && typ.Alias == nil

		if !typ.Interface && !typ.Custom {
			if typ.Alias != nil {
				if err := printAliasDefinition(ctx, env, f, typ); err != nil {
					return nil, kerr.Wrap("TRERIECOEP", err)
				}
				if kind, _ := typ.Kind(ctx); kind == system.KindValue {
					if err := printValueMethod(ctx, env, f, typ); err != nil {
						return nil, kerr.Wrap("PGDUJQVQGR", err)
					}
				}
			} else {
				if err := printStructDefinition(ctx, env, f, typ); err != nil {
					return nil, kerr.Wrap("XKRYMXUIJD", err)
				}
			}
		}

		if !typ.Interface && !isRule && !isNativeCollection {
			printInterfaceDefinition(ctx, env, f, typ)
			printInterfaceImplementation(ctx, env, f, typ)
		}

		if !isRule && !isNativeCollection {
			printInterfaceUnpacker(ctx, env, f, typ)
		}

		if !typ.Custom && !typ.Interface && !isNativeCollection {
			if err := printUnpacker(ctx, env, f, typ); err != nil {
				return nil, kerr.Wrap("YJNWUAUKXI", err)
			}
			if err := printRepacker(ctx, env, f, typ); err != nil {
				return nil, kerr.Wrap("NCFFXUHYNY", err)
			}
		}

	}
	printInitFunction(ctx, env, f, types)

	for _, name := range exports.Keys() {
		export, ok := exports.Get(name)
		if !ok {
			// ke: {"block": {"notest": true}}
			continue
		}
		if err := printExportFunction(ctx, env, f, export); err != nil {
			return nil, kerr.Wrap("YJHRXIASNO", err)
		}
	}

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("XKYHSDKBEP", err)
	}
	return buf.Bytes(), nil
}

func printExportFunction(ctx context.Context, env *envctx.Env, f *File, export *sysctx.SysExportInfo) error {
	/*
		func {export.Name}() *{typ} {
			ctx := system.NewContext(context.Background, "{env.Path}", {env.Aliases})
			o := new({typ})
			if err := o.Unpack(ctx, system.MustPackString("{export.JsonContents}"), false); err != nil {
				panic(err.Error())
			}
			return o
		}
	*/
	typ := Qual(export.TypePackage, system.GoName(export.TypeName))
	f.Func().Id(export.Name).Params().Op("*").Add(typ).Block(
		Id("ctx").Op(":=").Qual("kego.io/system", "NewContext").Call(
			Qual("context", "Background"),
			Lit(env.Path),
			Lit(env.Aliases),
		),
		Id("o").Op(":=").New(typ),
		If(
			Err().Op(":=").Id("o").Dot("Unpack").Call(
				Id("ctx"),
				Qual("kego.io/system", "MustPackString").Call(
					Lit(string(export.JsonContents)),
				),
				False(),
			),
			Err().Op("!=").Nil(),
		).Block(
			Panic(Err().Dot("Error").Call()),
		),
		Return(Id("o")),
	)
	return nil
}

func printRepacker(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) error {
	name := system.GoName(typ.Id.Name)

	var ptr, vIn *Statement
	if typ.PassedAsPointer(ctx) {
		ptr = Op("*")
		vIn = Parens(Op("*").Id("v"))
	} else {
		vIn = Id("v")
	}

	/*
		func (v [*]{name}) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {

		}
	*/

	var outerErr error
	f.Func().Params(
		Id("v").Add(ptr).Id(name),
	).Id("Repack").Params(
		Id("ctx").Qual("context", "Context"),
	).Params(
		Id("data").Interface(),
		Id("typePackage").String(),
		Id("typeName").String(),
		Id("jsonType").Qual("kego.io/system", "JsonType"),
		Err().Error(),
	).BlockFunc(func(g *Group) {
		/*
			if v == nil {
				return nil, "{typ.Id.Package}", "{typ.Id.Name}", system.J_NULL, nil
			}
		*/
		g.If(Id("v").Op("==").Nil()).Block(
			Return(
				Nil(),
				Lit(typ.Id.Package),
				Lit(typ.Id.Name),
				Qual("kego.io/system", "J_NULL"),
				Nil(),
			),
		)

		jtype := typ.NativeJsonType(ctx)
		var jsonQual Code
		switch jtype {
		case system.J_NUMBER:
			jsonQual = Qual("kego.io/system", "J_NUMBER")
		case system.J_STRING:
			jsonQual = Qual("kego.io/system", "J_STRING")
		case system.J_BOOL:
			jsonQual = Qual("kego.io/system", "J_BOOL")
		case system.J_MAP:
			jsonQual = Qual("kego.io/system", "J_MAP")
		case system.J_OBJECT:
			jsonQual = Qual("kego.io/system", "J_OBJECT")
		case system.J_ARRAY:
			jsonQual = Qual("kego.io/system", "J_ARRAY")
		case system.J_NULL:
			jsonQual = Qual("kego.io/system", "J_NULL")
		}

		kind, _ := typ.Kind(ctx)
		switch kind {
		case system.KindStruct:
			/*
				m := map[string]interface{}{}
			*/
			g.Id("m").Op(":=").Map(String()).Interface().Values()

			structType := typ
			if typ.Alias != nil {
				structType = system.WrapRule(ctx, typ.Alias).Parent
			}
			for _, embedRef := range structType.AllEmbeds() {
				embedName := system.GoName(embedRef.Name)

				/*
					if v.{embedName} != nil {
						ob, _, _, _, err := v.{embedName}.Repack(ctx)
						if err != nil {
							return nil, "", "", "", err
						}
						for key, val := range ob.(map[string]interface{}) {
							m[key] = val
						}
					}
				*/

				g.If(Id("v").Dot(embedName).Op("!=").Nil()).Block(
					List(Id("ob"), Id("_"), Id("_"), Id("_"), Err()).Op(":=").Id("v").Dot(embedName).Dot("Repack").Call(Id("ctx")),
					If(Err().Op("!=").Nil()).Block(
						Return(Nil(), Lit(""), Lit(""), Lit(""), Err()),
					),
					For(List(Id("key"), Id("val")).Op(":=").Range().Id("ob").Assert(Map(String()).Interface())).Block(
						Id("m").Index(Id("key")).Op("=").Id("val"),
					),
				)
			}
			for _, f := range structType.SortedFields() {
				fieldRule := system.WrapRule(ctx, f.Rule)
				fieldName := system.GoName(f.Name)
				fieldType := fieldRule.Parent
				kind, alias := fieldRule.Kind(ctx)
				switch {
				case kind == system.KindStruct || alias:

					/*
						if v.{fieldName} != nil {
							{repackCode}
							m["{f.Name}"] = ob0
						}
					*/
					g.If(Id("v").Dot(fieldName).Op("!=").Nil()).BlockFunc(func(g *Group) {
						if err := printRepackCode(ctx, env, g, Id("v").Dot(fieldName), "ob0", 0, f.Rule, true); err != nil {
							outerErr = kerr.Wrap("WSARHJIFHS", err)
							return
						}
						g.Id("m").Index(Lit(f.Name)).Op("=").Id("ob0")
					})

				case kind == system.KindValue:

					var lit Code
					switch fieldType.NativeJsonType(ctx) {
					case system.J_STRING:
						lit = Lit("")
					case system.J_NUMBER:
						lit = Lit(0.0)
					case system.J_BOOL:
						lit = Lit(false)
					}

					/*
						if v.{fieldName} != {lit} {
							{repackCode}
							m["{f.Name}"] = ob0
						}
					*/

					g.If(Id("v").Dot(fieldName).Op("!=").Add(lit)).BlockFunc(func(g *Group) {
						if err := printRepackCode(ctx, env, g, Id("v").Dot(fieldName), "ob0", 0, f.Rule, true); err != nil {
							outerErr = kerr.Wrap("YYDYVIMXPM", err)
							return
						}
						g.Id("m").Index(Lit(f.Name)).Op("=").Id("ob0")
					})

				case kind == system.KindArray ||
					kind == system.KindMap ||
					kind == system.KindInterface:

					/*
						if v.{fieldName} != nil {
							{repackCode}
							m["{f.Name}"] = ob0
						}
					*/

					g.If(Id("v").Dot(fieldName).Op("!=").Nil()).BlockFunc(func(g *Group) {
						if err := printRepackCode(ctx, env, g, Id("v").Dot(fieldName), "ob0", 0, f.Rule, true); err != nil {
							outerErr = kerr.Wrap("YSFPHQTBNA", err)
							return
						}
						g.Id("m").Index(Lit(f.Name)).Op("=").Id("ob0")
					})
				}
			}
			/*
				return m, "{typ.Id.Package}", "{typ.Id.Name}", {jsonType}, nil
			*/
			g.Return(
				Id("m"),
				Lit(typ.Id.Package),
				Lit(typ.Id.Name),
				jsonQual,
				Nil(),
			)
		case system.KindValue:

			outer := typ
			inner := typ
			if typ.Alias != nil {
				inner = system.WrapRule(ctx, typ.Alias).Parent
			}
			switch inner.NativeJsonType(ctx) {
			case system.J_STRING:
				/*
					if v != nil {
						return string([*]v), "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
					}
					return nil, "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
				*/
				g.If(Id("v").Op("!=").Nil()).Block(
					Return(
						String().Parens(Add(ptr).Id("v")),
						Lit(outer.Id.Package),
						Lit(outer.Id.Name),
						jsonQual,
						Nil(),
					),
				)
				g.Return(
					Nil(),
					Lit(outer.Id.Package),
					Lit(outer.Id.Name),
					jsonQual,
					Nil(),
				)

			case system.J_NUMBER:
				/*
					if v != nil {
						return float64([*]v), "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
					}
					return nil, "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
				*/
				g.If(Id("v").Op("!=").Nil()).Block(
					Return(
						Float64().Parens(Add(ptr).Id("v")),
						Lit(outer.Id.Package),
						Lit(outer.Id.Name),
						jsonQual,
						Nil(),
					),
				)
				g.Return(
					Nil(),
					Lit(outer.Id.Package),
					Lit(outer.Id.Name),
					jsonQual,
					Nil(),
				)

			case system.J_BOOL:
				/*
					if v != nil {
						return bool([*]v), "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
					}
					return nil, "{outer.Id.Package}", "{outer.Id.Name}", {jsonType}, nil
				*/
				g.If(Id("v").Op("!=").Nil()).Block(
					Return(
						Bool().Parens(Add(ptr).Id("v")),
						Lit(outer.Id.Package),
						Lit(outer.Id.Name),
						jsonQual,
						Nil(),
					),
				)
				g.Return(
					Nil(),
					Lit(outer.Id.Package),
					Lit(outer.Id.Name),
					jsonQual,
					Nil(),
				)
			}

		case system.KindArray, system.KindMap:

			/*
				{repackCode}
				return ob0, "{typ.Id.Package}", "{typ.Id.Name}", {jsonType}, nil
			*/

			if err := printRepackCode(ctx, env, g, vIn, "ob0", 0, typ.Alias, false); err != nil {
				outerErr = kerr.Wrap("SYKLQKLCEO", err)
				return
			}
			g.Return(
				Id("ob0"),
				Lit(typ.Id.Package),
				Lit(typ.Id.Name),
				jsonQual,
				Nil(),
			)
		}
	})
	if outerErr != nil {
		return outerErr
	}
	return nil
}

func printRepackCode(ctx context.Context, env *envctx.Env, g *Group, in *Statement, out string, depth int, f system.RuleInterface, inStruct bool) error {
	field := system.WrapRule(ctx, f)
	kind, alias := field.Kind(ctx)
	//repackerDef := g.SprintRef("kego.io/system", "Repacker")
	switch {
	case kind == system.KindInterface:
		valueVar := out + "_value"

		/*
			var {out} interface{}
			{valueVar}, pkg, name, typ, err := {in}.(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
		*/

		g.Var().Id(out).Interface()
		g.List(
			Id(valueVar),
			Id("pkg"),
			Id("name"),
			Id("typ"),
			Err(),
		).Op(":=").Add(in).Assert(
			Qual("kego.io/system", "Repacker"),
		).Dot("Repack").Call(Id("ctx"))

		g.If(Err().Op("!=").Nil()).Block(
			Return(Nil(), Lit(""), Lit(""), Lit(""), Err()),
		)

		/*
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "{field.Parent.Id.Package}", "{field.Parent.Id.Name}") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				{out} = map[string]interface{}{
					"type": typeVal,
					"value": {valueVar},
				}
			} else {
				{out} = {valueVar}
			}
		*/

		g.If(Qual("kego.io/system", "ShouldUseExplicitTypeNotation").Call(
			Id("pkg"),
			Id("name"),
			Id("typ"),
			Lit(field.Parent.Id.Package),
			Lit(field.Parent.Id.Name),
		)).Block(
			Id("typRef").Op(":=").Qual("kego.io/system", "NewReference").Call(Id("pkg"), Id("name")),
			List(Id("typeVal"), Err()).Op(":=").Id("typRef").Dot("ValueContext").Call(Id("ctx")),
			If(Err().Op("!=").Nil()).Block(
				Return(Nil(), Lit(""), Lit(""), Lit(""), Err()),
			),
			Id(out).Op("=").Map(String()).Interface().Values(Dict{
				Lit("type"):  Id("typeVal"),
				Lit("value"): Id(valueVar),
			}),
		).Else().Block(
			Id(out).Op("=").Id(valueVar),
		)

	case kind == system.KindStruct || (alias && inStruct):
		/*
			{out}, _, _, _, err := {in}.Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
		*/
		g.List(Id(out), Id("_"), Id("_"), Id("_"), Err()).Op(":=").Add(in).Dot("Repack").Call(Id("ctx"))
		g.If(Err().Op("!=").Nil()).Block(
			Return(Nil(), Lit(""), Lit(""), Lit(""), Err()),
		)

	case kind == system.KindValue:
		/*
			{out} := {in}
		*/
		g.Id(out).Op(":=").Add(in)

	case kind == system.KindArray:

		/*
			{out} := []interface{}{}
			for {iVar} := range {in} {
				{repackCode}
				{out} = append({out}, {childOut})
			}
		*/
		var outerErr error
		iVar := fmt.Sprintf("i%d", depth)
		g.Id(out).Op(":=").Index().Interface().Values()
		g.For(Id(iVar).Op(":=").Range().Add(in)).BlockFunc(func(g *Group) {
			childIn := Add(in).Index(Id(iVar))
			childDepth := depth + 1
			childOut := fmt.Sprintf("ob%d", childDepth)
			childRule, err := field.ItemsRule()
			if err != nil {
				outerErr = kerr.Wrap("VUKWDVGVAT", err)
				return
			}
			if err := printRepackCode(ctx, env, g, childIn, childOut, childDepth, childRule.Interface, true); err != nil {
				outerErr = kerr.Wrap("GDWUQWGFUI", err)
				return
			}
			g.Id(out).Op("=").Append(Id(out), Id(childOut))
		})
		if outerErr != nil {
			return outerErr
		}

	case kind == system.KindMap:
		/*
			{out} := map[string]interface{}{}
			for {iVar} := range {in} {
				{repackCode}
				{out}[{kVar}] = {childOut}
			}
		*/
		var outerErr error
		kVar := fmt.Sprintf("k%d", depth)
		g.Id(out).Op(":=").Map(String()).Interface().Values()
		g.For(Id(kVar).Op(":=").Range().Add(in)).BlockFunc(func(g *Group) {
			childIn := Add(in).Index(Id(kVar))
			childDepth := depth + 1
			childOut := fmt.Sprintf("ob%d", childDepth)
			childRule, err := field.ItemsRule()
			if err != nil {
				outerErr = kerr.Wrap("NYDJVRENGA", err)
				return
			}
			if err := printRepackCode(ctx, env, g, childIn, childOut, childDepth, childRule.Interface, true); err != nil {
				outerErr = kerr.Wrap("VNWOUDMDQC", err)
				return
			}
			g.Id(out).Index(Id(kVar)).Op("=").Id(childOut)
		})
		if outerErr != nil {
			return outerErr
		}
	}
	return nil
}

func printUnpacker(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) error {

	name := system.GoName(typ.Id.Name)

	/*
		func (v *{name}) Unpack(ctx context.Context, in system.Packed, iface bool) error {

		}
	*/

	var outerErr error

	f.Func().Params(Id("v").Op("*").Id(name)).Id("Unpack").Params(
		Id("ctx").Qual("context", "Context"),
		Id("in").Qual("kego.io/system", "Packed"),
		Id("iface").Bool(),
	).Error().BlockFunc(func(g *Group) {
		/*
			if in == nil || in.Type() == system.J_NULL {
				return nil
			}
		*/
		g.If(Id("in").Op("==").Nil().Op("||").Id("in").Dot("Type").Call().Op("==").Qual("kego.io/system", "J_NULL")).Block(
			Return(Nil()),
		)

		kind, _ := typ.Kind(ctx)
		switch kind {
		case system.KindStruct:
			structType := typ
			if typ.Alias != nil {
				structType = system.WrapRule(ctx, typ.Alias).Parent
			}

			for _, embedRef := range structType.AllEmbeds() {
				embedType, ok := system.GetTypeFromCache(ctx, embedRef.Package, embedRef.Name)
				if !ok {
					outerErr = kerr.New("IOEEVJCDPU", "Type %s not found", embedRef.String())
					return
				}
				embedId := system.GoName(embedRef.Name)

				/*
					if v.{embedName} == nil {
						v.{embedName} = new({embedType})
					}
				*/

				g.If(Id("v").Dot(embedId).Op("==").Nil()).Block(
					Id("v").Dot(embedId).Op("=").New(builder.Reference(embedType.Id)),
				)

				/*
					if err := v.{embedName}.Unpack(ctx, in, false); err != nil {
						return err
					}
				*/

				g.If(
					Err().Op(":=").Id("v").Dot(embedId).Dot("Unpack").Call(Id("ctx"), Id("in"), False()),
					Err().Op("!=").Nil(),
				).Block(
					Return(Err()),
				)

				if embedRef.Package == "kego.io/system" && embedRef.Name == "object" {
					/*
						if err := v.Object.InitializeType("{typ.Id.Package}", "{typ.Id.Name}"); err != nil {
							return err
						}
					*/
					g.If(
						Err().Op(":=").Id("v").Dot("Object").Dot("InitializeType").Call(
							Lit(typ.Id.Package),
							Lit(typ.Id.Name),
						),
						Err().Op("!=").Nil(),
					).Block(
						Return(Err()),
					)
				}
			}

			for _, f := range structType.SortedFields() {
				/*
					if field, ok := in.Map()["{f.Name}"]; ok && field.Type() != system.J_NULL {
						{printUnpackCode}
						v.{f.Name} = {out}
					}
				*/
				fieldName := system.GoName(f.Name)
				g.If(
					List(
						Id("field"),
						Id("ok"),
					).Op(":=").Id("in").Dot("Map").Call().Index(Lit(f.Name)),
					Id("ok").Op("&&").Id("field").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_NULL"),
				).BlockFunc(func(g *Group) {
					in := "field"
					out := "ob0"
					depth := 0
					if err := printUnpackCode(ctx, env, g, Id(in), out, depth, f.Rule); err != nil {
						outerErr = kerr.Wrap("QLARKEBDBJ", err)
						return
					}
					g.Id("v").Dot(fieldName).Op("=").Id(out)
				}).Do(func(s *Statement) {
					if dr, ok := f.Rule.(system.DefaultRule); ok && dr.GetDefault() != nil {
						/*
							else {
								{unpackCode}
								v.{f.Name} = {out}
							}
						*/
						s.Else().BlockFunc(func(g *Group) {
							b, err := json.Marshal(dr.GetDefault())
							if err != nil {
								outerErr = kerr.Wrap("DLOUEHXVJF", err)
								return
							}
							in := Qual("kego.io/system", "MustPackString").Call(Lit(string(b)))
							out := "ob0"
							depth := 0
							if err := printUnpackCode(ctx, env, g, in, out, depth, f.Rule); err != nil {
								outerErr = kerr.Wrap("UOWRFWSTNT", err)
								return
							}
							g.Id("v").Dot(fieldName).Op("=").Id(out)
						})
					}
				})
				if outerErr != nil {
					return
				}
			}

		case system.KindValue:
			/*
				if in.Type() == system.J_MAP {
					in = in.Map()["value"]
				}
			*/
			g.If(Id("in").Dot("Type").Call().Op("==").Qual("kego.io/system", "J_MAP")).Block(
				Id("in").Op("=").Id("in").Dot("Map").Call().Index(Lit("value")),
			)

			switch typ.NativeJsonType(ctx) {
			case system.J_BOOL:
				/*
					if in.Type() != system.J_BOOL {
						return fmt.Errorf("Invalid type %s while unpacking a bool.", in.Type())
					}
					*v = {name}(in.Bool())
				*/
				g.If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_BOOL")).Block(
					Return(Qual("fmt", "Errorf").Call(
						Lit("Invalid type %s while unpacking a bool."),
						Id("in").Dot("Type").Call(),
					)),
				)
				g.Op("*").Id("v").Op("=").Id(name).Call(Id("in").Dot("Bool").Call())

			case system.J_STRING:
				/*
					if in.Type() != system.J_STRING {
						return fmt.Errorf("Invalid type %s while unpacking a string.", in.Type())
					}
					*v = {name}(in.String())
				*/
				g.If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_STRING")).Block(
					Return(Qual("fmt", "Errorf").Call(
						Lit("Invalid type %s while unpacking a string."),
						Id("in").Dot("Type").Call(),
					)),
				)
				g.Op("*").Id("v").Op("=").Id(name).Call(Id("in").Dot("String").Call())

			case system.J_NUMBER:
				/*
					if in.Type() != system.J_NUMBER {
						return fmt.Errorf("Invalid type %s while unpacking a number.", in.Type())
					}
					*v = {name}(in.Number())
				*/
				g.If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_NUMBER")).Block(
					Return(Qual("fmt", "Errorf").Call(
						Lit("Invalid type %s while unpacking a number."),
						Id("in").Dot("Type").Call(),
					)),
				)
				g.Op("*").Id("v").Op("=").Id(name).Call(Id("in").Dot("Number").Call())
			default:
				panic(fmt.Sprintf("invalid type kind: %s, json native: %s", kind, typ.NativeJsonType(ctx)))
			}
		case system.KindArray:
			/*
				if in.Type() == system.J_MAP {
					in = in.Map()["value"]
				}
			*/
			g.If(Id("in").Dot("Type").Call().Op("==").Qual("kego.io/system", "J_MAP")).Block(
				Id("in").Op("=").Id("in").Dot("Map").Call().Index(Lit("value")),
			)

			/*
				if in.Type() != system.J_ARRAY {
					return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
				}
			*/
			g.If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_ARRAY")).Block(
				Return(Qual("fmt", "Errorf").Call(
					Lit("Invalid type %s while unpacking an array."),
					Id("in").Dot("Type").Call(),
				)),
			)

			/*
				{unpackCode}
				*v = {out}
			*/
			in := "in"
			out := "ob0"
			depth := 0
			if err := printUnpackCode(ctx, env, g, Id(in), out, depth, typ.Alias); err != nil {
				outerErr = kerr.Wrap("VELYRXXGMC", err)
				return
			}
			g.Op("*").Id("v").Op("=").Id(out)

		case system.KindMap:

			/*
				if iface {
					if in.Type() != system.J_MAP {
						return fmt.Errorf("Invalid type %s while unpacking a map.", in.Type())
					}
					in = in.Map()["value"]
				}
			*/

			g.If(Id("iface")).Block(
				If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_MAP")).Block(
					Return(Qual("fmt", "Errorf").Call(
						Lit("Invalid type %s while unpacking a map."),
						Id("in").Dot("Type").Call(),
					)),
				),
				Id("in").Op("=").Id("in").Dot("Map").Call().Index(Lit("value")),
			)

			/*
				if in.Type() != system.J_MAP {
					return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
				}
			*/

			g.If(Id("in").Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_MAP")).Block(
				Return(Qual("fmt", "Errorf").Call(
					Lit("Invalid type %s while unpacking an array."),
					Id("in").Dot("Type").Call(),
				)),
			)

			/*
				{unpackCode}
				*v = {out}
			*/

			in := "in"
			out := "ob0"
			depth := 0
			if err := printUnpackCode(ctx, env, g, Id(in), out, depth, typ.Alias); err != nil {
				outerErr = kerr.Wrap("TNEIUAFUNY", err)
				return
			}
			g.Op("*").Id("v").Op("=").Id(out)

		}

		g.Return(Nil())

	})
	if outerErr != nil {
		return outerErr
	}
	return nil
}

func printUnpackCode(ctx context.Context, env *envctx.Env, g *Group, in *Statement, out string, depth int, f system.RuleInterface) error {
	field := system.WrapRule(ctx, f)
	fieldType := field.Parent
	kind, alias := field.Kind(ctx)
	switch {
	case kind == system.KindStruct || alias:

		/*
			{out} := [*]new({fieldType.Id})
			if err := {out}.Unpack(ctx, {in}, false); err != nil {
				return err
			}
		*/

		g.Id(out).Op(":=").Do(func(s *Statement) {
			if !fieldType.PassedAsPointer(ctx) {
				s.Op("*")
			}
		}).New(builder.Reference(fieldType.Id))

		g.If(
			Err().Op(":=").Id(out).Dot("Unpack").Call(Id("ctx"), in, False()),
			Err().Op("!=").Nil(),
		).Block(
			Return(Err()),
		)

	case kind == system.KindValue:
		var funcQual Code
		switch fieldType.NativeJsonType(ctx) {
		case system.J_STRING:
			funcQual = Qual("kego.io/system", "UnpackString")
		case system.J_NUMBER:
			funcQual = Qual("kego.io/system", "UnpackNumber")
		case system.J_BOOL:
			funcQual = Qual("kego.io/system", "UnpackBool")
		default:
			return kerr.New("LSGUACQGHB", "Kind == KindValue but native json type==%s", fieldType.NativeJsonType(ctx))
		}
		/*
			{out}, err := system.{funcName}(ctx, {in})
			if err != nil {
				return err
			}
		*/

		g.List(Id(out), Err()).Op(":=").Add(funcQual).Call(Id("ctx"), in)
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)

	case kind == system.KindInterface:
		var unpackQual Code
		if fieldType.Interface {
			unpackQual = Qual(
				fieldType.Id.Package,
				"Unpack"+system.GoName(fieldType.Id.Name),
			)
		} else {
			unpackQual = Qual(
				fieldType.Id.Package,
				"Unpack"+system.GoInterfaceName(fieldType.Id.Name),
			)
		}
		/*
			{out}, err := {unpackName}(ctx, {in})
			if err != nil {
				return err
			}
		*/
		g.List(Id(out), Err()).Op(":=").Add(unpackQual).Call(Id("ctx"), in)
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)

	case kind == system.KindArray:
		/*
			if {in}.Type() != system.J_ARRAY {
				return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", {in}.Type())")
			}
		*/

		g.If(Add(in).Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_ARRAY")).Block(
			Return(Qual("fmt", "Errorf").Call(
				Lit("Unsupported json type %s found while unpacking into an array."),
				Add(in).Dot("Type").Call(),
			)),
		)

		/*
			{out} := {fieldType}{}
			for {iVar} := range {in}.Array() {
				{unpackCode}
				{out} = append({out}, {childOut})
			}
		*/

		fieldType, err := builder.TypeDefinition(ctx, f)
		if err != nil {
			return kerr.Wrap("UDPMSSFLTW", err)
		}
		iVar := fmt.Sprintf("i%d", depth)

		g.Id(out).Op(":=").Add(fieldType).Values()
		var outerErr error
		g.For(Id(iVar).Op(":=").Range().Add(in).Dot("Array").Call()).BlockFunc(func(block *Group) {
			childRule, err := field.ItemsRule()
			if err != nil {
				outerErr = kerr.Wrap("PJQBRUEHLD", err)
				return
			}
			childDepth := depth + 1
			childIn := Add(in).Dot("Array").Call().Index(Id(iVar))
			childOut := fmt.Sprintf("ob%d", childDepth)
			if err := printUnpackCode(ctx, env, block, childIn, childOut, childDepth, childRule.Interface); err != nil {
				outerErr = kerr.Wrap("XCHEJPDJDS", err)
				return
			}
			block.Id(out).Op("=").Append(Id(out), Id(childOut))
		})
		if outerErr != nil {
			return outerErr
		}

	case kind == system.KindMap:

		/*
			if {in}.Type() != system.J_MAP {
				return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", {in}.Type())")
			}
		*/

		g.If(Add(in).Dot("Type").Call().Op("!=").Qual("kego.io/system", "J_MAP")).Block(
			Return(Qual("fmt", "Errorf").Call(
				Lit("Unsupported json type %s found while unpacking into a map."),
				Add(in).Dot("Type").Call(),
			)),
		)

		fieldType, err := builder.TypeDefinition(ctx, f)
		if err != nil {
			return kerr.Wrap("IHNYODUIKG", err)
		}
		kVar := fmt.Sprintf("k%d", depth)

		/*
			{out} := {fieldType}{}
			for {kVar} := range {in}.Map() {
				{unpackCode}
				{out}[{kVar}] = {childOut}
			}
		*/

		g.Id(out).Op(":=").Add(fieldType).Values()
		var outerErr error
		g.For(Id(kVar).Op(":=").Range().Add(in).Dot("Map").Call()).BlockFunc(func(g *Group) {
			items, err := field.ItemsRule()
			if err != nil {
				outerErr = kerr.Wrap("FPOBYGVOPP", err)
				return
			}
			childDepth := depth + 1
			childIn := Add(in).Dot("Map").Call().Index(Id(kVar))
			childOut := fmt.Sprintf("ob%d", childDepth)
			if err := printUnpackCode(ctx, env, g, childIn, childOut, childDepth, items.Interface); err != nil {
				outerErr = kerr.Wrap("ONUJJGIKJE", err)
				return
			}
			g.Id(out).Index(Id(kVar)).Op("=").Id(childOut)
		})
		if outerErr != nil {
			return outerErr
		}
	}
	return nil
}

func printInterfaceUnpacker(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) {
	// if type kind is native, also accept bare json native
	// if type kind is array, also accept bare json array
	var interfaceName string
	typeName := system.GoName(typ.Id.Name)
	if typ.Interface {
		interfaceName = system.GoName(typ.Id.Name)
	} else {
		interfaceName = system.GoInterfaceName(typ.Id.Name)
	}
	/*
		func Unpack{interfaceName}(ctx context.Context, in system.Packed) ({interfaceName}, error) {
			switch in.Type() {
			case system.J_MAP:
				i, err := system.UnpackUnknownType(ctx, in, true, strconv.Quote(typ.Id.Package), strconv.Quote(typ.Id.Name)
				if err != nil {
					return nil, err
				}
				ob, ok := i.({interfaceName})
				if !ok {
					return nil, fmt.Errorf("%T does not implement {interfaceName}", i)
				}
				return ob, nil
			case system.{J_STRING|J_NUMBER|J_BOOL|J_ARRAY}:
				ob := new({typeName})
				if err := ob.Unpack(ctx, in, false); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("Unsupported json type %s when unpacking into {interfaceName}.", in.Type())
			}
		}
	*/
	f.Func().Id("Unpack"+interfaceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("in").Qual("kego.io/system", "Packed"),
	).Params(
		Id(interfaceName),
		Error(),
	).Block(
		Switch().Id("in").Dot("Type").Call().Block(
			Case(Qual("kego.io/system", "J_MAP")).Block(
				List(Id("i"), Err()).Op(":=").Qual("kego.io/system", "UnpackUnknownType").Call(
					Id("ctx"),
					Id("in"),
					True(),
					Lit(typ.Id.Package),
					Lit(typ.Id.Name),
				),
				If(Err().Op("!=").Nil()).Block(
					Return(Nil(), Err()),
				),
				List(Id("ob"), Id("ok")).Op(":=").Id("i").Assert(Id(interfaceName)),
				If(Op("!").Id("ok")).Block(
					Return(
						Nil(),
						Qual("fmt", "Errorf").Call(
							Lit("%T does not implement "+interfaceName),
							Id("i"),
						),
					),
				),
				Return(Id("ob"), Nil()),
			),
			Do(func(s *Statement) {
				var qual Code
				switch typ.NativeJsonType(ctx) {
				case system.J_STRING:
					qual = Qual("kego.io/system", "J_STRING")
				case system.J_NUMBER:
					qual = Qual("kego.io/system", "J_NUMBER")
				case system.J_BOOL:
					qual = Qual("kego.io/system", "J_BOOL")
				case system.J_ARRAY:
					qual = Qual("kego.io/system", "J_ARRAY")
				}
				if qual != nil {
					s.Case(qual).Block(
						Id("ob").Op(":=").New(Id(typeName)),
						If(
							Err().Op(":=").Id("ob").Dot("Unpack").Call(
								Id("ctx"),
								Id("in"),
								False(),
							),
							Err().Op("!=").Nil(),
						).Block(
							Return(Nil(), Err()),
						),
						Return(Id("ob"), Nil()),
					)
				}
			}),
			Default().Block(
				Return(
					Nil(),
					Qual("fmt", "Errorf").Call(
						Lit("Unsupported json type %s when unpacking into "+interfaceName+"."),
						Id("in").Dot("Type").Call(),
					),
				),
			),
		),
	)
}

func printInterfaceDefinition(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) {
	/*
		type {typ.Id.Name}Interface interface {
			Get{typ.Id.Name}(ctx context.Context)[*]{typ.Id.Name}
		}
	*/
	name := fmt.Sprint("Get", system.GoName(typ.Id.Name))
	f.Type().Id(system.GoInterfaceName(typ.Id.Name)).Interface(
		Id(name).Params(
			Id("ctx").Qual("context", "Context"),
		).Do(func(s *Statement) {
			if typ.PassedAsPointer(ctx) {
				s.Op("*")
			}
		}).Id(system.GoName(typ.Id.Name)),
	)
}

func printInterfaceImplementation(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) {
	/*
		func (o [*]{typ.Id.Name}) Get{typ.Id.Name}(ctx context.Context) [*]{typ.Id.Name} {
			return o
		}
	*/
	name := fmt.Sprint("Get", system.GoName(typ.Id.Name))
	f.Func().Params(
		Id("o").Do(func(s *Statement) {
			if typ.PassedAsPointer(ctx) {
				s.Op("*")
			}
		}).Id(system.GoName(typ.Id.Name)),
	).Id(name).Params(
		Id("ctx").Qual("context", "Context"),
	).Do(func(s *Statement) {
		if typ.PassedAsPointer(ctx) {
			s.Op("*")
		}
	}).Id(system.GoName(typ.Id.Name)).Block(
		Return(Id("o")),
	)
}

func printAliasDefinition(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) error {
	if typ.Description != "" {
		f.Comment(typ.Description)
	}
	aliasType, err := builder.AliasTypeDefinition(ctx, typ.Alias)
	if err != nil {
		return kerr.Wrap("FWOLIESYUA", err)
	}
	/*
		type {typ.Id.Name} {aliasType}
	*/
	f.Type().Id(system.GoName(typ.Id.Name)).Add(aliasType)
	return nil
}

func printValueMethod(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) error {
	name := system.GoName(typ.Id.Name)
	gotype, err := typ.NativeValueGolangType()
	if err != nil {
		return kerr.Wrap("LUHQRVBQRT", err)
	}
	/*
		func (o *{name}) Value() {gotype} {
			return {gotype}(*o)
		}
	*/
	f.Func().Params(
		Id("o").Op("*").Id(name),
	).Id("Value").Params().Add(gotype).Block(
		Return(
			Add(gotype).Parens(
				Op("*").Id("o"),
			),
		),
	)
	return nil
}

func printStructDefinition(ctx context.Context, env *envctx.Env, f *File, typ *system.Type) error {
	if typ.Description != "" {
		f.Comment(typ.Description)
	}
	var errOuter error
	f.Type().Id(system.GoName(typ.Id.Name)).StructFunc(func(g *Group) {
		if !typ.Basic {
			g.Op("*").Qual("kego.io/system", "Object")
		}
		embedsSortable := system.SortableReferences(typ.Embed)
		sort.Sort(embedsSortable)
		embeds := []*system.Reference(embedsSortable)
		for _, embed := range embeds {
			g.Op("*").Add(builder.Reference(embed))
		}

		for _, nf := range typ.SortedFields() {
			b := nf.Rule.(system.ObjectInterface).GetObject(nil)
			if b.Description != "" {
				g.Comment(b.Description)
			}
			descriptor, err := builder.FieldTypeDefinition(ctx, nf.Name, nf.Rule)
			if err != nil {
				errOuter = kerr.Wrap("GDSKJDEKQD", err)
				return
			}
			g.Id(system.GoName(nf.Name)).Add(descriptor)
		}
	})
	if errOuter != nil {
		return errOuter
	}
	return nil
}

func printInitFunction(ctx context.Context, env *envctx.Env, f *File, types *sysctx.SysTypes) {
	/*
		func init() {
			pkg := jsonctx.InitPackage("{env.Path}")
			pkg.SetHash({env.Hash})
			{...}
		}
	*/
	f.Func().Id("init").Params().BlockFunc(func(g *Group) {
		g.Id("pkg").Op(":=").Qual("kego.io/context/jsonctx", "InitPackage").Call(Lit(env.Path))
		g.Id("pkg").Dot("SetHash").Call(Lit(env.Hash))
		for _, name := range types.Keys() {
			t, ok := types.Get(name)
			if !ok {
				// ke: {"block": {"notest": true}}
				continue
			}
			typ := t.Type.(*system.Type)
			isRule := typ.Id.IsRule()

			if isRule {
				continue
			}

			interfaceFunc := Nil()
			isNativeCollection := typ.IsNativeCollection() && typ.Alias == nil
			if !isNativeCollection {
				var ifaceName string
				if typ.Interface {
					ifaceName = system.GoName(typ.Id.Name)
				} else {
					ifaceName = system.GoInterfaceName(typ.Id.Name)
				}
				/*
					{interfaceFunc:}
					func() reflect.Type {
						return reflect.TypeOf((*{ifaceName})(nil)).Elem()
					}
				*/
				interfaceFunc = Func().Params().Qual("reflect", "Type").Block(
					Return(
						Qual("reflect", "TypeOf").Call(
							Parens(Op("*").Id(ifaceName)).Parens(Nil()),
						).Dot("Elem").Call(),
					),
				)
			}

			newFunc := Nil()
			derefFunc := Nil()
			if typ.Interface {
				/*
					{newFunc:}
					func() interface{} {
						return (*{typ.Id.Name})(nil)
					}
				*/
				newFunc = Func().Params().Interface().Block(
					Return(
						Parens(Op("*").Id(system.GoName(typ.Id.Name))).Parens(Nil()),
					),
				)

			} else if !typ.IsNativeCollection() || typ.Alias != nil {
				/*
					{newFunc:}
					func() interface{} {
						return new({typ.Id.Name})
					}
				*/
				newFunc = Func().Params().Interface().Block(
					Return(
						New(Id(system.GoName(typ.Id.Name))),
					),
				)
				if !typ.PassedAsPointer(ctx) {
					/*
						{derefFunc:}
						func(in interface{}) interface{} {
							return *in.(*{typ.Id.Name})
						}
					*/
					derefFunc = Func().Params(Id("in").Interface()).Interface().Block(
						Return(
							Op("*").Id("in").Assert(Op("*").Id(system.GoName(typ.Id.Name))),
						),
					)
				}
			}
			/*
				{ruleFunc:}
				func() interface{} {
					return new({typ.Id.ChangeToRule().Name})
				}
			*/
			ruleFunc := Func().Params().Interface().Block(
				Return(
					New(Id(system.GoName(typ.Id.ChangeToRule().Name))),
				),
			)

			/*
				pkg.Init(
					"{typ.Id.Name}",
					{newFunc},
					{derefFunc},
					{ruleFunc},
					{interfaceFunc},
				)
			*/
			g.Id("pkg").Dot("Init").Call(
				Lit(typ.Id.Name),
				newFunc,
				derefFunc,
				ruleFunc,
				interfaceFunc,
			)
		}
	})
}

type InfoStruct struct {
	Path string
	Hash uint64
}
