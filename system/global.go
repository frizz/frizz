package system

import (
	"sort"
	"sync"

	"kego.io/kerr"
)

var registry struct {
	sync.RWMutex
	m map[Reference]Hashed
}

type Hashed struct {
	Object Object
	Hash   uint64
}

func Register(path string, name string, ob Object, hash uint64) error {
	registry.Lock()
	defer registry.Unlock()
	if registry.m == nil {
		registry.m = make(map[Reference]Hashed)
	}
	r := NewReference(path, name)
	if h, found := registry.m[r]; found && h.Hash != hash {
		return kerr.New("ANBTDVMCYE", nil, "Global %s already exists", r.String())
	}
	registry.m[r] = Hashed{ob, hash}
	return nil
}

func Unregister(path string, name string) {
	registry.Lock()
	defer registry.Unlock()
	if registry.m == nil {
		return
	}
	delete(registry.m, NewReference(path, name))
}

func GetGlobal(path string, name string) (hashed Hashed, found bool) {
	registry.RLock()
	defer registry.RUnlock()
	hashed, found = registry.m[NewReference(path, name)]
	return
}

// GetAllGlobalsInPackage returns all the globals in a given package, optionally
// filtering them by type. Use filter = nil to return all globals
func GetAllGlobalsInPackage(path string, filter *Reference) []Hashed {
	out := SortableHashed{}
	registry.RLock()
	defer registry.RUnlock()
	for ref, h := range registry.m {
		if ref.Package != path {
			continue
		}
		if filter != nil && h.Object.Object().Type != *filter {
			continue
		}
		out = append(out, h)
	}
	sort.Sort(out)
	return []Hashed(out)
}

type SortableHashed []Hashed

func (s SortableHashed) Len() int {
	return len(s)
}
func (s SortableHashed) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableHashed) Less(i, j int) bool {
	return s[i].Object.Object().Id.Value() < s[j].Object.Object().Id.Value()
}
