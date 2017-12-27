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

// compiles frizz.io/edit/bootstrap to javascript and saves the output to frizz.io/edit/assets/static/bootstrap.js

func main() {
	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/assets/assets/static")
	if err != nil {
		log.Fatal(err)
	}
	if err := bootstrap(dir); err != nil {
		log.Fatal(err)
	}
}

func bootstrap(dir string) error {
	fpath := filepath.Join(dir, "bootstrap.js")
	fmt.Println("writing " + fpath)
	b, err := edit.Compile(vos.Os(), false, false, true)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fpath, b, 0666); err != nil {
		return err
	}
	return nil
}
