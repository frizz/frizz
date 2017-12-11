package regenerate

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"

	"bytes"
	"crypto/sha1"

	"fmt"
	"io"

	"frizz.io/generate"
	"github.com/dave/jennifer/jen"
)

var packages = []string{
	"frizz.io/system",
	"frizz.io/common",
	"frizz.io/validators",
	"frizz.io/tests/packer",
	"frizz.io/tests/packer/silent",
	"frizz.io/tests/packer/silent/use",
	"frizz.io/tests/packer/sub",
	"frizz.io/tests/data",
	"frizz.io/tests/validation",
}

// Regenerate generates the packages that are comprahensively tested, and
// writes the SHA1 hash of the generated code to a source file in this package.
// The test runs the regeneration (without saving the source) and checks these
// hashes to ensure that the tests are running with the correct generated files.
func Regenerate(env vos.Env, out io.Writer, save bool) (map[string][20]byte, error) {
	hashes := map[string][20]byte{}

	scan, err := generate.Scan(env, packages...)
	if err != nil {
		return nil, err
	}

	for path, def := range scan.Packages {
		generated := &bytes.Buffer{}
		hasContent, err := def.Generate(generated)
		if err != nil {
			return nil, err
		}
		hashes[path] = sha1.Sum(generated.Bytes())

		if save {
			fname := filepath.Join(def.Dir, "generated.frizz.go")

			// notest

			var existing []byte
			if _, err := os.Stat(fname); err == nil {
				existing, err = ioutil.ReadFile(fname)
				if err != nil {
					return nil, errors.WithMessage(err, "reading existing source")
				}
			}

			if hasContent {
				if bytes.Equal(existing, generated.Bytes()) {
					fmt.Fprintf(out, "unchanged: %s\n", def.Path)
				} else {
					if err := ioutil.WriteFile(fname, generated.Bytes(), 0777); err != nil {
						return nil, errors.Wrap(err, "saving source")
					}
					fmt.Fprintf(out, "updated: %s\n", def.Path)
				}
			} else {
				if _, err := os.Stat(fname); err == nil {
					if err := os.Remove(fname); err != nil {
						return nil, errors.Wrap(err, "removing source")
					}
					fmt.Fprintf(out, "removed: %s\n", def.Path)
				}
			}
		}
	}
	return hashes, nil
}

// SaveHashes saves the hashes of the generated packages to a source file in
// this package. The test runs the regeneration again and compares these hashes
// to the new values.
func SaveHashes(env vos.Env, hashes map[string][20]byte) error {
	// notest
	f := jen.NewFile("regenerate")
	f.Var().Id("hashes").Op("=").Map(jen.String()).Index(jen.Lit(20)).Byte().Values(
		jen.DictFunc(func(d jen.Dict) {
			for name, hash := range hashes {
				d[jen.Lit(name)] = jen.ValuesFunc(func(g *jen.Group) {
					for _, b := range hash {
						g.Lit(int(b))
					}
				})
			}
		}),
	)
	dir, err := patsy.Dir(env, "frizz.io/tests/regenerate")
	if err != nil {
		return err
	}
	if err := f.Save(filepath.Join(dir, "hashes.frizz.go")); err != nil {
		return err
	}
	return nil
}
