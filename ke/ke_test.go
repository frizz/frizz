package ke_test

import (
	"testing"

	"path/filepath"

	"io/ioutil"

	"bytes"

	"golang.org/x/net/context"
	"kego.io/ke"
	"kego.io/kerr/assert"
	"kego.io/process/packages"
	"kego.io/system"
)

func TestKego(t *testing.T) {

	ctx := ke.NewContext(context.Background(), "kego.io/system", nil)

	_, err := ke.Open(ctx, "")
	assert.IsError(t, err, "CXIULJCEBE")

	systemDir, err := packages.GetDirFromPackage(ctx, "kego.io/system")
	assert.NoError(t, err)

	i, err := ke.Open(ctx, filepath.Join(systemDir, "type.json"))
	assert.NoError(t, err)
	_, ok := i.(*system.Type)
	assert.True(t, ok)

	b, err := ioutil.ReadFile(filepath.Join(systemDir, "type.json"))
	assert.NoError(t, err)

	var i1 interface{}
	err = ke.Unmarshal(ctx, b, &i1)
	assert.NoError(t, err)
	_, ok = i1.(*system.Type)
	assert.True(t, ok)

	r := &system.Reference{}
	err = ke.UnmarshalUntyped(ctx, []byte(`"type"`), r)
	assert.NoError(t, err)
	assert.Equal(t, *system.NewReference("kego.io/system", "type"), *r)

	var i2 interface{}
	buf := bytes.NewBuffer(b)
	err = ke.NewDecoder(ctx, buf).Decode(&i2)
	assert.NoError(t, err)
	_, ok = i2.(*system.Type)
	assert.True(t, ok)

	b1 := []byte{}
	bufw := bytes.NewBuffer(b1)
	err = ke.NewEncoder(bufw).Encode(system.NewReference("kego.io/system", "type"))
	assert.NoError(t, err)
	assert.Equal(t, "\"kego.io/system:type\"\n", bufw.String())

	b2, err := ke.Marshal(system.NewReference("kego.io/system", "type"))
	assert.NoError(t, err)
	assert.Equal(t, "\"kego.io/system:type\"", string(b2))

	b3, err := ke.MarshalContext(ctx, system.NewReference("kego.io/system", "type"))
	assert.NoError(t, err)
	assert.Equal(t, "\"type\"", string(b3))

}
