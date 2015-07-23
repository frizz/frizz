package kego_test

import (
	"testing"

	"kego.io"
	"kego.io/kerr/assert"
)

func TestKego(t *testing.T) {
	_, err := kego.Open("", "", map[string]string{})
	assert.NoError(t, err)
}
