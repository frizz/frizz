// debug is a command for hooking tests up to the delve debugger - because for some reason only
// commands (not tests) can be debugged
package main

// ke: {"notest":true}

import (
	"testing"
)

func main() {
	tests := []testing.InternalTest{{"TestGo", TestGo}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}

func TestGo(t *testing.T) {

}
