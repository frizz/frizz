package generate

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestEditor(t *testing.T) {
	cb := tests.Context("a.b/c").Alias("d.e/f", "g")
	b, err := Editor(cb.Ctx(), cb.Env())
	assert.NoError(t, err)
	assert.Contains(t, string(b), `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	"kego.io/editor/client"
	"kego.io/editor/client/console"
	_ "kego.io/system"
	_ "kego.io/system/editors"
)

func main() {
	if err := client.Start(); err != nil {
		console.Error(err.Error())
	}
}
`)
}

func TestEditorErr(t *testing.T) {
	cb := tests.Context("a.b/c").Alias("\"", "a")
	_, err := Editor(cb.Ctx(), cb.Env())
	assert.IsError(t, err, "CBTOLUQOGL")
	assert.HasError(t, err, "CRBYOUOHPG")
}
