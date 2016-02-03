package system

import (
	"reflect"
	"testing"

	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
	"kego.io/process/tests/unpacker"
)

func TestUnpackDefaultNativeTypeBool(t *testing.T) {
	testUnpackDefaultNativeTypeBool(t, unpacker.Unmarshal)
	testUnpackDefaultNativeTypeBool(t, unpacker.Unpack)
}
func testUnpackDefaultNativeTypeBool(t *testing.T, up unpacker.Interface) {

	data := `{
		"type": "a",
		"b": true
	}`

	type A struct {
		*Object
		B BoolInterface `json:"b"`
	}

	var i interface{}

	ctx := tests.Context("kego.io/system").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	err := up.Process(ctx, []byte(data), &i)
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, true, a.B.GetBool(nil).Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":true}`, string(b))

}

func TestNewBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.Value())

	b1 := NewBool(false)
	assert.False(t, b1.Value())
}

func TestBoolUnmarshalJSON(t *testing.T) {

	var b *Bool

	err := b.Unpack(envctx.Empty, json.Pack(nil))
	assert.IsError(t, err, "FXCQGNYKIJ")

	b = NewBool(false)

	err = b.Unpack(envctx.Empty, json.Pack(true))
	assert.NoError(t, err)
	assert.NotNil(t, b)
	assert.True(t, b.Value())

	err = b.Unpack(envctx.Empty, json.Pack(false))
	assert.NoError(t, err)
	assert.NotNil(t, b)
	assert.False(t, b.Value())

	err = b.Unpack(envctx.Empty, json.Pack("foo"))
	assert.IsError(t, err, "GXQGNEPJYS")

}
func TestBoolMarshalJSON(t *testing.T) {

	var b *Bool
	ba, err := b.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	b = NewBool(false)
	ba, err = b.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "false", string(ba))

	b = NewBool(true)
	ba, err = b.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "true", string(ba))

}
func TestBoolString(t *testing.T) {

	var b *Bool
	s := b.String()
	assert.Equal(t, "", s)

	b = NewBool(false)
	s = b.String()
	assert.Equal(t, "false", s)

	b = NewBool(true)
	s = b.String()
	assert.Equal(t, "true", s)

}

func TestBoolInterfaces(t *testing.T) {
	var nb NativeBool = NewBool(true)
	assert.True(t, nb.NativeBool())

	d := NewBool(true)
	var dr DefaultRule = &BoolRule{Default: d}
	assert.Equal(t, d, dr.GetDefault())
}
