package process

import (
	"testing"

	"kego.io/context/envctx"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestGetOptions(t *testing.T) {

	getTrue := func() *bool {
		val := true
		return &val
	}
	a := "a"
	f := FromFlags{
		Edit:     getTrue(),
		Update:   getTrue(),
		Log:      getTrue(),
		Validate: getTrue(),
		Debug:    getTrue(),
		Path:     &a,
	}
	d := f.getOptions()
	assert.True(t, d.Edit)
	assert.True(t, d.Update)
	assert.True(t, d.Log)
	assert.True(t, d.Validate)
	assert.True(t, d.Debug)
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
	assert.NoError(t, err)
	env := envctx.FromContext(ctx)
	assert.Equal(t, dirA, env.Dir)
	assert.Equal(t, pathA, env.Path)

	cb.OsWd("/")

	ctx, _, err = Initialise(ctx, &FromDefaults{
		Path: pathB,
	})
	env = envctx.FromContext(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dirB, env.Dir)
	assert.Equal(t, pathB, env.Path)

	_, _, err = Initialise(ctx, &FromDefaults{
		Path: "",
	})
	assert.IsError(t, err, "ADNJKTLAWY")
	assert.HasError(t, err, "CXOETFPTGM")
}
