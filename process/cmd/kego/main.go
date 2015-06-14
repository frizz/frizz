package main

import (
	"fmt"
	"os"

	"kego.io/process"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := process.GenerateFiles(process.F_MAIN); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.GenerateAndRunCmd(process.F_CMD_TYPES); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.GenerateAndRunCmd(process.F_CMD_VALIDATE); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("OK")
	os.Exit(0)
}
