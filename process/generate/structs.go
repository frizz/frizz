package generate

import (
	"fmt"
	"strconv"

	"sort"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/json"
	"kego.io/kerr"
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

	infoBytes, _ := json.MarshalPlain(InfoStruct{Path: env.Path, Hash: env.Hash})

	g.SetPackageComment("info:" + string(infoBytes))

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
			continue
		}
		typ := t.(*system.Type)

		isRule := typ.Id.IsRule()

		if typ.IsNativeCollection() {
			continue
		}

		if !typ.Interface && !typ.IsNativeValue() {
			if err := printStructDefinition(ctx, env, g, typ); err != nil {
				return nil, kerr.Wrap("XKRYMXUIJD", err)
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
		return nil, kerr.Wrap("XKYHSDKBEP", err)
	}
	return b, nil
}

func printInterfaceDefinition(env *envctx.Env, g *builder.Builder, typ *system.Type) {
	g.Println("type ", system.GoInterfaceName(typ.Id.Name), " interface {")
	{
		g.Println("Get",
			system.GoName(typ.Id.Name),
			"(ctx ",
			builder.Reference("golang.org/x/net/context", "Context", env.Path, g.Imports.Add),
			") *",
			system.GoName(typ.Id.Name))
	}
	g.Println("}")
}

func printInterfaceImplementation(env *envctx.Env, g *builder.Builder, typ *system.Type) {
	g.Println("func (o *",
		system.GoName(typ.Id.Name),
		") Get",
		system.GoName(typ.Id.Name),
		"(ctx ",
		builder.Reference("golang.org/x/net/context", "Context", env.Path, g.Imports.Add),
		") *",
		system.GoName(typ.Id.Name),
		" {")
	{
		g.Println("return o")
	}
	g.Println("}")
}

func printStructDefinition(ctx context.Context, env *envctx.Env, g *builder.Builder, typ *system.Type) error {
	if typ.Description != "" {
		g.Println("// ", typ.Description)
	}
	g.Println("type ", system.GoName(typ.Id.Name), " struct {")
	{
		if !typ.Basic {
			g.Println("*", builder.Reference("kego.io/system", system.GoName("object"), env.Path, g.Imports.Add))
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
			descriptor, err := builder.Type(ctx, nf.Name, nf.Rule, env.Path, g.Imports.Add)
			if err != nil {
				return kerr.Wrap("GDSKJDEKQD", err)
			}
			g.Println(system.GoName(nf.Name), " ", descriptor)
		}
	}
	g.Println("}")
	return nil
}

func printInitFunction(env *envctx.Env, g *builder.Builder, types *sysctx.TypeCache) {
	g.Println("func init() {")
	{
		g.Print("pkg := ")
		g.PrintFunctionCall(
			"kego.io/context/jsonctx",
			"InitPackage",
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

			g.PrintMethodCall(
				"pkg",
				"InitType",
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

type InfoStruct struct {
	Path string
	Hash uint64
}
