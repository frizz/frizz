package generate

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestEditor(t *testing.T) {
	cb := tests.Context("a.b/c").Alias("g", "d.e/f")
	b, err := Editor(cb.Ctx(), cb.Env())
	assert.NoError(t, err)
	assert.Contains(t, string(b), `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	"fmt"
	"kego.io/editor/client"
	_ "kego.io/system"
)

func main() {
	if err := client.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
`)
}

func TestEditorErr(t *testing.T) {
	cb := tests.Context("a.b/c").Alias("a", "\"")
	_, err := Editor(cb.Ctx(), cb.Env())
	assert.IsError(t, err, "CBTOLUQOGL")
	assert.HasError(t, err, "CRBYOUOHPG")
}
