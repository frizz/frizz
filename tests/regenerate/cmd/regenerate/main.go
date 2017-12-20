package main

import (
	"log"

	"os"

	"frizz.io/tests/regenerate"
	"github.com/dave/patsy/vos"
)

func main() {
	env := vos.Os()
	hashes, err := regenerate.Regenerate(env, os.Stdout, true)
	if err != nil {
		log.Fatal(err)
	}
	if err := regenerate.SaveHashes(env, hashes); err != nil {
		log.Fatal(err)
	}
}
