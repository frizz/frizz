package sysctx // import "frizz.io/context/sysctx"

// notest

import (
	"sync"

	"sort"

	"context"

	"frizz.io/context/envctx"
	"github.com/dave/kerr"
)

type SysCache struct {
	sync.RWMutex
	m map[string]*SysPackageInfo
}

type SysPackageInfo struct {
	*envctx.Env
	PackageBytes    []byte
	PackageFilename string
	Types           *SysTypes
	Files           *SysFiles
	Globals         *SysGlobals
	Exports         *SysExports
}

type SysTypes struct {
	sync.RWMutex
	types map[string]*SysTypeInfo
}

type SysTypeInfo struct {
	Name string
	File string
	Type interface{}
}

type SysFiles struct {
	sync.RWMutex
	files map[string]*SysFileInfo
}

type SysFileInfo struct {
	File  string
	Bytes []byte
}

type SysGlobals struct {
	sync.RWMutex
	globals map[string]*SysGlobalInfo
}

type SysGlobalInfo struct {
	Name string
	File string
}

type SysExports struct {
	sync.RWMutex
	exports map[string]*SysExportInfo
}

type SysExportInfo struct {
	Name         string
	TypeName     string
	TypePackage  string
	JsonContents []byte
}

func (c *SysCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}
func (c *SysCache) Set(path string) *SysPackageInfo {
	return c.SetEnv(&envctx.Env{Path: path, Aliases: map[string]string{}})
}
func (c *SysCache) SetEnv(env *envctx.Env) *SysPackageInfo {
	c.Lock()
	defer c.Unlock()
	p := &SysPackageInfo{
		Env:     env,
		Types:   &SysTypes{types: map[string]*SysTypeInfo{}},
		Files:   &SysFiles{files: map[string]*SysFileInfo{}},
		Globals: &SysGlobals{globals: map[string]*SysGlobalInfo{}},
		Exports: &SysExports{exports: map[string]*SysExportInfo{}},
	}
	c.m[env.Path] = p
	return p
}

func (c *SysCache) GetType(path string, name string) (interface{}, bool) {
	p, ok := c.Get(path)
	if !ok {
		return nil, false
	}
	t, ok := p.Types.Get(name)
	if !ok {
		return nil, false
	}
	return t.Type, true
}

func (c *SysCache) Get(path string) (*SysPackageInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	info, ok := c.m[path]
	return info, ok
}

func (c *SysCache) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysTypes) Set(id string, filename string, t interface{}) {
	c.Lock()
	defer c.Unlock()
	c.types[id] = &SysTypeInfo{
		Name: id,
		File: filename,
		Type: t,
	}
}

func (c *SysTypes) Get(id string) (*SysTypeInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	ti, ok := c.types[id]
	return ti, ok
}

func (c *SysTypes) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.types)
}

func (c *SysTypes) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.types {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysGlobals) Set(id string, filename string) {
	c.Lock()
	defer c.Unlock()
	c.globals[id] = &SysGlobalInfo{
		Name: id,
		File: filename,
	}
}

func (c *SysGlobals) Get(id string) (*SysGlobalInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.globals[id]
	return t, ok
}

func (c *SysGlobals) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.globals)
}

func (c *SysGlobals) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.globals {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysFiles) Set(id string, filename string, b []byte) {
	c.Lock()
	defer c.Unlock()
	c.files[id] = &SysFileInfo{
		File:  filename,
		Bytes: b,
	}
}

func (c *SysFiles) Get(id string) (*SysFileInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	s, ok := c.files[id]
	return s, ok
}

func (c *SysFiles) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.files)
}

func (c *SysFiles) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.files {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysExports) Set(id, typeName, typePackage string, jsonContents []byte) {
	c.Lock()
	defer c.Unlock()
	c.exports[id] = &SysExportInfo{
		Name:         id,
		TypeName:     typeName,
		TypePackage:  typePackage,
		JsonContents: jsonContents,
	}
}

func (c *SysExports) Get(id string) (*SysExportInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.exports[id]
	return t, ok
}

func (c *SysExports) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.exports)
}

func (c *SysExports) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.exports {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// cacheKey is the key for sysctx.Cache values in Contexts.  It is
// unexported; clients use sysctx.NewContext and sysctx.FromContext
// instead of using this key directly.
var cacheKey key = 0

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context) context.Context {
	c := &SysCache{m: map[string]*SysPackageInfo{}}
	return context.WithValue(ctx, cacheKey, c)
}

// FromContext returns the Cache value stored in ctx, and panics if it's not found.
func FromContext(ctx context.Context) *SysCache {
	e, ok := ctx.Value(cacheKey).(*SysCache)
	if !ok {
		panic(kerr.New("DGGUPAXMUP", "No sys cache in ctx").Error())
	}
	return e
}

func FromContextOrNil(ctx context.Context) *SysCache {
	e, ok := ctx.Value(cacheKey).(*SysCache)
	if ok {
		return e
	}
	return nil
}
