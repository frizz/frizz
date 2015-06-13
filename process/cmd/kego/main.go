package main

import (
	"kego.io/process"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	process.GenerateFiles(process.F_MAIN)
	process.GenerateAndRunCmd()
}
