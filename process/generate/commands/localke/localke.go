package localke // import "kego.io/process/generate/commands/localke"
import (
	"fmt"
	"os"

	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/editor/server"
	"kego.io/process"
)

// Run executes the local "ke" command that is generated in the data directory.
func Main(recursive bool, path string) {
	update := false
	ctx, err := process.InitialiseCommand(update, recursive, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	env, ok := envctx.FromContext(ctx)
	if !ok {
		fmt.Println("No env in ctx")
		os.Exit(1)
	}
	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		fmt.Println("No cmd in ctx")
		os.Exit(1)
	}
	if typesChanged, err := process.ValidateCommand(ctx); err != nil {
		if !typesChanged || !cmd.Verbose {
			// when ValidateCommand detects the types have changed, it
			// spawns a ke command. In verbose mode this will output any
			// error so we don't need to print the error
			fmt.Println(process.FormatError(err))
		}
		os.Exit(1)
	}
	if cmd.Edit {
		if err = server.Start(env.Path, cmd.Verbose, false); err != nil {
			fmt.Println(process.FormatError(err))
			os.Exit(1)
		}
	}
}
