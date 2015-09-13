package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestMapMarshal(t *testing.T) {

	data := `{
		"type": "a",
		"b": {"c": "d"},
		"e": "f"
	}`

	type A struct {
		*Base
		E String            `json:"e"`
		B map[string]String `json:"b"`
	}

	json.RegisterType("kego.io/system", "a", reflect.TypeOf(&A{}), 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system", "a")

	var i interface{}
	err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "d", a.B["c"].Value)

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","e":"f","b":{"c":"d"}}`, string(b))

}

func TestMapRule_Enforce(t *testing.T) {
	r := Map_rule{MaxItems: NewInt(2)}
	ok, message, err := r.Enforce(map[string]int{"foo": 1, "bar": 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(map[string]int{"foo": 1, "bar": 2, "baz": 3}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", message)
	assert.False(t, ok)

	r = Map_rule{MinItems: NewInt(2)}
	ok, message, err = r.Enforce(map[string]int{"foo": 1, "bar": 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(map[string]int{"foo": 1}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", message)
	assert.False(t, ok)

}
