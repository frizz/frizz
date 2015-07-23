package ke_test

import (
	"testing"

	"kego.io/ke"
	"kego.io/kerr/assert"
)

func TestKego(t *testing.T) {
	_, err := ke.Open("", "", map[string]string{})
	assert.Error(t, err)
}
