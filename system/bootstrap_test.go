package system

import (
	"testing"

	"github.com/dave/ktest/require"
	"github.com/stretchr/testify/assert"
	"kego.io/tests"
)

func TestFoo(t *testing.T) {
	cb := tests.Context("kego.io/system")
	typ := new(Type)
	in := Pack(map[string]interface{}{"type": "type", "id": "a"})
	err := typ.Unpack(cb.Ctx(), in, false)
	require.NoError(t, err)
	assert.Equal(t, "a", typ.Id.Name)
}
