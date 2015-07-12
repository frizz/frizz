package process

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"

	"sort"

	"kego.io/kerr"
	"kego.io/literal"
	"kego.io/process/generator"
	"kego.io/system"
)

func Generate(file fileType, set settings) (source []byte, err error) {
	b := bytes.NewBuffer(nil)
	switch file {
	case F_CMD_MAIN:
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
				if err := process.GenerateFiles(process.F_MAIN, set); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_CMD_TYPES:
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
				if err := process.GenerateFiles(process.F_TYPES, set); err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if err := process.GenerateFiles(process.F_GLOBALS, set); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_CMD_VALIDATE:
		g := generator.NewWithName(set.path, "main", b)
		g.Imports.Add("os")
		g.Imports.Add("fmt")
		g.Imports.Add("kego.io/process")
		g.Imports.Add("kego.io/kerr")
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
				set, err := process.InitialiseManually(false, ` + fmt.Sprint(set.recursive) + `, false, ` + strconv.Quote(set.path) + `)
				if err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if err := process.Validate(set); err != nil {
					if u, ok := err.(kerr.UniqueError); ok {
						if m, ok := u.Source().(process.ValidationError); ok {
							fmt.Println("Error:", m.Message)
							os.Exit(1)
						}
					}
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_MAIN:
		types := system.GetAllTypesInPackage(set.path)
		g, err := generator.New(set.path, b)
		if err != nil {
			return nil, kerr.New("WPIPODJIGV", err, "process.Generate", "generator.New")
		}
		if len(types) == 0 {
			g.Build()
			break
		}
		for _, typeDef := range types {
			typ := typeDef.Type
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
					descriptor, err := generator.Type(nf.Field, set.path, g.Imports.Add)
					if err != nil {
						return nil, kerr.New("GDSKJDEKQD", err, "process.Generate", "generator.Type")
					}
					g.Println(system.GoName(nf.Name), " ", descriptor)
				}
			}
			g.Println("}")
		}
		g.Println("func init() {")
		{
			for _, typeDef := range types {
				typ := typeDef.Type
				if typ.Interface {
					continue
				}
				jsonRegisterType := generator.Reference("kego.io/json", "RegisterType", set.path, g.Imports.Add)
				pkg := strconv.Quote(typ.Id.Package)
				name := strconv.Quote(typ.Id.Name)
				reflectTypeOf := generator.Reference("reflect", "TypeOf", set.path, g.Imports.Add)
				typOf := fmt.Sprintf("%s(&%s{})", reflectTypeOf, system.GoName(typ.Id.Name))
				// e.g.
				// json.RegisterType("kego.io/gallery/data", "@gallery", reflect.TypeOf(&Gallery_rule{}))
				// json.RegisterType("kego.io/gallery/data", "gallery", reflect.TypeOf(&Gallery{}))
				g.Printf("%s(%s, %s, %s, %#v)\n", jsonRegisterType, pkg, name, typOf, typeDef.Hash)
			}
		}
		g.Println("}")
		g.Build()
	case F_TYPES:
		types := system.GetAllTypesInPackage(set.path)
		typesPath := fmt.Sprintf("%s/types", set.path)
		g, err := generator.New(typesPath, b)
		if err != nil {
			return nil, kerr.New("KYIKJQXPMR", err, "process.Generate", "generator.New")
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
			systemRegisterType := generator.Reference("kego.io/system", "RegisterType", typesPath, g.Imports.Add)
			for _, typeDef := range types {
				t := typeDef.Type
				pkg := strconv.Quote(t.Id.Package)
				name := strconv.Quote(t.Id.Name)
				pointer := literal.Build(t, &pointers, typesPath, g.Imports.Add)
				// e.g.
				// system.RegisterType("kego.io/gallery/data", "gallery", ptr8728815248)
				// system.RegisterType("kego.io/gallery/data", "@gallery", ptr8728815360)
				line := fmt.Sprintf("%s(%s, %s, %s, %#v)", systemRegisterType, pkg, name, pointer.Name, typeDef.Hash)
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
	case F_GLOBALS:
		globals := system.GetAllGlobalsInPackage(set.path)
		if len(globals) == 0 {
			return
		}
		g, err := generator.New(set.path, b)
		if err != nil {
			return nil, kerr.New("WMGMMHHQGX", err, "process.Generate", "generator.New")
		}

		literals := []string{}
		pointers := []literal.Pointer{}
		for _, global := range globals {
			pointer := literal.Build(global, &pointers, set.path, g.Imports.Add)
			line := fmt.Sprint("var ", system.GoName(global.GetBase().Id.Name), " = ", pointer.Name)
			literals = append(literals, line)
		}
		for _, p := range pointers {
			g.Println("var ", p.Name, " = ", p.Source)
		}
		for _, s := range literals {
			g.Println(s)
		}
		g.Build()
	}
	source, err = format.Source(b.Bytes())
	if err != nil {
		err = kerr.New("CRBYOUOHPG", err, "process.Generate", "format.Source: %s", b.String())
	}
	return
}
