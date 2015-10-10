package process // import "kego.io/process"

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"bytes"

	"kego.io/kerr"
	"kego.io/process/generate"
	"kego.io/process/settings"
)

type sourceType string

const (
	S_STRUCTS sourceType = "structs"
	S_TYPES              = "types"
	S_EDITOR             = "editor"
)

type commandType string

const (
	C_STRUCTS commandType = "structs"
	C_TYPES               = "types"
	C_KE                  = "ke"
)

func FormatError(err error) string {
	source := kerr.Source(err)
	if m, ok := source.(ValidationError); ok {
		return fmt.Sprint("Error: ", m.Description)
	}
	if t, ok := source.(TypesChangedError); ok {
		return fmt.Sprint("Error: ", t.Description)
	}
	return err.Error()
}

func KeCommand(set *settings.Settings) error {

	for p, a := range set.Aliases {
		if set.Verbose {
			if set.Update {
				fmt.Print("Updating package ", a, "... ")
			} else {
				fmt.Print("Getting package ", a, "... ")
			}
		}
		params := []string{"get"}
		if set.Update {
			params = append(params, "-u")
		}
		if set.Verbose {
			params = append(params, "-v")
		}
		params = append(params, p)

		combined, stdout, stderr := logger(set.Verbose)
		cmd := exec.Command("go", params...)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return kerr.New("HHKSTQMAKG", err, "go get command: %s", combined.String())
		} else {
			if set.Verbose {
				fmt.Println("OK.")
			}
		}
	}

	if err := RunAllCommands(set); err != nil {
		return err
	}
	return nil
}

func RunAllCommands(set *settings.Settings) error {
	if err := RunCommand(C_STRUCTS, set); err != nil {
		return err
	}
	if err := RunCommand(C_TYPES, set); err != nil {
		return err
	}
	if err := RunCommand(C_KE, set); err != nil {
		return err
	}
	return nil
}

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run". When run, this command generates the extra types data in
// the "types" subpackage.
func RunCommand(file commandType, set *settings.Settings) error {

	var source []byte
	var err error
	switch file {
	case C_KE:
		source, err = generate.KeCommand(set)
	case C_STRUCTS:
		source, err = generate.StructsCommand(set)
	case C_TYPES:
		source, err = generate.TypesCommand(set)
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
			return ValidationError{kerr.New("ETWHPXTUVB", nil, errorMessage)}
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

// GenerateFiles generates the source code from templates and writes
// the files to the correct folders.
//
// file == F_MAIN: generated-structs.go in the root of the package.
//
// file == F_TYPES: generated-types.go containing advanced type information
// in the "types" sub package. Note that to generate this file, we need to
// have the main generated-structs.go compiled in, so we generate a temporary
// command and run it with "go run".
//
// file == F_GLOBALS: generated-globals.go in the root of the package.
//
// file == F_CMD_TYPES: this is the temporary command that we create in order to
// generate the types source file
//
// file == F_CMD_KE: this is the temporary command that we create in order
// to run the validation and start the editor
//
func Generate(file sourceType, set *settings.Settings) error {

	if set.Verbose {
		fmt.Print("Generating ", file, "... ")
	}

	if file == S_STRUCTS {
		hasFiles, err := HasSourceFiles(set)
		if err != nil {
			return kerr.New("GXGGDQVHHP", err, "ScanForKegoFiles")
		}
		if !hasFiles {
			return kerr.New("YSLREDFDLJ", nil, "No kego files found")
		}
	}

	if file == S_STRUCTS || file == S_TYPES {

		// We only tolerate unknown types when we're initially building the
		// struct files. At all other times, the generated structs should
		// provide all types.
		ignoreUnknownTypes := file == S_STRUCTS

		// When generating structs or types, we need to scan for types. All other runs will have
		// them compiled in the types sub-package.
		if err := ScanForTypes(ignoreUnknownTypes, set); err != nil {
			return kerr.New("XYIUHERDHE", err, "ScanForTypes")
		}
	}

	var outputDir string
	var filename string
	var source []byte
	var err error

	switch file {
	case S_STRUCTS:
		outputDir = set.Dir
		filename = "generated-structs.go"
		source, err = generate.Structs(set)
	case S_TYPES:
		outputDir = filepath.Join(set.Dir, "types")
		filename = "generated-types.go"
		source, err = generate.Types(set)
	case S_EDITOR:
		outputDir = filepath.Join(set.Dir, "editor")
		filename = "generated-editor.go"
		source, err = generate.Editor(set)
	}
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "generate: %s", file)
	}

	// We only backup in the system structs and types files because they are the only
	// generated files we ever need to roll back
	backup := set.Path == "kego.io/system" && (file == S_STRUCTS || file == S_TYPES)

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "save")
	} else {
		if set.Verbose {
			fmt.Println("OK.")
		}
	}

	return nil
}

func save(dir string, contents []byte, name string, backup bool) error {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.New("BPGOUIYPXO", err, "os.MkdirAll")
		}
	}

	file := filepath.Join(dir, name)

	if backup {
		backupPath := filepath.Join(dir, fmt.Sprintf("%s.backup", name))
		if _, err := os.Stat(backupPath); err == nil {
			os.Remove(backupPath)
		}
		if _, err := os.Stat(file); err == nil {
			os.Rename(file, backupPath)
		}
	} else {
		os.Remove(file)
	}

	if len(contents) == 0 {
		return nil
	}

	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return kerr.New("NWLWHSGJWP", err, "os.OpenFile (could not open output file)")
	}
	defer output.Close()

	if _, err := output.Write(contents); err != nil {
		return kerr.New("FBMGPRWQBL", err, "output.Write")
	}

	if err := output.Sync(); err != nil {
		return kerr.New("EGFNTMNKFX", err, "output.Sync")
	}

	return nil
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
