package generate

import (
	"fmt"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/generator"
	"kego.io/kerr"
)

func Editor(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	g := generator.New("main")

	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous("kego.io/system/types")
	g.Imports.Anonymous("kego.io/system/editors")
	g.Imports.Anonymous(env.Path)
	g.Imports.Anonymous(fmt.Sprint(env.Path, "/types"))
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
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
		clientStart := generator.Reference("kego.io/editor/client", "Start", env.Path, g.Imports.Add)
		g.Println("if err := ", clientStart, "(); err != nil {")
		{
			consoleError := generator.Reference("kego.io/editor/client/console", "Error", env.Path, g.Imports.Add)
			g.Println(consoleError, "(err.Error())")
		}
		g.Println("}")
	}
	g.Println("}")

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("CBTOLUQOGL", err, "Build")
	}
	return b, nil
}
