package main

import (
	"fmt"
	"os"

	"kego.io/process"
	_ "kego.io/system"
	//_ "kego.io/system/types"
)

func main() {
	dir, test, recursive, verbose, path, imports, err := process.Initialise()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.GenerateAndRunCmd(process.F_CMD_MAIN, dir, test, recursive, verbose, path, imports); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.GenerateAndRunCmd(process.F_CMD_TYPES, dir, test, recursive, verbose, path, imports); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.GenerateAndRunCmd(process.F_CMD_VALIDATE, dir, test, recursive, verbose, path, imports); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
