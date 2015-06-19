package system

import (
	"fmt"
	"strings"
	"sync"
)

var globals struct {
	sync.RWMutex
	m map[string]Object
}

func RegisterGlobal(name string, ob Object) {
	globals.Lock()
	defer globals.Unlock()
	if globals.m == nil {
		globals.m = make(map[string]Object)
	}
	globals.m[name] = ob
}

func UnregisterGlobal(name string) {
	globals.Lock()
	defer globals.Unlock()
	if globals.m == nil {
		return
	}
	delete(globals.m, name)
}

func GetGlobal(name string) (global Object, found bool) {
	globals.RLock()
	defer globals.RUnlock()
	global, found = globals.m[name]
	return
}

func GetAllGlobalsInPackage(path string) map[string]Object {
	out := map[string]Object{}
	globals.RLock()
	defer globals.RUnlock()
	for k, g := range globals.m {
		if strings.HasPrefix(k, fmt.Sprintf("%s:", path)) {
			out[k] = g
		}
	}
	return out
}
