package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestNativeGoType(t *testing.T) {
	n, err := nativeGoType("string")
	assert.NoError(t, err)
	assert.Equal(t, "string", n)

	n, err = nativeGoType("number")
	assert.NoError(t, err)
	assert.Equal(t, "float64", n)

	n, err = nativeGoType("bool")
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)

	_, err = nativeGoType("a")
	assert.IsError(t, err, "TXQIDRBJRH")
}

func TestTypeIsNativeType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n := y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("array")}
	n = y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("object")}
	n = y.IsNativeObject()
	assert.True(t, n)

	// Will also return true for any other value
	y = &Type{Native: NewString("a")}
	n = y.IsNativeObject()
	assert.True(t, n)
}

func TestTypeNativeValueGolangType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n, err := y.NativeValueGolangType()
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)
}

func TestTypeGoName(t *testing.T) {
	y := &Type{Object: &Object{Id: NewReference("a.b/c", "aa")}}
	n := y.GoName()
	assert.Equal(t, "Aa", n)
}
