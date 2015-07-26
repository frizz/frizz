package tests

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"fmt"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/internal/pkgtest"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

// TODO: this is WIP - currently stuck getting comments to line up.
// TODO: See http://stackoverflow.com/questions/31628613 for my question.
func TestNonDestructiveGeneration(t *testing.T) {

	t.Skip()

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	namespaceDir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespaceDir)

	path, _, _, err := pkgtest.CreateTemporaryPackage(namespaceDir, "a", map[string]string{
		"a.yaml": `
			type: system:type
			id: b
			fields:
				c:
					type: system:@string
				d:
					type: system:@number`,
		"b.go": `package a

		import "kego.io/system"

		// This is a struct comment
		type B struct {
			// This is a field comment
			C system.String
		}`,
	})

	verbose := false
	set, err := process.InitialiseManually(false, false, false, verbose, false, path)
	assert.NoError(t, err)

	err = process.GenerateStructs(set)
	assert.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(namespaceDir, "a", "b.go"))
	assert.NoError(t, err)

	assert.Contains(t, string(b), "package a")
	fmt.Println(string(b))

}
