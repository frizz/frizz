package sysctx // import "kego.io/context/sysctx"

// ke: {"package": {"notest":true}}

import (
	"sync"

	"sort"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/kerr"
)

type SysCache struct {
	sync.RWMutex
	m map[string]*PackageInfo
}

type PackageInfo struct {
	Environment  *envctx.Env
	PackageBytes []byte
	Types        *TypeCache
	TypeSource   *TypeSourceCache
	Globals      *GlobalCache
}

type TypeCache struct {
	sync.RWMutex
	m map[string]interface{}
}

type TypeSourceCache struct {
	sync.RWMutex
	m map[string][]byte
}

type GlobalCache struct {
	sync.RWMutex
	m map[string]GlobalInfo
}

type GlobalInfo struct {
	Name string
	File string
}

func (c *SysCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}
func (c *SysCache) Set(env *envctx.Env) *PackageInfo {
	c.Lock()
	defer c.Unlock()
	p := &PackageInfo{
		Environment: env,
		Types:       &TypeCache{m: map[string]interface{}{}},
		TypeSource:  &TypeSourceCache{m: map[string][]byte{}},
		Globals:     &GlobalCache{m: map[string]GlobalInfo{}},
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

func (c *SysCache) Get(path string) (*PackageInfo, bool) {
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

func (c *TypeCache) Set(id string, t interface{}) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *TypeCache) Get(id string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *TypeCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *TypeCache) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *GlobalCache) Set(id string, t GlobalInfo) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *GlobalCache) Get(id string) (GlobalInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *GlobalCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *GlobalCache) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *TypeSourceCache) Set(id string, b []byte) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = b
}

func (c *TypeSourceCache) Get(id string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.m[id]
	return t, ok
}

func (c *TypeSourceCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *TypeSourceCache) Keys() []string {
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
	c := &SysCache{m: map[string]*PackageInfo{}}
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
