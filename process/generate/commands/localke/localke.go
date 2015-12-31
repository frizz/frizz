package localke // import "kego.io/process/generate/commands/localke"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/editor/server"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/process"
)

// Main executes the local "ke" command that is generated in the data directory.
func Main(path string) {
	update := false
	ctx, cancel, err := process.Initialise(&process.FromFlags{
		Update: &update,
		Path:   &path,
	})
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

	changes, err := comparePackageHash(ctx, "kego.io/system")
	if !changes && err == nil {
		changes, err = comparePackageHash(ctx, env.Path)
	}
	if err != nil {
		fmt.Println(process.FormatError(err))
		wgctx.WaitAndExit(ctx, 1)
	}
	if changes {
		fmt.Println("Hash changed")
		wgctx.WaitAndExit(ctx, 1)
	}

	err = process.ValidateCommand(ctx)
	if err != nil {
		fmt.Println(process.FormatError(err))
		wgctx.WaitAndExit(ctx, 1)
	}
	/*
		if typesChanged {
			if cmd.Verbose {
				fmt.Println("Types have changed - rebuilding...")
			}
			if err := process.RunAllCommands(ctx); err != nil {
				if !cmd.Verbose {
					// in verbose mode, we have already written the output of the exec'ed ke command,
					// so we don't need to duplicate the error message.
					fmt.Println(process.FormatError(err))
				}
				wgctx.WaitAndExit(ctx, 1)
			}
			// if the types have changed, RunAllCommands will opened the server etc.
			wgctx.WaitAndExit(ctx, 0)
		}*/
	if cmd.Edit {
		if err = server.Start(ctx); err != nil {
			fmt.Println(process.FormatError(err))
			wgctx.WaitAndExit(ctx, 1)
		}
	}
	wgctx.WaitAndExit(ctx, 0)
}

func comparePackageHash(ctx context.Context, path string) (changes bool, err error) {

	cache := cachectx.FromContext(ctx)
	pcache, ok := cache.Get(path)
	if !ok {
		return false, kerr.New("NHXWLPHCHL", nil, "%s not found in ctx", path)
	}

	for aliasPath, _ := range pcache.Environment.Aliases {
		changes, err := comparePackageHash(ctx, aliasPath)
		if err != nil {
			return false, kerr.New("DGJTLHQOCQ", err, "comparePackageHash")
		}
		if changes {
			return true, nil
		}
	}

	h, ok := json.GetPackageHash(path)
	if !ok {
		return true, nil
	}
	if h != pcache.Environment.Hash {
		return true, nil
	}
	return false, nil
}
