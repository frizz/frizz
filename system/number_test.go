package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/envctx"
)

/*
func TestUnpackDefaultNativeTypeNumber(t *testing.T) {
	testUnpackDefaultNativeTypeNumber(t, unpacker.Unpack)
}
func testUnpackDefaultNativeTypeNumber(t *testing.T, up unpacker.Interface) {

	data := `{
		"type": "a",
		"b": 1.2
	}`

	type A struct {
		*Object
		B NumberInterface `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := up.Process(ctx, []byte(data), &i)
	require.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, 1.2, a.B.GetNumber(nil).Value())

	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":1.2}`, string(b))

}
*/

func TestNumberRule_Enforce(t *testing.T) {
	r := NumberRule{Rule: &Rule{Optional: false}, Minimum: NewNumber(1.5)}
	fail, messages, err := r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(2))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1))
	require.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 1.5", messages[0])
	assert.True(t, fail)

	r = NumberRule{Rule: &Rule{Optional: false}, Minimum: NewNumber(1.5), ExclusiveMinimum: true}
	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	require.NoError(t, err)
	assert.Equal(t, "Minimum (exclusive): value 1.5 must be greater than 1.5", messages[0])
	assert.True(t, fail)

	r = NumberRule{Rule: &Rule{Optional: false}, Maximum: NewNumber(1.5)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(2))
	require.NoError(t, err)
	assert.Equal(t, "Maximum: value 2 must not be greater than 1.5", messages[0])
	assert.True(t, fail)

	r = NumberRule{Rule: &Rule{Optional: false}, Maximum: NewNumber(1.5), ExclusiveMaximum: true}
	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	require.NoError(t, err)
	assert.Equal(t, "Maximum (exclusive): value 1.5 must be less than 1.5", messages[0])
	assert.True(t, fail)

	r = NumberRule{Rule: &Rule{Optional: false}, MultipleOf: NewNumber(1.5)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(0))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(1.5))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(3))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewNumber(4))
	require.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 1.5", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, "foo")
	assert.IsError(t, err, "FUGYGJVHYS")

}

func TestNewNumber(t *testing.T) {
	n := NewNumber(1.2)
	assert.NotNil(t, n)
	assert.Equal(t, 1.2, n.Value())
}

func TestNumberUnmarshalJSON(t *testing.T) {

	var n *Number

	err := n.Unpack(envctx.Empty, Pack(nil), false)
	require.NoError(t, err)
	assert.Nil(t, n)

	n = NewNumber(0.0)
	err = n.Unpack(envctx.Empty, Pack(1.2), false)
	require.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, 1.2, n.Value())

	n = NewNumber(0.0)
	err = n.Unpack(envctx.Empty, Pack(map[string]interface{}{
		"type":  "system:number",
		"value": 1.2,
	}), false)
	require.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, 1.2, n.Value())

	n = NewNumber(0.0)
	err = n.Unpack(envctx.Empty, Pack("foo"), false)
	assert.Error(t, err)

}

func TestNumberMarshalJSON(t *testing.T) {

	var n *Number

	ba, _, _, err := n.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, nil, ba)

	n = NewNumber(1.2)
	ba, _, _, err = n.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, 1.2, ba)

	n = NewNumber(1.0)
	ba, _, _, err = n.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, 1.0, ba)

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

	n = nil
	s = n.String()
	assert.Equal(t, "", s)

}

func TestNumberInterfaces(t *testing.T) {
	var nn NativeNumber = NewNumber(99.9)
	assert.Equal(t, 99.9, nn.NativeNumber())

	var dr DefaultRule = &NumberRule{Default: NewNumber(22.2)}
	assert.Equal(t, 22.2, dr.GetDefault().(*Number).Value())
}
