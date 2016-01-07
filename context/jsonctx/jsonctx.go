package jsonctx // import "kego.io/context/jsonctx"

import (
	"reflect"
	"sync"

	"sort"

	"fmt"

	"strings"

	"golang.org/x/net/context"
	"kego.io/kerr"
)

type JsonCache struct {
	Packages *PackageCache
	Dummies  *DummyCache
}

const RULE_PREFIX = "@"

func (c *JsonCache) GetType(path string, name string) (reflect.Type, bool) {
	rule := false
	if strings.HasPrefix(name, RULE_PREFIX) {
		rule = true
		name = name[1:]
	}

	p, ok := c.Packages.Get(path)
	if !ok {
		return nil, false
	}

	t, ok := p.Types.Get(name)
	if !ok {
		return nil, false
	}

	if rule {
		return t.Rule, true
	} else {
		return t.Type, true
	}
}
func (c *JsonCache) GetInterface(path string, name string) (reflect.Type, bool) {
	if strings.HasPrefix(name, RULE_PREFIX) {
		name = name[1:]
	}

	p, ok := c.Packages.Get(path)
	if !ok {
		return nil, false
	}

	t, ok := p.Types.Get(name)
	if !ok {
		return nil, false
	}

	return t.Iface, true
}
func (c *JsonCache) GetTypeByReflectType(typ reflect.Type) (path string, name string, found bool) {
	for _, p := range c.Packages.Keys() {
		pk, ok := c.Packages.Get(p)
		if !ok {
			continue
		}
		for _, n := range pk.Types.Keys() {
			def, ok := pk.Types.Get(n)
			if !ok {
				continue
			}
			if def.Type == typ || (def.Type.Kind() == reflect.Ptr && def.Type.Elem() == typ) {
				return p, n, true
			}
			if def.Rule != nil && (def.Rule == typ || (def.Rule.Kind() == reflect.Ptr && def.Rule.Elem() == typ)) {
				return p, "@" + n, true
			}
		}
	}
	return "", "", false
}
func (c *JsonCache) GetTypeByInterface(iface reflect.Type) (typ reflect.Type, found bool) {
	for _, p := range c.Packages.Keys() {
		pk, ok := c.Packages.Get(p)
		if !ok {
			continue
		}
		for _, n := range pk.Types.Keys() {
			def, ok := pk.Types.Get(n)
			if !ok {
				continue
			}
			if def.Iface == iface {
				return def.Type, true
			}
		}
	}
	return nil, false
}

type PackageCache struct {
	sync.RWMutex
	m map[string]*PackageInfo
}

type PackageInfo struct {
	Path  string
	Hash  uint64
	Types *TypeCache
}

type TypeCache struct {
	sync.RWMutex
	m map[string]*TypeInfo
}

type DummyCache struct {
	sync.RWMutex
	m map[reflect.Type]reflect.Type
}

type TypeInfo struct {
	Name  string
	Type  reflect.Type
	Rule  reflect.Type
	Iface reflect.Type
}

func (c *PackageCache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *PackageCache) Set(path string, hash uint64) *PackageInfo {
	c.Lock()
	defer c.Unlock()
	p := &PackageInfo{
		Path:  path,
		Hash:  hash,
		Types: &TypeCache{m: map[string]*TypeInfo{}},
	}
	c.m[path] = p
	return p
}

func (c *PackageCache) Get(path string) (*PackageInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	info, ok := c.m[path]
	return info, ok
}

func (c *PackageCache) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *TypeCache) Set(id string, t *TypeInfo) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *TypeCache) Get(id string) (*TypeInfo, bool) {
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

func (c *DummyCache) Set(iface reflect.Type, dummy reflect.Type) {
	c.Lock()
	defer c.Unlock()
	c.m[iface] = dummy
}

func (c *DummyCache) Get(iface reflect.Type) (reflect.Type, bool) {
	c.RLock()
	defer c.RUnlock()
	d, ok := c.m[iface]
	return d, ok
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// cacheKey is the key for jsonctx.Cache values in Contexts.  It is
// unexported; clients use jsonctx.NewContext and jsonctx.FromContext
// instead of using this key directly.
var jsonKey key = 0

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, auto bool, manual ...string) context.Context {

	if auto && len(manual) > 0 {
		panic(kerr.New("CMFIJSMYSB", nil, "if a list of manually imported packages is defined, auto must be false"))
	}

	pc := &PackageCache{m: map[string]*PackageInfo{}}
	packages.Lock()
	defer packages.Unlock()

	do := func(path string, pkg *packageInfo) {
		p := pc.Set(path, pkg.hash)
		for name, typ := range pkg.types {
			p.Types.Set(name, &TypeInfo{
				Name:  name,
				Type:  typ.typ,
				Rule:  typ.rule,
				Iface: typ.iface,
			})
		}
	}

	if auto {
		// In automatic mode, we import all packages
		for path, pkg := range packages.m {
			do(path, pkg)
		}
	} else {
		// If we specify a manual list of packages to import, we should import them
		for _, path := range manual {
			pkg, ok := packages.m[path]
			if !ok {
				panic(fmt.Errorf("Package not found %s", path))
			}
			do(path, pkg)
		}
	}

	dc := &DummyCache{m: map[reflect.Type]reflect.Type{}}
	if auto {
		for iface, dummy := range dummies.m {
			dc.Set(iface, dummy)
		}
	}

	jc := &JsonCache{
		Packages: pc,
		Dummies:  dc,
	}

	return context.WithValue(ctx, jsonKey, jc)
}

// FromContext returns the Cache value stored in ctx, and panics if it's not found.
func FromContext(ctx context.Context) *JsonCache {
	e, ok := ctx.Value(jsonKey).(*JsonCache)
	if !ok {
		panic(kerr.New("XUTUUVDMMX", nil, "No cache in ctx"))
	}
	return e
}
