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
	p = &packed{v: "a"}
	assert.Equal(t, "a", p.Interface())
}
