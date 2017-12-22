package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"fmt"

	"frizz.io/edit"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

// compiles frizz.io/edit/bootstrap to javascript and saves the output to frizz.io/edit/static/data/bootstrap.js

func main() {

	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/assets/data")
	if err != nil {
		log.Fatal(err)
	}

	fpath := filepath.Join(dir, "bootstrap.js")
	fmt.Println("writing " + fpath)

	b, err := edit.Compile(vos.Os(), false, false, true)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fpath, b, 0666); err != nil {
		log.Fatal(err)
	}
}
