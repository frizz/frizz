package main

import (
	"log"

	"frizz.io/tests/regenerate"

	"github.com/dave/patsy/vos"
)

func main() {
	env := vos.Os()

	hashes, err := regenerate.Regenerate(env, true)
	if err != nil {
		log.Fatal(err)
	}

	if err := regenerate.SaveHashes(env, hashes); err != nil {
		log.Fatal(err)
	}

}
