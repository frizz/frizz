package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestUnpackDefaultNativeTypeInt(t *testing.T) {
	testUnpackDefaultNativeTypeInt(t, unmarshalFunc)
	testUnpackDefaultNativeTypeInt(t, unpackFunc)
}
func testUnpackDefaultNativeTypeInt(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": 2
	}`

	type A struct {
		*Object
		B IntInterface `json:"b"`
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
	assert.Equal(t, 2, a.B.GetInt().Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":2}`, string(b))

}

func TestIntRule_Enforce(t *testing.T) {
	r := IntRule{Rule: &Rule{Optional: false}, Minimum: NewInt(2)}
	ok, message, err := r.Enforce(nil, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 2", message)
	assert.False(t, ok)

	r = IntRule{Rule: &Rule{Optional: false}, Maximum: NewInt(2)}
	ok, message, err = r.Enforce(nil, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value 3 must not be greater than 2", message)
	assert.False(t, ok)

	r = IntRule{Rule: &Rule{Optional: false}, MultipleOf: NewInt(3)}
	ok, message, err = r.Enforce(nil, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(0), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(6), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(4), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 3", message)
	assert.False(t, ok)

}

func TestNewInt(t *testing.T) {
	n := NewInt(2)
	assert.NotNil(t, n)
	assert.Equal(t, 2, n.Value())
}

func TestIntUnmarshalJSON(t *testing.T) {

	var i *Int

	err := i.Unpack(json.Pack(nil))
	assert.IsError(t, err, "JEJANRWFMH")

	i = NewInt(0)

	err = i.Unpack(json.Pack(2.0))
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, 2, i.Value())

	err = i.Unpack(json.Pack(-12.0))
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, -12, i.Value())

	err = i.Unpack(json.Pack("foo"))
	assert.IsError(t, err, "UJUBDGVYGF")

	err = i.Unpack(json.Pack(1.2))
	assert.HasError(t, err, "KVEOETSIJY")

}

func TestIntMarshalJSON(t *testing.T) {

	var i *Int
	ba, err := i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	i = NewInt(12)
	ba, err = i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "12", string(ba))

	i = NewInt(-101)
	ba, err = i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "-101", string(ba))

}

func TestIntString(t *testing.T) {

	var i *Int
	s := i.String()
	assert.Equal(t, "", s)

	i = NewInt(12)
	s = i.String()
	assert.Equal(t, "12", s)

	i = NewInt(-101)
	s = i.String()
	assert.Equal(t, "-101", s)

}
