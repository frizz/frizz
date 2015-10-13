package generate

import (
	"fmt"
	"strconv"

	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/process/settings"
)

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

func KeCommand(set *settings.Settings) (source []byte, err error) {

	g := generator.WithName(set.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/process")
	g.Imports.Add("kego.io/editor/server")
	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous("kego.io/system/types")
	g.Imports.Anonymous(set.Path)
	g.Imports.Anonymous(fmt.Sprint(set.Path, "/types"))
	for p, _ := range set.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Print(`
		func main() {
			update := false
			recursive := ` + fmt.Sprint(set.Recursive) + `
			path := ` + strconv.Quote(set.Path) + `
			set, err := process.InitialiseCommand(update, recursive, path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if typesChanged, err := process.ValidateCommand(set); err != nil {
				if !typesChanged || !set.Verbose {
					// when ValidateCommand detects the types have changed, it
					// spawns a ke command. In verbose mode this will output any
					// error so we don't need to print the error
					fmt.Println(process.FormatError(err))
				}
				os.Exit(1)
			}
			if set.Edit {
				if err = server.Start(set.Path, set.Verbose, false); err != nil {
					fmt.Println(process.FormatError(err))
					os.Exit(1)
				}
			}
		}`)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}
