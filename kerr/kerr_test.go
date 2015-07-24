package kerr_test

import (
	"testing"

	"fmt"

	"kego.io/kerr"
	"kego.io/kerr/assert"
)

func TestErr(t *testing.T) {
	e := kerr.New("FITIHGYHTR", nil, "a", "b %s", "c")
	assert.Equal(t, "\nFITIHGYHTR error in kerr_test.go:13 TestErr: b c.\n", e.Error())

	e = kerr.New("LBDLIDLDPE", fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "\nLBDLIDLDPE error in kerr_test.go:16 TestErr: c d returned an error: \na", e.Error())

	// Should remove a leading new-line from errors
	e = kerr.New("OHUKDAEMPT", fmt.Errorf("\na"), "b", "c %s", "d")
	assert.Equal(t, "\nOHUKDAEMPT error in kerr_test.go:20 TestErr: c d returned an error: \na", e.Error())

	e = kerr.New("TUPDJYPRNU", nil, "a", "b")
	assert.Equal(t, "TUPDJYPRNU", e.ErrorId())

}

func TestSource(t *testing.T) {
	source := fmt.Errorf("Foo")
	e := kerr.New("PBNMPOIILQ", source, "a", "b")
	s := kerr.Source(e)
	assert.NotNil(t, s)
	assert.Equal(t, "Foo", s.Error())
}

type TestError struct {
	kerr.Struct
	Message string
}

func TestCustomError(t *testing.T) {

	inner := TestError{kerr.New("JOUYKMBSBU", nil, "a", "b"), "c"}
	outer := kerr.New("GJXHQCHGUO", inner, "d", "e")
	source := kerr.Source(outer)
	_, ok := source.(TestError)
	assert.True(t, ok)

}
