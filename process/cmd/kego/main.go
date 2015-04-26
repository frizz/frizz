package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"kego.io/process"
	"kego.io/system"
)

func main() {
	var testFlag = flag.Bool("test", false, "test mode? e.g. don't write the files")
	var pathFlag = flag.String("path", "", "full package path e.g. github.com/foo/bar")
	flag.Parse()
	testMode := *testFlag
	packagePath := *pathFlag

	currentDir, err := filepath.Abs("")
	if err != nil {
		log.Fatalf("Error getting the current working directory:\n%v\n", err.Error())
	}

	if packagePath == "" {
		path, err := GetPackage(currentDir)
		if err != nil {
			log.Fatalf("Error while getting package path from current working directory. You can specify the full package path with the -path=github.com/foo\n%v\n", err.Error())
		}
		packagePath = path
	}

	parts := strings.Split(packagePath, string(os.PathSeparator))
	packageName := parts[len(parts)-1]

	imports := map[string]string{"system": "kego.io/system", "json": "kego.io/json"}

	types := map[string]*system.Type{}

	if err := process.Scan(currentDir, packageName, packagePath, imports, types); err != nil {
		panic(err)
	}

	if err := process.Generate(currentDir, packageName, packagePath, imports, types, testMode); err != nil {
		panic(err)
	}
}

func GetPackage(dir string) (string, error) {
	return getPackage(dir, os.Getenv("GOPATH"))
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
	return "", fmt.Errorf("Package not found for %s\n", dir)
}
