package process

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"kego.io/kerr"
)

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
		return nil
	}

	outputDir := filepath.Join(dir, "kego_temporary_cmd")
	outputName := "generated_cmd.go"
	outputPath := filepath.Join(outputDir, outputName)

	if err = save(outputDir, source, outputName); err != nil {
		return kerr.New("FRLCYFOWCJ", err, "process.GenerateAndRunCmd", "save")
	}

	cmd := exec.Command("go", "run", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return kerr.New("UDDSSMQRHA", err, "process.GenerateAndRunCmd", "cmd.Run")
	}

	return nil
}

type fileType string

const (
	F_MAIN  fileType = "main"
	F_TYPES fileType = "types"
	F_CMD   fileType = "cmd"
)

var testFlag = flag.Bool("test", false, "test mode? e.g. don't write the files")
var pathFlag = flag.String("path", "", "full package path e.g. github.com/foo/bar")

func GeneratorInit() {
	flag.Parse()
}

func parseOptions() (test bool, dir string, path string, imports map[string]string, err error) {

	test = *testFlag
	path = *pathFlag

	dir, err = filepath.Abs("")
	if err != nil {
		err = kerr.New("OKOLXAMBSJ", err, "process.GenerateFiles", "filepath.Abs")
		return
	}

	if path == "" {
		gopath := os.Getenv("GOPATH")
		path, err = getPackage(dir, gopath)
		if err != nil {
			err = kerr.New("PSRAWHQCPV", err, "process.GenerateFiles", "getPackageFromEnv")
			return
		}
	}

	imports = map[string]string{}

	return

}

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
		ignoreUnknownTypes = true
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

	if err = save(outputDir, source, "generated.go"); err != nil {
		return kerr.New("UONJTTSTWW", err, "process.GenerateFiles", "save")
	}

	return nil
}

func save(dir string, contents []byte, name string) error {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.New("BPGOUIYPXO", err, "process.save", "os.MkdirAll")
		}
	}
	file := filepath.Join(dir, name)
	backup(file, filepath.Join(dir, fmt.Sprintf("%s.backup", name)))

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

func backup(file string, backup string) {
	if _, err := os.Stat(backup); err == nil {
		os.Remove(backup)
	}
	if _, err := os.Stat(file); err == nil {
		os.Rename(file, backup)
	}
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
