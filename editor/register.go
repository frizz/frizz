package editor

import (
	"sync"

	"kego.io/system"
)

type Factory func(*Node) Editor

var registry struct {
	sync.RWMutex
	m map[system.Reference]Factory
}

func Register(typ system.Reference, factory Factory) {
	registry.Lock()
	defer registry.Unlock()
	if registry.m == nil {
		registry.m = make(map[system.Reference]Factory)
	}
	registry.m[typ] = factory
	return
}

func Get(typ system.Reference) Factory {
	registry.RLock()
	defer registry.RUnlock()
	if registry.m == nil {
		return nil
	}
	f, ok := registry.m[typ]
	if !ok {
		return nil
	}
	return f
}
