package generator // import "kego.io/process/generator"

import (
	"fmt"
	"io"
)

type generator struct {
	path       string
	name       string
	Imports    Imports
	statements []string
	writer     io.Writer
}

func New(path string, writer io.Writer) (*generator, error) {
	name := getPackageName(path)
	return &generator{path: path, name: name, Imports: Imports{}, writer: writer}, nil
}
func NewWithName(path string, name string, writer io.Writer) *generator {
	return &generator{path: path, name: name, Imports: Imports{}, writer: writer}
}

func (g *generator) Print(args ...interface{}) *generator {
	g.statements = append(g.statements, fmt.Sprint(args...))
	return g
}
func (g *generator) Printf(f string, args ...interface{}) *generator {
	g.statements = append(g.statements, fmt.Sprintf(f, args...))
	return g
}
func (g *generator) Println(args ...interface{}) *generator {
	g.statements = append(g.statements, fmt.Sprint(args...)+"\n")
	return g
}

func (g *generator) Build() {
	w := g.writer
	fmt.Fprintf(w, "package %s", g.name)
	if len(g.Imports) > 0 {
		fmt.Fprintf(w, "\n\n")
		fmt.Fprintf(w, "import (")
		// The packages should build in the same order to make our tests work reproducibly, so
		// we use the sorted packages() array rather than ranging over the map
		for _, path := range g.Imports.packages() {
			fmt.Fprintf(w, "\n")
			imp := g.Imports[path]
			alias := ""
			if imp.Alias != imp.Name {
				// if the alias is the same as the package name, we can omit the alias
				alias = imp.Alias
			}
			fmt.Fprintf(w, "%s \"%s\"", alias, path)
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, ")\n")
	}
	for _, statement := range g.statements {
		fmt.Fprintf(w, "%s", statement)
	}
}
