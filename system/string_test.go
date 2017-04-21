package system

import (
	"testing"

	"frizz.io/context/envctx"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestStringGetDefault(t *testing.T) {
	r := StringRule{Default: NewString("a")}
	i := r.GetDefault()
	s, ok := i.(*String)
	assert.True(t, ok)
	assert.Equal(t, "a", s.Value())
}

/*
func TestUnpackDefaultNativeTypeString(t *testing.T) {

	data := `{
		"type": "a",
		"b": "c"
	}`

	type A struct {
		*Object
		B StringInterface `json:"b"`
	}

	ctx := tests.Context("frizz.io/system").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.GetString(nil).Value())

	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"frizz.io/system:a","b":"c"}`, string(b))

}
*/

/*
func TestMarshal(t *testing.T) {

	data := `{
		"type": "a",
		"b": "c"
	}`

	type A struct {
		*Object
		B String `json:"b"`
	}

	ctx := tests.Context("frizz.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, "c", a.B.Value())

	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"frizz.io/system:a","b":"c"}`, string(b))

}
*/

/*
func TestMarshal1(t *testing.T) {

	type A struct {
		*Object
		B *String `json:"b"`
		C *String `json:"c"`
	}

	ctx := tests.Context("frizz.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	a := A{
		Object: &Object{
			Type: NewReference("frizz.io/system", "a"),
		},
		B: NewString("d"),
	}
	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"frizz.io/system:a","b":"d"}`, string(b))

}
*/

/*
func TestMarshal2(t *testing.T) {

	type A struct {
		*Object
		B *String `json:"b"`
	}

	ctx := tests.Context("frizz.io/system").Jtype("a", reflect.TypeOf(&A{})).Ctx()

	a := A{
		Object: &Object{
			Type: NewReference("frizz.io/system", "a"),
		},
		B: NewString("c"),
	}
	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"frizz.io/system:a","b":"c"}`, string(b))

	b, err = Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	ctx1 := tests.Context("d.e/f").JtypePath("frizz.io/system", "a", reflect.TypeOf(&A{})).Ctx()

	b, err = Marshal(ctx1, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"system:a","b":"c"}`, string(b))

}
*/

/*
func TestMarshal3(t *testing.T) {

	type A struct {
		*Object
		B *String `json:"b"`
	}
	ctx := tests.Context("c.d/e").Jtype("a", reflect.TypeOf(&A{})).Ctx()
	a := A{
		Object: &Object{
			Type: NewReference("c.d/e", "a"),
		},
		B: NewString("c"),
	}
	b, err := Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"c.d/e:a","b":"c"}`, string(b))

	b, err = Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"a","b":"c"}`, string(b))

	b, err = Marshal(tests.Context("f.g/h").JtypePath("c.d/e", "a", reflect.TypeOf(&A{})).Ctx(), a)
	// The json package doesn't use kerr throughout, so we can't use HasError
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "WGCDQQCFAD")

	b, err = Marshal(tests.Context("f.g/h").Alias("i", "c.d/e").JtypePath("c.d/e", "a", reflect.TypeOf(&A{})).Ctx(), a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"i:a","b":"c"}`, string(b))

}
*/

func TestNewString(t *testing.T) {
	s := NewString("a")
	assert.NotNil(t, s)
	assert.Equal(t, "a", s.Value())
	assert.Equal(t, "a", s.NativeString())
}

func TestStringUnmarshalJSON(t *testing.T) {

	var s *String

	err := s.Unpack(envctx.Empty, Pack(nil), false)
	require.NoError(t, err)
	assert.Nil(t, s)

	s = NewString("")
	err = s.Unpack(envctx.Empty, Pack(`foo "bar"`), false)
	require.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, `foo "bar"`, s.Value())

	s = NewString("")
	err = s.Unpack(envctx.Empty, Pack(map[string]interface{}{
		"type":  "system:string",
		"value": `foo "bar"`,
	}), false)
	require.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, `foo "bar"`, s.Value())

	s = NewString("")
	err = s.Unpack(envctx.Empty, Pack(1.0), false)
	assert.Error(t, err)

}

