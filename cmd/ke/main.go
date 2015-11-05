// ke is the main cli tool that is used to start the ke processor / editor.
package main // import "kego.io/cmd/ke"

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
	if err := process.KeCommand(set); err != nil {
		if !set.Verbose {
			// in verbose mode, we have already written the output of the exec'ed ke command,
			// so we don't need to duplicate the error message.
			fmt.Println(process.FormatError(err))
		}
		os.Exit(1)
	}
	os.Exit(0)
}
