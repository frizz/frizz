package generator

import (
	"testing"

	"bytes"

	"kego.io/kerr/assert"
)

func TestNew(t *testing.T) {
	b := bytes.NewBuffer(nil)
	g, err := New("a.b/c", b)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(g.Imports))
	assert.Equal(t, "a.b/c", g.path)
	assert.Equal(t, "c", g.name)
}

func TestNewWithName(t *testing.T) {
	b := bytes.NewBuffer(nil)
	g := NewWithName("a.b/c", "d", b)
	assert.Equal(t, 0, len(g.Imports))
	assert.Equal(t, "a.b/c", g.path)
	assert.Equal(t, "d", g.name)
}

func TestGenerator(t *testing.T) {
	b := bytes.NewBuffer(nil)
	g, err := New("a.b/c", b)
	assert.NoError(t, err)
	g.Build()
	assert.Equal(t, "package c", b.String())
	b.Reset()

	g.Imports.Anonymous("e.f/g")
	g.Build()
	assert.Equal(t, "package c\n\nimport (\n_ \"e.f/g\"\n)\n", b.String())
	b.Reset()

	alias := g.Imports.Add("h.i/j")
	assert.Equal(t, "j", alias)

	g.Build()
	assert.Equal(t, "package c\n\nimport (\n_ \"e.f/g\"\n \"h.i/j\"\n)\n", b.String())
	b.Reset()

	g.Print("foo")
	g.Println("bar")
	g.Printf("baz\n%s", "qux")
	g.Build()
	assert.Equal(t, "package c\n\nimport (\n_ \"e.f/g\"\n \"h.i/j\"\n)\nfoobar\nbaz\nqux", b.String())
}
