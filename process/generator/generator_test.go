package generator

import (
	"testing"

	"bytes"

	"kego.io/kerr/assert"
)

func TestPackageName(t *testing.T) {
	n := PackageName("a.b/c")
	assert.Equal(t, "c", n)

	// Double check we're not splitting on backslash
	n = PackageName(`a.b/c\d`)
	assert.Equal(t, `c\d`, n)
}

func TestNew(t *testing.T) {
	b := bytes.NewBuffer(nil)
	g := New("a.b/c", "d", b)
	assert.Equal(t, 0, len(g.imports))
	assert.Equal(t, "a.b/c", g.path)
	assert.Equal(t, "d", g.alias)
}

func TestGenerator(t *testing.T) {
	b := bytes.NewBuffer(nil)
	g := New("a.b/c", "d", b)
	g.Build()
	assert.Equal(t, "package d", b.String())
	b.Reset()

	g.AnonymousImport("e.f/g")
	g.Build()
	assert.Equal(t, "package d\n\nimport (\n_ \"e.f/g\"\n)\n", b.String())
	b.Reset()

	alias := g.Import("h.i/j")
	assert.Equal(t, "j", alias)

	g.Build()
	assert.Equal(t, "package d\n\nimport (\n_ \"e.f/g\"\nj \"h.i/j\"\n)\n", b.String())
	b.Reset()

	g.Print("foo")
	g.Println("bar")
	g.Printf("baz\n%s", "qux")
	g.Build()
	assert.Equal(t, "package d\n\nimport (\n_ \"e.f/g\"\nj \"h.i/j\"\n)\nfoobar\nbaz\nqux", b.String())
}
