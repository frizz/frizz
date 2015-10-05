package json

import (
	"fmt"
	"testing"

	"kego.io/kerr"
	"kego.io/kerr/assert"
)

type UnpackA struct {
	B UnpackValue
	C UnpackArray
	D UnpackObject
}
type UnpackValue string
type UnpackArray []string
type UnpackObject map[string]string

func (u *UnpackValue) Unpack(in Unpackable) error {
	if in.UpType() != J_STRING {
		return kerr.New("FUGBYYGFUL", nil, "Should be string")
	}
	*u = UnpackValue(fmt.Sprint(in.UpString(), " bar"))
	return nil
}

var _ Unpacker = (*UnpackValue)(nil)

func (u *UnpackArray) Unpack(in Unpackable) error {
	if in.UpType() != J_ARRAY {
		return kerr.New("MYARYRIJLL", nil, "Should be array")
	}
	out := []string{}
	for _, child := range in.UpArray() {
		if child.UpType() != J_STRING {
			return kerr.New("OJPLBQMLLE", nil, "Children should be strings")
		}
		out = append(out, child.UpString())
	}
	out = append(out, "bar")
	*u = UnpackArray(out)
	return nil
}

var _ Unpacker = (*UnpackArray)(nil)

func (u *UnpackObject) Unpack(in Unpackable) error {
	if in.UpType() != J_MAP {
		return kerr.New("TMXIJXIGQK", nil, "Should be map")
	}
	out := map[string]string{}
	for name, child := range in.UpMap() {
		if child.UpType() != J_STRING {
			return kerr.New("QCEPMOJQCS", nil, "Children should be strings")
		}
		out[name] = child.UpString()
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

func (u *UnpackContextValue) Unpack(in Unpackable, path string, aliases map[string]string) error {
	if in.UpType() != J_STRING {
		return kerr.New("CUFSESRMCX", nil, "Should be string")
	}
	*u = UnpackContextValue(fmt.Sprint(in.UpString(), " bar ", path, " ", aliases["d.e/f"]))
	return nil
}

var _ ContextUnpacker = (*UnpackContextValue)(nil)

func (u *UnpackContextArray) Unpack(in Unpackable, path string, aliases map[string]string) error {
	if in.UpType() != J_ARRAY {
		return kerr.New("UXJTIVMLLG", nil, "Should be array")
	}
	out := []string{}
	for _, child := range in.UpArray() {
		if child.UpType() != J_STRING {
			return kerr.New("WFLWABUHTJ", nil, "Children should be strings")
		}
		out = append(out, child.UpString())
	}
	out = append(out, fmt.Sprint("bar ", path, " ", aliases["d.e/f"]))
	*u = UnpackContextArray(out)
	return nil
}

var _ ContextUnpacker = (*UnpackContextArray)(nil)

func (u *UnpackContextObject) Unpack(in Unpackable, path string, aliases map[string]string) error {
	if in.UpType() != J_MAP {
		return kerr.New("ACCJBEXHYG", nil, "Should be map")
	}
	out := map[string]string{}
	for name, child := range in.UpMap() {
		if child.UpType() != J_STRING {
			return kerr.New("CCNJNICMAQ", nil, "Children should be strings")
		}
		out[name] = child.UpString()
	}
	out["baz"] = fmt.Sprint("qux ", path, " ", aliases["d.e/f"])
	*u = UnpackContextObject(out)
	return nil
}

var _ ContextUnpacker = (*UnpackContextObject)(nil)

func TestUnpackContext(t *testing.T) {
	in := `{"B":"foo","C":["foo"],"D":{"foo":"bar"}}`
	a := &UnpackContextA{}
	err := UnmarshalPlainContext([]byte(in), a, "a.b/c", map[string]string{"d.e/f": "g"})
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
