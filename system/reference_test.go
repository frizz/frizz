package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestUnpackDefaultNativeTypeReference(t *testing.T) {
	testUnpackDefaultNativeTypeReference(t, unmarshalFunc)
	testUnpackDefaultNativeTypeReference(t, unpackFunc)

	// needs types
	//testUnpackDefaultNativeTypeReference(t, repackFunc)
}
func testUnpackDefaultNativeTypeReference(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": "e:f"
	}`

	type A struct {
		*Object
		B ReferenceInterface `json:"b"`
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{"c.d/e": "e"})
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, NewReference("c.d/e", "f"), a.B.GetReference())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c.d/e:f"}`, string(b))

}

func TestReferenceRuleChangeTo(t *testing.T) {

	r := NewReference("a.b/c", "@d")
	r1 := r.ChangeToType()
	assert.Equal(t, "a.b/c:d", r1.Value())

	r = NewReference("a.b/c", "d")
	r1 = r.ChangeToType()
	assert.Equal(t, "a.b/c:d", r1.Value())

	r = NewReference("a.b/c", "$d")
	r1 = r.ChangeToType()
	assert.Equal(t, "a.b/c:d", r1.Value())

	r = NewReference("a.b/c", "@d")
	r1 = r.ChangeToRule()
	assert.Equal(t, "a.b/c:@d", r1.Value())

	r = NewReference("a.b/c", "d")
	r1 = r.ChangeToRule()
	assert.Equal(t, "a.b/c:@d", r1.Value())

	r = NewReference("a.b/c", "$d")
	r1 = r.ChangeToRule()
	assert.Equal(t, "a.b/c:@d", r1.Value())

	r = NewReference("a.b/c", "@d")
	r1 = r.ChangeToInterface()
	assert.Equal(t, "a.b/c:$d", r1.Value())

	r = NewReference("a.b/c", "d")
	r1 = r.ChangeToInterface()
	assert.Equal(t, "a.b/c:$d", r1.Value())

	r = NewReference("a.b/c", "$d")
	r1 = r.ChangeToInterface()
	assert.Equal(t, "a.b/c:$d", r1.Value())
}

func TestReferenceUnmarshalJson(t *testing.T) {

	reset := func() *Reference {
		// Let's pre-load with some values so we check that when we
		// load a null value, we clear all the fields
		r := NewReference("a.b/c", "d")
		return &r
	}

	r := reset()
	err := r.Unpack(json.NewJsonUnpacker(nil), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, r.Exists)
	assert.Equal(t, "", r.Package)
	assert.Equal(t, "", r.Name)
	assert.Equal(t, "", r.Value())

	r = reset()
	err = r.Unpack(json.NewJsonUnpacker("a.b/c:d"), "", map[string]string{})
	assert.EqualError(t, err, "Unknown package a.b/c")
	assert.False(t, r.Exists)

	r = reset()
	err = r.Unpack(json.NewJsonUnpacker("a.b/c:d"), "", map[string]string{"a.b/c": "c"})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Name)
	assert.Equal(t, "a.b/c:d", r.Value())

	r = reset()
	err = r.Unpack(json.NewJsonUnpacker("a.b/c:@d"), "", map[string]string{"a.b/c": "c"})
	assert.NoError(t, err)
	assert.True(t, r.Exists)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "@d", r.Name)
	assert.Equal(t, "a.b/c:@d", r.Value())

	r = reset()
	err = r.Unpack(json.NewJsonUnpacker("a:b"), "", map[string]string{})
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
		Object: &Object{Id: NewReference("a.b/c", "d"), Type: NewReference("kego.io/system", "type")},
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
