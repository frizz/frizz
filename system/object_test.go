package system

import (
	"testing"

	"kego.io/assert"
)

func TestObjectSetContext(t *testing.T) {
	o := &Object{
		Context:     &Context{Package: "a.b/c", Imports: map[string]string{"d": "e.f/g"}},
		Description: "h",
		Id:          "i",
		Type:        NewReference("j.k/l", "m")}
	o.SetContext("n.o/p", map[string]string{"q": "r.s/t"})
	assert.Equal(t, "n.o/p", o.Context.Package)
	assert.Equal(t, 1, len(o.Context.Imports))
	assert.Equal(t, "r.s/t", o.Context.Imports["q"])
}

func TestObjectGetType(t *testing.T) {
	ty := &Type{Object: &Object{Id: "foo"}}

	RegisterType("a.b/c:t", ty)

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("a.b/c:t")

	o := &Object{Type: NewReference("a.b/c", "t")}

	gt, ok := o.GetType()
	assert.True(t, ok)
	assert.Equal(t, "foo", gt.Id)

}
