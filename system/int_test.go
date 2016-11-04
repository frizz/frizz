package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/envctx"
)

func TestIntRule_Enforce(t *testing.T) {
	r := IntRule{Rule: &Rule{Optional: false}, Minimum: NewInt(2)}
	fail, messages, err := r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(2))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(1))
	require.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 2", messages[0])
	assert.True(t, fail)

	r = IntRule{Rule: &Rule{Optional: false}, Maximum: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(1))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(2))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	require.NoError(t, err)
	assert.Equal(t, "Maximum: value 3 must not be greater than 2", messages[0])
	assert.True(t, fail)

	r = IntRule{Rule: &Rule{Optional: false}, MultipleOf: NewInt(3)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(0))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(6))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(4))
	require.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 3", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, "foo")
	assert.IsError(t, err, "AISBHNCJXJ")

}

func TestNewInt(t *testing.T) {
	n := NewInt(2)
	assert.NotNil(t, n)
	assert.Equal(t, 2, n.Value())
}

func TestIntUnmarshalJSON(t *testing.T) {

	var i *Int

	err := i.Unpack(envctx.Empty, Pack(nil), false)
	require.NoError(t, err)
	assert.Nil(t, i)

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, Pack(2.0), false)
	require.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, 2, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, Pack(map[string]interface{}{
		"type":  "system:int",
		"value": 2.0,
	}), false)
	require.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, 2, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, Pack(-12.0), false)
	require.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, -12, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, Pack("foo"), false)
	assert.IsError(t, err, "UJUBDGVYGF")

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, Pack(1.2), false)
	assert.HasError(t, err, "KVEOETSIJY")

}

func TestIntRepack(t *testing.T) {

	var i *Int
	ba, _, _, _, err := i.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, nil, ba)

	i = NewInt(12)
	ba, _, _, _, err = i.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, 12.0, ba)

	i = NewInt(-101)
	ba, _, _, _, err = i.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, -101.0, ba)

}

func TestIntString(t *testing.T) {

	var i *Int
	s := i.String()
	s1 := i.GetString(nil)
	assert.Equal(t, "", s)
	assert.Equal(t, "", s1.String())

	i = NewInt(12)
	s = i.String()
	s1 = i.GetString(nil)
	assert.Equal(t, "12", s)
	assert.Equal(t, "12", s1.String())

	i = NewInt(-101)
	s = i.String()
	s1 = i.GetString(nil)
	assert.Equal(t, "-101", s)
	assert.Equal(t, "-101", s1.String())

	var n NativeNumber = NewInt(99)
	assert.Equal(t, 99.0, n.NativeNumber())

	var dr DefaultRule = &IntRule{Default: NewInt(22)}
	assert.Equal(t, 22, dr.GetDefault().(*Int).Value())

}
