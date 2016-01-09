package system

import (
	"reflect"
	"testing"

	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestMapMarshal(t *testing.T) {
	testMapMarshal(t, unmarshalFunc)
	testMapMarshal(t, unpackFunc)
}
func testMapMarshal(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": {"c": "d"},
		"e": "f"
	}`

	type A struct {
		*Object
		E *String            `json:"e"`
		B map[string]*String `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := unpacker(ctx, []byte(data), &i)
	assert.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "d", a.B["c"].Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","e":"f","b":{"c":"d"}}`, string(b))

}

func TestMapRule_Enforce(t *testing.T) {

	r := MapRule{MaxItems: NewInt(2)}
	ok, message, err := r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2, "baz": 3})
	assert.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", message)
	assert.False(t, ok)

	r = MapRule{MinItems: NewInt(2)}
	ok, message, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1})
	assert.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", message)
	assert.False(t, ok)

}
