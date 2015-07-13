package main

import (
	"fmt"
	"os"

	"kego.io/process"
	_ "kego.io/system"
)

func main() {
	set, err := process.InitialiseAutomatic()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.Generate(process.S_TYPES, set); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if set.Globals() {
		if err := process.Generate(process.S_GLOBALS, set); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
