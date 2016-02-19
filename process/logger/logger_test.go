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

func TestMultiWriter(t *testing.T) {
	var p, w1, w2 []byte
	pb := bytes.NewBuffer(p)
	w1b := bytes.NewBuffer(w1)
	w2b := bytes.NewBuffer(w2)
	mw := MultiWriter(pb, w1b, w2b)
	mw.Write([]byte("a"))
	mw.Write([]byte("b"))
	assert.Equal(t, "ab", pb.String())
	assert.Equal(t, "ab", w1b.String())
	assert.Equal(t, "ab", w2b.String())
}
