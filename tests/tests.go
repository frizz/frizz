package tests // import "kego.io/tests"

// ke: {"package": {"notest": true}}

import (
	"reflect"
	"sync"

	"strings"

	"fmt"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/context/vosctx"
	"kego.io/context/wgctx"
)

type ContextBuilder struct {
	ctx               context.Context
	gopathInitialized bool
	tempNamespace     string
	tempPackageDir    string
	tempDirs          []string
	cancel            context.CancelFunc
}

func New() *ContextBuilder {
	return From(context.Background())
}

func From(ctx context.Context) *ContextBuilder {
	ctx, cancel := context.WithCancel(ctx)
	cb := &ContextBuilder{ctx: ctx, cancel: cancel}
	return cb
}

func Context(path string) *ContextBuilder {

	env := &envctx.Env{
		Path:    path,
		Aliases: map[string]string{},
	}

	ctx := context.Background()
	ctx = envctx.NewContext(ctx, env)

	return From(ctx)
}

func (c *ContextBuilder) Cancel() *ContextBuilder {
	c.cancel()
	return c
}

type Printer interface {
	Print(ctx context.Context) string
}

func (c *ContextBuilder) Println(args ...interface{}) {
	var args1 []interface{}
	for _, a := range args {
		if n, ok := a.(Printer); ok {
			args1 = append(args1, n.Print(c.Ctx()))
			continue
		}
		args1 = append(args1, a)
		continue
	}
	fmt.Println(args1...)
}

func (c *ContextBuilder) Ctx() context.Context {
	return c.ctx
}
func (c *ContextBuilder) SetCtx(ctx context.Context) *ContextBuilder {
	c.ctx = ctx
	return c
}
func (c *ContextBuilder) Env() *envctx.Env {
	return envctx.FromContext(c.ctx)
}

func (c *ContextBuilder) initEnv() *envctx.Env {
	env := envctx.FromContextOrNil(c.ctx)
	if env == nil {
		env = &envctx.Env{Aliases: map[string]string{}}
		c.ctx = envctx.NewContext(c.ctx, env)
	}
	return env
}

func (c *ContextBuilder) initCmd() *cmdctx.Cmd {
	cmd := cmdctx.FromContextOrNil(c.ctx)
	if cmd == nil {
		cmd = &cmdctx.Cmd{}
		c.ctx = cmdctx.NewContext(c.ctx, cmd)
	}
	return cmd
}

func (c *ContextBuilder) initWg() *sync.WaitGroup {
	wg := wgctx.FromContextOrNil(c.ctx)
	if wg != nil {
		return wg
	}
	c.ctx = wgctx.NewContext(c.ctx)
	return wgctx.FromContext(c.ctx)
}

func (c *ContextBuilder) initJson() *jsonctx.JsonCache {
	jcache := jsonctx.FromContextOrNil(c.ctx)
	if jcache == nil {
		c.ctx = jsonctx.ManualContext(c.ctx)
		jcache = jsonctx.FromContext(c.ctx)
	}
	return jcache
}

func (c *ContextBuilder) initSys() *sysctx.SysCache {
	scache := sysctx.FromContextOrNil(c.ctx)
	if scache == nil {
		c.ctx = sysctx.NewContext(c.ctx)
		scache = sysctx.FromContext(c.ctx)
	}
	return scache
}
func (c *ContextBuilder) initSysPkg(path string) *sysctx.SysPackageInfo {
	scache := c.initSys()
	pi, ok := scache.Get(path)
	if !ok {
		pi = scache.Set(path)
	}
	return pi
}
func (c *ContextBuilder) initVos() vosctx.Vos {
	vos := vosctx.FromContextOrNil(c.ctx)
	if vos != nil {
		return vos
	}
	vos = &MockOs{
		EnvironmentVariables: map[string]string{},
	}
	c.ctx = vosctx.NewContext(c.ctx, vos)
	return vos
}

func (c *ContextBuilder) Cmd() *ContextBuilder {
	c.initCmd()
	return c
}

func (c *ContextBuilder) CmdUpdate(value bool) *ContextBuilder {
	cmd := c.initCmd()
	cmd.Update = value
	return c
}

func (c *ContextBuilder) Wg() *ContextBuilder {
	c.initWg()
	return c
}
func (c *ContextBuilder) Log() *ContextBuilder {
	cmd := c.initCmd()
	cmd.Log = true
	return c
}

func (c *ContextBuilder) Path(path string) *ContextBuilder {
	env := c.initEnv()
	env.Path = path
	//pi := c.initSysPkg(path)
	//copyEnv(env, pi.Env)
	return c
}

func (c *ContextBuilder) Alias(aliasName string, aliasPath string) *ContextBuilder {
	env := c.initEnv()
	env.Aliases[aliasName] = aliasPath
	//if env.Path != "" {
	//	pi := c.initSysPkg(env.Path)
	//	copyEnv(env, pi.Env)
	//}
	return c
}

func copyEnv(from *envctx.Env, to *envctx.Env) {
	to.Path = from.Path
	to.Recursive = from.Recursive
	to.Hash = from.Hash
	to.Aliases = map[string]string{}
	for n, p := range from.Aliases {
		to.Aliases[n] = p
	}
}

func (c *ContextBuilder) Dir(dir string) *ContextBuilder {
	env := c.initEnv()
	env.Dir = dir
	return c
}

func (c *ContextBuilder) Recursive(state bool) *ContextBuilder {
	env := c.initEnv()
	env.Recursive = state
	return c
}

func (c *ContextBuilder) Jempty() *ContextBuilder {
	c.initJson()
	return c
}

