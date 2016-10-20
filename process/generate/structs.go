package generate

import (
	"strconv"

	"sort"

	"context"

	"encoding/json"

	"fmt"

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

	g := builder.New(env.Path)

	infoBytes, _ := json.Marshal(InfoStruct{Path: env.Path, Hash: env.Hash})

	g.SetPackageComment("info:" + string(infoBytes))
	g.SetIntroComment(`ke: {"file": {"notest": true}}`)

	if types.Len() == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.Wrap("BBRLIODBKL", err)
		}
		return b, nil
	}
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
				if err := printAliasDefinition(ctx, env, g, typ); err != nil {
					return nil, kerr.Wrap("TRERIECOEP", err)
				}
			} else {
				if err := printStructDefinition(ctx, env, g, typ); err != nil {
					return nil, kerr.Wrap("XKRYMXUIJD", err)
				}
			}
		}

		if !typ.Interface && !isRule && !isNativeCollection {
			printInterfaceDefinition(ctx, env, g, typ)
			printInterfaceImplementation(ctx, env, g, typ)
		}

		if !isRule && !isNativeCollection {
			printInterfaceUnpacker(ctx, env, g, typ)
		}

		if !typ.Custom && !typ.Interface && !isNativeCollection {
			if err := printUnpacker(ctx, env, g, typ); err != nil {
				return nil, kerr.Wrap("YJNWUAUKXI", err)
			}
			if err := printRepacker(ctx, env, g, typ); err != nil {
				return nil, kerr.Wrap("NCFFXUHYNY", err)
			}
		}

	}
	printInitFunction(ctx, env, g, types)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.Wrap("XKYHSDKBEP", err)
	}
	return b, nil
}

