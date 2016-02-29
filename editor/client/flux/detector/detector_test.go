package detector

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestLoopDetector(t *testing.T) {
	s1 := &mockStore1{}
	s2 := &mockStore2{}
	d := New()

	found, store := d.RequestWait(s1, s2)
	assert.False(t, found)
	assert.Nil(t, store)

	found, store = d.RequestWait(s2, s1)
	assert.True(t, found)
	assert.Equal(t, s1, store)

	d.FinishedWait(s1)
	found, store = d.RequestWait(s2, s1)
	assert.False(t, found)
	assert.Nil(t, store)
}

func TestLoopDetectorInner(t *testing.T) {
	s1 := &mockStore1{}
	s2 := &mockStore2{}
	s3 := &mockStore3{}
	s4 := &mockStore4{}
	s5 := &mockStore5{}
	s6 := &mockStore6{}
	d := New()

	found, store := d.RequestWait(s1, s2)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.RequestWait(s2, s3)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.RequestWait(s3, s1)
	assert.True(t, found)
	assert.Equal(t, s1, store)
	found, store = d.RequestWait(s4, s5)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.RequestWait(s6, s4)
	assert.False(t, found)
	assert.Nil(t, store)
}

type mockStore1 struct{}
type mockStore2 struct{}
type mockStore3 struct{}
type mockStore4 struct{}
type mockStore5 struct{}
type mockStore6 struct{}
