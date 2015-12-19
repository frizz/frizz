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

	env := envctx.FromContext(ctx)

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("kego.io/process")
	g.Imports.Add("kego.io/process/generate/commands/structstypes")
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
			structstypes.Main(process.S_STRUCTS)
		}`)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}

func TypesCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("kego.io/process")
	g.Imports.Add("kego.io/process/generate/commands/structstypes")
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
			structstypes.Main(process.S_TYPES)
		}`)

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("GQBDCCNMIA", err, "Build")
	}
	return b, nil
}

func LocalKeCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("kego.io/process/generate/commands/localke")
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
			localke.Main(
				` + fmt.Sprint(cmd.Recursive) + `,
				` + strconv.Quote(env.Path) + `)
		}`)
	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}
