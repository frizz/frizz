package generator // import "kego.io/process/generator"

import (
	"fmt"
	"io"
	"strings"
)

type generator struct {
	path       string
	alias      string
	imports    imports
	statements []string
	writer     io.Writer
}

func New(path string, alias string, writer io.Writer) *generator {
	return &generator{path: path, alias: alias, imports: map[string]string{}, writer: writer}
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

func (g *generator) AnonymousImport(path string) {
	g.imports.add(path, true)
}

func (g *generator) Import(path string) string {
	return g.imports.add(path, false)
}

func (g *generator) Build() {
	w := g.writer
	fmt.Fprintf(w, "package %s", g.alias)
	if len(g.imports) > 0 {
		fmt.Fprintf(w, "\n\n")
		fmt.Fprintf(w, "import (")
		// The packages should build in the same order to make our tests work reproducibly, so
		// we use the sorted packages() array rather than ranging over the map
		for _, path := range g.imports.packages() {
			fmt.Fprintf(w, "\n")
			fmt.Fprintf(w, "%s \"%s\"", g.imports[path], path)
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, ")\n")
	}
	for _, statement := range g.statements {
		fmt.Fprintf(w, "%s", statement)
	}
}

// getPackageNameFromPath returns the name of the package, given the package
// path. Note this is golang packages not file system paths, so we always use
// forward slash instead of os.PathSeparator
// TODO: Lots of packages have a different name to the path...
// TODO: Work out what to do here.
func PackageName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
