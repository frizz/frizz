package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"

	"crypto/sha1"

	"frizz.io/system/generate"
	"github.com/dave/jennifer/jen"
)

func regenerate(env vos.Env, save bool) (map[string][20]byte, error) {
	packages := map[string][20]byte{
		"frizz.io/tests/unpacker":     {},
		"frizz.io/tests/unpacker/sub": {},
	}
	for path := range packages {

		dir, err := patsy.Dir(env, path)
		if err != nil {
			return nil, err
		}

		buf := &bytes.Buffer{}
		if err := generate.Generate(buf, env, path, dir); err != nil {
			return nil, err
		}

		packages[path] = sha1.Sum(buf.Bytes())

		if save {
			if err := ioutil.WriteFile(filepath.Join(dir, "generated.frizz.go"), buf.Bytes(), 0777); err != nil {
				return nil, err
			}
		}
	}
	return packages, nil
}

func main() {
	env := vos.Os()

	packages, err := regenerate(env, true)

	f := jen.NewFilePathName("frizz.io/tests/cmd/regenerate", "main")
	f.Var().Id("hashes").Op("=").Map(jen.String()).Index(jen.Lit(20)).Byte().Values(
		jen.DictFunc(func(d jen.Dict) {
			for name, hash := range packages {
				d[jen.Lit(name)] = jen.ValuesFunc(func(g *jen.Group) {
					for _, b := range hash {
						g.Lit(int(b))
					}
				})
			}
		}),
	)
	dir, err := patsy.Dir(env, "frizz.io/tests/cmd/regenerate")
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Save(filepath.Join(dir, "hashes.generated.go")); err != nil {
		log.Fatal(err)
	}
}
