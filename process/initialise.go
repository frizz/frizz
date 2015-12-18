package process

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/process/scan"
	"kego.io/system"
)

func InitialiseManually(edit bool, update bool, recursive bool, verbose bool, path string) (context.Context, error) {

	c := &cmdctx.Cmd{}
	e := &envctx.Env{}

	c.Edit = edit
	c.Update = update
	c.Recursive = recursive
	c.Verbose = verbose
	if path == "" {

		dir, err := os.Getwd()
		if err != nil {
			return nil, kerr.New("OKOLXAMBSJ", err, "os.Getwd")
		}
		c.Dir = dir

		pathFromDir, err := getPackagePath(c.Dir, os.Getenv("GOPATH"))
		if err != nil {
			return nil, kerr.New("PSRAWHQCPV", err, "getPackage")
		}
		e.Path = pathFromDir

	} else {

		e.Path = path

		out, err := exec.Command("go", "list", "-f", "{{.Dir}}", e.Path).CombinedOutput()
		if err == nil {
			c.Dir = strings.TrimSpace(string(out))
		} else {
			dir, err := getPackageDir(e.Path, os.Getenv("GOPATH"))
			if err != nil {
				return nil, kerr.New("GXTUPMHETV", err, "Can't find %s", e.Path)
			}
			c.Dir = dir
		}

	}

	// ScanForPackage is finding our Aliases, so it's ok to give it an
	// incomplete env context with just the path
	dummyCtx := context.Background()
	dummyCtx = envctx.NewContext(dummyCtx, e)
	dummyCtx = cmdctx.NewContext(dummyCtx, c)
	if err := scan.ScanForPackage(dummyCtx); err != nil {
		return nil, kerr.New("IAAETYCHSW", err, "scan.ScanForPackage")
	}
	p, ok := system.GetPackage(e.Path)
	if !ok {
		return nil, kerr.New("BHLJNCIWUJ", nil, "Package not found")
	}
	e.Aliases = p.Aliases

	ctx := context.Background()
	ctx = envctx.NewContext(ctx, e)
	ctx = cmdctx.NewContext(ctx, c)

	return ctx, nil
}

func InitialiseCommand(update bool, recursive bool, path string) (context.Context, error) {

	var editFlag = flag.Bool("e", false, "Edit: open the editor")
	var verboseFlag = flag.Bool("v", false, "Verbose")

	if !flag.Parsed() {
		flag.Parse()
	}

	ctx, err := InitialiseManually(*editFlag, update, recursive, *verboseFlag, path)
	if err != nil {
		return nil, kerr.New("UKAMOSMQST", err, "InitialiseManually")
	}

	return ctx, nil
}

func InitialiseAutomatic() (context.Context, error) {

	var editFlag = flag.Bool("e", false, "Edit: open the editor")
	var updateFlag = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
	var recursiveFlag = flag.Bool("r", false, "Recursive: scan subdirectories for objects")
	var verboseFlag = flag.Bool("v", false, "Verbose")

	if !flag.Parsed() {
		flag.Parse()
	}

	var firstArg = flag.Arg(0)

	ctx, err := InitialiseManually(*editFlag, *updateFlag, *recursiveFlag, *verboseFlag, firstArg)
	if err != nil {
		return nil, kerr.New("UKAMOSMQST", err, "InitialiseManually")
	}

	return ctx, nil

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
