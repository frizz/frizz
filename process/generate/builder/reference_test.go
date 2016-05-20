package builder

import (
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestReference(t *testing.T) {

	mockGetAlias := func(s string) string {
		return "[alias:" + s + "]"
	}
	s := Reference("a.b/c", "d", "e.f/g", mockGetAlias)
	assert.Equal(t, "[alias:a.b/c].d", s)

	s = Reference("a.b/c", "d", "a.b/c", mockGetAlias)
	assert.Equal(t, "d", s)

	s = Reference("kego.io/json", "String", "a.b/c", mockGetAlias)
	assert.Equal(t, "string", s)

	s = Reference("kego.io/json", "Number", "a.b/c", mockGetAlias)
	assert.Equal(t, "float64", s)

	s = Reference("kego.io/json", "Bool", "a.b/c", mockGetAlias)
	assert.Equal(t, "bool", s)
}