func TestStringMarshalJSON(t *testing.T) {

	var s *String

	ba, _, _, _, err := s.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, nil, ba)

	s = NewString(`foo "bar"`)
	ba, _, _, _, err = s.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, `foo "bar"`, ba)

}

func TestStringString(t *testing.T) {

	var s *String
	str := s.String()
	assert.Equal(t, "", str)

	s = NewString(`foo "bar"`)
	str = s.String()
	assert.Equal(t, `foo "bar"`, str)

}

func TestStringRule_Validate(t *testing.T) {
	r := &StringRule{}
	fail, messages, err := r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	r = &StringRule{PatternNot: NewString("[")}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.True(t, fail)
	assert.Equal(t, "PatternNot: regex does not compile: [", messages[0])

	r = &StringRule{Pattern: NewString("[")}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.True(t, fail)
	assert.Equal(t, "Pattern: regex does not compile: [", messages[0])

	r = &StringRule{MaxLength: NewInt(10)}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	r = &StringRule{MaxLength: NewInt(10), MinLength: NewInt(5)}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	r = &StringRule{MaxLength: NewInt(5), MinLength: NewInt(5)}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	r = &StringRule{MaxLength: NewInt(4), MinLength: NewInt(5)}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.True(t, fail)
	assert.Equal(t, "MaxLength 4 must not be less than MinLength 5", messages[0])
}

func TestStringRule_Enforce(t *testing.T) {
	r := StringRule{Rule: &Rule{Optional: false}, Equal: NewString("a"), MaxLength: NewInt(1)}
	fail, messages, err := r.Enforce(envctx.Empty, NewString("a"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Equal: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("b"))
	require.NoError(t, err)
	assert.Equal(t, "Equal: value \"b\" must equal 'a'", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("123456789012345678901234567890"))
	require.NoError(t, err)
	assert.Equal(t, "Equal: value \"12345678901234567...\" must equal 'a'", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, "a")
	assert.IsError(t, err, "SXFBXGQSEA")

	r = StringRule{Rule: &Rule{Optional: false}, MaxLength: NewInt(1)}
	fail, messages, err = r.Enforce(envctx.Empty, NewString("ab"))
	require.NoError(t, err)
	assert.Equal(t, "MaxLength: length of \"ab\" must not be greater than 1", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "MaxLength: value must exist", messages[0])
	assert.True(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, MinLength: NewInt(5)}
	fail, messages, err = r.Enforce(envctx.Empty, NewString("abcde"))
	require.NoError(t, err)
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("abcd"))
	require.NoError(t, err)
	assert.Equal(t, "MinLength: length of \"abcd\" must not be less than 5", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "MinLength: value must exist", messages[0])
	assert.True(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, Enum: []string{"a", "b"}}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Enum: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("a"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("c"))
	require.NoError(t, err)
	assert.Equal(t, "Enum: value \"c\" must be one of: [a b]", messages[0])
	assert.True(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`[`)}
	fail, messages, err = r.Enforce(envctx.Empty, NewString(""))
	require.NoError(t, err)
	assert.Equal(t, "Pattern: regex does not compile: [", messages[0])
	assert.True(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, PatternNot: NewString(`[`)}
	fail, messages, err = r.Enforce(envctx.Empty, NewString(""))
	require.NoError(t, err)
	assert.Equal(t, "PatternNot: regex does not compile: [", messages[0])
	assert.True(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, Pattern: NewString(`^foo\d`)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Pattern: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("a"))
	require.NoError(t, err)
	assert.Equal(t, "Pattern: value \"a\" must match ^foo\\d", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("foo1"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	r = StringRule{Rule: &Rule{Optional: false}, PatternNot: NewString(`^foo\d`)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "PatternNot: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("foo1"))
	require.NoError(t, err)
	require.Equal(t, 1, len(messages))
	assert.Equal(t, "PatternNot: value \"foo1\" must not match ^foo\\d", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewString("a"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

}
