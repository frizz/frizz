package main

import (
	"flag"

	"log"

	"frizz.io/system/generate"
)

func main() {
	packagePath := flag.Arg(0)
	if err := generate.Save(packagePath); err != nil {
		log.Fatal(err)
	}
}
