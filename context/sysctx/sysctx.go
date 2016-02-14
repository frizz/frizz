package sysctx // import "kego.io/context/sysctx"

// ke: {"package": {"notest": true}}

import (
	"sync"

	"sort"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/kerr"
)

type SysCache struct {
	sync.RWMutex
	m map[string]*SysPackageInfo
}

type SysPackageInfo struct {
	*envctx.Env
	PackageBytes []byte
	Types        *SysTypes
	Files        *SysFiles
	Globals      *SysGlobals
}

type SysTypes struct {
	sync.RWMutex
	m map[string]interface{}
}

type SysFiles struct {
	sync.RWMutex
	m map[string][]byte
}

type SysGlobals struct {
	sync.RWMutex
	m map[string]SysGlobalInfo
}

type SysGlobalInfo struct {
	Name string
	File string
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
		Types:   &SysTypes{m: map[string]interface{}{}},
		Files:   &SysFiles{m: map[string][]byte{}},
		Globals: &SysGlobals{m: map[string]SysGlobalInfo{}},
	}
	c.m[env.Path] = p
	return p
}

func (c *SysCache) GetType(path string, name string) (interface{}, bool) {
	p, ok := c.Get(path)
	if !ok {
		return nil, false
	}
	return p.Types.Get(name)
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

func (c *SysTypes) Set(id string, t interface{}) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *SysTypes) Get(id string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *SysTypes) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *SysTypes) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysGlobals) Set(id string, t SysGlobalInfo) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *SysGlobals) Get(id string) (SysGlobalInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *SysGlobals) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *SysGlobals) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *SysFiles) Set(id string, b []byte) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = b
}

func (c *SysFiles) Get(id string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *SysFiles) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *SysFiles) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
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
