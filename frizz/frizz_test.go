package frizz_test

import (
	"testing"

	"path/filepath"

	"io/ioutil"

	"context"

	"frizz.io/frizz"
	"frizz.io/process/packages"
	"frizz.io/system"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestFrizz(t *testing.T) {

	ctx := frizz.NewContext(context.Background(), "frizz.io/system", nil)

	_, err := frizz.Open(ctx, "")
	assert.IsError(t, err, "CXIULJCEBE")

	systemDir, err := packages.GetDirFromPackage(ctx, "frizz.io/system")
	require.NoError(t, err)

	i, err := frizz.Open(ctx, filepath.Join(systemDir, "type.json"))
	require.NoError(t, err)
	_, ok := i.(*system.Type)
	assert.True(t, ok)

	b, err := ioutil.ReadFile(filepath.Join(systemDir, "type.json"))
	require.NoError(t, err)

	var i1 interface{}
	err = frizz.Unmarshal(ctx, b, &i1)
	require.NoError(t, err)
	_, ok = i1.(*system.Type)
	assert.True(t, ok)

	r := &system.Reference{}
	err = frizz.Unmarshal(ctx, []byte(`"type"`), r)
	require.NoError(t, err)
	assert.Equal(t, *system.NewReference("frizz.io/system", "type"), *r)

	b3, err := frizz.Marshal(ctx, system.NewReference("frizz.io/system", "type"))
	require.NoError(t, err)
	assert.Equal(t, "{\"type\":\"reference\",\"value\":\"type\"}", string(b3))

	b4, err := frizz.MarshalIndent(ctx, &system.Package{Recursive: true}, "", " ")
	require.NoError(t, err)
	assert.Equal(t, "{\n \"recursive\": true\n}", string(b4))

}
