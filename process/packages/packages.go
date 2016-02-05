package packages // import "kego.io/process/packages"

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

	out, err := exec.Command("go", "list", "-f", "{{.Dir}}", packagePath).CombinedOutput()
	if err == nil {
		return strings.TrimSpace(string(out)), nil
	}

	dir, err := GetDirFromEmptyPackage(packagePath, vos.Getenv("GOPATH"))
	if err != nil {
		return "", kerr.Wrap("GXTUPMHETV", err)
	}
	return dir, nil

}

func GetDirFromEmptyPackage(path string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
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
	return "", kerr.New("CXOETFPTGM", "Package not found for %s", dir)
}

func GetCurrentGopath(ctx context.Context) string {
	vos := vosctx.FromContext(ctx)
	currentDir, _ := vos.Getwd()
	return getCurrentGopath(currentDir, vos.Getenv("GOPATH"))
}

func getCurrentGopath(currentDir string, gopathEnv string) string {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		if strings.HasPrefix(currentDir, gopath) {
			return gopath
		}
	}
	return gopaths[0]
}
