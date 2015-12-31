package pkgutils // import "kego.io/process/pkgutils"
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"kego.io/kerr"
)

func GetDirFromPackage(packagePath string) (string, error) {

	out, err := exec.Command("go", "list", "-f", "{{.Dir}}", packagePath).CombinedOutput()
	if err == nil {
		return strings.TrimSpace(string(out)), nil
	}

	dir, err := getDirFromEmptyPackage(packagePath, os.Getenv("GOPATH"))
	if err != nil {
		return "", kerr.New("GXTUPMHETV", err, "Can't find %s", packagePath)
	}
	return dir, nil

}

func getDirFromEmptyPackage(path string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		dir := filepath.Join(gopath, "src", path)
		if s, err := os.Stat(dir); err == nil && s.IsDir() {
			return dir, nil
		}
	}
	return "", kerr.New("SUTCWEVRXS", nil, "%s not found", path)
}

func GetPackageFromDir(dir string) (string, error) {
	return getPackageFromDir(dir, os.Getenv("GOPATH"))
}

func getPackageFromDir(dir string, gopathEnv string) (string, error) {
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
