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
	assert.Equal(t, "d", rp.Type)
	assert.Equal(t, "a.b/c:d", rp.Value)

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
	err := r.UnmarshalJSON([]byte("foo"), "", map[string]string{})
	assert.IsError(t, err, "BBWVFPNNTT")

	r = reset()
	err = r.UnmarshalJSON([]byte("null"), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, r.Exists)
	assert.Equal(t, "", r.Package)
	assert.Equal(t, "", r.Type)
	assert.Equal(t, "", r.Value)

	r = reset()
	err = r.UnmarshalJSON([]byte("\"a.b/c:d\""), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Type)
	assert.Equal(t, "a.b/c:d", r.Value)

	r = reset()
	err = r.UnmarshalJSON([]byte("\"a.b/c:@d\""), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "@d", r.Type)
	assert.Equal(t, "a.b/c:@d", r.Value)

	r = reset()
	err = r.UnmarshalJSON([]byte("\"a:b\""), "", map[string]string{})
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

func TestReferenceGoReference(t *testing.T) {
	r := NewReference("a.b/c", "d")
	s, err := r.GoReference("a.b/c", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "D", s)

	r = NewReference("a.b/c", "d")
	_, err = r.GoReference("", map[string]string{})
	assert.IsError(t, err, "LVHAQUOQGR")

	r = NewReference("kego.io/json", "string")
	s, err = r.GoReference("", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "string", s)

	r = NewReference("kego.io/json", "bool")
	s, err = r.GoReference("", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "bool", s)

	r = NewReference("kego.io/json", "number")
	s, err = r.GoReference("", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "float64", s)

	r = NewReference("kego.io/json", "a")
	s, err = r.GoReference("", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "json.A", s)
}

func TestReferenceGetType(t *testing.T) {

	ty := &Type{
		Base: &Base{Context: &Context{Package: "a.b/c"}, Id: "d", Type: NewReference("kego.io/system", "type")},
	}
	RegisterType("a.b/c:d", ty)
	defer UnregisterType("a.b/c:d")

	r := NewReference("a.b/c", "d")
	typ, ok := r.GetType()
	assert.True(t, ok)
	assert.Equal(t, "d", typ.Id)

	r = NewReference("a.b/c", "e")
	_, ok = r.GetType()
	assert.False(t, ok)

	r = Reference{}
	_, ok = r.GetType()
	assert.False(t, ok)
}
