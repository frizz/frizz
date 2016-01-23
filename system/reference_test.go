package system

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestUnpackDefaultNativeTypeReference(t *testing.T) {
	testUnpackDefaultNativeTypeReference(t, unmarshalFunc)
	testUnpackDefaultNativeTypeReference(t, unpackFunc)
}
func testUnpackDefaultNativeTypeReference(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"type": "a",
		"b": "e:f"
	}`

	type A struct {
		*Object
		B ReferenceInterface `json:"b"`
	}

	ctx := tests.Context("kego.io/system").Alias("c.d/e", "e").Jsystem().Jtype("a", reflect.TypeOf(&A{})).Ctx()

	var i interface{}
	err := unpacker(ctx, []byte(data), &i)
	assert.NoError(t, err)

	a, ok := i.(*A)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, a)
	assert.Equal(t, NewReference("c.d/e", "f"), a.B.GetReference(nil))

	b, err := json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"kego.io/system:a","b":"c.d/e:f"}`, string(b))

}

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

}

func TestReferenceUnmarshalJson(t *testing.T) {

	reset := func() *Reference {
		// Let's pre-load with some values so we check that when we
		// load a null value, we clear all the fields
		r := NewReference("a.b/c", "d")
		return r
	}

	r := reset()
	err := r.Unpack(envctx.Empty, json.Pack(nil))
	assert.IsError(t, err, "MOQVSKJXRB")

	r = reset()
	err = r.Unpack(envctx.Empty, json.Pack(1.0))
	assert.IsError(t, err, "RFLQSBPMYM")

	r = reset()
	err = r.Unpack(envctx.Empty, json.Pack("a.b/c:d"))
	assert.IsError(t, err, "MSXBLEIGVJ")
	assert.HasError(t, err, "KJSOXDESFD")
	p, ok := kerr.Source(err).(json.UnknownPackageError)
	assert.True(t, ok)
	assert.Equal(t, "a.b/c", p.UnknownPackage)

	ctx := tests.Context("").Alias("a.b/c", "c").Ctx()

	r = reset()
	err = r.Unpack(ctx, json.Pack("a.b/c:d"))
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "d", r.Name)
	assert.Equal(t, "a.b/c:d", r.Value())

	r = reset()
	err = r.Unpack(ctx, json.Pack("a.b/c:@d"))
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "a.b/c", r.Package)
	assert.Equal(t, "@d", r.Name)
	assert.Equal(t, "a.b/c:@d", r.Value())

	r = reset()
	err = r.Unpack(envctx.Empty, json.Pack("a:b"))
	assert.IsError(t, err, "MSXBLEIGVJ")
	assert.HasError(t, err, "DKKFLKDKYI")
	p, ok = kerr.Source(err).(json.UnknownPackageError)
	assert.True(t, ok)
	assert.Equal(t, "a", p.UnknownPackage)

}

func TestReferenceMarshalJson(t *testing.T) {

	var r *Reference
	b, err := r.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(b))

	r = NewReference("a.b/c", "d")
	b, err = r.MarshalJSON(envctx.Empty)
	assert.NoError(t, err)
	assert.Equal(t, "\"a.b/c:d\"", string(b))

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
	r := NewReference("", "")
	assert.Equal(t, "", r.Value())

	v, err := r.ValueContext(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "", v)

	cb := tests.Context("a.b/c")
	v, err = r.ValueContext(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "", v)

	r = NewReference("a.b/c", "d")
	v, err = r.ValueContext(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "d", v)

	r = NewReference("kego.io/json", "a")
	v, err = r.ValueContext(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "json:a", v)

	r = NewReference("kego.io/system", "a")
	v, err = r.ValueContext(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "system:a", v)

	cb.Alias("d.e/f", "g")
	r = NewReference("d.e/f", "h")
	v, err = r.ValueContext(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "g:h", v)

	r = NewReference("i.j/k", "l")
	v, err = r.ValueContext(cb.Ctx())
	assert.IsError(t, err, "WGCDQQCFAD")

}

func TestNewReferenceFromString(t *testing.T) {
	cb := tests.Context("a.b/c")
	r, err := NewReferenceFromString(cb.Ctx(), "d")
	assert.NoError(t, err)
	assert.Equal(t, NewReference("a.b/c", "d"), r)

	r, err = NewReferenceFromString(cb.Ctx(), "e:f")
	assert.IsError(t, err, "VXRGOQHWNB")

}
