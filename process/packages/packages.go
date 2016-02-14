package packages // import "kego.io/process/packages"

// ke: {"package": {"complete": true}}

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"os"

	"golang.org/x/net/context"
	"kego.io/context/vosctx"
	"kego.io/kerr"
)

func GetDirFromPackage(ctx context.Context, packagePath string) (string, error) {

	vos := vosctx.FromContext(ctx)

	exe := exec.Command("go", "list", "-f", "{{.Dir}}", packagePath)
	exe.Env = vos.Environ()
	out, err := exe.CombinedOutput()
	if err == nil {
		return strings.TrimSpace(string(out)), nil
	}

	dir, err := GetDirFromEmptyPackage(ctx, packagePath)
	if err != nil {
		return "", kerr.Wrap("GXTUPMHETV", err)
	}
	return dir, nil

}

func GetDirFromEmptyPackage(ctx context.Context, path string) (string, error) {
	vos := vosctx.FromContext(ctx)
	gopaths := filepath.SplitList(vos.Getenv("GOPATH"))
	for _, gopath := range gopaths {
		dir := filepath.Join(gopath, "src", path)
		if s, err := os.Stat(dir); err == nil && s.IsDir() {
			return dir, nil
		}
	}
	return "", NotFoundError{Struct: kerr.New("SUTCWEVRXS", "%s not found", path)}
}

type NotFoundError struct {
	kerr.Struct
}

func GetPackageFromDir(ctx context.Context, dir string) (string, error) {
	vos := vosctx.FromContext(ctx)
	gopaths := filepath.SplitList(vos.Getenv("GOPATH"))
	var savedError error
	for _, gopath := range gopaths {
		if strings.HasPrefix(dir, gopath) {
			gosrc := fmt.Sprintf("%s/src", gopath)
			relpath, err := filepath.Rel(gosrc, dir)
			if err != nil {
				// ke: {"block": {"notest": true}}
				// I don't *think* we can trigger this error if dir starts with gopath
				savedError = err
				continue
			}
			if relpath == "" {
				// ke: {"block": {"notest": true}}
				// I don't *think* we can trigger this either
				continue
			}
			// Remember we're returning a package path which uses forward slashes even on windows
			return filepath.ToSlash(relpath), nil
		}
	}
	if savedError != nil {
		// ke: {"block": {"notest": true}}
		return "", savedError
	}
	return "", kerr.New("CXOETFPTGM", "Package not found for %s", dir)
}

func GetCurrentGopath(ctx context.Context) string {
	vos := vosctx.FromContext(ctx)
	currentDir, _ := vos.Getwd()
	gopaths := filepath.SplitList(vos.Getenv("GOPATH"))
	for _, gopath := range gopaths {
		if strings.HasPrefix(currentDir, gopath) {
			return gopath
		}
	}
	return gopaths[0]
}
