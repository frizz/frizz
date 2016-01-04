package parse

import (
	"fmt"
	"strconv"

	"sort"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/envctx"
	"kego.io/generator"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

func Structs(ctx context.Context, env *envctx.Env) (source []byte, err error) {

	cache := cachectx.FromContext(ctx)

	pcache, ok := cache.Get(env.Path)
	if !ok {
		return nil, kerr.New("DQVQWTKRSK", nil, "%s not found in ctx", env.Path)
	}
	types := pcache.Types

	g := generator.New(env.Path)

	infoBytes, err := json.MarshalPlain(InfoStruct{Path: env.Path, Hash: env.Hash})
	if err != nil {
		return nil, kerr.New("HVFWIUVLSM", err, "MarshalPlain")
	}
	g.SetPackageComment("info:" + string(infoBytes))

	if types.Len() == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.New("BBRLIODBKL", err, "Build")
		}
		return b, nil
	}
	for _, name := range types.Keys() {
		t, ok := types.Get(name)
		if !ok {
			continue
		}
		typ := t.(*system.Type)

		isRule := typ.Id.IsRule()

		if typ.IsNativeCollection() {
			continue
		}

		if !typ.Interface && !typ.IsNativeValue() {
			if err := printStructDefinition(ctx, env, g, typ); err != nil {
				return nil, kerr.New("XKRYMXUIJD", err, "printNewStructDefinition")
			}
		}

		if !typ.Interface && !isRule {
			printInterfaceDefinition(env, g, typ)
			printInterfaceImplementation(env, g, typ)
		}

	}
	printInitFunction(env, g, types)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("XKYHSDKBEP", err, "Build")
	}
	return b, nil
}

func printInterfaceDefinition(env *envctx.Env, g *generator.Generator, typ *system.Type) {
	g.Println("type ", system.GoInterfaceName(typ.Id.Name), " interface {")
	{
		g.Println("Get",
			system.GoName(typ.Id.Name),
			"(ctx ",
			generator.Reference("golang.org/x/net/context", "Context", env.Path, g.Imports.Add),
			") *",
			system.GoName(typ.Id.Name))
	}
	g.Println("}")
}

func printInterfaceImplementation(env *envctx.Env, g *generator.Generator, typ *system.Type) {
	g.Println("func (o *",
		system.GoName(typ.Id.Name),
		") Get",
		system.GoName(typ.Id.Name),
		"(ctx ",
		generator.Reference("golang.org/x/net/context", "Context", env.Path, g.Imports.Add),
		") *",
		system.GoName(typ.Id.Name),
		" {")
	{
		g.Println("return o")
	}
	g.Println("}")
}

func printStructDefinition(ctx context.Context, env *envctx.Env, g *generator.Generator, typ *system.Type) error {
	if typ.Description != "" {
		g.Println("// ", typ.Description)
	}
	g.Println("type ", system.GoName(typ.Id.Name), " struct {")
	{
		if !typ.Basic {
			g.Println("*", generator.Reference("kego.io/system", system.GoName("object"), env.Path, g.Imports.Add))
		}

		embedsSortable := system.SortableReferences(typ.Embed)
		sort.Sort(embedsSortable)
		embeds := []*system.Reference(embedsSortable)
		for _, embed := range embeds {
			g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), env.Path, g.Imports.Add))
		}

		for _, nf := range typ.SortedFields() {
			b := nf.Rule.(system.ObjectInterface).GetObject(nil)
			if b.Description != "" {
				g.Println("// ", b.Description)
			}
			descriptor, err := generator.Type(ctx, nf.Name, nf.Rule, env.Path, g.Imports.Add)
			if err != nil {
				return kerr.New("GDSKJDEKQD", err, "generator.TypeNew")
			}
			g.Println(system.GoName(nf.Name), " ", descriptor)
		}
	}
	g.Println("}")
	return nil
}

func printInitFunction(env *envctx.Env, g *generator.Generator, types *cachectx.TypeCache) {
	g.Println("func init() {")
	{
		g.PrintFunctionCall(
			"kego.io/json",
			"RegisterPackage",
			strconv.Quote(env.Path),
			env.Hash,
		)
		g.Println("")

		for _, name := range types.Keys() {
			t, ok := types.Get(name)
			if !ok {
				continue
			}
			typ := t.(*system.Type)
			isRule := typ.Id.IsRule()

			if isRule {
				continue
			}

			if typ.IsNativeCollection() {
				continue
			}

			typeOf1 := ""
			if typ.Interface {
				typeOf1 = g.SprintFunctionCall(
					"reflect",
					"TypeOf",
					fmt.Sprintf("(*%s)(nil)", system.GoName(typ.Id.Name)),
				) + ".Elem()"
			} else {
				typeOf1 = g.SprintFunctionCall(
					"reflect",
					"TypeOf",
					fmt.Sprintf("(*%s)(nil)", system.GoName(typ.Id.Name)),
				)
			}

			typeOf2 := g.SprintFunctionCall(
				"reflect",
				"TypeOf",
				fmt.Sprintf("(*%s)(nil)", system.GoName(typ.Id.ChangeToRule().Name)),
			)

			typeOf3 := "nil"
			if !typ.Interface {
				typeOf3 = g.SprintFunctionCall(
					"reflect",
					"TypeOf",
					fmt.Sprintf("(*%s)(nil)", system.GoInterfaceName(typ.Id.Name)),
				) + ".Elem()"
			}

			g.PrintFunctionCall(
				"kego.io/json",
				"RegisterType",
				strconv.Quote(typ.Id.Package),
				strconv.Quote(typ.Id.Name),
				typeOf1,
				typeOf2,
				typeOf3,
			)
			g.Println("")
		}
	}
	g.Println("}")
}
