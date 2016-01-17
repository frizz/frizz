package kerr_test

import (
	"testing"

	"fmt"

	"kego.io/kerr"
	"kego.io/kerr/assert"
)

func TestErr(t *testing.T) {
	e := kerr.New("FITIHGYHTR", "b %s", "c")
	assert.Equal(t, "\nFITIHGYHTR error in kerr_ext_test.go:13 TestErr: b c.\n", e.Error())

	e = kerr.Wrap("LBDLIDLDPE", fmt.Errorf("a"))
	assert.Equal(t, "\nLBDLIDLDPE error in kerr_ext_test.go:16 TestErr: \na", e.Error())

	// Should remove a leading new-line from errors
	e = kerr.Wrap("OHUKDAEMPT", fmt.Errorf("\na"))
	assert.Equal(t, "\nOHUKDAEMPT error in kerr_ext_test.go:20 TestErr: \na", e.Error())

	e = kerr.New("TUPDJYPRNU", "b")
	assert.Equal(t, "TUPDJYPRNU", e.ErrorId())

}

func TestSource(t *testing.T) {
	source := fmt.Errorf("Foo")
	e := kerr.Wrap("PBNMPOIILQ", source)
	s := kerr.Source(e)
	assert.NotNil(t, s)
	assert.Equal(t, "Foo", s.Error())
}

type TestError struct {
	kerr.Struct
	Message string
}

func TestCustomError(t *testing.T) {

	inner := TestError{Struct: kerr.New("JOUYKMBSBU", "b"), Message: "c"}
	outer := kerr.Wrap("GJXHQCHGUO", inner)
	source := kerr.Source(outer)
	_, ok := source.(TestError)
	assert.True(t, ok)

}
