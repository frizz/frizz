package generate

import (
	"fmt"
	"strconv"

	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/literal"
	"kego.io/process/settings"
	"kego.io/system"
)

func Types(set *settings.Settings) (source []byte, err error) {

	filter := system.NewReference("kego.io/system", "type")
	types := system.GetAllGlobalsInPackage(set.Path, &filter)
	typesPath := fmt.Sprintf("%s/types", set.Path)
	g := generator.New(typesPath)

	if len(types) == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.New("JMHUIFWPWV", err, "Build")
		}
		return b, nil
	}
	if typesPath != "kego.io/system/types" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	for p, _ := range set.Aliases {
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

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("BATUJJFGHT", err, "Build")
	}
	return b, nil
}

func TypesCommand(set *settings.Settings) (source []byte, err error) {

	g := generator.WithName(set.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/process")
	g.Imports.Anonymous("kego.io/system")
	if set.Path != "kego.io/system" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	g.Imports.Anonymous(set.Path)
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
			if err := process.Generate(process.S_TYPES, set); err != nil {
				fmt.Println(process.FormatError(err))
				os.Exit(1)
			}
			if err := process.Generate(process.S_EDITOR, set); err != nil {
				fmt.Println(process.FormatError(err))
				os.Exit(1)
			}
		}`)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("GQBDCCNMIA", err, "Build")
	}
	return b, nil
}
