package validate // import "kego.io/process/generate/commands/validate"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process"
	"kego.io/process/validate"
)

// Main executes the validate command that is generated in the data directory.
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
		fmt.Println(err.Error())
		wgctx.WaitAndExit(ctx, 1) // Exit status 1: generic error
	}
	if changes {
		wgctx.WaitAndExit(ctx, 3) // Exit status 3: hash changed error
	}

	if err := validate.ValidatePackage(ctx); err != nil {
		if v, ok := kerr.Source(err).(validate.ValidationError); ok {
			fmt.Println(v.Description)
			wgctx.WaitAndExit(ctx, 4) // Exit status 4: validation error
		}
		fmt.Println(err.Error())
		wgctx.WaitAndExit(ctx, 1) // Exit status 1: generic error
	}
	wgctx.WaitAndExit(ctx, 0) // Exit status 0: success
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

	jcache := jsonctx.FromContext(ctx)

	h, ok := jcache.Packages.Get(path)
	if !ok {
		return true, nil
	}
	if h.Hash != pcache.Environment.Hash {
		return true, nil
	}
	return false, nil
}
