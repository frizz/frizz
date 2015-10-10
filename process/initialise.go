package process

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"kego.io/kerr"
	"kego.io/process/settings"
	"kego.io/system"
)

func InitialiseManually(edit bool, update bool, recursive bool, verbose bool, path string) (*settings.Settings, error) {
	set := settings.New()
	set.Edit = edit
	set.Update = update
	set.Recursive = recursive
	set.Verbose = verbose
	if path == "" {

		dir, err := os.Getwd()
		if err != nil {
			return nil, kerr.New("OKOLXAMBSJ", err, "os.Getwd")
		}
		set.Dir = dir

		pathFromDir, err := getPackagePath(set.Dir, os.Getenv("GOPATH"))
		if err != nil {
			return nil, kerr.New("PSRAWHQCPV", err, "getPackage")
		}
		set.Path = pathFromDir

	} else {

		set.Path = path

		out, err := exec.Command("go", "list", "-f", "{{.Dir}}", set.Path).CombinedOutput()
		if err == nil {
			set.Dir = strings.TrimSpace(string(out))
		} else {
			dir, err := getPackageDir(set.Path, os.Getenv("GOPATH"))
			if err != nil {
				return nil, kerr.New("GXTUPMHETV", err, "Can't find %s", set.Path)
			}
			set.Dir = dir
		}

	}

	if err := ScanForPackage(set); err != nil {
		return nil, kerr.New("IAAETYCHSW", err, "ScanForPackage")
	}
	p, ok := system.GetPackage(set.Path)
	if !ok {
		return nil, kerr.New("BHLJNCIWUJ", nil, "Package not found")
	}
	set.Aliases = p.Aliases

	return set, nil
}

func InitialiseCommand(update bool, recursive bool, path string) (*settings.Settings, error) {

	var editFlag = flag.Bool("e", false, "Edit: open the editor")
	var verboseFlag = flag.Bool("v", false, "Verbose")

	if !flag.Parsed() {
		flag.Parse()
	}

	set, err := InitialiseManually(*editFlag, update, recursive, *verboseFlag, path)
	if err != nil {
		return nil, kerr.New("UKAMOSMQST", err, "InitialiseManually")
	}

	return set, nil
}

func InitialiseAutomatic() (*settings.Settings, error) {

	var editFlag = flag.Bool("e", false, "Edit: open the editor")
	var updateFlag = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
	var recursiveFlag = flag.Bool("r", false, "Recursive: scan subdirectories for objects")
	var verboseFlag = flag.Bool("v", false, "Verbose")

	if !flag.Parsed() {
		flag.Parse()
	}

	var firstArg = flag.Arg(0)

	set, err := InitialiseManually(*editFlag, *updateFlag, *recursiveFlag, *verboseFlag, firstArg)
	if err != nil {
		return nil, kerr.New("UKAMOSMQST", err, "InitialiseManually")
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
	return "", kerr.New("CXOETFPTGM", nil, "Package not found for %s", dir)
}

func getPackageDir(path string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		dir := filepath.Join(gopath, "src", path)
		if s, err := os.Stat(dir); err == nil && s.IsDir() {
			return dir, nil
		}
	}
	return "", kerr.New("SUTCWEVRXS", nil, "%s not found", path)
}
