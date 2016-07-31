package server

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"golang.org/x/net/context"
)

func TestServer_Data(t *testing.T) {
	source := []byte(`
	package main
	import "fmt"
	func main(){
		fmt.Println("foo")
	}`)
	c, err := Compile(context.Background(), source, false)
	require.NoError(t, err)
	assert.True(t, len(c) > 100000) // should be ~700K

	m, err := Compile(context.Background(), source, true)
	require.NoError(t, err)
	assert.True(t, len(m) > 10000) // should be ~35K

}
