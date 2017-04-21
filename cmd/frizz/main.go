// frizz is the main cli tool that is used to start the frizz processor / editor.
package main // import "frizz.io/cmd/frizz"

// notest

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"context"

	"frizz.io/context/cmdctx"
	"frizz.io/context/envctx"
	"frizz.io/context/wgctx"
	"frizz.io/editor/server"
	"frizz.io/process"
	"frizz.io/process/validate"
	_ "frizz.io/system"
	"github.com/dave/kerr"
)

func main() {
	// This turns off focus reporting mode that will print “^[[O” and “^[[I”
	// when the terminal window gets / loses focus
	// http://superuser.com/questions/931873/o-and-i-appearing-on-iterm2-when-focus-lost
	fmt.Print("\033[?1004l")

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
		err := process.RunValidateCommand(ctx)
		if err != nil {
			if !cmd.Log {
				// in log mode, we have already written the output of the exec'ed frizz command,
				// so we don't need to duplicate the error message.
				if v, ok := kerr.Source(err).(validate.ValidationCommandError); ok {
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
