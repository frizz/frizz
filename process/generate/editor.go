package generate

import (
	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/process/generate/builder"
)

func Editor(ctx context.Context, env *envctx.Env) (source []byte, err error) {

	g := builder.New("main")

	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous(env.Path)
	for _, p := range env.Aliases {
		g.Imports.Anonymous(p)
	}
	/*
		func main() {
			if err := client.Start(); err != nil {
				console.Error(err.Error())
			}
		}
	*/
	g.Println("func main() {")
	{
		clientStart := builder.Reference("kego.io/editor/client", "Start", env.Path, g.Imports.Add)
		g.Println("if err := ", clientStart, "(); err != nil {")
		{
			fmtPrintln := builder.Reference("fmt", "Println", env.Path, g.Imports.Add)
			g.Println(fmtPrintln, "(err.Error())")
		}
		g.Println("}")
	}
	g.Println("}")

	b, err := g.Build()
	if err != nil {
		return nil, kerr.Wrap("CBTOLUQOGL", err)
	}
	return b, nil
}
