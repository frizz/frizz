package json

import (
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestPackedType(t *testing.T) {
	var p *packed
	assert.Equal(t, J_NULL, p.Type())
	p = &packed{v: nil}
	assert.Equal(t, J_NULL, p.Type())
	assert.Equal(t, "", p.String())
	assert.Equal(t, 0.0, p.Number())
	assert.Equal(t, false, p.Bool())
	p = &packed{v: "a"}
	assert.Equal(t, "a", p.Interface())
	p = PackString(`"s"`)
	assert.Equal(t, "s", p.Interface())
}
