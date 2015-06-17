package kerr_test

import (
	"testing"

	"fmt"

	"kego.io/kerr"
	"kego.io/kerr/assert"
)

func TestErr(t *testing.T) {
	e := kerr.New("FITIHGYHTR", nil, "a", "b %s", "c")
	assert.Equal(t, "\nFITIHGYHTR error in a: b c.\n", e.Error())

	e = kerr.New("LBDLIDLDPE", fmt.Errorf("a"), "b", "c %s", "d")
	assert.Equal(t, "\nLBDLIDLDPE error in b: c d returned an error: \na", e.Error())

	// Should remove a leading new-line from errors
	e = kerr.New("OHUKDAEMPT", fmt.Errorf("\na"), "b", "c %s", "d")
	assert.Equal(t, "\nOHUKDAEMPT error in b: c d returned an error: \na", e.Error())

	e = kerr.New("TUPDJYPRNU", nil, "a", "b")
	assert.Equal(t, "TUPDJYPRNU", e.Unique())

}
