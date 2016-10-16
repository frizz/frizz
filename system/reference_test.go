package system

import (
	"sort"
	"testing"

	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/envctx"
	"kego.io/packer"
	"kego.io/tests"
)

func TestReferenceRule_Validate(t *testing.T) {
	r := &ReferenceRule{}
	fail, messages, err := r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.False(t, fail)
	assert.Equal(t, 0, len(messages))

	r = &ReferenceRule{Pattern: NewString("[")}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.True(t, fail)
	assert.Equal(t, "Pattern: regex does not compile: [", messages[0])

	r = &ReferenceRule{PatternNot: NewString("[")}
	fail, messages, err = r.Validate(envctx.Empty)
	require.NoError(t, err)
	assert.True(t, fail)
	assert.Equal(t, "PatternNot: regex does not compile: [", messages[0])
}

func TestReferenceRule_Enforce(t *testing.T) {

	r := ReferenceRule{Rule: &Rule{Optional: false}, PatternNot: NewString(`[`)}
	fail, messages, err := r.Enforce(envctx.Empty, NewReference("", ""))
	require.NoError(t, err)
	assert.Equal(t, "PatternNot: regex does not compile: [", messages[0])
	assert.True(t, fail)

	r = ReferenceRule{Rule: &Rule{Optional: false}, Pattern: NewString(`[`)}
	fail, messages, err = r.Enforce(envctx.Empty, NewReference("", ""))
	require.NoError(t, err)
	assert.Equal(t, "Pattern: regex does not compile: [", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, "")
	require.IsError(t, err, "BYDVGGETWW")

	r = ReferenceRule{Rule: &Rule{Optional: false}, Pattern: NewString(`^foo\d`)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "Pattern: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewReference("", "a"))
	require.NoError(t, err)
	assert.Equal(t, "Pattern: value \"a\" must match ^foo\\d", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewReference("", "foo1"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	r = ReferenceRule{Rule: &Rule{Optional: false}, PatternNot: NewString(`^foo\d`)}
	fail, messages, err = r.Enforce(envctx.Empty, nil)
	require.NoError(t, err)
	assert.Equal(t, "PatternNot: value must exist", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewReference("", "foo1"))
	require.NoError(t, err)
	assert.Equal(t, "PatternNot: value \"foo1\" must not match ^foo\\d", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, NewReference("", "a"))
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

}

/*
func TestUnpackDefaultNativeTypeReference(t *testing.T) {
	testUnpackDefaultNativeTypeReference(t, unpacker.Unpack)
}
func testUnpackDefaultNativeTypeReference(t *testing.T, up unpacker.Interface) {

	data := `{
		"type": "a",
		"b": "e:f"
	}`

	type A struct {
		*Object
		B ReferenceInterface `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Alias("e", "c.d/e").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := up.Process(ctx, []byte(data), &i)
	require.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, NewReference("c.d/e", "f"), a.B.GetReference(nil))

	b, err := packer.Marshal(ctx, a)
	require.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c.d/e:f"}`, string(b))

}
*/

func TestReferenceRuleChangeTo(t *testing.T) {

	r := NewReference("a.b/c", "@d")
	r1 := r.ChangeToType()
	assert.Equal(t, "a.b/c:d", r1.Value())

	r = NewReference("a.b/c", "d")
	r1 = r.ChangeToType()
	assert.Equal(t, "a.b/c:d", r1.Value())

	r = NewReference("a.b/c", "@d")
	r1 = r.ChangeToRule()
	assert.Equal(t, "a.b/c:@d", r1.Value())

	r = NewReference("a.b/c", "d")
	r1 = r.ChangeToRule()
	assert.Equal(t, "a.b/c:@d", r1.Value())

	r = NewReference("", "d")
	r1 = r.ChangeToRule()
	assert.Equal(t, *NewReference("", ""), r1)

	r = NewReference("a.b/c", "")
	r1 = r.ChangeToRule()
	assert.Equal(t, *NewReference("", ""), r1)

}

func TestReferenceUnmarshal(t *testing.T) {

	reset := func() *Reference {
		// Let's pre-load with some values so we check that when we
		// load a null value, we clear all the fields
		r := NewReference("a.b/c", "d")
		return r
	}

	r := reset()
	err := r.Unpack(envctx.Empty, packer.Pack(nil), false)
	assert.IsError(t, err, "MOQVSKJXRB")

	r = reset()
	err = r.Unpack(envctx.Empty, packer.Pack(1.0), false)
	assert.IsError(t, err, "RFLQSBPMYM")

	r = reset()
	err = r.Unpack(envctx.Empty, packer.Pack("a.b/c:d"), false)
	assert.IsError(t, err, "MSXBLEIGVJ")
	assert.HasError(t, err, "KJSOXDESFD")
	p, ok := kerr.Source(err).(packer.UnknownPackageError)
	assert.True(t, ok)
	assert.Equal(t, "a.b/c", p.UnknownPackage)

	ctx := tests.Context("").Alias("c", "a.b/c").Ctx()

	r = reset()
	err = r.Unpack(ctx, packer.Pack("a.b/c:d"), false)
	require.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Name)
	assert.Equal(t, "a.b/c:d", r.Value())

	r = reset()
	err = r.Unpack(ctx, packer.Pack(map[string]interface{}{
		"type":  "system:reference",
		"value": "a.b/c:d",
	}), false)
	require.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Name)
	assert.Equal(t, "a.b/c:d", r.Value())

	r = reset()
	err = r.Unpack(ctx, packer.Pack("a.b/c:@d"), false)
	require.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "@d", r.Name)
	assert.Equal(t, "a.b/c:@d", r.Value())

	r = reset()
	err = r.Unpack(envctx.Empty, packer.Pack("a:b"), false)
	assert.IsError(t, err, "MSXBLEIGVJ")
	assert.HasError(t, err, "DKKFLKDKYI")
	p, ok = kerr.Source(err).(packer.UnknownPackageError)
	assert.True(t, ok)
	assert.Equal(t, "a", p.UnknownPackage)

	r = reset()
	err = r.UnmarshalInterface(envctx.Empty, 1)
	require.NoError(t, err)
	assert.Equal(t, *NewReference("", ""), *r)

	r = reset()
	err = r.UnmarshalInterface(envctx.Empty, "")
	require.NoError(t, err)
	assert.Equal(t, *NewReference("", ""), *r)

	r = reset()
	err = r.UnmarshalInterface(envctx.Empty, "a.b/c:d")
	assert.IsError(t, err, "ETLPLMMWCC")

	r = reset()
	err = r.UnmarshalInterface(tests.Context("a.b/c").Ctx(), "a.b/c:d")
	require.NoError(t, err)
	assert.Equal(t, *NewReference("a.b/c", "d"), *r)
}

