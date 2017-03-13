package tests_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"context"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"kego.io/context/envctx"
	"kego.io/process"
	"kego.io/tests"
)

func TestDefaultInterfaceNativeType(t *testing.T) {

	cb := tests.New()
	defer cb.Cleanup()

	_, err := runKe(cb, "a", map[string]string{
		"foo.yaml": `
			type: system:type
			id: foo
			fields:
				baz:
					type: system:@string
					interface: true`,
		"bar.yaml": `
			type: foo
			id: bar
			baz: qux`,
	})
	require.NoError(t, err)

}
func TestNeedsDummyRule(t *testing.T) {

	cb := tests.New()
	defer cb.Cleanup()

	_, err := runKe(cb, "a", map[string]string{
		"foo.yaml": `
			type: system:type
			id: foo
			fields:
				bar:
					type: system:@string
			rule:
				type: system:type
				embed: ["system:rule"]
				fields:
					default:
						type: "@foo"`,
	})
	require.NoError(t, err)

}

func TestInt(t *testing.T) {

	cb := tests.New()
	defer cb.Cleanup()

	_, err := runKe(cb, "a", map[string]string{
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

	cb := tests.New()
	defer cb.Cleanup()

	_, err := runKe(cb, "a", map[string]string{
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
	assert.Contains(t, err.Error(), "Equal: value \"http\" must equal 'https'")

}

func TestRules(t *testing.T) {

	cb := tests.New()
	defer cb.Cleanup()

	_, err := runKe(cb, "a", map[string]string{
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
	require.NoError(t, err)

	_, err = runKe(cb, "b", map[string]string{
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
	assert.Contains(t, err.Error(), "Equal: value \"bar\" must equal 'foo'")

}

//TODO: Fix this test (we removed globals, so we must scan for globals)
/*
func TestImport(t *testing.T) {

	namespaceDir, err := tests.CreateTemporaryNamespace()
	require.NoError(t, err)
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
	require.NoError(t, err)

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
				"github.com/dave/ktest/assert"
			)
			func TestMain(t *testing.T) {
				assert.Equal(t, "Heading", Site.Title.Text.Value())
			}`,
	})
	require.NoError(t, err)
}*/

func runKe(cb *tests.ContextBuilder, name string, files map[string]string) (string, error) {

	tests := false
	for name, _ := range files {
		if strings.HasSuffix(name, "_test.go") {
			// if we add a xxx_test.go file we should also run "go test"
			tests = true
		}
	}

	path, _ := cb.RealGopath().TempPackage(name, files)

	ctx, _, err := process.Initialise(context.Background(), &process.Options{
		Path: path,
	})
	if err != nil {
		return "", err
	}

	env := envctx.FromContext(ctx)

	if err := process.Generate(ctx, env); err != nil {
		return "", err
	}

	if err := process.RunValidateCommand(ctx); err != nil {
		return "", err
	}

	if tests {
		if out, err := exec.Command("go", "test", path).CombinedOutput(); err != nil {
			return "", fmt.Errorf("%s", string(out))
		}
	}

	return env.Path, nil
}
