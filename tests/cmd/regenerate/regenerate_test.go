package main

import (
	"testing"

	"github.com/dave/patsy/vos"
)

func Test(t *testing.T) {
	env := vos.Os()
	packages, err := regenerate(env, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(packages) != len(hashes) {
		t.Fatalf("%d packages generated as last regenerate, % now", len(hashes), len(packages))
	}
	for path, hash := range packages {
		if hashes[path] != hash {
			t.Fatalf("generated code for %s changed since last regenerate", path)
		}
	}
}
