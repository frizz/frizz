package main

import (
	"fmt"
	"os"

	"kego.io/process"
	_ "kego.io/system"
)

func main() {
	dir, test, recursive, verbose, path, imports, err := process.Initialise()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.KegoCmd(dir, test, recursive, verbose, path, imports); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
