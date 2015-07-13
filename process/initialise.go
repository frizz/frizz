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

type settings struct {
	dir       string
	update    bool
	recursive bool
	verbose   bool
	globals   bool
	path      string
	aliases   map[string]string
}

func (s settings) Path() string {
	return s.path
}
func (s settings) Globals() bool {
	return s.globals
}

func InitialiseManually(update bool, recursive bool, verbose bool, globals bool, path string) (settings, error) {
	set := settings{}
	set.update = update
	set.recursive = recursive
	set.verbose = verbose
	set.globals = globals
	if path == "" {

		dir, err := os.Getwd()
		if err != nil {
			return settings{}, kerr.New("OKOLXAMBSJ", err, "process.InitialiseManually", "os.Getwd")
		}
		set.dir = dir

		pathFromDir, err := getPackagePath(set.dir, os.Getenv("GOPATH"))
		if err != nil {
			return settings{}, kerr.New("PSRAWHQCPV", err, "process.InitialiseManually", "getPackage")
		}
		set.path = pathFromDir

	} else {

		set.path = path

		out, err := exec.Command("go", "list", "-f", "{{.Dir}}", set.path).CombinedOutput()
		if err == nil {
			set.dir = strings.TrimSpace(string(out))
		} else {
			dir, err := getPackageDir(set.path, os.Getenv("GOPATH"))
			if err != nil {
				return settings{}, kerr.New("GXTUPMHETV", err, "process.InitialiseManually", "Can't find %s", set.path)
			}
			set.dir = dir
		}

	}

	aliases, err := ScanForAliases(set)
	if err != nil {
		return settings{}, kerr.New("IAAETYCHSW", err, "process.InitialiseManually", "ScanForImports")
	}
	set.aliases = aliases

	return set, nil
}

func InitialiseAutomatic() (settings, error) {

	var pathFlag = flag.String("p", "", "Package: full package path e.g. github.com/foo/bar")
	var updateFlag = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
	var recursiveFlag = flag.Bool("r", false, "Recursive: scan subdirectories for objects")
	var globalsFlag = flag.Bool("g", false, "Globals: generate Go literals for all global objects")
	var verboseFlag = flag.Bool("v", false, "Verbose")

	if !flag.Parsed() {
		flag.Parse()
	}

	set, err := InitialiseManually(*updateFlag, *recursiveFlag, *verboseFlag, *globalsFlag, *pathFlag)
	if err != nil {
		return settings{}, kerr.New("UKAMOSMQST", err, "process.InitialiseAutomatic", "InitialiseManually")
	}

	return set, nil

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
