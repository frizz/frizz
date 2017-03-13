package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

type st struct {
	*Store
}

func (s *st) Handle(payload *Payload) (finished bool) { return true }

func TestStore_Init(t *testing.T) {
	s := &st{Store: &Store{}}
	s.Init(s)
	assert.NotNil(t, s.self)
	assert.Equal(t, s.self, s)
}
