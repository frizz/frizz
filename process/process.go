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

func KeCommand(set settings) error {

	for p, a := range set.aliases {
		if set.verbose {
			if set.update {
				fmt.Print("Updating package ", a, "... ")
			} else {
				fmt.Print("Getting package ", a, "... ")
			}
		}
		params := []string{"get"}
		if set.update {
			params = append(params, "-u")
		}
		if set.verbose {
			params = append(params, "-v")
		}
		params = append(params, p)
		if out, err := exec.Command("go", params...).CombinedOutput(); err != nil {
			return kerr.New("HHKSTQMAKG", err, "go get command: %s", out)
		} else {
			if set.verbose {
				fmt.Println("OK.")
			}
		}
	}

	if err := RunAllCommands(set); err != nil {
		return err
	}
	return nil
}

func RunAllCommands(set settings) error {
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
func RunCommand(file commandType, set settings) error {

	source, err := GenerateCommand(file, set)
	if err != nil {
		return kerr.New("SPRFABSRWK", err, "Generate")
	}

	outputDir, err := ioutil.TempDir(set.dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	keCommandPath := filepath.Join(set.dir, "ke")

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "save")
	}

	if file == C_KE {
		if set.verbose {
			fmt.Print("Building ", file, " command... ")
		}
		out, err := exec.Command("go", "build", "-o", keCommandPath, outputPath).CombinedOutput()
		if err != nil {
			return kerr.New("OEPAEEYKIS", err, "go build: %s", out)
		} else {
			if set.verbose {
				fmt.Println("OK.")
			}
		}
		if set.verbose {
			fmt.Print(string(out))
		}
	}

	command := ""
	params := []string{}

	if set.verbose {
		fmt.Println("Running", file, "command...")
	}
	if file == C_KE {
		command = keCommandPath
		if set.verbose {
			params = append(params, "-v")
		}
		if set.edit {
			params = append(params, "-e")
		}
	} else {
		command = "go"
		params = []string{"run", outputPath}

		if set.update {
			params = append(params, "-u")
		}
		if set.recursive {
			params = append(params, "-r")
		}
		if set.verbose {
			params = append(params, "-v")
		}
		if set.edit {
			params = append(params, "-e")
		}

		params = append(params, set.path)
	}

	out := bytes.Buffer{}
	var wErr io.Writer
	var wOut io.Writer
	if set.verbose {
		wErr = MultiWriter(os.Stderr, &out)
		wOut = MultiWriter(os.Stdout, &out)
	} else {
		wErr = &out
		wOut = &out
	}

	cmd := exec.Command(command, params...)
	cmd.Stdout = wOut
	cmd.Stderr = wErr
	if err := cmd.Run(); err != nil {
		if file == C_KE {
			errorMessage := strings.TrimSpace(out.String())
			if strings.HasPrefix(errorMessage, "Error: ") {
				errorMessage = errorMessage[7:]
			}
			return ValidationError{kerr.New("ETWHPXTUVB", nil, errorMessage)}
		}
		return kerr.New("UDDSSMQRHA", err, "cmd.Run: %s", out.String())
	}

	return nil
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
func Generate(file sourceType, set settings) error {

	if set.verbose {
		fmt.Print("Generating ", file, "... ")
	}

	outputDir := set.dir
	if file == S_TYPES {
		outputDir = filepath.Join(set.dir, "types")
	} else if file == S_EDITOR {
		outputDir = filepath.Join(set.dir, "editor")
	}

	// We only tolerate unknown types when we're initially building the struct files. At all other
	// times, the generated structs should provide all types.
	ignoreUnknownTypes := file == S_STRUCTS

	if file == S_STRUCTS {
		hasFiles, err := ScanForKegoFiles(set)
		if err != nil {
			return kerr.New("GXGGDQVHHP", err, "ScanForKegoFiles")
		}
		if !hasFiles {
			return kerr.New("YSLREDFDLJ", nil, "No kego files found")
		}
	}

	if file == S_STRUCTS || file == S_TYPES {
		// When generating structs or types, we need to scan for types. All other runs will have
		// them compiled in the types sub-package.
		if err := ScanForTypes(ignoreUnknownTypes, set); err != nil {
			return kerr.New("XYIUHERDHE", err, "ScanForTypes")
		}
	}

	source, err := GenerateSource(file, set)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "GenerateSource")
	}

	// We only backup in the system structs and types files because they are the only
	// generated files we ever need to roll back
	backup := set.path == "kego.io/system" && (file == S_STRUCTS || file == S_TYPES)

	// Filenames:
	// generated-globals.go
	// generated-structs.go
	// generated-types.go
	// generated-editor.go
	filename := fmt.Sprintf("generated-%s.go", file)

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "save")
	} else {
		if set.verbose {
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
