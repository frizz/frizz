package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestBaseSetContext(t *testing.T) {
	o := &Base{
		Context:     &Context{Package: "a.b/c", Imports: map[string]string{"d": "e.f/g"}},
		Description: "h",
		Id:          NewReference("n.o/p", "i"),
		Type:        NewReference("j.k/l", "m")}
	o.SetContext("n.o/p", map[string]string{"q": "r.s/t"})
	assert.Equal(t, "n.o/p", o.Context.Package)
	assert.Equal(t, 1, len(o.Context.Imports))
	assert.Equal(t, "r.s/t", o.Context.Imports["q"])
}

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
