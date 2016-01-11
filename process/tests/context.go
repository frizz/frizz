package tests

import (
	"reflect"
	"sync"

	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
)

type ContextBuilder struct {
	ctx       context.Context
	env       *envctx.Env
	cmd       *cmdctx.Cmd
	cache     *cachectx.PackageCache
	cacheInfo *cachectx.PackageInfo
	jcache    *jsonctx.JsonCache
	wg        *sync.WaitGroup
}

func Context(path string) *ContextBuilder {

	env := &envctx.Env{
		Path:    path,
		Aliases: map[string]string{},
	}

	ctx := context.Background()
	ctx = envctx.NewContext(ctx, env)

	return &ContextBuilder{ctx: ctx, env: env}
}

func ContextFrom(ctx context.Context) *ContextBuilder {
	cb := &ContextBuilder{
		ctx:    ctx,
		env:    envctx.FromContextOrNil(ctx),
		cmd:    cmdctx.FromContextOrNil(ctx),
		cache:  cachectx.FromContextOrNil(ctx),
		jcache: jsonctx.FromContextOrNil(ctx),
	}
	if cb.env != nil && cb.cache != nil {
		ci, ok := cb.cache.Get(cb.env.Path)
		if ok {
			cb.cacheInfo = ci
		}
	}
	return cb
}

func (c *ContextBuilder) Ctx() context.Context {
	return c.ctx
}

func (c *ContextBuilder) Path(path string) *ContextBuilder {
	c.env.Path = path
	return c
}

func (c *ContextBuilder) Alias(aliasPath string, aliasName string) *ContextBuilder {

	c.env.Aliases[aliasPath] = aliasName
	if c.cacheInfo != nil {
		c.cacheInfo.Environment.Aliases[aliasPath] = aliasName
	}

	return c
}

func (c *ContextBuilder) Cmd() *ContextBuilder {

	if c.cmd != nil {
		return c
	}

	c.ctx = cmdctx.NewContext(c.ctx, &cmdctx.Cmd{})
	c.cmd = cmdctx.FromContext(c.ctx)

	return c
}

func (c *ContextBuilder) Wg() *ContextBuilder {

	if c.wg != nil {
		return c
	}

	c.ctx = wgctx.NewContext(c.ctx)
	c.wg = wgctx.FromContext(c.ctx)

	return c
}

func (c *ContextBuilder) Dir(dir string) *ContextBuilder {

	if c.cmd == nil {
		c.Cmd()
	}

	c.cmd.Dir = dir

	return c

}

func (c *ContextBuilder) Json() *ContextBuilder {

	if c.jcache != nil {
		return c
	}

	c.ctx = jsonctx.NewContext(c.ctx)
	c.jcache = jsonctx.FromContext(c.ctx)

	return c
}

func (c *ContextBuilder) Jauto() *ContextBuilder {
	c.ctx = jsonctx.NewContext(c.ctx)
	return c
}

func (c *ContextBuilder) Jtype(name string, typ reflect.Type) *ContextBuilder {
	return c.JtypePath(c.env.Path, name, typ, nil)
}
func (c *ContextBuilder) JtypeRule(name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	return c.JtypePathRule(c.env.Path, name, typ, rule)
}
func (c *ContextBuilder) JtypePath(path string, name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	return c.JtypePathRule(path, name, typ, nil)
}

func (c *ContextBuilder) JtypePathRule(path string, name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {

	if c.jcache == nil {
		c.Json()
	}

	p, ok := c.jcache.Packages.Get(path)
	if !ok {
		p = c.jcache.Packages.Set(path, 0)
	}
	isrule := false
	if strings.HasPrefix(name, jsonctx.RULE_PREFIX) {
		if rule != nil {
			panic(kerr.New("UBFYEAGXHJ", nil, "rule specified!"))
		}
		isrule = true
		name = name[1:]
	}
	if isrule {
		p.Types.Set(name, &jsonctx.TypeInfo{
			Name: name,
			Rule: typ,
		})
	} else {
		p.Types.Set(name, &jsonctx.TypeInfo{
			Name: name,
			Type: typ,
			Rule: rule,
		})
	}

	return c

}

func (c *ContextBuilder) Cache() *ContextBuilder {

	if c.cache != nil {
		return c
	}

	c.ctx = cachectx.NewContext(c.ctx)
	c.cache = cachectx.FromContext(c.ctx)
	c.cacheInfo = c.cache.Set(c.env)

	return c
}

// Cauto runs parse.Parse, but to stop an import cycle we pass the parse.Parse function in as a
// parameter
func (c *ContextBuilder) Cauto(p func(context.Context, string, []string) (*cachectx.PackageInfo, error)) *ContextBuilder {
	if c.cache == nil {
		c.Cache()
	}
	if c.cmd == nil {
		c.Cmd()
	}
	pcache, err := p(c.ctx, c.env.Path, []string{})
	if err != nil {
		panic(kerr.New("KCEEBYAXGB", err, "parse.Parse"))
	}
	c.cacheInfo = pcache
	return c
}

func (c *ContextBuilder) Ctype(name string, typ interface{}) *ContextBuilder {
	return c.CtypePath(c.env.Path, name, typ)
}

func (c *ContextBuilder) CtypePath(path string, name string, typ interface{}) *ContextBuilder {

	if c.cache == nil {
		c.Cache()
	}

	p, ok := c.cache.Get(path)
	if !ok {
		p = c.cache.Set(&envctx.Env{Path: path})
	}

	p.Types.Set(name, typ)

	return c
}

func (c *ContextBuilder) AllPath(path string, name string, typ reflect.Type, rule reflect.Type, typTyp interface{}, ruleTyp interface{}) *ContextBuilder {
	c.JtypePathRule(path, name, typ, rule)
	c.CtypePath(path, name, typTyp)
	c.CtypePath(path, jsonctx.RULE_PREFIX+name, ruleTyp)
	return c
}

func (c *ContextBuilder) Dummy(iface reflect.Type, dummy reflect.Type) *ContextBuilder {

	if c.jcache == nil {
		c.Json()
	}
	c.jcache.Dummies.Set(iface, dummy)

	return c
}
