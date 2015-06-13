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
	F_MAIN  fileType = "main"
	F_TYPES          = "types"
	F_CMD            = "cmd"
)

var generatorTestFlag = flag.Bool("test", false, "test mode? e.g. don't write the files")
var generatorPathFlag = flag.String("path", "", "full package path e.g. github.com/foo/bar")

func parseOptions() (test bool, dir string, path string, imports map[string]string, err error) {

	if !flag.Parsed() {
		flag.Parse()
	}

	test = *generatorTestFlag
	path = *generatorPathFlag

	dir, err = os.Getwd()
	if err != nil {
		err = kerr.New("OKOLXAMBSJ", err, "process.GenerateFiles", "os.Getwd")
		return
	}

	if path == "" {
		path, err = getPackage(dir, os.Getenv("GOPATH"))
		if err != nil {
			err = kerr.New("PSRAWHQCPV", err, "process.GenerateFiles", "getPackage")
			return
		}
	}

	imports = map[string]string{}

	return

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
// file == F_CMD: this is the temporary command that we create in order to
// generate the types source file
//
func GenerateFiles(file fileType) error {

	test, dir, path, imports, err := parseOptions()
	if err != nil {
		return kerr.New("MHLOQTSAOX", err, "process.GenerateFiles", "parseOptions")
	}

	outputDir := dir
	if file == F_TYPES {
		outputDir = filepath.Join(dir, "types")
	}

	ignoreUnknownTypes := true
	if file == F_TYPES {
		ignoreUnknownTypes = false
	}

	if err := Scan(dir, ignoreUnknownTypes, path, imports); err != nil {
		return kerr.New("XYIUHERDHE", err, "process.GenerateFiles", "Scan")
	}

	source, err := Generate(file, path, imports)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "process.GenerateFiles", "Generate")
	}

	if test {
		fmt.Printf("%s\n", source)
		return nil
	}

	if err = save(outputDir, source, "generated.go", true); err != nil {
		return kerr.New("UONJTTSTWW", err, "process.GenerateFiles", "save")
	}

	return nil
}

// This creates a temporary folder in the package, in which the go source
// for a command is generated. This command is then compiled and run with
// "go run". When run, this command generates the extra types data in
// the "types" subpackage.
func GenerateAndRunCmd() error {
	test, dir, path, imports, err := parseOptions()
	if err != nil {
		return kerr.New("NGXAEJCFSA", err, "process.GenerateAndRunCmd", "parseOptions")
	}

	source, err := Generate(F_CMD, path, imports)
	if err != nil {
		return kerr.New("SPRFABSRWK", err, "process.GenerateAndRunCmd", "Generate")
	}

	if test {
		fmt.Printf("%s\n", source)
	}

	outputDir, err := ioutil.TempDir(dir, "")
	if err != nil {
		return kerr.New("HWOPVXYMCT", err, "process.GenerateAndRunCmd", "ioutil.TempDir")
	}
	defer os.RemoveAll(outputDir)
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	if err = save(outputDir, source, outputName, false); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "process.GenerateAndRunCmd", "save")
	}

	params := []string{"run", outputPath}
	if *generatorTestFlag {
		params = append(params, "-test")
	}
	if *generatorPathFlag != "" {
		params = append(params, fmt.Sprintf("-path=%s", strconv.Quote(*generatorPathFlag)))
	}
	cmd := exec.Command("go", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return kerr.New("UDDSSMQRHA", err, "process.GenerateAndRunCmd", "cmd.Run")
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
