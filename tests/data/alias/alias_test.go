package alias

import (
	"testing"

	"github.com/davelondon/ktest/require"

	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/tests"
)

func Test(t *testing.T) {
	cb := tests.Context("kego.io/tests/data/alias").Jauto().Sauto(parser.Parse)
	s := `{
  "type": "main",
  "id": "data",
  "a": {"a": {"type":"simple", "js": "a0"}, "b": {"type":"simple", "js": "b0"}}
}`
	var m Main
	err := system.Unmarshal(cb.Ctx(), []byte(s), &m)
	require.NoError(t, err)
	require.Equal(t, "a0", m.A["a"].Js)
	require.Equal(t, "b0", m.A["b"].Js)

}
