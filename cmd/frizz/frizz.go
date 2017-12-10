package main

import (
	"flag"

	"log"

	"frizz.io/generate"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	packagePath := flag.Arg(0)
	if packagePath == "" {
		log.Fatal(errors.New("must specify package path"))
	}
	if err := generate.Save(packagePath); err != nil {
		log.Fatal(err)
	}
}
