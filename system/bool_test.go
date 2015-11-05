package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestUnpackDefaultNativeTypeBool(t *testing.T) {
	testUnpackDefaultNativeTypeBool(t, unmarshalFunc)
	testUnpackDefaultNativeTypeBool(t, unpackFunc)
}
func testUnpackDefaultNativeTypeBool(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": true
	}`

	type A struct {
		*Object
		B BoolInterface `json:"b"`
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
	assert.Equal(t, true, a.B.GetBool().Value())

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

	err := b.Unpack(json.Pack(nil))
	assert.IsError(t, err, "FXCQGNYKIJ")

	b = NewBool(false)

	err = b.Unpack(json.Pack(true))
	assert.NoError(t, err)
	assert.NotNil(t, b)
	assert.True(t, b.Value())

	err = b.Unpack(json.Pack(false))
	assert.NoError(t, err)
	assert.NotNil(t, b)
	assert.False(t, b.Value())

	err = b.Unpack(json.Pack("foo"))
	assert.IsError(t, err, "GXQGNEPJYS")

}
func TestBoolMarshalJSON(t *testing.T) {

	var b *Bool
	ba, err := b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	b = NewBool(false)
	ba, err = b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "false", string(ba))

	b = NewBool(true)
	ba, err = b.MarshalJSON()
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
