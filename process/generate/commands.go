package generate

import (
	"strconv"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/generator"
	"kego.io/kerr"
)

func ValidateCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	g := generator.WithName(env.Path, "main")
	g.Imports.Add("kego.io/process/generate/commands/validate")
	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous(env.Path)
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
	}
	g.Print(`
		func main() {
			validate.Main(` + strconv.Quote(env.Path) + `)
		}`)
	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "Build")
	}
	return b, nil
}
