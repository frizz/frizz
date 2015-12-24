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
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process/scan"
	"kego.io/system"
)

type optionsSpec interface {
	getOptions() FromDefaults
}

type FromFlags struct {
	Edit    *bool
	Update  *bool
	Verbose *bool
	Path    *string
	Debug   *bool
}

type FromDefaults struct {
	Edit    bool
	Update  bool
	Verbose bool
	Path    string
	Debug   bool
}

func (f FromDefaults) getOptions() FromDefaults {
	return f
}

func (f FromFlags) getOptions() FromDefaults {

	var edit, update, verbose, debug *bool
	var path *string
	if f.Edit == nil {
		edit = flag.Bool("e", false, "Edit: open the editor")
	} else {
		edit = f.Edit
	}
	if f.Update == nil {
		update = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
	} else {
		update = f.Update
	}
	if f.Verbose == nil {
		verbose = flag.Bool("v", false, "Verbose")
	} else {
		verbose = f.Verbose
	}
	if f.Debug == nil {
		debug = flag.Bool("d", false, "Debug: don't close the server when the connection is closed")
	} else {
		debug = f.Debug
	}
	if !flag.Parsed() {
		flag.Parse()
	}
	if f.Path == nil {
		p := flag.Arg(0)
		path = &p
	} else {
		path = f.Path
	}

	return FromDefaults{
		Edit:    *edit,
		Update:  *update,
		Verbose: *verbose,
		Path:    *path,
		Debug:   *debug,
	}
}

func Initialise(overrides optionsSpec) (context.Context, context.CancelFunc, error) {
	if overrides == nil {
		overrides = FromFlags{}
	}
	options := overrides.getOptions()

	ctx, cancel := context.WithCancel(context.Background())
	ctx = wgctx.NewContext(ctx)

	c := &cmdctx.Cmd{}
	e := &envctx.Env{}

	c.Edit = options.Edit
	c.Update = options.Update
	c.Verbose = options.Verbose
	c.Debug = options.Debug
	if options.Path == "" {

		dir, err := os.Getwd()
		if err != nil {
			return nil, nil, kerr.New("OKOLXAMBSJ", err, "os.Getwd")
		}
		c.Dir = dir

		pathFromDir, err := getPackagePath(c.Dir, os.Getenv("GOPATH"))
		if err != nil {
			return nil, nil, kerr.New("PSRAWHQCPV", err, "getPackage")
		}
		e.Path = pathFromDir

	} else {

		e.Path = options.Path

		out, err := exec.Command("go", "list", "-f", "{{.Dir}}", e.Path).CombinedOutput()
		if err == nil {
			c.Dir = strings.TrimSpace(string(out))
		} else {
			dir, err := getPackageDir(e.Path, os.Getenv("GOPATH"))
			if err != nil {
				return nil, nil, kerr.New("GXTUPMHETV", err, "Can't find %s", e.Path)
			}
			c.Dir = dir
		}

	}

	// ScanForPackage is finding our Aliases, so it's ok to give it an
	// incomplete env context with just the path
	dummyCtx := ctx
	dummyCtx = envctx.NewContext(dummyCtx, e)
	dummyCtx = cmdctx.NewContext(dummyCtx, c)
	if err := scan.ScanForPackage(dummyCtx); err != nil {
		return nil, nil, kerr.New("IAAETYCHSW", err, "scan.ScanForPackage")
	}
	p, ok := system.GetPackage(e.Path)
	if !ok {
		return nil, nil, kerr.New("BHLJNCIWUJ", nil, "Package not found")
	}
	e.Aliases = p.Aliases
	e.Recursive = p.Recursive

	ctx = envctx.NewContext(ctx, e)
	ctx = cmdctx.NewContext(ctx, c)

	return ctx, cancel, nil
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
