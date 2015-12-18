package ke_test

import (
	"testing"

	"kego.io/context/envctx"
	"kego.io/ke"
	"kego.io/kerr/assert"
)

func TestKego(t *testing.T) {
	_, err := ke.Open(envctx.Empty, "")
	assert.Error(t, err)
}
