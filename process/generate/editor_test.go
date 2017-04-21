package generate

import (
	"testing"

	"frizz.io/tests"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestEditor(t *testing.T) {
	cb := tests.Context("a.b/c").Alias("g", "d.e/f")
	b, err := Editor(cb.Ctx(), cb.Env())
	require.NoError(t, err)
	assert.Contains(t, string(b), `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	fmt "fmt"
	client "frizz.io/editor/client"
	_ "frizz.io/system"
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
	assert.IsError(t, err, "NEUSQXAQVL")
}
