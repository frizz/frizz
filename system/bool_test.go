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

	// needs types
	//testUnpackDefaultNativeTypeBool(t, repackFunc)
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
	assert.Equal(t, true, a.B.GetBool().Value)

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":true}`, string(b))

}

func TestNewBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)
}

func TestBoolUnmarshalJSON(t *testing.T) {

	var b Bool

	err := b.Unpack(json.NewJsonUnpacker(true))
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)

	err = b.Unpack(json.NewJsonUnpacker(false))
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.False(t, b.Value)

	err = b.Unpack(json.NewJsonUnpacker(nil))
	assert.NoError(t, err)
	assert.False(t, b.Exists)
	assert.False(t, b.Value)

	err = b.Unpack(json.NewJsonUnpacker("foo"))
	assert.IsError(t, err, "GXQGNEPJYS")

}
func TestBoolMarshalJSON(t *testing.T) {

	b := Bool{Exists: false, Value: false}
	ba, err := b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	b = Bool{Exists: true, Value: false}
	ba, err = b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "false", string(ba))

	b = Bool{Exists: true, Value: true}
	ba, err = b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "true", string(ba))

}
func TestBoolString(t *testing.T) {

	b := Bool{Exists: false, Value: false}
	s := b.String()
	assert.Equal(t, "", s)

	b = Bool{Exists: true, Value: false}
	s = b.String()
	assert.Equal(t, "false", s)

	b = Bool{Exists: true, Value: true}
	s = b.String()
	assert.Equal(t, "true", s)

}
