package tests

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/internal/pkgtest"
)

func TestRules(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	_, err = runKego(namespace, "a", map[string]string{
		"gallery.yaml": `
			type: system:type
			id: gallery
			fields:
				str:
					type: system:@string
					rules:
						-
							selector: ":root"
							type: system:@string
							equal: foo`,
		"faces.yaml": `
			type: gallery
			id: faces
			str: foo`,
	})
	assert.NoError(t, err)

	_, err = runKego(namespace, "b", map[string]string{
		"gallery.yaml": `
			type: system:type
			id: gallery
			fields:
				str:
					type: system:@string
					rules:
						-
							selector: ":root"
							type: system:@string
							equal: foo`,
		"faces.yaml": `
			type: gallery
			id: faces
			str: bar`,
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Equal: value must equal 'foo'")

}

func TestImport(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	namespaceDir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespaceDir)

	path, err := runKego(namespaceDir, "elements", map[string]string{
		"title.yaml": `
			type: system:type
			id: title
			fields:
				text:
					type: system:@string`,
		"page.yaml": `
			type: system:type
			id: page
			fields:
				title:
					type: "@title"`,
	})
	assert.NoError(t, err)

	_, err = runKego(namespaceDir, "site", map[string]string{
		"site.yaml": `
			type: elements:page
			id: site
			title:
				type: elements:title
				text: Heading`,
		"imports.yaml": fmt.Sprintf(`
			type: system:imports
			imports:
				%s: elements`, path),
		"site_test.go": `
			package site
			import (
				"testing"
				"kego.io/kerr/assert"
			)
			func TestMain(t *testing.T) {
				assert.Equal(t, "Heading", Site.Title.Text.Value)
			}`,
	})
	assert.NoError(t, err)
}

func runKego(namespace string, name string, files map[string]string) (string, error) {

	runTests := false
	for name, contents := range files {
		if strings.HasSuffix(name, ".yaml") {
			// our Go files are tab indented, but yaml files don't like tabs.
			files[name] = strings.Replace(contents, "\t", "    ", -1)
		}
		if strings.HasSuffix(name, "_test.go") {
			// if we add a xxx_test.go file we should also run "go test"
			runTests = true
		}
	}

	path, _, err := pkgtest.CreateTemporaryPackage(namespace, name, files)
	if err != nil {
		return "", err
	}

	set, err := process.InitialiseManually(false, false, false, true, path)
	if err != nil {
		return "", err
	}

	if err := process.KegoCmd(set); err != nil {
		return "", err
	}

	if runTests {
		if out, err := exec.Command("go", "test", path).CombinedOutput(); err != nil {
			return "", fmt.Errorf("%s", string(out))
		}
	}

	return set.Path(), nil
}
