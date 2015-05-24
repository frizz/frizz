package uerr_test

import (
	"testing"

	"fmt"

	"kego.io/assert"
	"kego.io/uerr"
)

func TestErr(t *testing.T) {
	e := uerr.New("FITIHGYHTR", nil, "a", "b %s", "c")
	assert.Equal(t, "\nFITIHGYHTR error in a: b c.\n", e.Error())

	e = uerr.New("LBDLIDLDPE", fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "\nLBDLIDLDPE error in b: c d returned an error: \na", e.Error())

	// Should remove a leading new-line from errors
	e = uerr.New("OHUKDAEMPT", fmt.Errorf("\na"), "b", "c %s", "d")
	assert.Equal(t, "\nOHUKDAEMPT error in b: c d returned an error: \na", e.Error())

	e = uerr.New("TUPDJYPRNU", nil, "a", "b")
	assert.Equal(t, "TUPDJYPRNU", e.Unique())

}
