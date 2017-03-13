package server

import (
	"testing"

	"context"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
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
