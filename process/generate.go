package process

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"

	"sort"

	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/literal"
	"kego.io/system"
)

func GenerateSource(file sourceType, set settings) (source []byte, err error) {
	b := bytes.NewBuffer(nil)
	switch file {
	case S_STRUCTS:
		filter := system.NewReference("kego.io/system", "type")
		types := system.GetAllGlobalsInPackage(set.path, &filter)
		g, err := generator.New(set.path, b)
		if err != nil {
			return nil, kerr.New("WPIPODJIGV", err, "generator.New")
		}
		if len(types) == 0 {
			g.Build()
			break
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
					g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), set.path, g.Imports.Add))
				}

				if !typ.Basic {
					g.Println("*", generator.Reference("kego.io/system", "Object_base", set.path, g.Imports.Add))
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
						g.Println("*", generator.Reference(t.Base.Id.Package, system.GoName(t.Base.Id.Name), set.path, g.Imports.Add))
					}
				}

				for _, nf := range typ.SortedFields() {
					b := nf.Rule.(system.Object).Object()
					if b.Description != "" {
						g.Println("// ", b.Description)
					}
					descriptor, err := generator.Type(nf.Name, nf.Rule, set.path, g.Imports.Add)
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
				jsonRegisterType := generator.Reference("kego.io/json", "Register", set.path, g.Imports.Add)
				reflectTypeOf := generator.Reference("reflect", "TypeOf", set.path, g.Imports.Add)
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
		g.Build()
	case S_EDITOR:
		g, err := generator.New("main", b)
		if err != nil {
			return nil, kerr.New("TMBFYTQHXU", err, "generator.New")
		}
		g.Imports.Anonymous("kego.io/system")
		g.Imports.Anonymous("kego.io/system/types")
		g.Imports.Anonymous(set.path)
		g.Imports.Anonymous(fmt.Sprint(set.path, "/types"))
		for p, _ := range set.aliases {
			g.Imports.Anonymous(p)
			g.Imports.Anonymous(fmt.Sprint(p, "/types"))
		}
		/*
			func main() {
				if err := client.Start("XXX"); err != nil {
					console.Error(err.Error())
				}
			}
		*/
		g.Println("func main() {")
		{
			clientStart := generator.Reference("kego.io/editor/client", "Start", set.path, g.Imports.Add)
			g.Println("if err := ", clientStart, "(", strconv.Quote(set.path), "); err != nil {")
			{
				consoleError := generator.Reference("kego.io/js/console", "Error", set.path, g.Imports.Add)
				g.Println(consoleError, "(err.Error())")
			}
			g.Println("}")
		}
		g.Println("}")
		g.Build()
	case S_TYPES:
		filter := system.NewReference("kego.io/system", "type")
		types := system.GetAllGlobalsInPackage(set.path, &filter)
		typesPath := fmt.Sprintf("%s/types", set.path)
		g, err := generator.New(typesPath, b)
		if err != nil {
			return nil, kerr.New("KYIKJQXPMR", err, "generator.New")
		}
		if len(types) == 0 {
			g.Build()
			break
		}
		if typesPath != "kego.io/system/types" {
			g.Imports.Anonymous("kego.io/system/types")
		}
		for p, _ := range set.aliases {
			g.Imports.Anonymous(fmt.Sprint(p, "/types"))
		}
		g.Println("func init() {")
		{
			registers := []string{}
			pointers := []literal.Pointer{}
			systemRegisterType := generator.Reference("kego.io/system", "Register", typesPath, g.Imports.Add)
			for _, hashed := range types {
				t := hashed.Object.(*system.Type)
				pkg := strconv.Quote(t.Id.Package)
				name := strconv.Quote(t.Id.Name)
				pointer := literal.Build(t, &pointers, typesPath, g.Imports.Add)
				// e.g.
				// system.Register("kego.io/gallery/data", "gallery", ptr8728815248)
				// system.Register("kego.io/gallery/data", "@gallery", ptr8728815360)
				line := fmt.Sprintf("%s(%s, %s, %s, %#v)", systemRegisterType, pkg, name, pointer.Name, hashed.Hash)
				registers = append(registers, line)
			}
			for _, p := range pointers {
				g.Println(p.Name, " := ", p.Source)
			}
			for _, s := range registers {
				g.Println(s)
			}
		}
		g.Println("}")
		g.Build()
	}
	source, err = format.Source(b.Bytes())
	if err != nil {
		err = kerr.New("CRBYOUOHPG", err, "format.Source: %s", b.String())
	}
	return
}

func GenerateCommand(file commandType, set settings) (source []byte, err error) {
	b := bytes.NewBuffer(nil)
	switch file {
	case C_STRUCTS:
		g := generator.NewWithName(set.path, "main", b)
		g.Imports.Add("os")
		g.Imports.Add("fmt")
		g.Imports.Add("kego.io/process")
		g.Imports.Anonymous("kego.io/system")
		if set.path != "kego.io/system" {
			g.Imports.Anonymous("kego.io/system/types")
		}
		for p, _ := range set.aliases {
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
		g.Build()
	case C_TYPES:
		g := generator.NewWithName(set.path, "main", b)
		g.Imports.Add("os")
		g.Imports.Add("fmt")
		g.Imports.Add("kego.io/process")
		g.Imports.Anonymous("kego.io/system")
		if set.path != "kego.io/system" {
			g.Imports.Anonymous("kego.io/system/types")
		}
		g.Imports.Anonymous(set.path)
		for p, _ := range set.aliases {
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
				if err := process.Generate(process.S_TYPES, set); err != nil {
					fmt.Println(process.FormatError(err))
			        os.Exit(1)
				}
				if err := process.Generate(process.S_EDITOR, set); err != nil {
					fmt.Println(process.FormatError(err))
			        os.Exit(1)
				}
			}`)
		g.Build()
	case C_KE:
		g := generator.NewWithName(set.path, "main", b)
		g.Imports.Add("os")
		g.Imports.Add("fmt")
		g.Imports.Add("kego.io/process")
		g.Imports.Add("kego.io/editor/server")
		g.Imports.Anonymous("kego.io/system")
		g.Imports.Anonymous("kego.io/system/types")
		g.Imports.Anonymous(set.path)
		g.Imports.Anonymous(fmt.Sprint(set.path, "/types"))
		for p, _ := range set.aliases {
			g.Imports.Anonymous(p)
			g.Imports.Anonymous(fmt.Sprint(p, "/types"))
		}
		g.Print(`
			func main() {
				update := false
				recursive := ` + fmt.Sprint(set.recursive) + `
				path := ` + strconv.Quote(set.path) + `
				set, err := process.InitialiseCommand(update, recursive, path)
				if err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if typesChanged, err := process.ValidateCommand(set); err != nil {
					if !typesChanged || !set.Verbose() {
						// when ValidateCommand detects the types have changed, it
						// spawns a ke command. In verbose mode this will output any
						// error so we don't need to print the error
						fmt.Println(process.FormatError(err))
					}
					os.Exit(1)
				}
				if set.Edit() {
					if err = server.Start(set.Path(), set.Verbose()); err != nil {
						fmt.Println(process.FormatError(err))
						os.Exit(1)
					}
				}
			}`)
		g.Build()
	}
	source, err = format.Source(b.Bytes())
	if err != nil {
		err = kerr.New("CRBYOUOHPG", err, "format.Source: %s", b.String())
	}
	return
}
