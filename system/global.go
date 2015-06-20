package system

import "sync"

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

func GetAllGlobalsInPackage(path string) map[Reference]Object {
	out := map[Reference]Object{}
	globals.RLock()
	defer globals.RUnlock()
	for ref, g := range globals.m {
		if ref.Package == path {
			out[ref] = g
		}
	}
	return out
}
