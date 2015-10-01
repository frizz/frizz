package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestReferenceRuleToParentType(t *testing.T) {

	r := NewReference("a.b/c", "@d")
	rp, err := r.RuleToParentType()
	assert.NoError(t, err)
	assert.Equal(t, "a.b/c", rp.Package)
	assert.Equal(t, "d", rp.Name)
	assert.Equal(t, "a.b/c:d", rp.Value())

	r = Reference{}
	rp, err = r.RuleToParentType()
	assert.IsError(t, err, "OSQKOWGVWX")

	r = NewReference("a.b/c", "d")
	rp, err = r.RuleToParentType()
	assert.IsError(t, err, "HBKCDXQBYG")

}

func TestReferenceUnmarshalJson(t *testing.T) {

	reset := func() *Reference {
		// Let's pre-load with some values so we check that when we
		// load a null value, we clear all the fields
		r := NewReference("a.b/c", "d")
		return &r
	}

	r := reset()
	err := r.Unpack(nil, "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, r.Exists)
	assert.Equal(t, "", r.Package)
	assert.Equal(t, "", r.Name)
	assert.Equal(t, "", r.Value())

	r = reset()
	err = r.Unpack("a.b/c:d", "", map[string]string{})
	assert.EqualError(t, err, "Unknown package a.b/c")
	assert.False(t, r.Exists)

	r = reset()
	err = r.Unpack("a.b/c:d", "", map[string]string{"a.b/c": "c"})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Name)
	assert.Equal(t, "a.b/c:d", r.Value())

	r = reset()
	err = r.Unpack("a.b/c:@d", "", map[string]string{"a.b/c": "c"})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "@d", r.Name)
	assert.Equal(t, "a.b/c:@d", r.Value())

	r = reset()
	err = r.Unpack("a:b", "", map[string]string{})
	assert.False(t, r.Exists)
}

func TestReferenceMarshalJson(t *testing.T) {
	r := NewReference("a.b/c", "d")
	b, err := r.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "\"a.b/c:d\"", string(b))

	r = Reference{}
	b, err = r.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(b))
}

func TestReferenceGetType(t *testing.T) {

	ty := &Type{
		Object_base: &Object_base{Id: NewReference("a.b/c", "d"), Type: NewReference("kego.io/system", "type")},
	}
	Register("a.b/c", "d", ty, 0)
	defer Unregister("a.b/c", "d")

	r := NewReference("a.b/c", "d")
	typ, ok := r.GetType()
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:d", typ.Id.Value())

	r = NewReference("a.b/c", "e")
	_, ok = r.GetType()
	assert.False(t, ok)

	r = Reference{}
	_, ok = r.GetType()
	assert.False(t, ok)
}
