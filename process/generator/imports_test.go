package generator

import (
	"testing"

	"fmt"

	"kego.io/kerr/assert"
)

func TestImports(t *testing.T) {
	i := imports{}
	i.add("a.b/c", false)
	assert.Equal(t, "c", i["a.b/c"])
	i.add("a.b/c", true)
	assert.Equal(t, 1, len(i))
	i.add("d.e/f", true)
	assert.Equal(t, "_", i["d.e/f"])
	i.add("d.e/f", true)
	assert.Equal(t, 2, len(i))
	i.add("g.h/f", true)
	assert.Equal(t, "_", i["d.e/f"])
	assert.Equal(t, "_", i["g.h/f"])
	i.add("i.j/f", false)
	assert.Equal(t, "f", i["i.j/f"])
	i.add("k.l/f", false)
	assert.Equal(t, "f1", i["k.l/f"])
	i.add("m.n/f", false)
	assert.Equal(t, "f2", i["m.n/f"])
	i.add("o.p/map", false)
	assert.Equal(t, "map1", i["o.p/map"])
	i.add("q.r/map", false)
	assert.Equal(t, "map2", i["q.r/map"])

	i.add("a.a/a", false)
	p := i.packages()
	assert.Equal(t, 9, len(p))
	assert.Equal(t, "[a.a/a a.b/c d.e/f g.h/f i.j/f k.l/f m.n/f o.p/map q.r/map]", fmt.Sprintf("%v", p))

}
