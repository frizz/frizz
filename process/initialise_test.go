package process

import (
	"testing"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"kego.io/context/envctx"
	"kego.io/tests"
)

func TestGetOptions(t *testing.T) {

	getTrue := func() *bool {
		val := true
		return &val
	}
	a := "a"
	b := 2
	f := Flags{
		Edit:     getTrue(),
		Update:   getTrue(),
		Log:      getTrue(),
		Validate: getTrue(),
		Debug:    getTrue(),
		Path:     &a,
		Port:     &b,
	}
	d := f.getOptions()
	assert.True(t, d.Edit)
	assert.True(t, d.Update)
	assert.True(t, d.Log)
	assert.True(t, d.Validate)
	assert.True(t, d.Debug)
	assert.Equal(t, 2, d.Port)
	assert.Equal(t, "a", d.Path)

}

func TestInitialise(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"a.go":   "package a",
	})
	pathB, dirB := cb.TempPackage("b", map[string]string{
		"b.json": `{"type": "system:type", "id": "b"}`,
		"b.go":   "package b",
	})

	cb.OsWd(dirA)

	ctx, _, err := Initialise(cb.Ctx(), nil)
	require.NoError(t, err)
	env := envctx.FromContext(ctx)
	assert.Equal(t, dirA, env.Dir)
	assert.Equal(t, pathA, env.Path)

	cb.OsWd("/")

	ctx, _, err = Initialise(ctx, &Options{
		Path: pathB,
	})
	env = envctx.FromContext(ctx)
	require.NoError(t, err)
	assert.Equal(t, dirB, env.Dir)
	assert.Equal(t, pathB, env.Path)

	_, _, err = Initialise(ctx, &Options{
		Path: "",
	})
	assert.IsError(t, err, "ADNJKTLAWY")
	assert.HasErrorExternal(t, err, "CXOETFPTGM")
}
