package generate

import (
	"fmt"
	"strconv"

	"kego.io/generator"
	"kego.io/kerr"
	"kego.io/process/settings"
)

func Editor(set *settings.Settings) (source []byte, err error) {

	g := generator.New("main")

	g.Imports.Anonymous("kego.io/system")
	g.Imports.Anonymous("kego.io/system/types")
	g.Imports.Anonymous(set.Path)
	g.Imports.Anonymous(fmt.Sprint(set.Path, "/types"))
	for p, _ := range set.Aliases {
		g.Imports.Anonymous(p)
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	/*
		func main() {
			if err := client.Start("XXX"); err != nil {
				console.Error(err.Error())
			}
		}
	*/
	g.Println("func main() {")
	{
		clientStart := generator.Reference("kego.io/editor/client", "Start", set.Path, g.Imports.Add)
		g.Println("if err := ", clientStart, "(", strconv.Quote(set.Path), "); err != nil {")
		{
			consoleError := generator.Reference("kego.io/js/console", "Error", set.Path, g.Imports.Add)
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
