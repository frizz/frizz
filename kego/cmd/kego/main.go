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
	if err := process.KegoCmd(set); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
