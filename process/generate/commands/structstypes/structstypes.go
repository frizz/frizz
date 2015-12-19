package structstypes // import "kego.io/process/generate/commands/structstypes"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"kego.io/process"
)

// Run executes the structs or types generation process. This binary is temporary and
// is deleted after it has run.
func Main(sourceType process.SourceType) {
	ctx, cancel, err := process.Initialise(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()

	if err := process.Generate(ctx, sourceType); err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
}
