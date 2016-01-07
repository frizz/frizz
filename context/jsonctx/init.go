package jsonctx

import (
	"reflect"
	"sync"
)

var packages struct {
	sync.RWMutex
	m map[string]*packageInfo
}

type packageInfo struct {
	path  string
	hash  uint64
	types map[string]*typeInfo
}

type typeInfo struct {
	typ   reflect.Type
	rule  reflect.Type
	iface reflect.Type
}

func InitPackage(path string, hash uint64) *packageInfo {
	packages.Lock()
	defer packages.Unlock()
	if packages.m == nil {
		packages.m = map[string]*packageInfo{}
	}
	p := &packageInfo{
		path:  path,
		hash:  hash,
		types: map[string]*typeInfo{},
	}
	packages.m[path] = p
	return p
}
func (p *packageInfo) InitType(name string, typ reflect.Type, rule reflect.Type, iface reflect.Type) {
	p.types[name] = &typeInfo{
		typ:   typ,
		rule:  rule,
		iface: iface,
	}
}

var dummies struct {
	sync.RWMutex
	m map[reflect.Type]reflect.Type
}

func InitDummy(t reflect.Type, dummy reflect.Type) {
	dummies.Lock()
	defer dummies.Unlock()
	if dummies.m == nil {
		dummies.m = make(map[reflect.Type]reflect.Type)
	}
	dummies.m[t] = dummy
}
