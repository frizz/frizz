package process

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process/generate"
	"kego.io/process/validate"
)

const validateCommand = ".localke/validate"

func RunValidateCommand(ctx context.Context) (success bool, err error) {
	env := envctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	if _, err := os.Stat(validateCommandPath); err != nil {
		return false, nil
	}

	return runValidateCommand(ctx, true)
}

func runValidateCommand(ctx context.Context, repeatOnFail bool) (success bool, err error) {

	wgctx.Add(ctx, "runValidateCommand")
	defer wgctx.Done(ctx, "runValidateCommand")

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	if cmd.Log {
		fmt.Print("Running validate command... ")
	}

	params := []string{}
	if cmd.Log {
		// we don't log
		//params = append(params, "-l")
	}
	combined, stdout, stderr := logger(cmd.Log)

	exe := exec.Command(validateCommandPath, params...)
	exe.Stdout = stdout
	exe.Stderr = stderr
	if err := exe.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0. This works on both Unix and Windows.
			// Although package syscall is generally platform dependent, WaitStatus is defined for
			// both Unix and Windows and in both cases has an ExitStatus() method with the same
			// signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				switch status.ExitStatus() {
				case 3:
					// Exit status 3 = hash changed
					if cmd.Log {
						fmt.Println("Command is out of date.")
					}
					return false, nil
				case 4:
					// Exit status 4 = validation error
					return true, validate.ValidationError{Struct: kerr.New("ETWHPXTUVB", nil, strings.TrimSpace(combined.String()))}
				default:
					if repeatOnFail {
						if cmd.Log {
							fmt.Println("Validate command returned error, rebuilding...")
						}
						if err := BuildAndRunLocalCommand(ctx); err != nil {
							return true, kerr.New("HOHQEISLMI", err, "BuildAndRunLocalCommand (retry)")
						}
						return true, nil
					}
					return true, kerr.New("DTTHRRJSSF", err, "Run")
				}
			}
		} else {
			return true, kerr.New("GFTBBSYEXU", err, "Run")
		}
	}
	if cmd.Log {
		fmt.Println("OK.")
	}
	return true, nil
}

// BuildAndRunLocalCommand creates a temporary folder in the package, in which the go source for the
// local command is generated. This command is then compiled before executing the binary.
func BuildAndRunLocalCommand(ctx context.Context) error {

	wgctx.Add(ctx, "BuildAndRunLocalCommand")
	defer wgctx.Done(ctx, "BuildAndRunLocalCommand")

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	validateCommandPath := filepath.Join(env.Dir, validateCommand)

	source, err := generate.ValidateCommand(ctx)
	if err != nil {
		return kerr.New("SPRFABSRWK", err, "generate.ValidateCommand")
	}

	outputDir, err := ioutil.TempDir(env.Dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "save")
	}

	if cmd.Log {
		fmt.Print("Building validate command... ")
	}

	combined, stdout, stderr := logger(cmd.Log)
	exe := exec.Command("go", "build", "-o", validateCommandPath, outputPath)
	exe.Stdout = stdout
	exe.Stderr = stderr

	if err := exe.Run(); err != nil {
		return kerr.New("OEPAEEYKIS", err, "go build: %s", combined.String())
	}
	if cmd.Log {
		fmt.Println("OK.")
	}

	if cmd.Log {
		fmt.Print(combined.String())
	}

	success, err := runValidateCommand(ctx, false)
	if err != nil {
		return kerr.New("PPBPPXQMWV", err, "runValidateCommand")
	}
	if !success {
		return kerr.New("VGVLPAJAQN", nil, "runLocalCommand returned hash changed after build")
	}

	return nil
}

func logger(log bool) (combined *bytes.Buffer, stdout io.Writer, stderr io.Writer) {
	combined = &bytes.Buffer{}
	if log {
		stderr = MultiWriter(os.Stderr, combined)
		stdout = MultiWriter(os.Stdout, combined)
	} else {
		stderr = combined
		stdout = combined
	}
	return
}

// MultiWriter creates a writer that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
func MultiWriter(primary io.Writer, writers ...io.Writer) io.Writer {
	w := make([]io.Writer, len(writers))
	copy(w, writers)
	return &multiWriter{primary: primary, writers: w}
}

type multiWriter struct {
	primary io.Writer
	writers []io.Writer
}

func (t *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		w.Write(p)
	}
	return t.primary.Write(p)
}
