package system

import (
	"testing"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"kego.io/context/envctx"
)

func TestNewBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.Value())

	b1 := NewBool(false)
	assert.False(t, b1.Value())
}

func TestBoolUnmarshalJSON(t *testing.T) {

	var b *Bool

	err := b.Unpack(envctx.Empty, Pack(nil), false)
	require.NoError(t, err)
	assert.Nil(t, b)

	b = NewBool(false)
	err = b.Unpack(envctx.Empty, Pack(true), false)
	require.NoError(t, err)
	assert.NotNil(t, b)
	assert.True(t, b.Value())

	b = NewBool(false)
	err = b.Unpack(envctx.Empty, Pack(map[string]interface{}{
		"type":  "system:bool",
		"value": true,
	}), false)
	require.NoError(t, err)
	assert.NotNil(t, b)
	assert.True(t, b.Value())

	b = NewBool(true)
	err = b.Unpack(envctx.Empty, Pack(false), false)
	require.NoError(t, err)
	assert.NotNil(t, b)
	assert.False(t, b.Value())

	b = NewBool(false)
	err = b.Unpack(envctx.Empty, Pack("foo"), false)
	assert.Error(t, err)

}
func TestBoolRepack(t *testing.T) {

	var b *Bool
	ba, _, _, _, err := b.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, nil, ba)

	b = NewBool(false)
	ba, _, _, _, err = b.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, false, ba)

	b = NewBool(true)
	ba, _, _, _, err = b.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, true, ba)

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
