package system

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestObjectGetType(t *testing.T) {
	ty := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}

	o := &Object{Type: NewReference("a.b/c", "t")}

	gt, ok := o.Type.GetType(tests.Context("a.b/c").Ctype("t", ty).Ctx())
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:foo", gt.Id.Value())

}
