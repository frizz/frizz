package system

import (
	"testing"

	"frizz.io/context/envctx"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

/*
func TestMapMarshal(t *testing.T) {
	testMapMarshal(t, unpacker.Unpack)
}
func testMapMarshal(t *testing.T, up unpacker.Interface) {

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

	ctx := tests.Context("frizz.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "d", a.B["c"].Value())

	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"frizz.io/system:a","e":"f","b":{"c":"d"}}`, string(b))

}
*/

func TestMapRule_Enforce(t *testing.T) {

	r := MapRule{Keys: &IntRule{Rule: &Rule{}}}
	fail, messages, err := r.Enforce(envctx.Empty, map[string]int{"foo": 1})
	require.IsError(t, err, "WDKAXPCRJB")

	r = MapRule{Keys: &StringRule{Rule: &Rule{}, MaxLength: NewInt(1)}}
	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1})
	require.NoError(t, err)
	assert.Equal(t, 1, len(messages))
	assert.True(t, fail)
	assert.Equal(t, "MaxLength: length of \"foo\" must not be greater than 1", messages[0])

	r = MapRule{MaxItems: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2, "baz": 3})
	require.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", messages[0])
	assert.True(t, fail)

	r = MapRule{MinItems: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1})
	require.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", messages[0])
	assert.True(t, fail)

	_, _, err = r.Enforce(envctx.Empty, "a")
	assert.IsError(t, err, "NFWPLTOJLP")

	r = MapRule{}
	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1})
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	var cr CollectionRule = &MapRule{Items: &StringRule{Equal: NewString("a")}}
	assert.Equal(t, "a", cr.GetItemsRule().(*StringRule).Equal.Value())

}
