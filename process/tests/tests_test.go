package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/pkgtest"
)

func TestInt(t *testing.T) {

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
				images:
					type: system:@map
					items:
						type: "@photo"
			rules:
				-
					selector: "{photo} .width"
					type: system:@int
					minimum: 800`,
		"rectangle.yaml": `
			type: system:type
			id: rectangle
			fields:
				width:
					type: system:@int
				height:
					type: system:@int`,
		"photo.yaml": `
			type: system:type
			id: photo
			fields:
				size:
					type: "@rectangle"`,
		"faces.yaml": `
			type: gallery
			id: faces
			images:
				foo:
					type: photo
					size:
						type: rectangle
						width: 500
						height: 500`,
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Minimum: value 500 must not be less than 800")

}

func TestSelector(t *testing.T) {

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
				images:
					type: system:@map
					items:
						type: "@photo"
						rules:
							-
								selector: ".protocol"
								type: system:@string
								equal: https`,
		"photo.yaml": `
			type: system:type
			id: photo
			fields:
				protocol:
					type: system:@string
					default: http
					optional: true`,
		"faces.yaml": `
			type: gallery
			id: faces
			images:
				foo:
					type: photo
					protocol: http`,
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Equal: value must equal 'https'")

}

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

//TODO: Fix this test (we removed globals, so we must scan for globals)
/*
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
		"site.yaml": fmt.Sprintf(`
			type: system:package
			import:
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
}*/

func runKego(namespace string, name string, files map[string]string) (string, error) {

	path, _, tests, err := pkgtest.CreateTemporaryPackage(namespace, name, files)
	if err != nil {
		return "", err
	}

	verbose := false
	set, err := process.InitialiseManually(false, false, false, verbose, path)
	if err != nil {
		return "", err
	}

	if err := process.KeCommand(set); err != nil {
		return "", err
	}

	if tests {
		if out, err := exec.Command("go", "test", path).CombinedOutput(); err != nil {
			return "", fmt.Errorf("%s", string(out))
		}
	}

	return set.Path(), nil
}
