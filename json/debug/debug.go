package main

import "testing"

// This file is so that we can hook tests up to the debugger - because only apps can be debugged

func main() {
	tests := []testing.InternalTest{{"TestDebug", TestDebug}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}

func TestDebug(t *testing.T) {
	// copy test in here
}
