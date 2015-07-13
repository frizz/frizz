package process // import "kego.io/process"

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"kego.io/kerr"
)

type sourceType string

const (
	S_MAIN    sourceType = "main"
	S_TYPES              = "types"
	S_GLOBALS            = "globals"
)

type commandType string

const (
	C_MAIN     commandType = "main"
	C_TYPES                = "types"
	C_VALIDATE             = "validate"
)

func KegoCmd(set settings) error {

	for p, a := range set.aliases {
		if set.verbose {
			if set.update {
				fmt.Print("Updating package ", a, "...")
			} else {
				fmt.Print("Getting package ", a, "...")
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
			return kerr.New("HHKSTQMAKG", err, "process.KegoCmd", "go get command: %s", out)
		} else {
			if set.verbose {
				fmt.Println(" OK.")
			}
		}
	}

	if err := RunCommand(C_MAIN, set); err != nil {
		return err
	}
	if err := RunCommand(C_TYPES, set); err != nil {
		return err
	}
	if err := RunCommand(C_VALIDATE, set); err != nil {
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
		return kerr.New("SPRFABSRWK", err, fmt.Sprintf("process.RunCommand %s", file), "Generate")
	}

	outputDir, err := ioutil.TempDir(set.dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, fmt.Sprintf("process.RunCommand %s", file), "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	validateCommandPath := filepath.Join(set.dir, "validate")

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, fmt.Sprintf("process.RunCommand %s", file), "save")
	}

	if file == C_VALIDATE {
		if set.verbose {
			fmt.Print("Building ", file, "...")
		}
		out, err := exec.Command("go", "build", "-o", validateCommandPath, outputPath).CombinedOutput()
		if err != nil {
			return kerr.New("OEPAEEYKIS", err, fmt.Sprintf("process.RunCommand %s", file), "go build: cd%s", out)
		} else {
			if set.verbose {
				fmt.Println(" OK.")
			}
		}
		if set.verbose {
			fmt.Print(string(out))
		}
	}

	if set.verbose {
		fmt.Print("Running ", file, " command...")
	}

	command := ""
	params := []string{}

	if file == C_VALIDATE {
		command = validateCommandPath
	} else {
		command = "go"
		params = []string{"run", outputPath}
	}

	params = append(params, fmt.Sprintf("-p=%s", set.path))

	if set.update {
		params = append(params, "-u")
	}
	if set.recursive {
		params = append(params, "-r")
	}
	if set.verbose {
		params = append(params, "-v")
	}
	if set.globals {
		params = append(params, "-g")
	}
	out, err := exec.Command(command, params...).CombinedOutput()
	if err != nil {
		return kerr.New("UDDSSMQRHA", err, fmt.Sprintf("process.RunCommand %s", file), "cmd.Run: %s", out)
	} else {
		if set.verbose {
			fmt.Println(" OK.")
		}
	}

	if set.verbose {
		fmt.Print(string(out))
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
// file == F_CMD_VALIDATE: this is the temporary command that we create in order
// to run the validation
//
func Generate(file sourceType, set settings) error {

	if set.verbose {
		fmt.Print("Generating ", file, "...")
	}

	outputDir := set.dir
	if file == S_TYPES {
		outputDir = filepath.Join(set.dir, "types")
	}

	ignoreUnknownTypes := true
	if file == S_TYPES {
		ignoreUnknownTypes = false
	}

	if file == S_MAIN || file == S_TYPES {
		// If type == F_GLOBALS, we have already generated and imported the types, so
		// there is no need to scan.
		if err := ScanForTypes(ignoreUnknownTypes, set); err != nil {
			return kerr.New("XYIUHERDHE", err, "process.Generate", "ScanForTypes")
		}
	} else {
		// However, we need to scan for the globals.
		if err := ScanForGlobals(set); err != nil {
			return kerr.New("JQLAQVKLAN", err, "process.Generate", "ScanForGlobals")
		}
	}

	source, err := GenerateSource(file, set)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "process.Generate", "Generate")
	}

	// We only backup in the system types because they are the only
	// generated files you'll ever need to roll back
	backup := set.path == "kego.io/system" || set.path == "kego.io/system/types"

	var filename string
	if file == S_GLOBALS {
		filename = "generated-globals.go"
	} else if file == S_MAIN {
		filename = "generated-structs.go"
	} else if file == S_TYPES {
		filename = "generated-types.go"
	}

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "process.Generate", "save")
	} else {
		if set.verbose {
			fmt.Println(" OK.")
		}
	}

	return nil
}

func save(dir string, contents []byte, name string, backup bool) error {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.New("BPGOUIYPXO", err, "process.save", "os.MkdirAll")
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
		return kerr.New("NWLWHSGJWP", err, "process.save", "os.OpenFile (could not open output file)")
	}
	defer output.Close()

	if _, err := output.Write(contents); err != nil {
		return kerr.New("FBMGPRWQBL", err, "process.save", "output.Write")
	}

	if err := output.Sync(); err != nil {
		return kerr.New("EGFNTMNKFX", err, "process.save", "output.Sync")
	}

	return nil
}
