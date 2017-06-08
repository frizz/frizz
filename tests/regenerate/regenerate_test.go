package regenerate

import (
	"testing"

	"github.com/dave/patsy/vos"
)

func TestRegenerate(t *testing.T) {
	env := vos.Os()
	hashes, err := Regenerate(env, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(hashes) != len(hashes) {
		t.Fatalf("%d packages generated as last regenerate, %d now", len(hashes), len(hashes))
	}
	for path, hash := range hashes {
		if hashes[path] != hash {
			t.Fatalf("generated code for %s changed since last regenerate", path)
		}
	}
}
