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

func Generate(file fileType, path string, imports map[string]string) (source []byte, err error) {
	b := bytes.NewBuffer(nil)
	switch file {
	case F_CMD_MAIN:
		g := generator.New(path, "main", b)
		g.Import("os")
		g.Import("fmt")
		g.Import("kego.io/process")
		g.AnonymousImport("kego.io/system")
		if path != "kego.io/system" {
			g.AnonymousImport("kego.io/system/types")
		}
		for p, _ := range imports {
			g.AnonymousImport(p)
			g.AnonymousImport(fmt.Sprint(p, "/types"))
		}
		g.Print(`
			func main() {
				dir, test, recursive, verbose, path, imports, err := process.Initialise()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err := process.GenerateFiles(process.F_MAIN, dir, test, recursive, verbose, path, imports); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_CMD_TYPES:
		g := generator.New(path, "main", b)
		g.Import("os")
		g.Import("fmt")
		g.Import("kego.io/process")
		g.AnonymousImport("kego.io/system")
		if path != "kego.io/system" {
			g.AnonymousImport("kego.io/system/types")
		}
		g.AnonymousImport(path)
		for p, _ := range imports {
			g.AnonymousImport(p)
			g.AnonymousImport(fmt.Sprint(p, "/types"))
		}
		g.Print(`
			func main() {
				dir, test, recursive, verbose, path, imports, err := process.Initialise()
				if err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if err := process.GenerateFiles(process.F_TYPES, dir, test, recursive, verbose, path, imports); err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if err := process.GenerateFiles(process.F_GLOBALS, dir, test, recursive, verbose, path, imports); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_CMD_VALIDATE:
		g := generator.New(path, "main", b)
		g.Import("os")
		g.Import("fmt")
		g.Import("kego.io/process")
		g.Import("kego.io/kerr")
		g.AnonymousImport("kego.io/system")
		g.AnonymousImport("kego.io/system/types")
		g.AnonymousImport(path)
		g.AnonymousImport(fmt.Sprint(path, "/types"))
		for p, _ := range imports {
			g.AnonymousImport(p)
			g.AnonymousImport(fmt.Sprint(p, "/types"))
		}
		g.Print(`
			func main() {
				dir, _, recursive, verbose, path, imports, err := process.Initialise()
				if err != nil {
					fmt.Println(err)
			        os.Exit(1)
				}
				if err := process.Validate(dir, recursive, verbose, path, imports); err != nil {
					if u, ok := err.(kerr.UniqueError); ok {
						if m, ok := u.Source().(process.ValidationError); ok {
							fmt.Println("Error: ", m.Message)
							os.Exit(1)
						}
					}
					fmt.Println(err)
					os.Exit(1)
				}
			}`)
		g.Build()
	case F_MAIN:
		types := system.GetAllTypesInPackage(path)
		g := generator.New(path, generator.PackageName(path), b)
		if len(types) == 0 {
			g.Build()
			break
		}
		for _, typ := range types {
			if typ.Interface || typ.IsNativeValue() || typ.Exclude {
				continue
			}
			if typ.Description != "" {
				g.Println("// ", typ.Description)
			}
			g.Println("type ", system.GoName(typ.Id.Name), " struct {")
			{
				if !typ.Basic {
					g.Println("*", generator.Reference("kego.io/system", "Base", path, g.Import))
				}
				embeds := system.SortableReferences(typ.Embed)
				sort.Sort(embeds)
				for _, embed := range embeds {
					g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), path, g.Import))
				}
				for _, nf := range typ.SortedFields() {
					b := nf.Field.(system.Object).GetBase()
					if b.Description != "" {
						g.Println("// ", b.Description)
					}
					descriptor, err := generator.Type(nf.Field, path, g.Import)
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
			for _, typ := range types {
				if typ.Interface {
					continue
				}
				jsonRegisterType := generator.Reference("kego.io/json", "RegisterType", path, g.Import)
				pkg := strconv.Quote(typ.Id.Package)
				name := strconv.Quote(typ.Id.Name)
				reflectTypeOf := generator.Reference("reflect", "TypeOf", path, g.Import)
				typOf := fmt.Sprintf("%s(&%s{})", reflectTypeOf, system.GoName(typ.Id.Name))
				// e.g.
				// json.RegisterType("kego.io/gallery/data", "@gallery", reflect.TypeOf(&Gallery_rule{}))
				// json.RegisterType("kego.io/gallery/data", "gallery", reflect.TypeOf(&Gallery{}))
				g.Printf("%s(%s, %s, %s)\n", jsonRegisterType, pkg, name, typOf)
			}
		}
		g.Println("}")
		g.Build()
	case F_TYPES:
		types := system.GetAllTypesInPackage(path)
		path := fmt.Sprintf("%s/types", path)
		g := generator.New(path, "types", b)
		if len(types) == 0 {
			g.Build()
			break
		}
		if path != "kego.io/system/types" {
			g.AnonymousImport("kego.io/system/types")
		}
		for p, _ := range imports {
			g.AnonymousImport(fmt.Sprint(p, "/types"))
		}
		g.Println("func init() {")
		{
			registers := []string{}
			pointers := []literal.Pointer{}
			systemRegisterType := generator.Reference("kego.io/system", "RegisterType", path, g.Import)
			for _, t := range types {
				pkg := strconv.Quote(t.Id.Package)
				name := strconv.Quote(t.Id.Name)
				pointer := literal.Build(t, &pointers, path, g.Import)
				// e.g.
				// system.RegisterType("kego.io/gallery/data", "gallery", ptr8728815248)
				// system.RegisterType("kego.io/gallery/data", "@gallery", ptr8728815360)
				line := fmt.Sprintf("%s(%s, %s, %s)", systemRegisterType, pkg, name, pointer.Name)
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
		globals := system.GetAllGlobalsInPackage(path)
		if len(globals) == 0 {
			return
		}
		g := generator.New(path, generator.PackageName(path), b)

		literals := []string{}
		pointers := []literal.Pointer{}
		for _, global := range globals {
			pointer := literal.Build(global, &pointers, path, g.Import)
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
