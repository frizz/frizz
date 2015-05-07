package system

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	e := Err(nil, "a", "b %s", "c")
	assert.Equal(t, "Error in a: b c.", e.Error())

	e = Err(fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "Error in b: c d returned an error: \na\n", e.Error())

	RETURN_ERRORS = true
	e = Err(fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "a", e.Error())
}
