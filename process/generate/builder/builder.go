package builder // import "kego.io/process/generate/builder"

// ke: {"package": {"complete": true}}

import (
	"bytes"
	"fmt"
	"go/format"

	"strings"

	"github.com/davelondon/kerr"
)

type Builder struct {
	path       string
	name       string
	comment    string
	intro      string
	Imports    Imports
	statements []string
	buffer     *bytes.Buffer
}

func New(path string) *Builder {
	name := getPackageName(path)
	return &Builder{path: path, name: name, Imports: Imports{}, buffer: bytes.NewBuffer(nil)}
}
func WithName(path string, name string) *Builder {
	return &Builder{path: path, name: name, Imports: Imports{}, buffer: bytes.NewBuffer(nil)}
}

func (g *Builder) SetPackageComment(comment string) {
	g.comment = comment
}

func (g *Builder) SetIntroComment(comment string) {
	g.intro = comment
}

func (g *Builder) Print(args ...interface{}) *Builder {
	g.statements = append(g.statements, fmt.Sprint(args...))
	return g
}
func (g *Builder) Printf(f string, args ...interface{}) *Builder {
	g.statements = append(g.statements, fmt.Sprintf(f, args...))
	return g
}
func (g *Builder) Println(args ...interface{}) *Builder {
	g.statements = append(g.statements, fmt.Sprint(args...)+"\n")
	return g
}
func (g *Builder) PrintMethodCall(name string, method string, args ...interface{}) *Builder {
	g.statements = append(g.statements, g.SprintMethodCall(name, method, args...))
	return g
}
func (g *Builder) SprintMethodCall(name string, method string, args ...interface{}) string {
	funcName := fmt.Sprint(name, ".", method)
	argsList := []string{}
	for _, arg := range args {
		argsList = append(argsList, fmt.Sprint(arg))
	}
	return fmt.Sprintf("%s(%s)", funcName, strings.Join(argsList, ", "))
}
func (g *Builder) PrintFunctionCall(path string, name string, args ...interface{}) *Builder {
	g.statements = append(g.statements, g.SprintFunctionCall(path, name, args...))
	return g
}
func (g *Builder) SprintFunctionCall(path string, name string, args ...interface{}) string {
	funcName := Reference(path, name, g.path, g.Imports.Add)
	argsList := []string{}
	for _, arg := range args {
		argsList = append(argsList, fmt.Sprint(arg))
	}
	return fmt.Sprintf("%s(%s)", funcName, strings.Join(argsList, ", "))
}

func (g *Builder) Build() ([]byte, error) {
	b := g.buffer
	if len(g.comment) > 0 {
		fmt.Fprintf(b, "// %s\n", g.comment)
	}
	fmt.Fprintf(b, "package %s", g.name)

	if len(g.intro) > 0 {
		fmt.Fprintf(b, "\n\n")
		fmt.Fprintf(b, "// %s\n", g.intro)
	}

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
		return nil, kerr.Wrap("CRBYOUOHPG", err)
	}
	return source, nil
}
