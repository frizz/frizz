package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestObjectGetType(t *testing.T) {
	ty := &Type{Base: &Base{Id: NewReference("a.b/c", "foo")}}

	RegisterType("a.b/c", "t", ty)

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("a.b/c", "t")

	o := &Base{Type: NewReference("a.b/c", "t")}

	gt, ok := o.Type.GetType()
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:foo", gt.Id.Value())

}
