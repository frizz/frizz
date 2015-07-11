package process

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"strconv"

	"kego.io/kerr"
)

type fileType string

const (
	F_MAIN         fileType = "main"
	F_TYPES                 = "types"
	F_GLOBALS               = "globals"
	F_CMD_MAIN              = "cmd_main"
	F_CMD_TYPES             = "cmd_types"
	F_CMD_VALIDATE          = "cmd_validate"
)

var generatorUpdateFlag = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
var generatorPathFlag = flag.String("p", "", "Package: full package path e.g. github.com/foo/bar")
var generatorRecursiveFlag = flag.Bool("r", false, "Recursive: scan subdirectories for objects")
var generatorVerboseFlag = flag.Bool("v", false, "Verbose")

func Initialise() (dir string, update bool, recursive bool, verbose bool, path string, aliases map[string]string, err error) {

	if !flag.Parsed() {
		flag.Parse()
	}

	update = *generatorUpdateFlag
	path = *generatorPathFlag
	recursive = *generatorRecursiveFlag
	verbose = *generatorVerboseFlag

	dir, err = os.Getwd()
	if err != nil {
		err = kerr.New("OKOLXAMBSJ", err, "process.Initialise", "os.Getwd")
		return
	}

	if path == "" {
		path, err = getPackage(dir, os.Getenv("GOPATH"))
		if err != nil {
			err = kerr.New("PSRAWHQCPV", err, "process.Initialise", "getPackage")
			return
		}
	}

	aliases, err = ScanForAliases(dir, recursive, path)
	if err != nil {
		err = kerr.New("IAAETYCHSW", err, "process.Initialise", "ScanForImports")
		return
	}

	for p, _ := range aliases {
		params := []string{}
		if update {
			params = []string{"get", "-u", p}
		} else {
			params = []string{"get", p}
		}
		var out []byte
		if out, err = exec.Command("go", params...).CombinedOutput(); err != nil {
			err = kerr.New("HHKSTQMAKG", err, "process.Initialize", "go get command: %s", out)
			return
		}
	}

	return

}

func KegoCmd(dir string, update bool, recursive bool, verbose bool, path string, aliases map[string]string) error {
	if err := GenerateAndRunCmd(F_CMD_MAIN, dir, update, recursive, verbose, path, aliases); err != nil {
		return err
	}
	if err := GenerateAndRunCmd(F_CMD_TYPES, dir, update, recursive, verbose, path, aliases); err != nil {
		return err
	}
	if err := GenerateAndRunCmd(F_CMD_VALIDATE, dir, update, recursive, verbose, path, aliases); err != nil {
		return err
	}
	return nil
}

// GenerateFiles generates the source code from templates and writes
// the files to the correct folders.
//
// file == F_MAIN: the generated.go in the root of the package.
//
// file == F_TYPES: the generated.go containing advanced type information
// in the "types" sub package. Note that to generate this file, we need to
// have the main generated.go compiled in, so we generate a temporary
// command and run it with "go run".
//
// file == F_CMD_TYPES: this is the temporary command that we create in order to
// generate the types source file
//
// file == F_CMD_VALIDATE: this is the temporary command that we create in order
// to run the validation
//
func GenerateFiles(file fileType, dir string, update bool, recursive bool, verbose bool, path string, aliases map[string]string) error {

	if verbose {
		fmt.Println("Generating", file)
	}

	outputDir := dir
	if file == F_TYPES {
		outputDir = filepath.Join(dir, "types")
	}

	ignoreUnknownTypes := true
	if file == F_TYPES {
		ignoreUnknownTypes = false
	}

	if file == F_MAIN || file == F_TYPES {
		// If type == F_GLOBALS, we have already generated and imported the types, so
		// there is no need to scan.
		if err := ScanForTypes(dir, ignoreUnknownTypes, recursive, path, aliases); err != nil {
			return kerr.New("XYIUHERDHE", err, "process.GenerateFiles", "ScanForTypes")
		}
	} else {
		// However, we need to scan for the globals.
		if err := ScanForGlobals(dir, recursive, path, aliases); err != nil {
			return kerr.New("JQLAQVKLAN", err, "process.GenerateFiles", "ScanForGlobals")
		}
	}

	source, err := Generate(file, path, aliases)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "process.GenerateFiles", "Generate")
	}

	// We only backup in the system types because they are the only
	// generated files you'll ever need to roll back
	backup := path == "kego.io/system" || path == "kego.io/system/types"

	filename := "generated.go"
	if file == F_GLOBALS {
		filename = "globals.go"
	}

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "process.GenerateFiles", "save")
	}

	return nil
}

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run". When run, this command generates the extra types data in
// the "types" subpackage.
func GenerateAndRunCmd(file fileType, dir string, update bool, recursive bool, verbose bool, path string, aliases map[string]string) error {

	if verbose {
		fmt.Println("Generating", file)
	}

	source, err := Generate(file, path, aliases)
	if err != nil {
		return kerr.New("SPRFABSRWK", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "Generate")
	}

	outputDir, err := ioutil.TempDir(dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "save")
	}

	if file == F_CMD_VALIDATE {
		cmd := exec.Command("go", "build", "-o", "validate", outputPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return kerr.New("OEPAEEYKIS", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "cmd.Run (go build):\n%s", out)
		}
		if verbose {
			fmt.Print(string(out))
		}
	}

	command := ""
	params := []string{}

	if file == F_CMD_VALIDATE {
		command = "./validate"
	} else {
		command = "go"
		params = []string{"run", outputPath}
	}

	if *generatorPathFlag != "" {
		params = append(params, fmt.Sprintf("-p=%s", strconv.Quote(*generatorPathFlag)))
	}
	if update {
		params = append(params, "-u")
	}
	if recursive {
		params = append(params, "-r")
	}
	if verbose {
		params = append(params, "-v")
	}
	cmd := exec.Command(command, params...)

	if verbose {
		fmt.Println("Running", file)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return kerr.New("UDDSSMQRHA", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "cmd.Run: \n%s", out)
	}

	if verbose {
		fmt.Print(string(out))
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

func getPackage(dir string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
	var savedError error
	for _, gopath := range gopaths {
		if strings.HasPrefix(dir, gopath) {
			gosrc := fmt.Sprintf("%s/src", gopath)
			relpath, err := filepath.Rel(gosrc, dir)
			if err != nil {
				savedError = err
				continue
			}
			if relpath == "" {
				continue
			}
			return relpath, nil
		}
	}
	if savedError != nil {
		return "", savedError
	}
	return "", kerr.New("CXOETFPTGM", nil, "process.getPackage", "Package not found for %s", dir)
}
