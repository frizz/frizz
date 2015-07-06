package process

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"

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
				for _, embed := range typ.Embed {
					g.Println("*", generator.Reference(embed.Package, system.GoName(embed.Name), path, g.Import))
				}
				for name, field := range typ.Fields {
					b := field.(system.Object).GetBase()
					if b.Description != "" {
						g.Println("// ", b.Description)
					}
					descriptor, err := generator.Type(field, path, g.Import)
					if err != nil {
						return nil, kerr.New("GDSKJDEKQD", err, "process.Generate", "generator.Type")
					}
					g.Println(system.GoName(name), " ", descriptor)
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
			typesSource := map[system.Reference]string{}
			pointersOrder := []string{}
			pointersMap := map[string]string{}
			for r, t := range types {
				typesSource[r] = literal.Build(t, pointersMap, &pointersOrder, path, g.Import)
			}
			pointers := orderPointers(pointersOrder, pointersMap)
			for _, p := range pointers {
				g.Println(p.Name, " := ", p.Source)
			}
			for ref, source := range typesSource {
				systemRegisterType := generator.Reference("kego.io/system", "RegisterType", path, g.Import)
				pkg := strconv.Quote(ref.Package)
				name := strconv.Quote(ref.Name)
				// e.g.
				// system.RegisterType("kego.io/gallery/data", "gallery", ptr8728815248)
				// system.RegisterType("kego.io/gallery/data", "@gallery", ptr8728815360)
				g.Printf("%s(%s, %s, %s)\n", systemRegisterType, pkg, name, source)
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

		globalsSource := map[system.Reference]string{}
		pointersOrder := []string{}
		pointersMap := map[string]string{}
		for ref, global := range globals {
			globalsSource[ref] = literal.Build(global, pointersMap, &pointersOrder, path, g.Import)
		}
		pointers := orderPointers(pointersOrder, pointersMap)
		for _, p := range pointers {
			g.Println("var ", p.Name, " = ", p.Source)
		}
		for ref, source := range globalsSource {
			g.Println("var ", system.GoName(ref.Name), " = ", source)
		}
		g.Build()
	}
	source, err = format.Source(b.Bytes())
	if err != nil {
		err = kerr.New("CRBYOUOHPG", err, "process.Generate", "format.Source")
	}
	return
}

func orderPointers(pointersOrder []string, pointersMap map[string]string) []pointer {
	pointers := []pointer{}
	for _, name := range pointersOrder {
		pointers = append(pointers, pointer{name, pointersMap[name]})
	}
	return pointers
}

type pointer struct {
	Name   string
	Source string
}
