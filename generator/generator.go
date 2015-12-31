package generator // import "kego.io/generator"

import (
	"bytes"
	"fmt"
	"go/format"

	"strings"

	"kego.io/kerr"
)

type Generator struct {
	path       string
	name       string
	comment    string
	Imports    Imports
	statements []string
	buffer     *bytes.Buffer
}

func New(path string) *Generator {
	name := getPackageName(path)
	return &Generator{path: path, name: name, Imports: Imports{}, buffer: bytes.NewBuffer(nil)}
}
func WithName(path string, name string) *Generator {
	return &Generator{path: path, name: name, Imports: Imports{}, buffer: bytes.NewBuffer(nil)}
}

func (g *Generator) SetPackageComment(comment string) {
	g.comment = comment
}

func (g *Generator) Print(args ...interface{}) *Generator {
	g.statements = append(g.statements, fmt.Sprint(args...))
	return g
}
func (g *Generator) Printf(f string, args ...interface{}) *Generator {
	g.statements = append(g.statements, fmt.Sprintf(f, args...))
	return g
}
func (g *Generator) Println(args ...interface{}) *Generator {
	g.statements = append(g.statements, fmt.Sprint(args...)+"\n")
	return g
}
func (g *Generator) PrintFunctionCall(path string, name string, args ...interface{}) *Generator {
	g.statements = append(g.statements, g.SprintFunctionCall(path, name, args...))
	return g
}
func (g *Generator) SprintFunctionCall(path string, name string, args ...interface{}) string {
	funcName := Reference(path, name, g.path, g.Imports.Add)
	argsList := []string{}
	for _, arg := range args {
		argsList = append(argsList, fmt.Sprint(arg))
	}
	return fmt.Sprintf("%s(%s)", funcName, strings.Join(argsList, ", "))
}

func (g *Generator) Build() ([]byte, error) {
	b := g.buffer
	if len(g.comment) > 0 {
		fmt.Fprintf(b, "// %s\n", g.comment)
	}
	fmt.Fprintf(b, "package %s", g.name)
	if len(g.Imports) > 0 {
		fmt.Fprintf(b, "\n\n")
		fmt.Fprintf(b, "import (")
		// The packages should build in the same order to make our tests work reproducibly, so
		// we use the sorted packages() array rather than ranging over the map
		for _, path := range g.Imports.packages() {
			fmt.Fprintf(b, "\n")
			imp := g.Imports[path]
			alias := ""
			if imp.Alias != imp.Name {
				// if the alias is the same as the package name, we can omit the alias
				alias = imp.Alias
			}
			fmt.Fprintf(b, "%s \"%s\"", alias, path)
		}
		fmt.Fprintf(b, "\n")
		fmt.Fprintf(b, ")\n")
	}
	for _, statement := range g.statements {
		fmt.Fprintf(b, "%s", statement)
	}

	source, err := format.Source(b.Bytes())
	if err != nil {
		return nil, kerr.New("CRBYOUOHPG", err, "format.Source: %s", b.String())
	}
	return source, nil
}
