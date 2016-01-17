package json

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

type UnpackA struct {
	B	UnpackValue
	C	UnpackArray
	D	UnpackObject
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
	B	UnpackContextValue
	C	UnpackContextArray
	D	UnpackContextObject
}
type UnpackContextValue string
type UnpackContextArray []string
type UnpackContextObject map[string]string

func (u *UnpackContextValue) Unpack(ctx context.Context, in Packed) error {

	env := envctx.FromContext(ctx)

	if in.Type() != J_STRING {
		return kerr.New("CUFSESRMCX", "Should be string")
	}
	*u = UnpackContextValue(fmt.Sprint(in.String(), " bar ", env.Path, " ", env.Aliases["d.e/f"]))
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
	out = append(out, fmt.Sprint("bar ", env.Path, " ", env.Aliases["d.e/f"]))
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
	out["baz"] = fmt.Sprint("qux ", env.Path, " ", env.Aliases["d.e/f"])
	*u = UnpackContextObject(out)
	return nil
}

var _ Unpacker = (*UnpackContextObject)(nil)

func TestUnpackContext(t *testing.T) {
	in := `{"B":"foo","C":["foo"],"D":{"foo":"bar"}}`
	a := &UnpackContextA{}
	err := UnmarshalUntyped(tests.Context("a.b/c").Alias("d.e/f", "g").Ctx(), []byte(in), a)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "foo bar a.b/c g", string(a.B))
	assert.Equal(t, 2, len(a.C))
	assert.Equal(t, "foo", a.C[0])
	assert.Equal(t, "bar a.b/c g", a.C[1])
	assert.Equal(t, 2, len(a.D))
	assert.Equal(t, "bar", a.D["foo"])
	assert.Equal(t, "qux a.b/c g", a.D["baz"])
}
