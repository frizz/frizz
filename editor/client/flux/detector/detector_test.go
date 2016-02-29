package detector

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestLoopDetector(t *testing.T) {
	s1 := &mockStore{}
	s2 := &mockStore{}
	d := New()
	found, store := d.RequestWait(s1, s2)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.RequestWait(s2, s1)
	assert.True(t, found)
	assert.Equal(t, s1, store)
}

type mockStore struct{}
