package jsonctx

import (
	"reflect"
	"sync"
)

var packages struct {
	sync.RWMutex
	m map[string]*packageInfo
}

// The types map is never used
type packageInfo struct {
	path  string
	hash  uint64
	types map[string]*typeInfo
}

type typeInfo struct {
	typ      reflect.Type
	rule     reflect.Type
	iface    reflect.Type
	newFunc  func() interface{}
	ruleFunc func() interface{}
}

func InitPackage(path string, hash uint64) *packageInfo {
	packages.Lock()
	defer packages.Unlock()
	if packages.m == nil {
		packages.m = map[string]*packageInfo{}
	}
	if p, ok := packages.m[path]; ok {
		return p
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
	if p.types[name] == nil {
		p.types[name] = &typeInfo{}
	}
	p.types[name].typ = typ
	p.types[name].rule = rule
	p.types[name].iface = iface
}
func (p *packageInfo) InitNew(name string, newFunc, ruleFunc func() interface{}) {
	if p.types[name] == nil {
		p.types[name] = &typeInfo{}
	}
	p.types[name].newFunc = newFunc
	p.types[name].ruleFunc = ruleFunc
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
