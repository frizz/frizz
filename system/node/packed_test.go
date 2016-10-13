package node

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/packer"
)

func TestPackedType(t *testing.T) {
	var p *packed
	assert.Equal(t, packer.J_NULL, p.Type())
	p = &packed{n: nil}
	assert.Equal(t, packer.J_NULL, p.Type())
	p = &packed{n: &Node{JsonType: packer.J_OBJECT}}
	assert.Equal(t, packer.J_MAP, p.Type())
	p = &packed{n: &Node{JsonType: packer.J_STRING}}
	assert.Equal(t, packer.J_STRING, p.Type())
}

func TestPackedString(t *testing.T) {
	p := &packed{n: &Node{ValueString: "a"}}
	assert.Equal(t, "a", p.String())
}

func TestPackedNumber(t *testing.T) {
	p := &packed{n: &Node{ValueNumber: 2.0}}
	assert.Equal(t, 2.0, p.Number())
}

func TestPackedBool(t *testing.T) {
	p := &packed{n: &Node{ValueBool: true}}
	assert.Equal(t, true, p.Bool())
}

func TestPackedArray(t *testing.T) {
	n1 := &Node{ValueString: "n1"}
	n2 := &Node{ValueString: "n2"}
	p := &packed{n: &Node{Array: []*Node{n1, n2}}}
	a := p.Array()
	assert.Equal(t, 2, len(a))
	assert.Equal(t, "n1", a[0].String())
	assert.Equal(t, "n2", a[1].String())
}

func TestPackedMap(t *testing.T) {
	n1 := &Node{ValueString: "n1"}
	n2 := &Node{ValueString: "n2", Missing: true}
	n3 := &Node{ValueString: "n3"}
	p := &packed{n: &Node{Map: map[string]*Node{"n1": n1, "n2": n2, "n3": n3}}}
	m := p.Map()
	assert.Equal(t, 2, len(m))
	assert.Equal(t, "n1", m["n1"].String())
	assert.Equal(t, "n3", m["n3"].String())
}
func TestPackedInterface(t *testing.T) {
	p := &packed{n: &Node{Value: "a"}}
	assert.Equal(t, "a", p.Interface())
}
