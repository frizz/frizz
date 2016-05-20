package flux

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
)

func TestDispatcherMock_Dispatch(t *testing.T) {
	d := &DispatcherMock{}
	c := d.Dispatch("a")
	require.Equal(t, 1, len(d.Log))
	assert.Equal(t, "a", d.Log[0])
	waitFor(t, c, false, "")
}
