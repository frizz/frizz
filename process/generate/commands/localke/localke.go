package localke // import "kego.io/process/generate/commands/localke"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"kego.io/context/cmdctx"
	"kego.io/editor/server"
	"kego.io/process"
)

// Run executes the local "ke" command that is generated in the data directory.
func Main(recursive bool, path string) {
	update := false
	ctx, cancel, err := process.Initialise(&process.FromFlags{
		Update:    &update,
		Recursive: &recursive,
		Path:      &path,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		fmt.Println("No cmd in ctx")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		if cmd.Verbose {
			fmt.Println("Received OS interrupt - exiting.")
		}
		cancel()
	}()

	typesChanged, err := process.ValidateCommand(ctx)
	if err != nil {
		if !typesChanged || !cmd.Verbose {
			// when ValidateCommand detects the types have changed, it
			// spawns a ke command. In verbose mode this will output any
			// error so we don't need to print the error
			fmt.Println(process.FormatError(err))
		}
		os.Exit(1)
	}
	// if the types have changed, we have spawned a sub process to rebuild them, and this will have
	// opened the server etc.
	if typesChanged {
		os.Exit(0)
	}
	if cmd.Edit {
		if err = server.Start(ctx); err != nil {
			fmt.Println(process.FormatError(err))
			os.Exit(1)
		}
	}
}
