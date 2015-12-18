package generate

import (
	"fmt"
	"strconv"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/generator"
	"kego.io/kerr"
)

func StructsCommand(ctx context.Context) (source []byte, err error) {

	env, ok := envctx.FromContext(ctx)
	if !ok {
		return nil, kerr.New("NEFDVFABBD", nil, "No env in ctx")
	}

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/process")
	g.Imports.Anonymous("kego.io/system")
	if env.Path != "kego.io/system" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Print(`
		func main() {
			ctx, err := process.InitialiseAutomatic()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err := process.Generate(ctx, process.S_STRUCTS); err != nil {
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

func TypesCommand(ctx context.Context) (source []byte, err error) {

	env, ok := envctx.FromContext(ctx)
	if !ok {
		return nil, kerr.New("TNXQEXJDVH", nil, "No env in ctx")
	}

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/process")
	g.Imports.Anonymous("kego.io/system")
	if env.Path != "kego.io/system" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	g.Imports.Anonymous(env.Path)
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Print(`
		func main() {
			ctx, err := process.InitialiseAutomatic()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err := process.Generate(ctx, process.S_TYPES); err != nil {
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

func KeCommand(ctx context.Context) (source []byte, err error) {

	env, ok := envctx.FromContext(ctx)
	if !ok {
		return nil, kerr.New("BIGGOLIGKS", nil, "No env in ctx")
	}

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		return nil, kerr.New("OGLQBVLGSV", nil, "No cmd in ctx")
	}

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("os")
	g.Imports.Add("fmt")
	g.Imports.Add("kego.io/context/cmdctx")
	g.Imports.Add("kego.io/context/envctx")
	g.Imports.Add("kego.io/process")
	g.Imports.Add("kego.io/editor/server")
	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous("kego.io/system/types")
	g.Imports.Anonymous(env.Path)
	g.Imports.Anonymous(fmt.Sprint(env.Path, "/types"))
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Print(`
		func main() {
			update := false
			recursive := ` + fmt.Sprint(cmd.Recursive) + `
			path := ` + strconv.Quote(env.Path) + `
			ctx, err := process.InitialiseCommand(update, recursive, path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			env, ok := envctx.FromContext(ctx)
			if !ok {
				fmt.Println("No env in ctx")
				os.Exit(1)
			}
			cmd, ok := cmdctx.FromContext(ctx)
			if !ok {
				fmt.Println("No cmd in ctx")
				os.Exit(1)
			}
			if typesChanged, err := process.ValidateCommand(ctx); err != nil {
				if !typesChanged || !cmd.Verbose {
					// when ValidateCommand detects the types have changed, it
					// spawns a ke command. In verbose mode this will output any
					// error so we don't need to print the error
					fmt.Println(process.FormatError(err))
				}
				os.Exit(1)
			}
			if cmd.Edit {
				if err = server.Start(env.Path, cmd.Verbose, false); err != nil {
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
