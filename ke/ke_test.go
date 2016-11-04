package ke_test

import (
	"testing"

	"path/filepath"

	"io/ioutil"

	"context"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/ke"
	"kego.io/process/packages"
	"kego.io/system"
)

func TestKego(t *testing.T) {

	ctx := ke.NewContext(context.Background(), "kego.io/system", nil)

	_, err := ke.Open(ctx, "")
	assert.IsError(t, err, "CXIULJCEBE")

	systemDir, err := packages.GetDirFromPackage(ctx, "kego.io/system")
	require.NoError(t, err)

	i, err := ke.Open(ctx, filepath.Join(systemDir, "type.json"))
	require.NoError(t, err)
	_, ok := i.(*system.Type)
	assert.True(t, ok)

	b, err := ioutil.ReadFile(filepath.Join(systemDir, "type.json"))
	require.NoError(t, err)

	var i1 interface{}
	err = ke.Unmarshal(ctx, b, &i1)
	require.NoError(t, err)
	_, ok = i1.(*system.Type)
	assert.True(t, ok)

	r := &system.Reference{}
	err = ke.Unmarshal(ctx, []byte(`"type"`), r)
	require.NoError(t, err)
	assert.Equal(t, *system.NewReference("kego.io/system", "type"), *r)

	b3, err := ke.Marshal(ctx, system.NewReference("kego.io/system", "type"))
	require.NoError(t, err)
	assert.Equal(t, "{\"type\":\"reference\",\"value\":\"type\"}", string(b3))

	b4, err := ke.MarshalIndent(ctx, &system.Package{Recursive: true}, "", " ")
	require.NoError(t, err)
	assert.Equal(t, "{\n \"recursive\": true\n}", string(b4))

}
