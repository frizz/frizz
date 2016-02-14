package command // import "kego.io/process/validate/command"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process"
	"kego.io/process/validate"
)

// ValidateMain is called by the generated validate command.
// ke: {"func": {"notest": true}}
func ValidateMain(path string) {

	// Using process.Flags as the options means that the non-specified options are read from the
	// command flags.
	update := false
	options := &process.Flags{
		Update: &update,
		Path:   &path,
	}

	ctx, cancel, err := process.Initialise(context.Background(), options)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// The log function outputs errors and messages to the console. We pass it in to permit testing.
	log := func(message string) {
		fmt.Println(message)
	}

	// The interrupt chanel signals that Ctrl+C or other OS interrupts have happened.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, syscall.SIGTERM)

	exitStatus := validateMain(ctx, cancel, log, interrupt)

	wgctx.WaitAndExit(ctx, exitStatus)

}
func validateMain(ctx context.Context, cancel context.CancelFunc, log func(string), interrupt chan os.Signal) int {

	go func() {
		<-interrupt
		cancel()
	}()

	env := envctx.FromContext(ctx)

	changes, err := comparePackageHash(ctx, "kego.io/system")
	if !changes && err == nil {
		changes, err = comparePackageHash(ctx, env.Path)
	}
	if err != nil {
		log(err.Error())
		return 1 // Exit status 1: generic error
	}
	if changes {
		return 3 // Exit status 3: hash changed error
	}

	if err := validate.ValidatePackage(ctx); err != nil {
		if v, ok := kerr.Source(err).(validate.ValidationError); ok {
			log(v.Description)
			return 4 // Exit status 4: validation error
		}
		log(err.Error())
		return 1 // Exit status 1: generic error
	}
	return 0 // Exit status 0: success
}

func comparePackageHash(ctx context.Context, path string) (changes bool, err error) {

	scache := sysctx.FromContext(ctx)
	spi, ok := scache.Get(path)
	if !ok {
		return false, kerr.New("NHXWLPHCHL", "%s not found in sys ctx", path)
	}

	for aliasPath, _ := range spi.Aliases {
		changes, err := comparePackageHash(ctx, aliasPath)
		if err != nil {
			return false, kerr.Wrap("DGJTLHQOCQ", err)
		}
		if changes {
			return true, nil
		}
	}

	jcache := jsonctx.FromContext(ctx)
	jpi, ok := jcache.Packages.Get(path)
	if !ok {
		return true, nil
	}

	// pcache.Environment.Hash is computed after parsing all the data files.
	// h.Hash is in generated.go (jsonctx.InitPackage), and correct when the types were generated.
	if jpi.Hash != spi.Hash {
		return true, nil
	}
	return false, nil
}
