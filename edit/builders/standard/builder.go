package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"frizz.io/edit/common"

	"fmt"

	"runtime"
	"strings"

	"os"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

// compiles frizz.io/edit/bootstrap to javascript and saves the output to frizz.io/edit/static/data/bootstrap.js

func main() {

	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/assets/data")
	if err != nil {
		log.Fatal(err)
	}

	if err := stdlib(dir); err != nil {
		log.Fatal(fmt.Sprintf("%+v\n", err))
	}
}

func stdlib(dir string) error {
	fmt.Println("writing standard library")
	out := filepath.Join(dir, "go", "src")
	if err := os.RemoveAll(out); err != nil {
		return errors.WithStack(err)
	}
	root := filepath.Join(runtime.GOROOT(), "src")
	if err := addDir(root, root, out); err != nil {
		return err
	}
	return nil
}

func addDir(base, dir, out string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.WithStack(err)
	}
	relDir, err := filepath.Rel(base, dir)
	if err != nil {
		return errors.WithStack(err)
	}
	p := common.Package{}
	for _, fi := range files {
		if fi.IsDir() {
			if err := addDir(base, filepath.Join(dir, fi.Name()), out); err != nil {
				return err
			}
			continue
		}
		if !strings.HasSuffix(fi.Name(), ".go") || strings.HasSuffix(fi.Name(), "_test.go") {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(dir, fi.Name()))
		if err != nil {
			return errors.WithStack(err)
		}
		p[fi.Name()] = b
	}
	if len(p) > 0 {
		outBytes, err := p.ToBytes()
		if err != nil {
			return errors.WithStack(err)
		}
		outDir := filepath.Join(out, relDir)
		if err := os.MkdirAll(outDir, 0777); err != nil {
			return errors.WithStack(err)
		}
		if err := ioutil.WriteFile(filepath.Join(outDir, "source.bin"), outBytes, 0666); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
