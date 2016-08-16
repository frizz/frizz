package process

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"strings"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/process/generate"
	"kego.io/process/logger"
	"kego.io/process/validate"
)

const validateCommand = ".localke/validate"

type HashError struct {
	kerr.Struct
}

func RunValidateCommand(ctx context.Context) (err error) {
	env := envctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	build := false
	repeat := true
	if _, err := os.Stat(validateCommandPath); os.IsNotExist(err) {
		build = true
		repeat = false
	}

	if err = runValidateCommand(ctx, build, repeat); err != nil {
		return kerr.Wrap("KFNIOHWCBT", err)
	}
	return nil
}

func runValidateCommand(ctx context.Context, build bool, repeat bool) (err error) {

	wgctx.Add(ctx, "runValidateCommand")
	defer wgctx.Done(ctx, "runValidateCommand")

	if build {
		if err := buildValidateCommand(ctx); err != nil {
			return kerr.Wrap("FIJHGMEEUM", err)
		}
	}

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	cmd.Print("Running validate command... ")

	combined, stdout, stderr := logger.Logger(cmd.Log)

	hashChanged := false

	exe := exec.Command(validateCommandPath)
	exe.Stdout = stdout
	exe.Stderr = stderr
	if err = exe.Run(); err != nil {

		exiterr, ok := err.(*exec.ExitError)
		if !ok {
			goto Repeat
		}

		// The program has exited with an exit code != 0. This works on both Unix and Windows.
		// Although package syscall is generally platform dependent, WaitStatus is defined for
		// both Unix and Windows and in both cases has an ExitStatus() method with the same
		// signature.
		status, ok := exiterr.Sys().(syscall.WaitStatus)
		if !ok {
			// ke: {"block": {"notest": true}}
			goto Repeat
		}

		switch status.ExitStatus() {
		case 3:
			// Exit status 3 = hash changed
			hashChanged = true
			goto Repeat
		case 4:
			// Exit status 4 = validation error
			return validate.ValidationCommandError{Struct: kerr.New("ETWHPXTUVB", strings.TrimSpace(combined.String()))}
		default:
			// ke: {"block": {"notest": true}}
			goto Repeat
		}
	}
	cmd.Println("OK.")
	return nil

Repeat:
	if repeat {
		if hashChanged {
			cmd.Println("Types have changed since last run. Rebuilding...")
		} else {
			// ke: {"block": {"notest": true}}
			cmd.Println("Command returned an error. Rebuilding...")
		}
		if err := runValidateCommand(ctx, true, false); err != nil {
			return kerr.Wrap("HOHQEISLMI", err)
		}
		return nil
	}
	return kerr.Wrap("DTTHRRJSSF", err)
}

// buildValidateCommand creates a temporary folder in the package, in which the go source for the
// local command is generated. This command is then compiled.
func buildValidateCommand(ctx context.Context) error {

	wgctx.Add(ctx, "buildValidateCommand")
	defer wgctx.Done(ctx, "buildValidateCommand")

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	source, err := generate.ValidateCommand(ctx)
	if err != nil {
		return kerr.Wrap("SPRFABSRWK", err)
	}

	outputDir, err := ioutil.TempDir(env.Dir, "temporary")
	if err != nil {
		return kerr.Wrap("HWOPVXYMCT", err)
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.Wrap("FRLCYFOWCJ", err)
	}

	cmd.Print("Building validate command... ")

	combined, stdout, stderr := logger.Logger(cmd.Log)
	exe := exec.Command("go", "build", "-o", validateCommandPath, outputPath)
	exe.Stdout = stdout
	exe.Stderr = stderr

	if err := exe.Run(); err != nil {
		return kerr.Wrap("OEPAEEYKIS", err)
	}
	cmd.Println("OK.")
	cmd.Print(combined.String())

	return nil
}
