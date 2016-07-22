package client

import (
	"testing"

	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/davelondon/ktest/assert"
	"kego.io/context/sysctx"
	"kego.io/editor/shared"
	"kego.io/system"
	"kego.io/tests"
)

func TestRegisterTypes(t *testing.T) {
	cb := tests.New().Sempty().Jauto()

	imports := map[string]shared.ImportInfo{
		"a": {
			Path: "b",
			Types: map[string]shared.TypeInfo{
				"c": {Bytes: []byte(`{"type": "system:type", "id": "d"}`)},
			},
		},
		"g": {
			Path: "h",
			Types: map[string]shared.TypeInfo{
				"i": {Bytes: []byte(`{"type": "system:type", "id": "j"}`)},
			},
		},
	}
	pi, err := registerTypes(cb.Ctx(), "b", imports)
	assert.NoError(t, err)

	ti, ok := pi.Types.Get("d")
	assert.True(t, ok)
	assert.Equal(t, "b:d", ti.Type.(*system.Type).Id.String())

	sc := sysctx.FromContext(cb.Ctx())

	pi, ok = sc.Get("b")
	assert.True(t, ok)

	ti, ok = pi.Types.Get("d")
	assert.True(t, ok)
	assert.Equal(t, "b:d", ti.Type.(*system.Type).Id.String())

	pi, ok = sc.Get("h")
	assert.True(t, ok)

	ti, ok = pi.Types.Get("j")
	assert.True(t, ok)
	assert.Equal(t, "h:j", ti.Type.(*system.Type).Id.String())

}

func TestGetInfo(t *testing.T) {
	_, err := getInfo("foo")
	assert.IsError(t, err, "UTKDDLYKKH")

	info := &shared.Info{
		Path:    "a",
		Aliases: map[string]string{"b": "c", "d": "e"},
		Data:    map[string]string{"f": "g", "h": "i"},
		Package: []byte("j"),
		Imports: map[string]shared.ImportInfo{
			"k": {
				Path:    "l",
				Aliases: map[string]string{"m": "n", "o": "p"},
				Types: map[string]shared.TypeInfo{
					"q": {Bytes: []byte("r")},
					"s": {Bytes: []byte("t")},
				},
			},
			"u": {
				Path:    "v",
				Aliases: map[string]string{"w": "x", "y": "y"},
				Types: map[string]shared.TypeInfo{
					"aa": {Bytes: []byte("ab")},
					"ac": {Bytes: []byte("ad")},
				},
			},
		},
	}
	buf := bytes.NewBuffer([]byte{})
	err = gob.NewEncoder(buf).Encode(info)
	assert.NoError(t, err)
	val := base64.StdEncoding.EncodeToString(buf.Bytes())

	i, err := getInfo(val)
	assert.NoError(t, err)
	assert.Equal(t, info, i)

}
