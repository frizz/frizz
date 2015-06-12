package main

import (
	"kego.io/process"

	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	process.GeneratorInit()
	process.GenerateFiles(process.F_TYPES)
}
