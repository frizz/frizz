package jsonctx // import "kego.io/context/jsonctx"

// ke: {"package": {"notest": true}}

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
	Packages *JsonPackages
	Dummies  *JsonDummies
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
			if def.Type != nil && (def.Type == typ || (def.Type.Kind() == reflect.Ptr && def.Type.Elem() == typ)) {
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

type JsonPackages struct {
	sync.RWMutex
	m map[string]*JsonPackageInfo
}

type JsonPackageInfo struct {
	Path  string
	Hash  uint64
	Types *JsonTypes
}

type JsonTypes struct {
	sync.RWMutex
	m map[string]*JsonTypeInfo
}

type JsonDummies struct {
	sync.RWMutex
	m map[reflect.Type]reflect.Type
}

type JsonTypeInfo struct {
	Name  string
	Type  reflect.Type
	Rule  reflect.Type
	Iface reflect.Type
}

func (c *JsonPackages) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *JsonPackages) Set(path string, hash uint64) *JsonPackageInfo {
	c.Lock()
	defer c.Unlock()
	p := &JsonPackageInfo{
		Path:  path,
		Hash:  hash,
		Types: &JsonTypes{m: map[string]*JsonTypeInfo{}},
	}
	c.m[path] = p
	return p
}

func (c *JsonPackages) Get(path string) (*JsonPackageInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	info, ok := c.m[path]
	return info, ok
}

func (c *JsonPackages) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *JsonTypes) Set(id string, t *JsonTypeInfo) {
	c.Lock()
	defer c.Unlock()
	c.m[id] = t
}

func (c *JsonTypes) Get(id string) (*JsonTypeInfo, bool) {
	c.RLock()
	defer c.RUnlock()
	if strings.HasPrefix(id, RULE_PREFIX) {
		panic(kerr.New("XULXFINYQJ", "Type name given to TypeCache.Get should not be a rule").Error())
	}
	t, ok := c.m[id]
	return t, ok
}

func (c *JsonTypes) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

func (c *JsonTypes) Keys() []string {
	out := []string{}
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (c *JsonDummies) Set(iface reflect.Type, dummy reflect.Type) {
	c.Lock()
	defer c.Unlock()
	c.m[iface] = dummy
}

func (c *JsonDummies) Get(iface reflect.Type) (reflect.Type, bool) {
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

// AutoContext creates a new context, and imports all types.
func AutoContext(ctx context.Context) context.Context {
	jc := newJsonCache()
	jc.Dummies.InitAuto()
	jc.Packages.InitAuto()
	return context.WithValue(ctx, jsonKey, jc)
}

// ManualContext creates a new context. This should only be used in internal tests. For normal
// usage, use AutoContext
func ManualContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, jsonKey, newJsonCache())
}

func newJsonCache() *JsonCache {
	return &JsonCache{
		Packages: &JsonPackages{m: map[string]*JsonPackageInfo{}},
		Dummies:  &JsonDummies{m: map[reflect.Type]reflect.Type{}},
	}
}

func (pc *JsonPackages) InitAuto() {
	// In automatic mode, we import all packages
	packages.RLock()
	defer packages.RUnlock()
	for path, pkg := range packages.m {
		pc.imp(path, pkg)
	}
}
func (pc *JsonPackages) InitManual(packagesToInit ...string) {
	// If we specify a manual list of packages to import, we should import them
	packages.RLock()
	defer packages.RUnlock()
	for _, path := range packagesToInit {
		pkg, ok := packages.m[path]
		if !ok {
			panic(fmt.Errorf("Package not found %s", path))
		}
		pc.imp(path, pkg)
	}
}
func (pc *JsonPackages) imp(path string, pkg *packageInfo) {
	p := pc.Set(path, pkg.hash)
	for name, typ := range pkg.types {
		p.Types.Set(name, &JsonTypeInfo{
			Name:  name,
			Type:  typ.typ,
			Rule:  typ.rule,
			Iface: typ.iface,
		})
	}
}

func (dc *JsonDummies) InitAuto() {
	dummies.RLock()
	defer dummies.RUnlock()
	for iface, dummy := range dummies.m {
		dc.Set(iface, dummy)
	}
}

func newContext(ctx context.Context, autoPackages bool, autoDummies bool, manualPackages ...string) context.Context {
	pc := &JsonPackages{m: map[string]*JsonPackageInfo{}}
	if autoPackages {
		pc.InitAuto()
	} else {
		pc.InitManual(manualPackages...)
	}

	dc := &JsonDummies{m: map[reflect.Type]reflect.Type{}}
	if autoDummies {
		dc.InitAuto()
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
		panic(kerr.New("XUTUUVDMMX", "No json cache in ctx").Error())
	}
	return e
}

func FromContextOrNil(ctx context.Context) *JsonCache {
	e, ok := ctx.Value(jsonKey).(*JsonCache)
	if ok {
		return e
	}
	return nil
}
