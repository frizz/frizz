package builder

import (
	"testing"

	"fmt"

	"kego.io/kerr/assert"
)

func TestImports(t *testing.T) {
	i := Imports{}
	i.add("a.b/c", false)
	assert.Equal(t, "c", i["a.b/c"].Alias)
	i.add("a.b/c", true)
	assert.Equal(t, 1, len(i))
	i.add("d.e/f", true)
	assert.Equal(t, "_", i["d.e/f"].Alias)
	i.add("d.e/f", true)
	assert.Equal(t, 2, len(i))
	i.add("g.h/f", true)
	assert.Equal(t, "_", i["d.e/f"].Alias)
	assert.Equal(t, "_", i["g.h/f"].Alias)
	i.add("i.j/f", false)
	assert.Equal(t, "f", i["i.j/f"].Alias)
	i.add("k.l/f", false)
	assert.Equal(t, "f1", i["k.l/f"].Alias)
	i.add("m.n/f", false)
	assert.Equal(t, "f2", i["m.n/f"].Alias)
	i.add("o.p/map", false)
	assert.Equal(t, "map1", i["o.p/map"].Alias)
	i.add("q.r/map", false)
	assert.Equal(t, "map2", i["q.r/map"].Alias)

	i.add("a.a/a", false)
	p := i.packages()
	assert.Equal(t, 9, len(p))
	assert.Equal(t, "[a.a/a a.b/c d.e/f g.h/f i.j/f k.l/f m.n/f o.p/map q.r/map]", fmt.Sprintf("%v", p))

	i.add("s.t/u", true)
	assert.Equal(t, "_", i["s.t/u"].Alias)
	i.add("s.t/u", false)
	assert.Equal(t, "u", i["s.t/u"].Alias)

}

func TestImports1(t *testing.T) {
	i := Imports{}
	i.Anonymous("a.b/c")

}
