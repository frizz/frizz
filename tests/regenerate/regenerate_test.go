package regenerate

import (
	"testing"

	"bytes"

	"github.com/dave/patsy/vos"
)

func TestRegenerate(t *testing.T) {
	env := vos.Os()
	buf := &bytes.Buffer{}
	now, err := Regenerate(env, buf, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(now) != len(hashes) {
		t.Fatalf("%d packages generated as last regenerate, %d now", len(hashes), len(hashes))
	}
	for path, hash := range hashes {
		if now[path] != hash {
			t.Fatalf("generated code for %s changed since last regenerate", path)
		}
	}
}
