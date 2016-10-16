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
	newFunc       func() interface{}
	ruleFunc      func() interface{}
	dummyFunc     func() interface{}
	interfaceFunc func() reflect.Type
}

func (p *packageInfo) SetHash(hash uint64) {
	p.hash = hash
}

func InitPackage(path string) *packageInfo {
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
		types: map[string]*typeInfo{},
	}
	packages.m[path] = p
	return p
}

func (p *packageInfo) Init(name string, newFunc, ruleFunc func() interface{}, interfaceFunc func() reflect.Type) {
	if p.types[name] == nil {
		p.types[name] = &typeInfo{}
	}
	p.types[name].newFunc = newFunc
	p.types[name].ruleFunc = ruleFunc
	p.types[name].interfaceFunc = interfaceFunc
}

func (p *packageInfo) Dummy(name string, dummyFunc func() interface{}) {
	if p.types[name] == nil {
		p.types[name] = &typeInfo{}
	}
	p.types[name].dummyFunc = dummyFunc
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
