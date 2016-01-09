package system

import (
	"reflect"
	"testing"

	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestUnpackDefaultNativeTypeNumber(t *testing.T) {
	testUnpackDefaultNativeTypeNumber(t, unmarshalFunc)
	testUnpackDefaultNativeTypeNumber(t, unpackFunc)
}
func testUnpackDefaultNativeTypeNumber(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": 1.2
	}`

	type A struct {
		*Object
		B NumberInterface `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := unpacker(ctx, []byte(data), &i)
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, 1.2, a.B.GetNumber(nil).Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":1.2}`, string(b))

}

func TestNumberRule_Enforce(t *testing.T) {
	r := NumberRule{Rule: &Rule{Optional: false}, Minimum: NewNumber(1.5)}
	ok, message, err := r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(2))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1))
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 1.5", message)
	assert.False(t, ok)

	r = NumberRule{Rule: &Rule{Optional: false}, Minimum: NewNumber(1.5), ExclusiveMinimum: true}
	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	assert.NoError(t, err)
	assert.Equal(t, "Minimum (exclusive): value 1.5 must be greater than 1.5", message)
	assert.False(t, ok)

	r = NumberRule{Rule: &Rule{Optional: false}, Maximum: NewNumber(1.5)}
	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(2))
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value 2 must not be greater than 1.5", message)
	assert.False(t, ok)

	r = NumberRule{Rule: &Rule{Optional: false}, Maximum: NewNumber(1.5), ExclusiveMaximum: true}
	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	assert.NoError(t, err)
	assert.Equal(t, "Maximum (exclusive): value 1.5 must be less than 1.5", message)
	assert.False(t, ok)

	r = NumberRule{Rule: &Rule{Optional: false}, MultipleOf: NewNumber(1.5)}
	ok, message, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(0))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(3))
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, NewNumber(4))
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 1.5", message)
	assert.False(t, ok)

}

func TestNewNumber(t *testing.T) {
	n := NewNumber(1.2)
	assert.NotNil(t, n)
	assert.Equal(t, 1.2, n.Value())
}

func TestNumberUnmarshalJSON(t *testing.T) {

	var n *Number

	err := n.Unpack(envctx.Empty, json.Pack(nil))
	assert.IsError(t, err, "WHREWCCODC")

	n = NewNumber(0.0)

	err = n.Unpack(envctx.Empty, json.Pack(1.2))
	assert.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, 1.2, n.Value())

	err = n.Unpack(envctx.Empty, json.Pack("foo"))
	assert.IsError(t, err, "YHXBFTONCW")

}

func TestNumberMarshalJSON(t *testing.T) {

	var n *Number

	ba, err := n.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	n = NewNumber(1.2)
	ba, err = n.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "1.2", string(ba))

	n = NewNumber(1.0)
	ba, err = n.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "1", string(ba))

}

func TestNumberString(t *testing.T) {

	n := NewNumber(0.0)
	s := n.String()
	assert.Equal(t, "0", s)

	n = NewNumber(1.2)
	s = n.String()
	assert.Equal(t, "1.2", s)

	n = NewNumber(1.0)
	s = n.String()
	assert.Equal(t, "1", s)

}
