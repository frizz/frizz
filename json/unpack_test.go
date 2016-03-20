package json

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestGetTypeFromField(t *testing.T) {
	//(us *unpackStruct) getTypeFromField(ctx context.Context, in Packed, iface reflect.Value) (reflect.Type, error) {

	cb := tests.New()
	us := &unpackStruct{}

	//in := &packed{v: map[string]interface{}{"a": "b"}}
	in := &packed{v: ""}
	var v1 string
	_, err := us.getTypeFromField(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.IsError(t, err, "LCJRIHJXFU")

	in = &packed{v: map[string]interface{}{"a": "b"}}
	_, err = us.getTypeFromField(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.IsError(t, err, "RMMVQNVHTU")

	in = &packed{v: map[string]interface{}{"type": 1.0}}
	_, err = us.getTypeFromField(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.IsError(t, err, "RPBSKPRLJQ")

}

func TestUnpackObject(t *testing.T) {
	cb := tests.New()
	us := &unpackStruct{}

	in := &packed{v: "a"}
	var v1 string
	err := us.unpackObject(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.IsError(t, err, "PBAXKEKVTA")

	in = &packed{v: map[string]interface{}{"a": "b"}}
	var v2 UnpackObject
	err = us.unpackObject(cb.Ctx(), in, reflect.ValueOf(&v2))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(v2))
	assert.Equal(t, "b", v2["a"])
	assert.Equal(t, "qux", v2["baz"])

	cb.Path("a").Jempty()
	// TODO: We insist on a type attribute, but ignore it if no type is found. Is this a bug? Maybe
	// TODO: we should only throw the error for no type if the interface has methods? Or maybe we
	// TODO: should return an error if no typs is found?
	in = &packed{v: map[string]interface{}{"type": "foo", "a": "b"}}
	var v3 interface{}
	err = us.unpackObject(cb.Ctx(), in, reflect.ValueOf(&v3))
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"type": "foo", "a": "b"}, v3)

	in = &packed{v: map[string]interface{}{"type": "foo", "a": "b"}}
	var v4 map[int]interface{}
	err = us.unpackObject(cb.Ctx(), in, reflect.ValueOf(&v4))
	assert.IsError(t, err, "TXNQGFVHOT")

	assert.SkipError("SRULCNWOWM")

}

func TestUnpackArray(t *testing.T) {
	cb := tests.New()
	us := &unpackStruct{}

	in := &packed{v: "a"}
	var v1 string
	err := us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.IsError(t, err, "PVJMVPULMY")

	in = &packed{v: []interface{}{"a"}}
	var v2 UnpackArray
	err = us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v2))
	assert.NoError(t, err)
	assert.Equal(t, UnpackArray{"a", "bar"}, v2)

	in = &packed{v: []interface{}{"a"}}
	var v3 string
	err = us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v3))
	assert.IsError(t, err, "AODAOUPIED")

	in = &packed{v: []interface{}{"a"}}
	v4 := []interface{}{"", "", ""}
	err = us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v4))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(v4))

	in = &packed{v: []interface{}{"a"}}
	v5 := [3]interface{}{"b", "c", "d"}
	err = us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v5))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(v5))
	assert.Equal(t, "a", v5[0])
	assert.Equal(t, nil, v5[1])
	assert.Equal(t, nil, v5[2])

	in = &packed{v: []interface{}{}}
	v6 := []interface{}{"b", "c", "d"}
	err = us.unpackArray(cb.Ctx(), in, reflect.ValueOf(&v6))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(v6))

}

func TestSetDefaultNativeValueUnpack(t *testing.T) {
	cb := tests.New().Jempty()
	type s struct{}
	err := setDefaultNativeValueUnpack(cb.Ctx(), reflect.ValueOf(s{}), Pack([]byte{}))
	assert.IsError(t, err, "YSBBTCVOUU")
}

