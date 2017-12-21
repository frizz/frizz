package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"frizz.io/edit"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

// compiles frizz.io/edit/bootstrap to javascript and saves the output to frizz.io/edit/static/data/bootstrap.js

func main() {
	b, err := edit.Compile(vos.Os(), false, false, true)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/static/data")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "bootstrap.js"), b, 0666); err != nil {
		log.Fatal(err)
	}
}
