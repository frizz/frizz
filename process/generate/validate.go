package generate

import (
	"strconv"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/process/generate/builder"
)

func ValidateCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	g := builder.WithName(env.Path, "main")
	g.Imports.Add("kego.io/process/validate/command")
	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous(env.Path)
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
	}
	g.Print(`
		func main() {
			command.ValidateMain(` + strconv.Quote(env.Path) + `)
		}`)
	b, err := g.Build()
	if err != nil {
		return nil, kerr.Wrap("IIIRBBXASR", err)
	}
	return b, nil
}
