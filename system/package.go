package system

import (
	"sync"

	"kego.io/kerr"
)

type ExtendedPackage struct {
	*Package
	Path	string
	Hash	uint64
}

var packages struct {
	sync.RWMutex
	m	map[string]*ExtendedPackage
}

func EmptyPackage() *Package {
	return &Package{
		Object: &Object{
			Type: NewReference("kego.io/system", "package"),
		},
		Aliases:	map[string]string{},
		Recursive:	false,
	}
}

func RegisterPackage(path string, p *Package, hash uint64) error {
	packages.Lock()
	defer packages.Unlock()
	if packages.m == nil {
		packages.m = make(map[string]*ExtendedPackage)
	}
	if pa, found := packages.m[path]; found && pa.Hash != hash {
		return kerr.New("RVCNUWAGAQ", "Package %s already exists", pa.Path)
	}
	packages.m[path] = &ExtendedPackage{Package: p, Path: path, Hash: hash}
	return nil
}

func UnregisterPackage(path string) {
	packages.Lock()
	defer packages.Unlock()
	if packages.m == nil {
		return
	}
	delete(packages.m, path)
}

func GetPackage(path string) (pkg *ExtendedPackage, found bool) {
	packages.RLock()
	defer packages.RUnlock()
	pkg, found = packages.m[path]
	return
}