func TestReferenceMarshalJson(t *testing.T) {

	var r *Reference
	b, _, _, err := r.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, nil, b)

	r = NewReference("a.b/c", "d")
	b, _, _, err = r.Repack(envctx.Empty)
	require.NoError(t, err)
	assert.Equal(t, "a.b/c:d", b)

}

func TestReferenceGetType(t *testing.T) {

	ty := &Type{
		Object: &Object{Id: NewReference("a.b/c", "d"), Type: NewReference("kego.io/system", "type")},
	}

	ctx := tests.Context("a.b/c").Stype("d", ty).Ctx()

	r := NewReference("a.b/c", "d")
	typ, ok := r.GetType(ctx)
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:d", typ.Id.Value())

	r = NewReference("a.b/c", "e")
	_, ok = r.GetType(ctx)
	assert.False(t, ok)

	r = &Reference{}
	_, ok = r.GetType(ctx)
	assert.False(t, ok)
}

func TestReferenceValue(t *testing.T) {

	r := NewReference("a.b/c", "d")
	assert.Equal(t, "a.b/c:d", r.Value())
	assert.Equal(t, "a.b/c:d", r.String())
	assert.Equal(t, "a.b/c:d", r.NativeString())

	r = NewReference("", "")
	assert.Equal(t, "", r.Value())
	assert.Equal(t, "", r.String())
	assert.Equal(t, "", r.NativeString())

	v, err := r.ValueContext(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "", v)

	cb := tests.Context("a.b/c")
	v, err = r.ValueContext(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, "", v)

	r = NewReference("a.b/c", "d")
	v, err = r.ValueContext(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, "d", v)

	r = NewReference("kego.io/json", "a")
	v, err = r.ValueContext(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, "json:a", v)

	r = NewReference("kego.io/system", "a")
	v, err = r.ValueContext(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, "system:a", v)

	cb.Alias("g", "d.e/f")
	r = NewReference("d.e/f", "h")
	v, err = r.ValueContext(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, "g:h", v)

	r = NewReference("i.j/k", "l")
	v, err = r.ValueContext(cb.Ctx())
	assert.IsError(t, err, "WGCDQQCFAD")

}

func TestNewReferenceFromString(t *testing.T) {
	cb := tests.Context("a.b/c")
	r, err := NewReferenceFromString(cb.Ctx(), "d")
	require.NoError(t, err)
	assert.Equal(t, NewReference("a.b/c", "d"), r)

	r, err = NewReferenceFromString(cb.Ctx(), "e:f")
	assert.IsError(t, err, "VXRGOQHWNB")

}

/*
func TestGetReflectType(t *testing.T) {
	a := ""
	r := reflect.TypeOf(a)
	cb := tests.Context("a.b/c").Jtype("d", r)
	ref := NewReference("a.b/c", "d")
	rt, ok := ref.GetReflectType(cb.Ctx())
	assert.True(t, ok)
	assert.Equal(t, r, rt)

	ref = NewReference("a.b/c", "e")
	rt, ok = ref.GetReflectType(cb.Ctx())
	assert.False(t, ok)
}
*/

/*
func TestGetReflectInterface(t *testing.T) {
	a := ""
	r := reflect.TypeOf(a)
	cb := tests.Context("a.b/c").Jiface("d", r)
	ref := NewReference("a.b/c", "d")
	rt, ok := ref.GetReflectInterface(cb.Ctx())
	assert.True(t, ok)
	assert.Equal(t, r, rt)

	ref = NewReference("a.b/c", "e")
	rt, ok = ref.GetReflectInterface(cb.Ctx())
	assert.False(t, ok)
}
*/

func TestGetDefault(t *testing.T) {
	rr := &ReferenceRule{Default: NewReference("a.b/c", "d")}
	d := rr.GetDefault()
	dr, ok := d.(*Reference)
	assert.True(t, ok)
	assert.Equal(t, *NewReference("a.b/c", "d"), *dr)
}

func TestSortableReferences(t *testing.T) {
	ra := []*Reference{NewReference("e.f/g", "h"), NewReference("a.b/c", "d")}
	sr := SortableReferences(ra)
	sort.Sort(sr)
	ra = []*Reference(sr)
	assert.Equal(t, 2, len(ra))
	assert.Equal(t, *NewReference("a.b/c", "d"), *ra[0])
	assert.Equal(t, *NewReference("e.f/g", "h"), *ra[1])
}

func TestReference_Label(t *testing.T) {
	var r *Reference
	assert.Equal(t, "", r.Label(nil))

	r = NewReference("a", "b")
	assert.Equal(t, "b", r.Label(nil))
}
