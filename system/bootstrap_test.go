package system

import (
	"testing"

	"frizz.io/tests"
	"github.com/dave/ktest/require"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	cb := tests.Context("frizz.io/system")
	typ := new(Type)
	in := Pack(map[string]interface{}{"type": "type", "id": "a"})
	err := typ.Unpack(cb.Ctx(), in, false)
	require.NoError(t, err)
	assert.Equal(t, "a", typ.Id.Name)
}
