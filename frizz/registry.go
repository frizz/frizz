package frizz

import (
	"sync"
)

type Registry struct {
	*sync.RWMutex
	types map[registryKey]registryType
}

type registryKey struct {
	Path string
	Name string
}

type registryType func(*Root, Stack, interface{}) (interface{}, error)

func (r *Registry) Get(path, name string) (registryType, bool) {
	r.RLock()
	defer r.RUnlock()
	t, ok := r.types[registryKey{path, name}]
	return t, ok
}

func (r *Registry) Set(path, name string, typ registryType) {
	r.Lock()
	defer r.Unlock()
	r.types[registryKey{path, name}] = typ
}

var DefaultRegistry = &Registry{
	RWMutex: &sync.RWMutex{},
	types:   make(map[registryKey]registryType),
}
