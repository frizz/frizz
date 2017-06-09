package system

import (
	"context"
	"sync"
)

type ContextKeyType int

const (
	UnpackContextKey   ContextKeyType = 0
	RegistryContextKey ContextKeyType = 1
)

func NewContext(path string) context.Context {
	u := &UnpackContext{
		RWMutex: &sync.RWMutex{},
		imports: make(map[string]string),
		Path:    path,
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, UnpackContextKey, u)
	ctx = context.WithValue(ctx, RegistryContextKey, Registry)
	return ctx
}

type UnpackContext struct {
	*sync.RWMutex
	Path    string
	imports map[string]string
}

func (u *UnpackContext) Get(alias string) (string, bool) {
	u.RLock()
	defer u.RUnlock()
	path, ok := u.imports[alias]
	return path, ok
}

func (u *UnpackContext) Set(alias, path string) {
	u.Lock()
	defer u.Unlock()
	u.imports[alias] = path
}

type RegistryContextInterface interface {
	Get(RegistryTypeKey) (RegistryType, bool)
	Set(RegistryTypeKey, RegistryType)
}

type RegistryContext struct {
	*sync.RWMutex
	types map[RegistryTypeKey]RegistryType
}

type RegistryTypeKey struct {
	Path string
	Name string
}

type RegistryType struct {
	Unpacker func(context.Context, interface{}) (interface{}, error)
}

func (r *RegistryContext) Get(k RegistryTypeKey) (RegistryType, bool) {
	r.RLock()
	defer r.RUnlock()
	t, ok := r.types[k]
	return t, ok
}

func (r *RegistryContext) Set(key RegistryTypeKey, typ RegistryType) {
	r.Lock()
	defer r.Unlock()
	r.types[key] = typ
}

var Registry = &RegistryContext{
	RWMutex: &sync.RWMutex{},
	types:   make(map[RegistryTypeKey]RegistryType),
}
