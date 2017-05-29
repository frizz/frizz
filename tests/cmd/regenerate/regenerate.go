package main

import (
	"log"

	"frizz.io/system/generate"
)

func main() {
	if err := generate.Save("frizz.io/tests/unpacker"); err != nil {
		log.Fatal(err)
	}
}
