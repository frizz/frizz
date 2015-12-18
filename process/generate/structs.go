package generate

import (
	"fmt"
	"strconv"

	"sort"

	"strings"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/system"
)

func Structs(ctx context.Context) (source []byte, err error) {

	env, ok := envctx.FromContext(ctx)
	if !ok {
		return nil, kerr.New("FCLEJBEQWQ", nil, "No env in ctx")
	}

	types := system.GetAllGlobalsInPackage(env.Path, system.NewReference("kego.io/system", "type"))
	g := generator.New(env.Path)

	if len(types) == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.New("BBRLIODBKL", err, "Build")
		}
		return b, nil
	}
	for _, hashed := range types {
		typ := hashed.Object.(*system.Type)
		isRule := strings.HasPrefix(typ.Id.Name, system.RULE_PREFIX)

		if typ.IsNativeCollection() {
			continue
		}

		if !typ.Interface && !typ.IsNativeValue() {
			if err := printStructDefinition(ctx, g, typ, env.Path); err != nil {
				return nil, kerr.New("XKRYMXUIJD", err, "printNewStructDefinition")
			}
		}

		if !typ.Interface && !isRule {
			printInterfaceDefinition(g, typ)
			printInterfaceImplementation(g, typ)
		}

	}
	printInitFunction(g, types)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("XKYHSDKBEP", err, "Build")
	}
	return b, nil
}

func printInterfaceDefinition(g *generator.Generator, typ *system.Type) {
	g.Println("type ", system.GoInterfaceName(typ.Id.Name), " interface {")
	{
		g.Println("Get", system.GoName(typ.Id.Name), "() *", system.GoName(typ.Id.Name))
	}
	g.Println("}")
}

func printInterfaceImplementation(g *generator.Generator, typ *system.Type) {
	g.Println("func (o *", system.GoName(typ.Id.Name), ") Get", system.GoName(typ.Id.Name), "() *", system.GoName(typ.Id.Name), " {")
	{
		g.Println("return o")
	}
	g.Println("}")
}

func printStructDefinition(ctx context.Context, g *generator.Generator, typ *system.Type, path string) error {
	if typ.Description != "" {
		g.Println("// ", typ.Description)
	}
	g.Println("type ", system.GoName(typ.Id.Name), " struct {")
	{
		if !typ.Basic {
			g.Println("*", generator.Reference("kego.io/system", system.GoName("object"), path, g.Imports.Add))
		}

		embedsSortable := system.SortableReferences(typ.Embed)
		sort.Sort(embedsSortable)
		embeds := []*system.Reference(embedsSortable)
		for _, embed := range embeds {
			g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), path, g.Imports.Add))
		}

		for _, nf := range typ.SortedFields() {
			b := nf.Rule.(system.ObjectInterface).GetObject()
			if b.Description != "" {
				g.Println("// ", b.Description)
			}
			descriptor, err := generator.Type(ctx, nf.Name, nf.Rule, path, g.Imports.Add)
			if err != nil {
				return kerr.New("GDSKJDEKQD", err, "generator.TypeNew")
			}
			g.Println(system.GoName(nf.Name), " ", descriptor)
		}
	}
	g.Println("}")
	return nil
}

func printInitFunction(g *generator.Generator, types []system.Hashed) {
	g.Println("func init() {")
	{
		for _, hashed := range types {

			typ := hashed.Object.(*system.Type)
			isRule := strings.HasPrefix(typ.Id.Name, system.RULE_PREFIX)

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

			typeOf2 := "nil"
			if !typ.Interface && !isRule {
				typeOf2 = g.SprintFunctionCall(
					"reflect",
					"TypeOf",
					fmt.Sprintf("(*%s)(nil)", system.GoInterfaceName(typ.Id.Name)),
				) + ".Elem()"
			}

			g.PrintFunctionCall(
				"kego.io/json",
				"Register",
				strconv.Quote(typ.Id.Package),
				strconv.Quote(typ.Id.Name),
				typeOf1,
				typeOf2,
				hashed.Hash,
			)
			g.Println("")
		}
	}
	g.Println("}")
}