func printRepacker(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) error {
	name := system.GoName(typ.Id.Name)
	contextPkg := g.Imports.Add("context")
	g.Println("func (v *", name, ") Repack(ctx ", contextPkg, ".Context) (data interface{}, typePackage string, typeName string, err error) {")
	{
		g.Println("if v == nil {")
		{
			g.Println("return nil, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
		}
		g.Println("}")
		kind, _ := typ.Kind(ctx)
		switch kind {
		case system.KindStruct:
			g.Println("m := map[string]interface{}{}")
			structType := typ
			if typ.Alias != nil {
				structType = system.WrapRule(ctx, typ.Alias).Parent
			}
			for _, embedRef := range structType.AllEmbeds() {
				embedName := system.GoName(embedRef.Name)
				g.Println("if v.", embedName, " != nil {")
				{
					g.Println("ob, _, _, err := v.", embedName, ".Repack(ctx)")
					g.Println("if err != nil {")
					{
						g.Println(`return nil, "", "", err`)
					}
					g.Println("}")
					g.Println("for key, val := range ob.(map[string]interface{}) {")
					{
						g.Println("m[key] = val")
					}
					g.Println("}")
				}
				g.Println("}")
			}
			for n, f := range structType.Fields {
				fieldRule := system.WrapRule(ctx, f)
				fieldName := system.GoName(n)
				fieldType := fieldRule.Parent
				kind, alias := fieldRule.Kind(ctx)
				switch {
				case kind == system.KindStruct || alias:
					g.Println("if v.", fieldName, " != nil {")
					{
						if err := printRepackCode(ctx, env, g, "v."+fieldName, "ob0", 0, f, true); err != nil {
							return kerr.Wrap("WSARHJIFHS", err)
						}
						g.Println("m[", strconv.Quote(n), "] = ", "ob0")
					}
					g.Println("}")
				case kind == system.KindValue:
					switch fieldType.NativeJsonType(ctx) {
					case system.J_STRING:
						g.Println("if v.", fieldName, " != \"\" {")
					case system.J_NUMBER:
						g.Println("if v.", fieldName, " != 0.0 {")
					case system.J_BOOL:
						g.Println("if v.", fieldName, " != false {")
					}
					{
						if err := printRepackCode(ctx, env, g, "v."+fieldName, "ob0", 0, f, true); err != nil {
							return kerr.Wrap("YYDYVIMXPM", err)
						}
						g.Println("m[", strconv.Quote(n), "] = ", "ob0")
					}
					g.Println("}")
				case kind == system.KindArray:
					g.Println("if v.", fieldName, " != nil {")
					{
						if err := printRepackCode(ctx, env, g, "v."+fieldName, "ob0", 0, f, true); err != nil {
							return kerr.Wrap("YSFPHQTBNA", err)
						}
						g.Println("m[", strconv.Quote(n), "] = ", "ob0")
					}
					g.Println("}")
				}
			}
			g.Println("return m, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
		case system.KindValue:
			if typ.Alias != nil {
				typ = system.WrapRule(ctx, typ.Alias).Parent
			}
			switch typ.NativeJsonType(ctx) {
			case system.J_STRING:
				g.Println("if v != nil {")
				{
					g.Println("return string(*v), ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
				}
				g.Println("}")
				g.Println("return nil, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
			case system.J_NUMBER:
				g.Println("if v != nil {")
				{
					g.Println("return float64(*v), ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
				}
				g.Println("}")
				g.Println("return nil, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
			case system.J_BOOL:
				g.Println("if v != nil {")
				{
					g.Println("return bool(*v), ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
				}
				g.Println("}")
				g.Println("return nil, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
			}
		case system.KindArray:
			if err := printRepackCode(ctx, env, g, "(*v)", "ob0", 0, typ.Alias, false); err != nil {
				return kerr.Wrap("SYKLQKLCEO", err)
			}
			g.Println("return ob0, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
		case system.KindMap:
			if err := printRepackCode(ctx, env, g, "(*v)", "ob0", 0, typ.Alias, false); err != nil {
				return kerr.Wrap("HBSGXLGKCD", err)
			}
			g.Println("return ob0, ", strconv.Quote(typ.Id.Package), ", ", strconv.Quote(typ.Id.Name), ", nil")
		}
	}
	g.Println("}")
	return nil
}

func printRepackCode(ctx context.Context, env *envctx.Env, g *builder.Builder, in string, out string, depth int, f system.RuleInterface, inStruct bool) error {
	fieldRule := system.WrapRule(ctx, f)
	kind, alias := fieldRule.Kind(ctx)
	repackerDef := g.SprintRef("kego.io/system", "Repacker")
	switch {
	case kind == system.KindInterface:
		valueVar := out + "_value"
		g.Println(valueVar, ", pkg, name, err := ", in, ".(", repackerDef, ").Repack(ctx)")
		g.Println("if err != nil {")
		{
			g.Println(`return nil, "", "", err`)
		}
		g.Println("}")
		newReferenceDef := g.SprintRef("kego.io/system", "NewReference")
		g.Println("typRef := ", newReferenceDef, "(pkg, name)")
		g.Println("typeVal, err := typRef.ValueContext(ctx)")
		g.Println("if err != nil {")
		{
			g.Println(`return nil, "", "", err`)
		}
		g.Println("}")
		g.Println(out, " := map[string]interface{}{}")
		g.Println(out, "[\"type\"] = typeVal")
		g.Println(out, "[\"value\"] = ", valueVar)
	case kind == system.KindStruct || (alias && inStruct):
		g.Println(out, ", _, _, err := ", in, ".Repack(ctx)")
		g.Println("if err != nil {")
		{
			g.Println(`return nil, "", "", err`)
		}
		g.Println("}")
	case kind == system.KindValue:
		g.Println(out, " := ", in)
	case kind == system.KindArray:
		iVar := fmt.Sprintf("i%d", depth)
		g.Println(out, " := []interface{}{}")
		g.Println("for ", iVar, " := range ", in, " {")
		{
			childIn := in + "[" + iVar + "]"
			childDepth := depth + 1
			childOut := fmt.Sprintf("ob%d", childDepth)
			childRule, err := fieldRule.ItemsRule()
			if err != nil {
				return kerr.Wrap("VUKWDVGVAT", err)
			}
			if err := printRepackCode(ctx, env, g, childIn, childOut, childDepth, childRule.Interface, true); err != nil {
				return kerr.Wrap("GDWUQWGFUI", err)
			}
			g.Println(out, " = append(", out, ", ", childOut, ")")
		}
		g.Println("}")
	case kind == system.KindMap:
		kVar := fmt.Sprintf("k%d", depth)
		g.Println(out, " := map[string]interface{}{}")
		g.Println("for ", kVar, " := range ", in, " {")
		{
			childIn := in + "[" + kVar + "]"
			childDepth := depth + 1
			childOut := fmt.Sprintf("ob%d", childDepth)
			childRule, err := fieldRule.ItemsRule()
			if err != nil {
				return kerr.Wrap("NYDJVRENGA", err)
			}
			if err := printRepackCode(ctx, env, g, childIn, childOut, childDepth, childRule.Interface, true); err != nil {
				return kerr.Wrap("VNWOUDMDQC", err)
			}
			g.Println(out, "[", kVar, "] = ", childOut)
		}
		g.Println("}")
	}
	return nil
}

func printUnpacker(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) error {
	name := system.GoName(typ.Id.Name)
	fmtPkg := g.Imports.Add("fmt")
	g.Println("func (v *", name, ") Unpack(ctx ", g.SprintRef("context", "Context"), ", in ", g.SprintRef("kego.io/system", "Packed"), ", iface bool) error {")
	{

		g.Println("if in == nil || in.Type() == ", g.SprintRef("kego.io/system", "J_NULL"), " {")
		{
			g.Println("return nil")
		}
		g.Println("}")

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
					return kerr.New("IOEEVJCDPU", "Type %s not found", embedRef.String())
				}
				embedName := system.GoName(embedRef.Name)
				embedTypeName := builder.Reference(embedType.Id.Package, system.GoName(embedType.Id.Name), env.Path, g.Imports.Add)
				g.Println("if v.", embedName, " == nil {")
				{
					g.Println("v.", embedName, " = new(", embedTypeName, ")")
				}
				g.Println("}")
				g.Println("if err := v.", embedName, ".Unpack(ctx, in, false); err != nil {")
				{
					g.Println("return err")
				}
				g.Println("}")
			}
			for n, f := range structType.Fields {
				g.Println("if field, ok := in.Map()[", strconv.Quote(n), "]; ok && field.Type() != ", g.SprintRef("kego.io/system", "J_NULL"), " {")
				{
					in := "field"
					out := "ob0"
					depth := 0
					if err := printUnpackCode(ctx, env, g, in, out, depth, f); err != nil {
						return kerr.Wrap("QLARKEBDBJ", err)
					}
					g.Println("v.", system.GoName(n), " = ", out)
				}
				if dr, ok := f.(system.DefaultRule); ok && dr.GetDefault() != nil {
					g.Println("} else {")
					{
						b, err := json.Marshal(dr.GetDefault())
						if err != nil {
							return kerr.Wrap("DLOUEHXVJF", err)
						}
						in := g.SprintFunctionCall("kego.io/system", "Pack", string(b))
						out := "ob0"
						depth := 0
						if err := printUnpackCode(ctx, env, g, in, out, depth, f); err != nil {
							return kerr.Wrap("UOWRFWSTNT", err)
						}
						g.Println("v.", system.GoName(n), " = ", out)
					}
					g.Println("}")
				} else {
					g.Println("}")
				}
			}
		case system.KindValue:
			g.Println("if in.Type() == ", g.SprintRef("kego.io/system", "J_MAP"), " {")
			{
				g.Println("in = in.Map()[\"value\"]")
			}
			g.Println("}")
			switch typ.NativeJsonType(ctx) {
			case system.J_BOOL:
				g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_BOOL"), " {")
				{
					g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking a bool.", in.Type())`)
				}
				g.Println("}")
				g.Println("*v = ", name, "(in.Bool())")
			case system.J_STRING:
				g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_STRING"), " {")
				{
					g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking a string.", in.Type())`)
				}
				g.Println("}")
				g.Println("*v = ", name, "(in.String())")
			case system.J_NUMBER:
				g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_NUMBER"), " {")
				{
					g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking a number.", in.Type())`)
				}
				g.Println("}")
				g.Println("*v = ", name, "(in.Number())")
			default:
				panic(fmt.Sprintf("invalid type kind: %s, json native: %s", kind, typ.NativeJsonType(ctx)))
			}
		case system.KindArray:
			g.Println("if in.Type() == ", g.SprintRef("kego.io/system", "J_MAP"), " {")
			{
				g.Println("in = in.Map()[\"value\"]")
			}
			g.Println("}")

			g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_ARRAY"), " {")
			{
				g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking an array.", in.Type())`)
			}
			g.Println("}")

			in := "in"
			out := "ob0"
			depth := 0
			if err := printUnpackCode(ctx, env, g, in, out, depth, typ.Alias); err != nil {
				return kerr.Wrap("VELYRXXGMC", err)
			}
			g.Println("*v = ", out)
		case system.KindMap:
			g.Println("if iface {")
			{
				g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_MAP"), " {")
				{
					g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking a map.", in.Type())`)
				}
				g.Println("}")
				g.Println("in = in.Map()[\"value\"]")
			}
			g.Println("}")

			g.Println("if in.Type() != ", g.SprintRef("kego.io/system", "J_MAP"), " {")
			{
				g.Println("return ", fmtPkg, `.Errorf("Invalid type %s while unpacking an array.", in.Type())`)
			}
			g.Println("}")

			in := "in"
			out := "ob0"
			depth := 0
			if err := printUnpackCode(ctx, env, g, in, out, depth, typ.Alias); err != nil {
				return kerr.Wrap("TNEIUAFUNY", err)
			}
			g.Println("*v = ", out)
		}
	}
	g.Println("return nil")
	g.Println("}")
	return nil
}

func printUnpackCode(ctx context.Context, env *envctx.Env, g *builder.Builder, in string, out string, depth int, f system.RuleInterface) error {
	field := system.WrapRule(ctx, f)
	fieldType := field.Parent
	fieldTypeName := builder.Reference(fieldType.Id.Package, system.GoName(fieldType.Id.Name), env.Path, g.Imports.Add)
	kind, alias := field.Kind(ctx)
	switch {
	case kind == system.KindStruct || alias:
		ptr := fieldType.PassedAsPointerInverseString(ctx)
		g.Println(out, " := ", ptr, "new(", fieldTypeName, ")")
		g.Println("if err := ", out, ".Unpack(ctx, ", in, ", false); err != nil {")
		{
			g.Println("return err")
		}
		g.Println("}")
		//if store {
		//	g.Println("v.", fieldName, " = ob")
		//}
	case kind == system.KindValue:
		var funcName string
		switch fieldType.NativeJsonType(ctx) {
		case system.J_STRING:
			funcName = "UnpackString"
		case system.J_NUMBER:
			funcName = "UnpackNumber"
		case system.J_BOOL:
			funcName = "UnpackBool"
		default:
			return kerr.New("LSGUACQGHB", "Kind == KindValue but native json type==%s", fieldType.NativeJsonType(ctx))
		}
		funcRef := g.SprintRef("kego.io/system", funcName)
		g.Println(out, ", err := ", funcRef, "(ctx, ", in, ")")
		g.Println("if err != nil {")
		{
			g.Println("return err")
		}
		g.Println("}")
		//if store {
		//	g.Println("v.", fieldName, " = ob")
		//}
	case kind == system.KindInterface:
		var interfaceName string
		if fieldType.Interface {
			interfaceName = system.GoName(fieldType.Id.Name)
		} else {
			interfaceName = system.GoInterfaceName(fieldType.Id.Name)
		}
		unpackFuncName := builder.Reference(fieldType.Id.Package, "Unpack"+interfaceName, env.Path, g.Imports.Add)
		g.Println(out, ", err := ", unpackFuncName, "(ctx, ", in, ")")
		g.Println("if err != nil {")
		{
			g.Println("return err")
		}
		g.Println("}")
		//if store {
		//	g.Println("v.", fieldName, " = ob")
		//}
	case kind == system.KindArray:
		fmtPkg := g.Imports.Add("fmt")
		g.Println("if ", in, ".Type() != ", g.SprintRef("kego.io/system", "J_ARRAY"), " {")
		{
			g.Println("return ", fmtPkg, ".Errorf(\"Unsupported json type %s found while unpacking into an array.\", ", in, ".Type())")
		}
		g.Println("}")
		fieldType, err := builder.TypeDefinition(ctx, f, env.Path, g.Imports.Add)
		if err != nil {
			return kerr.Wrap("UDPMSSFLTW", err)
		}
		g.Println(out, " := ", fieldType, "{}")
		iVar := fmt.Sprintf("i%d", depth)
		g.Println("for ", iVar, " := range ", in, ".Array() {")
		{
			childRule, err := field.ItemsRule()
			if err != nil {
				return kerr.Wrap("PJQBRUEHLD", err)
			}
			childDepth := depth + 1
			childIn := in + ".Array()[" + iVar + "]"
			childOut := fmt.Sprintf("ob%d", childDepth)
			if err := printUnpackCode(ctx, env, g, childIn, childOut, childDepth, childRule.Interface); err != nil {
				return kerr.Wrap("XCHEJPDJDS", err)
			}
			g.Println(out, " = append(", out, ", ", childOut, ")")
		}
		g.Println("}")
	case kind == system.KindMap:
		fmtPkg := g.Imports.Add("fmt")
		g.Println("if ", in, ".Type() != ", g.SprintRef("kego.io/system", "J_MAP"), " {")
		{
			g.Println("return ", fmtPkg, ".Errorf(\"Unsupported json type %s found while unpacking into a map.\", ", in, ".Type())")
		}
		g.Println("}")
		fieldType, err := builder.TypeDefinition(ctx, f, env.Path, g.Imports.Add)
		if err != nil {
			return kerr.Wrap("IHNYODUIKG", err)
		}
		g.Println(out, " := ", fieldType, "{}")
		kVar := fmt.Sprintf("k%d", depth)
		g.Println("for ", kVar, " := range ", in, ".Map() {")
		{
			items, err := field.ItemsRule()
			if err != nil {
				return kerr.Wrap("FPOBYGVOPP", err)
			}
			childDepth := depth + 1
			childIn := in + ".Map()[" + kVar + "]"
			childOut := fmt.Sprintf("ob%d", childDepth)
			if err := printUnpackCode(ctx, env, g, childIn, childOut, childDepth, items.Interface); err != nil {
				return kerr.Wrap("ONUJJGIKJE", err)
			}
			g.Println(out, "[", kVar, "] = ", childOut)
		}
		g.Println("}")
	}
	return nil
}

func printInterfaceUnpacker(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) {
	// if type kind is native, also accept bare json native
	// if type kind is array, also accept bare json array
	typeName := system.GoName(typ.Id.Name)
	var interfaceName string
	if typ.Interface {
		interfaceName = system.GoName(typ.Id.Name)
	} else {
		interfaceName = system.GoInterfaceName(typ.Id.Name)
	}
	fmtPkg := g.Imports.Add("fmt")
	g.Println("func Unpack", interfaceName, "(ctx ", g.SprintRef("context", "Context"), ", in ", g.SprintRef("kego.io/system", "Packed"), ") (", interfaceName, ", error) {")
	{
		g.Println("switch in.Type() {")
		{
			g.Println("case ", g.SprintRef("kego.io/system", "J_MAP"), ":")
			{
				unknownTypeFunc := g.SprintFunctionCall("kego.io/system", "UnpackUnknownType", "ctx", "in", "true", strconv.Quote(typ.Id.Package), strconv.Quote(typ.Id.Name))
				g.Println("i, err := ", unknownTypeFunc)
				g.Println("if err != nil {")
				{
					g.Println("return nil, err")
				}
				g.Println("}")
				g.Println("ob, ok := i.(", interfaceName, ")")
				g.Println("if !ok {")
				{
					g.Println("return nil, ", fmtPkg, `.Errorf("%T does not implement `, interfaceName, `", i)`)
				}
				g.Println("}")
				g.Println("return ob, nil")
			}
			switch typ.NativeJsonType(ctx) {
			// We don't include J_MAP in here.
			case system.J_STRING, system.J_NUMBER, system.J_BOOL, system.J_ARRAY:
				switch typ.NativeJsonType(ctx) {
				case system.J_STRING:
					g.Println("case ", g.SprintRef("kego.io/system", "J_STRING"), ":")
				case system.J_NUMBER:
					g.Println("case ", g.SprintRef("kego.io/system", "J_NUMBER"), ":")
				case system.J_BOOL:
					g.Println("case ", g.SprintRef("kego.io/system", "J_BOOL"), ":")
				case system.J_ARRAY:
					g.Println("case ", g.SprintRef("kego.io/system", "J_ARRAY"), ":")
				}
				g.Println("ob := new(", typeName, ")")
				g.Println("if err := ob.Unpack(ctx, in, false); err != nil {")
				{
					g.Println("return nil, err")
				}
				g.Println("}")
				g.Println("return ob, nil")
			}
			g.Println("default:")
			{
				g.Println("return nil, ", fmtPkg, `.Errorf("Unsupported json type %s when unpacking into `, interfaceName, `.", in.Type())`)
			}
		}
		g.Println("}")
	}
	g.Println("}")
}

func printInterfaceDefinition(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) {

	g.Println("type ", system.GoInterfaceName(typ.Id.Name), " interface {")
	{
		g.Println("Get",
			system.GoName(typ.Id.Name),
			"(ctx ",
			builder.Reference("context", "Context", env.Path, g.Imports.Add),
			") ",
			typ.PassedAsPointerString(ctx),
			system.GoName(typ.Id.Name))
	}
	g.Println("}")
}

func printInterfaceImplementation(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) {

	g.Println("func (o ",
		typ.PassedAsPointerString(ctx),
		system.GoName(typ.Id.Name),
		") Get",
		system.GoName(typ.Id.Name),
		"(ctx ",
		builder.Reference("context", "Context", env.Path, g.Imports.Add),
		") ",
		typ.PassedAsPointerString(ctx),
		system.GoName(typ.Id.Name),
		" {")
	{
		g.Println("return o")
	}
	g.Println("}")
}

func printAliasDefinition(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) error {
	if typ.Description != "" {
		g.Println("// ", typ.Description)
	}
	aliasType, err := builder.AliasTypeDefinition(ctx, typ.Alias, env.Path, g.Imports.Add)
	if err != nil {
		return kerr.Wrap("FWOLIESYUA", err)
	}
	g.Println("type ", system.GoName(typ.Id.Name), " ", aliasType)
	return nil
}

func printStructDefinition(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) error {
	if typ.Description != "" {
		g.Println("// ", typ.Description)
	}
	g.Println("type ", system.GoName(typ.Id.Name), " struct {")
	{
		if !typ.Basic {
			g.Println("*", g.SprintRef("kego.io/system", system.GoName("object")))
		}

		embedsSortable := system.SortableReferences(typ.Embed)
		sort.Sort(embedsSortable)
		embeds := []*system.Reference(embedsSortable)
		for _, embed := range embeds {
			g.Println("*", builder.Reference(embed.Package, system.GoName(embed.Name), env.Path, g.Imports.Add))
		}

		for _, nf := range typ.SortedFields() {
			b := nf.Rule.(system.ObjectInterface).GetObject(nil)
			if b.Description != "" {
				g.Println("// ", b.Description)
			}
			descriptor, err := builder.FieldTypeDefinition(ctx, nf.Name, nf.Rule, env.Path, g.Imports.Add)
			if err != nil {
				return kerr.Wrap("GDSKJDEKQD", err)
			}
			g.Println(system.GoName(nf.Name), " ", descriptor)
		}
	}
	g.Println("}")
	return nil
}

func printInitFunction(ctx context.Context, env *envctx.Env, g *builder.Builder, types *sysctx.SysTypes) {
	g.Println("func init() {")
	{
		g.Print("pkg := ")
		g.PrintFunctionCall(
			"kego.io/context/jsonctx",
			"InitPackage",
			strconv.Quote(env.Path),
		)
		g.Println("")
		g.PrintMethodCall("pkg", "SetHash", env.Hash)
		g.Println("")
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

			interfaceFunc := "nil"
			isNativeCollection := typ.IsNativeCollection() && typ.Alias == nil
			if !isNativeCollection {
				var ifaceName string
				if typ.Interface {
					ifaceName = system.GoName(typ.Id.Name)
				} else {
					ifaceName = system.GoInterfaceName(typ.Id.Name)
				}
				reflectTypeof := g.SprintFunctionCall(
					"reflect",
					"TypeOf",
					fmt.Sprintf("(*%s)(nil)", ifaceName),
				) + ".Elem()"
				reflectType := builder.Reference("reflect", "Type", env.Path, g.Imports.Add)
				interfaceFunc = "func() " + reflectType + " { return " + reflectTypeof + " }"
			}

			newFunc := "nil"
			if typ.Interface {
				newFunc = "func() interface{} { return (*" + system.GoName(typ.Id.Name) + ")(nil) }"
			} else if !typ.IsNativeCollection() || typ.Alias != nil {
				newFunc = "func() interface{} { return " + typ.PassedAsPointerInverseString(ctx) + "new(" + system.GoName(typ.Id.Name) + ")}"
			}

			g.PrintMethodCall(
				"pkg",
				"Init",
				strconv.Quote(typ.Id.Name),
				newFunc,
				"func() interface{} { return new("+system.GoName(typ.Id.ChangeToRule().Name)+")}",
				interfaceFunc,
			)

			g.Println("")
		}
	}
	g.Println("}")
}

type InfoStruct struct {
	Path string
	Hash uint64
}
