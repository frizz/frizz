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
	g.Imports.Add("kego.io/process/generate/commands/structscmd")
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
			structscmd.Run()
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
	g.Imports.Add("kego.io/process/generate/commands/typescmd")
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
			typescmd.Run()
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
	g.Imports.Add("kego.io/process/generate/commands/kecmd")
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
			kecmd.Run(
				` + fmt.Sprint(cmd.Recursive) + `,
				` + strconv.Quote(env.Path) + `)
		}`)
	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}
