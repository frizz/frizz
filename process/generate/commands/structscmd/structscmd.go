package structscmd // import "kego.io/process/generate/commands/structscmd"

import (
	"fmt"
	"os"

	"kego.io/process"
)

func Run() {
	ctx, err := process.InitialiseAutomatic()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.Generate(ctx, process.S_STRUCTS); err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
}
