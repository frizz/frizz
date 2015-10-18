package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestUnpackDefaultNativeTypeString(t *testing.T) {
	testUnpackDefaultNativeTypeString(t, unmarshalFunc)
	testUnpackDefaultNativeTypeString(t, unpackFunc)

	// needs types
	//testUnpackDefaultNativeTypeString(t, repackFunc)
}
func testUnpackDefaultNativeTypeString(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": "c"
	}`

	type A struct {
		*Object
		B StringInterface `json:"b"`
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.GetString().Value)

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

}

func TestMarshal(t *testing.T) {
	testMarshal(t, unmarshalFunc)
	testMarshal(t, unpackFunc)

	// needs types
	//testMarshal(t, repackFunc)
}
func testMarshal(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": "c"
	}`

	type A struct {
		*Object
		B String `json:"b"`
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.Value)

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

}

func TestMarshal1(t *testing.T) {

	type A struct {
		*Object
		B String `json:"b"`
		C String `json:"c"`
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")

	a := A{
		Object: &Object{
			Type: NewReference("kego.io/system", "a"),
		},
		B: NewString("d"),
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"d"}`, string(b))

}

func TestMarshal2(t *testing.T) {

	type A struct {
		*Object
		B String `json:"b"`
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")

	a := A{
		Object: &Object{
			Type: NewReference("kego.io/system", "a"),
		},
		B: NewString("c"),
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

	b, err = json.MarshalContext(a, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	b, err = json.MarshalContext(a, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"system:a","b":"c"}`, string(b))

}

func TestMarshal3(t *testing.T) {

	type A struct {
		*Object
		B String `json:"b"`
	}

	json.Register("c.d/e", "a", reflect.TypeOf(&A{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("c.d/e", "a")

	a := A{
		Object: &Object{
			Type: NewReference("c.d/e", "a"),
		},
		B: NewString("c"),
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"c.d/e:a","b":"c"}`, string(b))

	b, err = json.MarshalContext(a, "c.d/e", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	b, err = json.MarshalContext(a, "f.g/h", map[string]string{})
	// The json package doesn't use kerr throughout, so we can't use HasError
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "WGCDQQCFAD")

	b, err = json.MarshalContext(a, "f.g/h", map[string]string{"c.d/e": "i"})
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"i:a","b":"c"}`, string(b))

}

func TestNewString(t *testing.T) {
	s := NewString("a")
	assert.True(t, s.Exists)
	assert.Equal(t, "a", s.Value)
}

func TestStringUnmarshalJSON(t *testing.T) {

	var s String

	err := s.Unpack(json.NewJsonUnpacker(`foo "bar"`))
	assert.NoError(t, err)
	assert.True(t, s.Exists)
	assert.Equal(t, `foo "bar"`, s.Value)

	err = s.Unpack(json.NewJsonUnpacker(nil))
	assert.NoError(t, err)
	assert.False(t, s.Exists)
	assert.Equal(t, "", s.Value)

	err = s.Unpack(json.NewJsonUnpacker(1.0))
	assert.IsError(t, err, "IXASCXOPMG")

}

func TestStringMarshalJSON(t *testing.T) {

	s := String{Exists: false, Value: ""}
	ba, err := s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	s = String{Exists: true, Value: `foo "bar"`}
	ba, err = s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `"foo \"bar\""`, string(ba))

}

func TestStringString(t *testing.T) {

	s := String{Exists: false, Value: ""}
	str := s.String()
	assert.Equal(t, "", str)

	s = String{Exists: true, Value: `foo "bar"`}
	str = s.String()
	assert.Equal(t, `foo "bar"`, str)

}

func TestStringRule_Validate(t *testing.T) {
	r := &StringRule{}
	ok, message, err := r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{Pattern: NewString("[")}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "Pattern: regex does not compile: [", message)

	r = &StringRule{MaxLength: NewInt(10)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(10), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(5), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(4), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "MaxLength 4 must not be less than MinLength 5", message)
}

func TestStringRule_Enforce(t *testing.T) {
	r := StringRule{Rule: &Rule{Optional: false}, Equal: NewString("a"), MaxLength: NewInt(1)}
	ok, message, err := r.Enforce(NewString("a"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Equal: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewString("b"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Equal: value must equal 'a'", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce("a", "", map[string]string{})
	assert.IsError(t, err, "SXFBXGQSEA")

	r = StringRule{Rule: &Rule{Optional: false}, MaxLength: NewInt(1)}
	ok, message, err = r.Enforce(NewString("ab"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxLength: length must not be greater than 1", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxLength: value must exist", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Enum: []string{"a", "b"}}
	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Enum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewString("a"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewString("c"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Enum: value must be one of: [a b]", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`[`)}
	ok, message, err = r.Enforce(NewString(""), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: regex does not compile: [", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`^foo\d`)}
	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewString("a"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: value must match ^foo\\d", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewString("foo1"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

}