func TestUnpackLiteral(t *testing.T) {
	cb := tests.New()
	us := &unpackStruct{}
	in := &packed{v: "a"}
	upv := UnpackValue("b")
	err := us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&upv))
	assert.NoError(t, err)
	// UnpackValue.Unpack adds "bar" to the end of whatever it unpacks
	assert.Equal(t, "a bar", string(upv))

	in = &packed{v: nil}
	var s *string
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&s))
	assert.NoError(t, err)
	assert.Equal(t, (*string)(nil), s)

	in = &packed{v: ""}
	var b *bool
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&b))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal string into Go value of type bool", err.Error())

	in = &packed{v: true}
	var i interface{}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&i))
	assert.NoError(t, err)
	assert.Equal(t, true, i)

	in = &packed{v: true}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&s))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal bool into Go value of type string", err.Error())

	in = &packed{v: ""}
	v1 := []int{}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v1))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal string into Go value of type []int", err.Error())

	in = &packed{v: "Zm9v"}
	v2 := []uint8{}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v2))
	assert.NoError(t, err)
	assert.Equal(t, "foo", string(v2))

	in = &packed{v: "a"}
	var v3 interface{}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v3))
	assert.NoError(t, err)
	assert.Equal(t, "a", v3)

	in = &packed{v: 2.0}
	var v4 interface{}
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v4))
	assert.NoError(t, err)
	assert.Equal(t, 2.0, v4)

	in = &packed{v: 2.0}
	v5 := false
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v5))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal number into Go value of type bool", err.Error())

	in = &packed{v: 2.0}
	v6 := NumberLiteral("")
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v6))
	assert.NoError(t, err)
	assert.Equal(t, "2", string(v6))

	in = &packed{v: 2.0}
	var v7 int
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v7))
	assert.NoError(t, err)
	assert.Equal(t, 2, v7)

	in = &packed{v: 100000000000000000000.0}
	var v8 int
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v8))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal number 100000000000000000000 into Go value of type int", err.Error())

	in = &packed{v: 2.0}
	var v9 uint
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v9))
	assert.NoError(t, err)
	assert.Equal(t, 2, int(v9))

	in = &packed{v: 100000000000000000000.0}
	var v10 uint
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v10))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal number 100000000000000000000 into Go value of type uint", err.Error())

	in = &packed{v: 2.0}
	var v11 float32
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v11))
	assert.NoError(t, err)
	assert.Equal(t, 2, int(v11))

	in = &packed{v: 10000000000000000000000000000000000000000.0}
	var v12 float32
	err = us.unpackLiteral(cb.Ctx(), in, reflect.ValueOf(&v12))
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal number 10000000000000000000000000000000000000000 into Go value of type float32", err.Error())

}

func TestUnpackFragment(t *testing.T) {
	cb := tests.New()
	var v interface{}
	err := UnpackFragment(cb.Ctx(), &packed{v: "a"}, &v, reflect.TypeOf(""))
	assert.NoError(t, err)
	assert.Equal(t, "a", v)
}

func TestUnpackErrors(t *testing.T) {
	err := Unpack(nil, nil, nil)
	assert.IsError(t, err, "QAINAWPLSU")
	err = Unpack(nil, &packed{v: ""}, nil)
	assert.IsError(t, err, "XOOUKLGORQ")
}

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

type UnpackA struct {
	B UnpackValue
	C UnpackArray
	D UnpackObject
}
type UnpackValue string
type UnpackArray []string
type UnpackObject map[string]string

func (u *UnpackValue) Unpack(ctx context.Context, in Packed) error {
	if in.Type() != J_STRING {
		return kerr.New("FUGBYYGFUL", "Should be string")
	}
	*u = UnpackValue(fmt.Sprint(in.String(), " bar"))
	return nil
}

var _ Unpacker = (*UnpackValue)(nil)

func (u *UnpackArray) Unpack(ctx context.Context, in Packed) error {
	if in.Type() != J_ARRAY {
		return kerr.New("MYARYRIJLL", "Should be array")
	}
	out := []string{}
	for _, child := range in.Array() {
		if child.Type() != J_STRING {
			return kerr.New("OJPLBQMLLE", "Children should be strings")
		}
		out = append(out, child.String())
	}
	out = append(out, "bar")
	*u = UnpackArray(out)
	return nil
}

