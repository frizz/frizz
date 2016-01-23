package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestEmptyPackage(t *testing.T) {
	ep := EmptyPackage()
	assert.NotNil(t, ep.Object)
	assert.Equal(t, NewReference("kego.io/system", "package"), ep.Type)
	assert.NotNil(t, ep.Aliases)
}
