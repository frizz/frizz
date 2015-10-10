package generate

import (
	"fmt"
	"strconv"

	"sort"

	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/process/settings"
	"kego.io/system"
)

func Structs(set *settings.Settings) (source []byte, err error) {

	filter := system.NewReference("kego.io/system", "type")
	types := system.GetAllGlobalsInPackage(set.Path, &filter)
	g := generator.New(set.Path)

	if len(types) == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.New("BBRLIODBKL", err, "Build")
		}
		return b, nil
	}
	for _, hashed := range types {
		typ := hashed.Object.(*system.Type)
		if typ.Interface || typ.IsNativeValue() {
			continue
		}
		if typ.Description != "" {
			g.Println("// ", typ.Description)
		}
		g.Println("type ", system.GoName(typ.Id.Name), " struct {")
		{
			embedsSortable := system.SortableReferences(typ.Embed)
			sort.Sort(embedsSortable)
			embeds := []system.Reference(embedsSortable)
			for _, embed := range embeds {
				g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), set.Path, g.Imports.Add))
			}

			if !typ.Basic {
				g.Println("*", generator.Reference("kego.io/system", "Object_base", set.Path, g.Imports.Add))
			}
			interfacesSortable := system.SortableReferences(typ.Is)
			sort.Sort(interfacesSortable)
			interfaces := []system.Reference(interfacesSortable)
			for _, iface := range interfaces {
				h, ok := system.GetGlobal(iface.Package, iface.Name)
				if !ok {
					return nil, kerr.New("UHIRQKNEEV", nil, "Can't find type %s", iface)
				}
				t, ok := h.Object.(*system.Type)
				if !ok {
					return nil, kerr.New("AOMVIJBFJA", nil, "%T is not a *system.Type", h.Object)
				}
				if t.Base != nil {
					g.Println("*", generator.Reference(t.Base.Id.Package, system.GoName(t.Base.Id.Name), set.Path, g.Imports.Add))
				}
			}

			for _, nf := range typ.SortedFields() {
				b := nf.Rule.(system.Object).Object()
				if b.Description != "" {
					g.Println("// ", b.Description)
				}
				descriptor, err := generator.Type(nf.Name, nf.Rule, set.Path, g.Imports.Add)
				if err != nil {
					return nil, kerr.New("GDSKJDEKQD", err, "generator.Type")
				}
				g.Println(system.GoName(nf.Name), " ", descriptor)
			}
		}
		g.Println("}")
	}
	g.Println("func init() {")
	{
		for _, hashed := range types {
			typ := hashed.Object.(*system.Type)
			jsonRegisterType := generator.Reference("kego.io/json", "Register", set.Path, g.Imports.Add)
			reflectTypeOf := generator.Reference("reflect", "TypeOf", set.Path, g.Imports.Add)
			reflectTypeOfParams := ""
			if typ.Interface {
				// reflect.TypeOf((*Foo)(nil)).Elem()
				reflectTypeOfParams = fmt.Sprintf("((*%s)(nil)).Elem()", system.GoName(typ.Id.Name))
			} else if typ.Native.Value == "object" {
				// reflect.TypeOf(&Foo{})
				reflectTypeOfParams = fmt.Sprintf("(&%s{})", system.GoName(typ.Id.Name))
			} else {
				// reflect.TypeOf(Foo{})
				reflectTypeOfParams = fmt.Sprintf("(%s{})", system.GoName(typ.Id.Name))
			}
			typOf := fmt.Sprint(reflectTypeOf, reflectTypeOfParams)
			// e.g.
			// json.Register("kego.io/foo", "@gallery", reflect.TypeOf(&Gallery_rule{}), 0x123)
			// json.Register("kego.io/foo", "gallery", reflect.TypeOf(&Gallery{}), 0x123)
			// json.Register("kego.io/foo", "string", reflect.TypeOf(String{}), 0x123)
			// json.Register("kego.io/foo", "iface", reflect.TypeOf((*Iface)(nil)).Elem(), 0x123)
			g.Printf("%s(%s, %s, %s, %#v)\n", jsonRegisterType, strconv.Quote(typ.Id.Package), strconv.Quote(typ.Id.Name), typOf, hashed.Hash)
		}
	}
	g.Println("}")

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("XKYHSDKBEP", err, "Build")
	}
	return b, nil
}

func StructsCommand(set *settings.Settings) (source []byte, err error) {

	g := generator.WithName(set.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/process")
	g.Imports.Anonymous("kego.io/system")
	if set.Path != "kego.io/system" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	for p, _ := range set.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Print(`
		func main() {
			set, err := process.InitialiseAutomatic()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err := process.Generate(process.S_STRUCTS, set); err != nil {
				fmt.Println(process.FormatError(err))
				os.Exit(1)
			}
		}`)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}
