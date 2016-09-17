package system

import (
	"reflect"
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/tests"
	"kego.io/tests/unpacker"
)

func TestUnpackDefaultNativeTypeInt(t *testing.T) {
	testUnpackDefaultNativeTypeInt(t, unpacker.Unmarshal)
	testUnpackDefaultNativeTypeInt(t, unpacker.Unpack)
	testUnpackDefaultNativeTypeInt(t, unpacker.Decode)
}
func testUnpackDefaultNativeTypeInt(t *testing.T, up unpacker.Interface) {

	data := `{
		"type": "a",
		"b": 2
	}`

	type A struct {
		*Object
		B IntInterface `json:"b"`
	}

	var i interface{}

	ctx := tests.Context("kego.io/system").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	err := up.Process(ctx, []byte(data), &i)
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, 2, a.B.GetInt(nil).Value())

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":2}`, string(b))

}

func TestIntRule_Enforce(t *testing.T) {
	r := IntRule{Rule: &Rule{Optional: false}, Minimum: NewInt(2)}
	fail, messages, err := r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(2))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(1))
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 2", messages[0])
	assert.True(t, fail)

	r = IntRule{Rule: &Rule{Optional: false}, Maximum: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(1))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(2))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value 3 must not be greater than 2", messages[0])
	assert.True(t, fail)

	r = IntRule{Rule: &Rule{Optional: false}, MultipleOf: NewInt(3)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(0))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(3))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(6))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewInt(4))
	assert.NoError(t, err)
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

	err := i.Unpack(envctx.Empty, json.Pack(nil))
	assert.IsError(t, err, "JEJANRWFMH")

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, json.Pack(2.0))
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, 2, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, json.Pack(map[string]interface{}{
		"type":  "system:int",
		"value": 2.0,
	}))
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, 2, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, json.Pack(-12.0))
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, -12, i.Value())

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, json.Pack("foo"))
	assert.IsError(t, err, "UJUBDGVYGF")

	i = NewInt(0)
	err = i.Unpack(envctx.Empty, json.Pack(1.2))
	assert.HasError(t, err, "KVEOETSIJY")

}

func TestIntMarshalJSON(t *testing.T) {

	var i *Int
	ba, err := i.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	i = NewInt(12)
	ba, err = i.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "12", string(ba))

	i = NewInt(-101)
	ba, err = i.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "-101", string(ba))

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
