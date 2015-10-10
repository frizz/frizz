package generator

import (
	"testing"

	"kego.io/kerr/assert"
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

func TestGenerator(t *testing.T) {
	g := New("a.b/c")
	b, err := g.Build()
	assert.NoError(t, err)
	assert.Equal(t, "package c\n", string(b))
	g.buffer.Reset()

	g.Imports.Anonymous("e.f/g")
	b, err = g.Build()
	assert.NoError(t, err)
	assert.Equal(t, "package c\n\nimport (\n\t_ \"e.f/g\"\n)\n", string(b))
	g.buffer.Reset()

	alias := g.Imports.Add("h.i/j")
	assert.Equal(t, "j", alias)

	b, err = g.Build()
	assert.NoError(t, err)
	assert.Equal(t, "package c\n\nimport (\n\t_ \"e.f/g\"\n\t\"h.i/j\"\n)\n", string(b))
	g.buffer.Reset()

	g.Print("var ")
	g.Println("foo string")
	g.Printf("var bar int\n%s", "var baz bool")
	b, err = g.Build()
	assert.NoError(t, err)
	assert.Equal(t, "package c\n\nimport (\n\t_ \"e.f/g\"\n\t\"h.i/j\"\n)\n\nvar foo string\nvar bar int\nvar baz bool\n", string(b))
}
