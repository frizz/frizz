package system

import (
	"sort"
	"sync"

	"reflect"

	"kego.io/json"
	"kego.io/kerr"
)

var registry struct {
	sync.RWMutex
	m map[Reference]Hashed
}

type Hashed struct {
	Object ObjectInterface
	Hash   uint64
}

func Register(path string, name string, ob ObjectInterface, hash uint64) error {
	registry.Lock()
	defer registry.Unlock()
	if registry.m == nil {
		registry.m = make(map[Reference]Hashed)
	}
	r := *NewReference(path, name)
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
	delete(registry.m, *NewReference(path, name))
}

func GetGlobal(path string, name string) (hashed Hashed, found bool) {
	registry.RLock()
	defer registry.RUnlock()
	hashed, found = registry.m[*NewReference(path, name)]
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
		if filter != nil && *h.Object.GetObject(nil).Type != *filter {
			continue
		}
		out = append(out, h)
	}
	sort.Sort(out)
	return []Hashed(out)
}

func GetAllTypesThatImplementInterface(typ *Type, rule bool) []Hashed {
	registry.RLock()
	defer registry.RUnlock()
	out := SortableHashed{}

	if typ.Interface && rule || !typ.Interface && !rule {
		// we can't try to use a rule interface for an interface type, and we can't use a non
		// interface type that doesn't have a rule type
		return nil
	}

	var reflectType reflect.Type
	if typ.Interface {
		rt, _, ok := json.GetType(typ.Id.Package, typ.Id.Name)
		if !ok {
			return nil
		}
		reflectType = rt
	} else { // => rule == true
		rt, _, ok := json.GetTypeInterface(typ.Id.Package, typ.Id.Name)
		if !ok {
			return nil
		}
		reflectType = rt
	}

	for _, h := range registry.m {
		if t, ok := h.Object.(*Type); ok && t.Implements(reflectType) {
			out = append(out, h)
		}
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
	return s[i].Object.GetObject(nil).Id.Value() < s[j].Object.GetObject(nil).Id.Value()
}
