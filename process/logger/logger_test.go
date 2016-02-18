package logger

import (
	"bytes"
	"testing"

	"kego.io/kerr/assert"
)

func TestLogger(t *testing.T) {
	_, o, e := Logger(false)
	_, ok := o.(*bytes.Buffer)
	assert.True(t, ok)
	_, ok = e.(*bytes.Buffer)
	assert.True(t, ok)

	_, o, e = Logger(true)
	_, ok = o.(*multiWriter)
	assert.True(t, ok)
	_, ok = e.(*multiWriter)
	assert.True(t, ok)
}
