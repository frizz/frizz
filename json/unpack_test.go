package json

import (
	"reflect"
	"testing"

	"kego.io/kerr/assert"
)

func TestSetType(t *testing.T) {

	a := true
	type b struct {
		C string
	}

	var i interface{}

	testSetType(t, reflect.ValueOf(&i), reflect.TypeOf(&a), "*bool")
	testSetType(t, reflect.ValueOf(&i), reflect.TypeOf(true), "bool")
	testSetType(t, reflect.ValueOf(&i), reflect.TypeOf(&b{}), "*json.b")
	testSetType(t, reflect.ValueOf(&i), reflect.TypeOf(b{}), "json.b")

}

func testSetType(t *testing.T, val reflect.Value, typ reflect.Type, expect string) {
	err := setType(val, typ)
	assert.NoError(t, err)
	assert.Equal(t, expect, val.Elem().Elem().Type().String())
}
