package main

import (
	"fmt"
	"os"

	"kego.io/process"
	_ "kego.io/system"
)

func main() {
	dir, update, recursive, verbose, path, aliases, err := process.Initialise()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.KegoCmd(dir, update, recursive, verbose, path, aliases); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
