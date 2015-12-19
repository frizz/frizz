// ke is the main cli tool that is used to start the ke processor / editor.
package main // import "kego.io/cmd/ke"

import (
	"fmt"
	"os"

	"kego.io/context/cmdctx"
	"kego.io/process"
	_ "kego.io/system"
)

func main() {
	ctx, err := process.Initialise(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		fmt.Println("No cmd in ctx")
		os.Exit(1)
	}

	if err := process.KeCommand(ctx); err != nil {
		if !cmd.Verbose {
			// in verbose mode, we have already written the output of the exec'ed ke command,
			// so we don't need to duplicate the error message.
			fmt.Println(process.FormatError(err))
		}
		os.Exit(1)
	}
	os.Exit(0)
}
