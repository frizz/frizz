package uerr

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	e := New("FITIHGYHTR", nil, "a", "b %s", "c")
	assert.Equal(t, "Error in a: b c.", e.Error())

	e = New("LBDLIDLDPE", fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "Error in b: c d returned an error: \na\n", e.Error())

	e = New("TUPDJYPRNU", nil, "a", "b")
	assert.Equal(t, "TUPDJYPRNU", e.Unique())

}