var _ Unpacker = (*UnpackArray)(nil)

func (u *UnpackObject) Unpack(ctx context.Context, in Packed) error {
	if in.Type() != J_MAP {
		return kerr.New("TMXIJXIGQK", "Should be map")
	}
	out := map[string]string{}
	for name, child := range in.Map() {
		if child.Type() != J_STRING {
			return kerr.New("QCEPMOJQCS", "Children should be strings")
		}
		out[name] = child.String()
	}
	out["baz"] = "qux"
	*u = UnpackObject(out)
	return nil
}

var _ Unpacker = (*UnpackObject)(nil)

func TestUnpack(t *testing.T) {
	in := `{"B":"foo","C":["foo"],"D":{"foo":"bar"}}`
	a := &UnpackA{}
	err := UnmarshalPlain([]byte(in), a)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "foo bar", string(a.B))
	assert.Equal(t, 2, len(a.C))
	assert.Equal(t, "foo", a.C[0])
	assert.Equal(t, "bar", a.C[1])
	assert.Equal(t, 2, len(a.D))
	assert.Equal(t, "bar", a.D["foo"])
	assert.Equal(t, "qux", a.D["baz"])

}

type UnpackContextA struct {
	B UnpackContextValue
	C UnpackContextArray
	D UnpackContextObject
}
type UnpackContextValue string
type UnpackContextArray []string
type UnpackContextObject map[string]string

func (u *UnpackContextValue) Unpack(ctx context.Context, in Packed) error {

	env := envctx.FromContext(ctx)

	if in.Type() != J_STRING {
		return kerr.New("CUFSESRMCX", "Should be string")
	}
	*u = UnpackContextValue(fmt.Sprint(in.String(), " bar ", env.Path, " ", env.Aliases["g"]))
	return nil
}

var _ Unpacker = (*UnpackContextValue)(nil)

func (u *UnpackContextArray) Unpack(ctx context.Context, in Packed) error {

	env := envctx.FromContext(ctx)

	if in.Type() != J_ARRAY {
		return kerr.New("UXJTIVMLLG", "Should be array")
	}
	out := []string{}
	for _, child := range in.Array() {
		if child.Type() != J_STRING {
			return kerr.New("WFLWABUHTJ", "Children should be strings")
		}
		out = append(out, child.String())
	}
	out = append(out, fmt.Sprint("bar ", env.Path, " ", env.Aliases["g"]))
	*u = UnpackContextArray(out)
	return nil
}

var _ Unpacker = (*UnpackContextArray)(nil)

func (u *UnpackContextObject) Unpack(ctx context.Context, in Packed) error {

	env := envctx.FromContext(ctx)

	if in.Type() != J_MAP {
		return kerr.New("ACCJBEXHYG", "Should be map")
	}
	out := map[string]string{}
	for name, child := range in.Map() {
		if child.Type() != J_STRING {
			return kerr.New("CCNJNICMAQ", "Children should be strings")
		}
		out[name] = child.String()
	}
	out["baz"] = fmt.Sprint("qux ", env.Path, " ", env.Aliases["g"])
	*u = UnpackContextObject(out)
	return nil
}

var _ Unpacker = (*UnpackContextObject)(nil)

func TestUnpackContext(t *testing.T) {
	in := `{"B":"foo","C":["foo"],"D":{"foo":"bar"}}`
	a := &UnpackContextA{}
	err := UnmarshalUntyped(tests.Context("a.b/c").Alias("g", "d.e/f").Ctx(), []byte(in), a)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "foo bar a.b/c d.e/f", string(a.B))
	assert.Equal(t, 2, len(a.C))
	assert.Equal(t, "foo", a.C[0])
	assert.Equal(t, "bar a.b/c d.e/f", a.C[1])
	assert.Equal(t, 2, len(a.D))
	assert.Equal(t, "bar", a.D["foo"])
	assert.Equal(t, "qux a.b/c d.e/f", a.D["baz"])
}
