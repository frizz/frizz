package main

import (
	"testing"

	_ "kego.io/system/types"
)

// This file is so that we can hook tests up to the debugger - because only apps can be debugged

func main() {
	tests := []testing.InternalTest{{"TestGo", TestGo}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}

func TestGo(t *testing.T) {

}
