package process

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process/generate"
	"kego.io/process/validate"
)

type commandType string

const (
	C_STRUCTS  commandType = "structs"
	C_TYPES                = "types"
	C_LOCAL_KE             = "ke"
)

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run", or in the case of the ke command, we "go build" before
// executing the binary.
func Run(ctx context.Context, file commandType) error {

	wgctx.FromContext(ctx).Add(1)
	defer wgctx.FromContext(ctx).Done()

	cmd := cmdctx.FromContext(ctx)
	env := envctx.FromContext(ctx)

	var source []byte
	var err error
	switch file {
	case C_STRUCTS:
		source, err = generate.StructsCommand(ctx)
	case C_TYPES:
		source, err = generate.TypesCommand(ctx)
	case C_LOCAL_KE:
		source, err = generate.LocalKeCommand(ctx)
	}
	if err != nil {
		return kerr.New("SPRFABSRWK", err, "generate command: %s", file)
	}

	outputDir, err := ioutil.TempDir(cmd.Dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	keCommandPath := filepath.Join(cmd.Dir, "ke")

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "save")
	}

	if file == C_LOCAL_KE {
		if cmd.Verbose {
			fmt.Print("Building ", file, " command... ")
		}

		combined, stdout, stderr := logger(cmd.Verbose)
		exe := exec.Command("go", "build", "-o", keCommandPath, outputPath)
		exe.Stdout = stdout
		exe.Stderr = stderr

		if err := exe.Run(); err != nil {
			return kerr.New("OEPAEEYKIS", err, "go build: %s", combined.String())
		}
		if cmd.Verbose {
			fmt.Println("OK.")
		}

		if cmd.Verbose {
			fmt.Print(combined.String())
		}
	}

	command := ""
	params := []string{}

	if cmd.Verbose {
		fmt.Println("Running", file, "command...")
	}
	if file == C_LOCAL_KE {
		command = keCommandPath
		if cmd.Verbose {
			params = append(params, "-v")
		}
		if cmd.Edit {
			params = append(params, "-e")
		}
	} else {
		command = "go"
		params = []string{"run", outputPath}

		if cmd.Update {
			params = append(params, "-u")
		}
		if cmd.Verbose {
			params = append(params, "-v")
		}
		if cmd.Edit {
			params = append(params, "-e")
		}

		params = append(params, env.Path)
	}

	combined, stdout, stderr := logger(cmd.Verbose)

	exe := exec.Command(command, params...)
	exe.Stdout = stdout
	exe.Stderr = stderr
	if err := exe.Run(); err != nil {
		if file == C_LOCAL_KE {
			errorMessage := strings.TrimSpace(combined.String())
			if strings.HasPrefix(errorMessage, "Error: ") {
				errorMessage = errorMessage[7:]
			}
			return validate.ValidationError{kerr.New("ETWHPXTUVB", nil, errorMessage)}
		}
		return CommandError{kerr.New("UDDSSMQRHA", nil, strings.TrimSpace(combined.String()))}
	}

	return nil
}

type CommandError struct {
	kerr.Struct
}

func logger(verbose bool) (combined *bytes.Buffer, stdout io.Writer, stderr io.Writer) {
	combined = &bytes.Buffer{}
	if verbose {
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