func (c *ContextBuilder) Jauto() *ContextBuilder {
	jcache := c.initJson()
	jcache.Packages.InitAuto()
	jcache.Dummies.InitAuto()
	return c
}

func (c *ContextBuilder) Jsystem() *ContextBuilder {
	jcache := c.initJson()
	jcache.Packages.InitManual("kego.io/json", "kego.io/system")
	jcache.Dummies.InitAuto()
	return c
}

func (c *ContextBuilder) Jtype(name string, typ reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRule(env.Path, name, typ, nil)
}
func (c *ContextBuilder) JtypeIface(name string, typ reflect.Type, iface reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRuleIface(env.Path, name, typ, nil, iface)
}
func (c *ContextBuilder) Jiface(name string, iface reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRuleIface(env.Path, name, nil, nil, iface)
}
func (c *ContextBuilder) JtypeRule(name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRule(env.Path, name, typ, rule)
}
func (c *ContextBuilder) JtypePath(path string, name string, typ reflect.Type) *ContextBuilder {
	return c.JtypePathRule(path, name, typ, nil)
}

func (c *ContextBuilder) JtypePathRule(path string, name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	return c.JtypePathRuleIface(path, name, typ, rule, nil)
}

func (c *ContextBuilder) Jpkg(path string, hash uint64) *ContextBuilder {
	if path == "" {
		panic("must specify path")
	}
	jcache := c.initJson()
	p, ok := jcache.Packages.Get(path)
	if ok {
		p.Hash = hash
		return c
	}
	jcache.Packages.Set(path, hash)
	return c
}

func (c *ContextBuilder) JtypePathRuleIface(path string, name string, typ reflect.Type, rule reflect.Type, iface reflect.Type) *ContextBuilder {
	if path == "" {
		panic("must specify path")
	}

	jcache := c.initJson()

	p, ok := jcache.Packages.Get(path)
	if !ok {
		p = jcache.Packages.Set(path, 0)
	}
	isrule := false
	if strings.HasPrefix(name, jsonctx.RULE_PREFIX) {
		if rule != nil {
			panic(kerr.New("UBFYEAGXHJ", "rule specified!").Error())
		}
		isrule = true
		name = name[1:]
	}
	if isrule {
		p.Types.Set(name, &jsonctx.JsonTypeInfo{
			Name:  name,
			Rule:  typ,
			Iface: iface,
		})
	} else {
		p.Types.Set(name, &jsonctx.JsonTypeInfo{
			Name:  name,
			Type:  typ,
			Rule:  rule,
			Iface: iface,
		})
	}

	return c

}

func (c *ContextBuilder) Sempty() *ContextBuilder {
	c.initSys()
	return c
}

func (c *ContextBuilder) Spkg(path string) *ContextBuilder {
	scache := c.initSys()
	if _, ok := scache.Get(path); !ok {
		scache.Set(path)
	}
	return c
}

// Sauto runs parser.Parse, but to stop an import cycle we pass the parser.Parse function in as a
// parameter
func (c *ContextBuilder) Sauto(parseParse func(context.Context, string) (*sysctx.SysPackageInfo, error)) *ContextBuilder {
	env := c.initEnv()
	c.initSys()
	c.initCmd()
	pi, err := parseParse(c.ctx, env.Path)
	if err != nil {
		panic(err.Error())
	}
	env.Hash = pi.Hash
	return c
}

// Ssystem runs parser.Parse on the system package, but to stop an import cycle we pass the
// parser.Parse function in as a parameter
func (c *ContextBuilder) Ssystem(parseParse func(context.Context, string) (*sysctx.SysPackageInfo, error)) *ContextBuilder {
	c.initEnv()
	c.initSys()
	c.initCmd()
	c.Jsystem() // need json types for system in order to parse
	_, err := parseParse(c.ctx, "kego.io/system")
	if err != nil {
		panic(err.Error())
	}
	return c
}

func (c *ContextBuilder) Stype(name string, typ interface{}) *ContextBuilder {
	env := c.initEnv()
	return c.StypePath(env.Path, name, typ)
}

func (c *ContextBuilder) StypePath(path string, name string, typ interface{}) *ContextBuilder {

	if path == "" {
		panic("path not specified")
	}

	scache := c.initSys()

	p, ok := scache.Get(path)
	if !ok {
		p = scache.Set(path)
	}

	p.Types.Set(name, "", typ)

	return c
}

func (c *ContextBuilder) AllPath(path string, name string, typ reflect.Type, rule reflect.Type, typTyp interface{}, ruleTyp interface{}) *ContextBuilder {
	c.JtypePathRule(path, name, typ, rule)
	c.StypePath(path, name, typTyp)
	c.StypePath(path, jsonctx.RULE_PREFIX+name, ruleTyp)
	return c
}

func (c *ContextBuilder) Dummy(iface reflect.Type, dummy reflect.Type) *ContextBuilder {
	jcache := c.initJson()
	jcache.Dummies.Set(iface, dummy)
	return c
}

func (c *ContextBuilder) OsVar(name string, value string) *ContextBuilder {
	vos := c.initVos()
	m, ok := vos.(*MockOs)
	if !ok {
		panic("must me *MockEnv")
	}
	m.EnvironmentVariables[name] = value
	return c
}

func (c *ContextBuilder) OsWd(dir string) *ContextBuilder {
	vos := c.initVos()
	m, ok := vos.(*MockOs)
	if !ok {
		panic("must me *MockEnv")
	}
	m.WorkingDirectory = dir
	return c
}
