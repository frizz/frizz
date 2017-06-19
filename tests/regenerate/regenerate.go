package regenerate

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"

	"crypto/sha1"

	"frizz.io/frizz/generate"
	"github.com/dave/jennifer/jen"
)

var packages = []string{
	"frizz.io/system",
	"frizz.io/common",
	"frizz.io/validators",
	"frizz.io/tests/unpacker",
	"frizz.io/tests/unpacker/sub",
	"frizz.io/tests/validation",
}

// Regenerate generates the packages that are comprahensively tested, and
// writes the SHA1 hash of the generated code to a source file in this package.
// The test runs the regeneration (without saving the source) and checks these
// hashes to ensure that the tests are running with the correct generated files.
func Regenerate(env vos.Env, save bool) (map[string][20]byte, error) {
	hashes := map[string][20]byte{}
	for _, path := range packages {

		dir, err := patsy.Dir(env, path)
		if err != nil {
			return nil, err
		}

		buf := &bytes.Buffer{}
		if err := generate.Generate(buf, env, path, dir); err != nil {
			return nil, err
		}

		hashes[path] = sha1.Sum(buf.Bytes())

		if save {
			// notest
			if err := ioutil.WriteFile(filepath.Join(dir, "packer.frizz.go"), buf.Bytes(), 0777); err != nil {
				return nil, err
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
