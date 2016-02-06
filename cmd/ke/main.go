// ke is the main cli tool that is used to start the ke processor / editor.
package main // import "kego.io/cmd/ke"

// ke: {"notest":true}

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/editor/server"
	"kego.io/kerr"
	"kego.io/process"
	"kego.io/process/validate"
	_ "kego.io/system"
)

func main() {
	ctx, cancel, err := process.Initialise(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()

	if err := process.GenerateAll(ctx, env.Path, map[string]bool{}); err != nil {
		fmt.Println(err.Error())
		wgctx.WaitAndExit(ctx, 1)
	}

	if cmd.Validate {
		success, err := process.RunValidateCommand(ctx)
		if !success && err == nil {
			err = process.BuildAndRunLocalCommand(ctx)
		}
		if err != nil {
			if !cmd.Log {
				// in log mode, we have already written the output of the exec'ed ke command,
				// so we don't need to duplicate the error message.
				if v, ok := kerr.Source(err).(validate.ValidationError); ok {
					fmt.Println(v.Description)
					wgctx.WaitAndExit(ctx, 4) // Exit code 4: validation error
				}
				fmt.Println(err.Error())
			}
			wgctx.WaitAndExit(ctx, 1)
		}
	}

	if cmd.Edit {
		if err = server.Start(ctx, cancel); err != nil {
			fmt.Println(err.Error())
			wgctx.WaitAndExit(ctx, 1)
		}
	}

	wgctx.WaitAndExit(ctx, 0)

}
