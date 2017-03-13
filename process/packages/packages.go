package packages // import "kego.io/process/packages"

// ke: {"package": {"complete": true}}

import (
	"context"

	"github.com/dave/gopackages"
	"kego.io/context/vosctx"
)

func GetDirFromPackage(ctx context.Context, packagePath string) (string, error) {
	vos := vosctx.FromContext(ctx)
	return gopackages.GetDirFromPackage(vos.Environ(), vos.Getenv("GOPATH"), packagePath)
}

func GetDirFromEmptyPackage(ctx context.Context, path string) (string, error) {
	vos := vosctx.FromContext(ctx)
	return gopackages.GetDirFromEmptyPackage(vos.Getenv("GOPATH"), path)
}

func GetPackageFromDir(ctx context.Context, dir string) (string, error) {
	vos := vosctx.FromContext(ctx)
	return gopackages.GetPackageFromDir(vos.Getenv("GOPATH"), dir)
}

func GetCurrentGopath(ctx context.Context) string {
	vos := vosctx.FromContext(ctx)
	currentDir, _ := vos.Getwd()
	return gopackages.GetCurrentGopath(vos.Getenv("GOPATH"), currentDir)
}
