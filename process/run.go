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

	"kego.io/kerr"
	"kego.io/process/generate"
	"kego.io/process/settings"
	"kego.io/process/validate"
)

type commandType string

const (
	C_STRUCTS commandType = "structs"
	C_TYPES               = "types"
	C_KE                  = "ke"
)

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run", or in the case of the ke command, we "go build" before
// executing the binary.
func Run(file commandType, set *settings.Settings) error {

	var source []byte
	var err error
	switch file {
	case C_STRUCTS:
		source, err = generate.StructsCommand(set)
	case C_TYPES:
		source, err = generate.TypesCommand(set)
	case C_KE:
		source, err = generate.KeCommand(set)
	}
	if err != nil {
		return kerr.New("SPRFABSRWK", err, "generate command: %s", file)
	}

	outputDir, err := ioutil.TempDir(set.Dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	keCommandPath := filepath.Join(set.Dir, "ke")

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "save")
	}

	if file == C_KE {
		if set.Verbose {
			fmt.Print("Building ", file, " command... ")
		}

		combined, stdout, stderr := logger(set.Verbose)
		cmd := exec.Command("go", "build", "-o", keCommandPath, outputPath)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return kerr.New("OEPAEEYKIS", err, "go build: %s", combined.String())
		}
		if set.Verbose {
			fmt.Println("OK.")
		}

		if set.Verbose {
			fmt.Print(combined.String())
		}
	}

	command := ""
	params := []string{}

	if set.Verbose {
		fmt.Println("Running", file, "command...")
	}
	if file == C_KE {
		command = keCommandPath
		if set.Verbose {
			params = append(params, "-v")
		}
		if set.Edit {
			params = append(params, "-e")
		}
	} else {
		command = "go"
		params = []string{"run", outputPath}

		if set.Update {
			params = append(params, "-u")
		}
		if set.Recursive {
			params = append(params, "-r")
		}
		if set.Verbose {
			params = append(params, "-v")
		}
		if set.Edit {
			params = append(params, "-e")
		}

		params = append(params, set.Path)
	}

	combined, stdout, stderr := logger(set.Verbose)

	cmd := exec.Command(command, params...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		if file == C_KE {
			errorMessage := strings.TrimSpace(combined.String())
			if strings.HasPrefix(errorMessage, "Error: ") {
				errorMessage = errorMessage[7:]
			}
			return validate.ValidationError{kerr.New("ETWHPXTUVB", nil, errorMessage)}
		}
		return kerr.New("UDDSSMQRHA", err, "cmd.Run: %s", combined.String())
	}

	return nil
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
