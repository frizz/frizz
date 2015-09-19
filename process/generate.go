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
			if typ.Interface || typ.IsNativeValue() || typ.Exclude {
				continue
			}
			if typ.Description != "" {
				g.Println("// ", typ.Description)
			}
			g.Println("type ", system.GoName(typ.Id.Name), " struct {")
			{
				if !typ.Basic {
					g.Println("*", generator.Reference("kego.io/system", "Base", set.path, g.Imports.Add))
				}
				embeds := system.SortableReferences(typ.Embed)
				sort.Sort(embeds)
				for _, embed := range embeds {
					g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), set.path, g.Imports.Add))
				}
				for _, nf := range typ.SortedFields() {
					b := nf.Field.(system.Object).GetBase()
					if b.Description != "" {
						g.Println("// ", b.Description)
					}
					descriptor, err := generator.Type(nf.Name, nf.Field, set.path, g.Imports.Add)
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
				if typ.Interface {
					continue
				}
				jsonRegisterType := generator.Reference("kego.io/json", "Register", set.path, g.Imports.Add)
				pkg := strconv.Quote(typ.Id.Package)
				name := strconv.Quote(typ.Id.Name)
				reflectTypeOf := generator.Reference("reflect", "TypeOf", set.path, g.Imports.Add)
				typOf := fmt.Sprintf("%s(&%s{})", reflectTypeOf, system.GoName(typ.Id.Name))
				// e.g.
				// json.Register("kego.io/gallery/data", "@gallery", reflect.TypeOf(&Gallery_rule{}))
				// json.Register("kego.io/gallery/data", "gallery", reflect.TypeOf(&Gallery{}))
				g.Printf("%s(%s, %s, %s, %#v)\n", jsonRegisterType, pkg, name, typOf, hashed.Hash)
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
				if err := process.ValidateCommand(set); err != nil {
					fmt.Println(process.FormatError(err))
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
