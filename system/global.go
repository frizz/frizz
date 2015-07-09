package system

import (
	"sort"
	"sync"
)

var globals struct {
	sync.RWMutex
	m map[Reference]Object
}

func RegisterGlobal(path string, name string, ob Object) {
	globals.Lock()
	defer globals.Unlock()
	if globals.m == nil {
		globals.m = make(map[Reference]Object)
	}
	globals.m[NewReference(path, name)] = ob
}

func UnregisterGlobal(path string, name string) {
	globals.Lock()
	defer globals.Unlock()
	if globals.m == nil {
		return
	}
	delete(globals.m, NewReference(path, name))
}

func GetGlobal(path string, name string) (global Object, found bool) {
	globals.RLock()
	defer globals.RUnlock()
	global, found = globals.m[NewReference(path, name)]
	return
}

func GetAllGlobalsInPackage(path string) []Object {
	out := SortableObjects{}
	globals.RLock()
	defer globals.RUnlock()
	for ref, g := range globals.m {
		if ref.Package == path {
			out = append(out, g)
		}
	}
	sort.Sort(out)
	return []Object(out)
}

type SortableObjects []Object

func (s SortableObjects) Len() int {
	return len(s)
}
func (s SortableObjects) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableObjects) Less(i, j int) bool {
	return s[i].GetBase().Id.Value() < s[j].GetBase().Id.Value()
}
