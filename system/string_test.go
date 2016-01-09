package system

import (
	"reflect"
	"testing"

	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestUnpackDefaultNativeTypeString(t *testing.T) {
	testUnpackDefaultNativeTypeString(t, unmarshalFunc)
	testUnpackDefaultNativeTypeString(t, unpackFunc)
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

	ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := unpacker(ctx, []byte(data), &i)
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.GetString(nil).Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

}

func TestMarshal(t *testing.T) {
	testMarshal(t, unmarshalFunc)
	testMarshal(t, unpackFunc)
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

	ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := unpacker(ctx, []byte(data), &i)
	assert.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

}

func TestMarshal1(t *testing.T) {

	type A struct {
		*Object
		B *String `json:"b"`
		C *String `json:"c"`
	}

	//ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

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
		B *String `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	a := A{
		Object: &Object{
			Type: NewReference("kego.io/system", "a"),
		},
		B: NewString("c"),
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c"}`, string(b))

	b, err = json.MarshalContext(ctx, a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	ctx1 := tests.Context("d.e/f").JtypePath("kego.io/system", "a", reflect.TypeOf(&A{})).Ctx()

	b, err = json.MarshalContext(ctx1, a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"system:a","b":"c"}`, string(b))

}

func TestMarshal3(t *testing.T) {

	type A struct {
		*Object
		B *String `json:"b"`
	}

	a := A{
		Object: &Object{
			Type: NewReference("c.d/e", "a"),
		},
		B: NewString("c"),
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"c.d/e:a","b":"c"}`, string(b))

	b, err = json.MarshalContext(tests.Context("c.d/e").Jtype("a", reflect.TypeOf(&A{})).Ctx(), a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	b, err = json.MarshalContext(tests.Context("f.g/h").JtypePath("c.d/e", "a", reflect.TypeOf(&A{})).Ctx(), a)
	// The json package doesn't use kerr throughout, so we can't use HasError
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "WGCDQQCFAD")

	b, err = json.MarshalContext(tests.Context("f.g/h").Alias("c.d/e", "i").JtypePath("c.d/e", "a", reflect.TypeOf(&A{})).Ctx(), a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"i:a","b":"c"}`, string(b))

}

func TestNewString(t *testing.T) {
	s := NewString("a")
	assert.NotNil(t, s)
	assert.Equal(t, "a", s.Value())
}

func TestStringUnmarshalJSON(t *testing.T) {

	var s *String

	err := s.Unpack(envctx.Empty, json.Pack(nil))
	assert.IsError(t, err, "PWTAHLCCWR")

	s = NewString("")

	err = s.Unpack(envctx.Empty, json.Pack(`foo "bar"`))
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, `foo "bar"`, s.Value())

	err = s.Unpack(envctx.Empty, json.Pack(1.0))
	assert.IsError(t, err, "IXASCXOPMG")

}

func TestStringMarshalJSON(t *testing.T) {

	var s *String

	ba, err := s.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	s = NewString(`foo "bar"`)
	ba, err = s.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, `"foo \"bar\""`, string(ba))

}

func TestStringString(t *testing.T) {

	var s *String
	str := s.String()
	assert.Equal(t, "", str)

	s = NewString(`foo "bar"`)
	str = s.String()
	assert.Equal(t, `foo "bar"`, str)

}

func TestStringRule_Validate(t *testing.T) {
	r := &StringRule{}
	ok, message, err := r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{Pattern: NewString("[")}
	ok, message, err = r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "Pattern: regex does not compile: [", message)

	r = &StringRule{MaxLength: NewInt(10)}
	ok, message, err = r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(10), MinLength: NewInt(5)}
	ok, message, err = r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(5), MinLength: NewInt(5)}
	ok, message, err = r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &StringRule{MaxLength: NewInt(4), MinLength: NewInt(5)}
	ok, message, err = r.Validate(envctx.Empty)
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "MaxLength 4 must not be less than MinLength 5", message)
}

func TestStringRule_Enforce(t *testing.T) {
	r := StringRule{Rule: &Rule{Optional: false}, Equal: NewString("a"), MaxLength: NewInt(1)}
	ok, message, err := r.Enforce(envctx.Empty, NewString("a"))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Equal: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewString("b"))
	assert.NoError(t, err)
	assert.Equal(t, "Equal: value must equal 'a'", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, "a")
	assert.IsError(t, err, "SXFBXGQSEA")

	r = StringRule{Rule: &Rule{Optional: false}, MaxLength: NewInt(1)}
	ok, message, err = r.Enforce(envctx.Empty, NewString("ab"))
	assert.NoError(t, err)
	assert.Equal(t, "MaxLength: length must not be greater than 1", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "MaxLength: value must exist", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Enum: []string{"a", "b"}}
	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Enum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewString("a"))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewString("c"))
	assert.NoError(t, err)
	assert.Equal(t, "Enum: value must be one of: [a b]", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`[`)}
	ok, message, err = r.Enforce(envctx.Empty, NewString(""))
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: regex does not compile: [", message)
	assert.False(t, ok)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`^foo\d`)}
	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewString("a"))
	assert.NoError(t, err)
	assert.Equal(t, "Pattern: value must match ^foo\\d", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewString("foo1"))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

}
