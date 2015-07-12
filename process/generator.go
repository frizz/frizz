package process

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func SetPathFlag(path string) {
	*generatorPathFlag = path
}

type settings struct {
	dir       string
	update    bool
	recursive bool
	verbose   bool
	path      string
	aliases   map[string]string
}

func (s settings) Path() string {
	return s.path
}

func Initialise() (settings, error) {

	if !flag.Parsed() {
		flag.Parse()
	}

	set := settings{}
	set.update = *generatorUpdateFlag
	set.recursive = *generatorRecursiveFlag
	set.verbose = *generatorVerboseFlag

	if *generatorPathFlag == "" {

		dir, err := os.Getwd()
		if err != nil {
			return settings{}, kerr.New("OKOLXAMBSJ", err, "process.Initialise", "os.Getwd")
		}
		set.dir = dir

		path, err := getPackagePath(set.dir, os.Getenv("GOPATH"))
		if err != nil {
			return settings{}, kerr.New("PSRAWHQCPV", err, "process.Initialise", "getPackage")
		}
		set.path = path

	} else {

		set.path = *generatorPathFlag

		out, err := exec.Command("go", "list", "-f", "{{.Dir}}", set.path).CombinedOutput()
		if err == nil {
			set.dir = strings.TrimSpace(string(out))
		} else {
			dir, err := getPackageDir(set.path, os.Getenv("GOPATH"))
			if err != nil {
				return settings{}, kerr.New("GXTUPMHETV", err, "process.Initialise", "Can't find %s", set.path)
			}
			set.dir = dir
		}

	}

	aliases, err := ScanForAliases(set)
	if err != nil {
		return settings{}, kerr.New("IAAETYCHSW", err, "process.Initialise", "ScanForImports")
	}
	set.aliases = aliases

	return set, nil

}

func KegoCmd(set settings) error {

	for p, _ := range set.aliases {
		if set.verbose {
			if set.update {
				fmt.Println("Updating package:", p)
			} else {
				fmt.Println("Getting package:", p)
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
		}
	}

	if err := GenerateAndRunCmd(F_CMD_MAIN, set); err != nil {
		return err
	}
	if err := GenerateAndRunCmd(F_CMD_TYPES, set); err != nil {
		return err
	}
	if err := GenerateAndRunCmd(F_CMD_VALIDATE, set); err != nil {
		return err
	}
	return nil
}

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run". When run, this command generates the extra types data in
// the "types" subpackage.
func GenerateAndRunCmd(file fileType, set settings) error {

	if set.verbose {
		fmt.Println("Generating", file)
	}

	source, err := Generate(file, set.path, set.aliases)
	if err != nil {
		return kerr.New("SPRFABSRWK", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "Generate")
	}

	outputDir, err := ioutil.TempDir(set.dir, "temporary")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	validateCommandPath := filepath.Join(set.dir, "validate")

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "save")
	}

	if file == F_CMD_VALIDATE {
		cmd := exec.Command("go", "build", "-o", validateCommandPath, outputPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return kerr.New("OEPAEEYKIS", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "cmd.Run (go build):\n%s", out)
		}
		if set.verbose {
			fmt.Print(string(out))
		}
	}

	if set.verbose {
		fmt.Println("Running", file)
	}

	command := ""
	params := []string{}

	if file == F_CMD_VALIDATE {
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
	out, err := exec.Command(command, params...).CombinedOutput()
	if err != nil {
		return kerr.New("UDDSSMQRHA", err, fmt.Sprintf("process.GenerateAndRunCmd %s", file), "cmd.Run: %s", out)
	}

	if set.verbose {
		fmt.Print(string(out))
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
func GenerateFiles(file fileType, set settings) error {

	if set.verbose {
		fmt.Println("Generating", file)
	}

	outputDir := set.dir
	if file == F_TYPES {
		outputDir = filepath.Join(set.dir, "types")
	}

	ignoreUnknownTypes := true
	if file == F_TYPES {
		ignoreUnknownTypes = false
	}

	if file == F_MAIN || file == F_TYPES {
		// If type == F_GLOBALS, we have already generated and imported the types, so
		// there is no need to scan.
		if err := ScanForTypes(ignoreUnknownTypes, set); err != nil {
			return kerr.New("XYIUHERDHE", err, "process.GenerateFiles", "ScanForTypes")
		}
	} else {
		// However, we need to scan for the globals.
		if err := ScanForGlobals(set); err != nil {
			return kerr.New("JQLAQVKLAN", err, "process.GenerateFiles", "ScanForGlobals")
		}
	}

	source, err := Generate(file, set.path, set.aliases)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "process.GenerateFiles", "Generate")
	}

	// We only backup in the system types because they are the only
	// generated files you'll ever need to roll back
	backup := set.path == "kego.io/system" || set.path == "kego.io/system/types"

	filename := "generated.go"
	if file == F_GLOBALS {
		filename = "globals.go"
	}

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "process.GenerateFiles", "save")
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

func getPackagePath(dir string, gopathEnv string) (string, error) {
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

func getPackageDir(path string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		dir := filepath.Join(gopath, "src", path)
		if s, err := os.Stat(dir); err == nil && s.IsDir() {
			return dir, nil
		}
	}
	return "", kerr.New("SUTCWEVRXS", nil, "process.getDir", "%s not found", path)
}
