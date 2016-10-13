package builder

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
)

func TestNew(t *testing.T) {
	g := New("a.b/c")
	assert.Equal(t, 0, len(g.Imports))
	assert.Equal(t, "a.b/c", g.path)
	assert.Equal(t, "c", g.name)
}

func TestNewWithName(t *testing.T) {
	g := WithName("a.b/c", "d")
	assert.Equal(t, 0, len(g.Imports))
	assert.Equal(t, "a.b/c", g.path)
	assert.Equal(t, "d", g.name)
}

func TestBuilder(t *testing.T) {
	g := New("a.b/c")
	b, err := g.Build()
	require.NoError(t, err)
	assert.Equal(t, "package c\n", string(b))
	g.buffer.Reset()

	g.Imports.Anonymous("e.f/g")
	b, err = g.Build()
	require.NoError(t, err)
	assert.Equal(t, "package c\n\nimport (\n\t_ \"e.f/g\"\n)\n", string(b))
	g.buffer.Reset()

	alias := g.Imports.Add("h.i/j")
	assert.Equal(t, "j", alias)

	b, err = g.Build()
	require.NoError(t, err)
	assert.Equal(t, "package c\n\nimport (\n\t_ \"e.f/g\"\n\t\"h.i/j\"\n)\n", string(b))
	g.buffer.Reset()

	g.SetPackageComment("comment")
	g.SetIntroComment("intro")
	g.Print("var ")
	g.Println("foo string")
	g.Printf("var bar int\n%s", "var baz bool")
	g.Println("")
	g.Println("func foo() {")
	g.PrintFunctionCall("k.l/m", "n", "o", "p")
	g.Println("")
	g.PrintMethodCall("a", "b", "c", "d")
	g.Println("}")
	b, err = g.Build()
	require.NoError(t, err)
	assert.Contains(t, string(b), `// comment
package c

// intro

import (
	_ "e.f/g"
	"h.i/j"
	"k.l/m"
)

var foo string
var bar int
var baz bool

func foo() {
	m.n(o, p)
	a.b(c, d)
}
`)

	g.Print("dfskjsdf")
	_, err = g.Build()
	assert.IsError(t, err, "CRBYOUOHPG") // Error formatting source

}
