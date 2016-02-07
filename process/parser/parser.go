package parser // import "kego.io/process/parser"

import (
	"fmt"

	"github.com/surge/cityhash"

	"path/filepath"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/process/packages"
	"kego.io/process/scanner"
	"kego.io/system"
)

func Parse(ctx context.Context, path string) (*sysctx.PackageInfo, error) {
	return parse(ctx, path, []string{})
}
func parse(ctx context.Context, path string, queue []string) (*sysctx.PackageInfo, error) {

	scache := sysctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	for _, q := range queue {
		if q == path {
			return nil, kerr.New("SCSCFJPPHD", "Circular import %v -> %v", queue, path)
		}
	}

	if _, found := scache.Get("kego.io/json"); !found {
		system.RegisterJsonTypes(ctx)
	}

	hash := &PackageHasher{Path: path, Aliases: map[string]string{}, Types: map[string]uint64{}}

	importPackage := func(importPath string, importAlias string) error {
		if _, found := scache.Get(importPath); !found {
			_, err := parse(ctx, importPath, append(queue, path))
			if err != nil {
				return kerr.Wrap("RIARRSCMVE", err)
			}
		}
		hash.Aliases[importPath] = importAlias
		return nil
	}

	env, err := ScanForEnv(ctx, path)
	if err != nil {
		return nil, kerr.Wrap("GJRHNGGWFD", err)
	}

	// Always scan the system package first if we don't have it already
	if path != "kego.io/system" {
		if err := importPackage("kego.io/system", "system"); err != nil {
			return nil, kerr.Wrap("ORRCDNUPOX", err)
		}
	}

	for aliasPath, aliasName := range env.Aliases {
		if aliasPath == "kego.io/system" || aliasName == "system" {
			return nil, kerr.New("EWMLNJDXKC", "Illegal import %s", aliasName)
		}
		if err := importPackage(aliasPath, aliasName); err != nil {
			return nil, kerr.Wrap("NOVMGYKHHI", err)
		}
	}

	pcache := scache.Set(env)

	cmd.Printf("Parsing %s...", path)

	if err := scanForTypes(ctx, env, pcache, hash); err != nil {
		return nil, kerr.Wrap("VFUNPHUFHD", err)
	}

	cmd.Println(" OK.")

	h, err := hash.Hash()
	if err != nil {
		return nil, kerr.Wrap("MIODRYNEJQ", err)
	}
	env.Hash = h

	return pcache, nil
}

func ScanForEnv(ctx context.Context, path string) (env *envctx.Env, err error) {
	// Scan the local directory for as system:package object
	dir, err := packages.GetDirFromPackage(ctx, path)
	if err != nil {
		return nil, kerr.Wrap("LASRFKILIH", err)
	}

	env = &envctx.Env{Path: path, Aliases: map[string]string{}, Recursive: false, Dir: dir}

	pkg, err := scanForPackage(ctx, env)
	if err != nil {
		return nil, kerr.Wrap("DTAEUSHJTQ", err)
	}

	if pkg != nil {
		env.Aliases = pkg.Aliases
		env.Recursive = pkg.Recursive
	}
	return env, nil
}

type objectStub struct {
	Id   *system.Reference `json:"id"`
	Type *system.Reference `json:"type"`
}

func scanForTypes(ctx context.Context, env *envctx.Env, cache *sysctx.PackageInfo, hash *PackageHasher) error {

	// While we're scanning for types, we should use a custom unpacking env, because the env from
	// the context is the one of the local package.

	files := scanner.ScanDirToFiles(ctx, env.Dir, env.Recursive)
	bytes := scanner.ScanFilesToBytes(ctx, files)
	localContext := envctx.NewContext(ctx, env)
	for b := range bytes {
		if b.Err != nil {
			return kerr.Wrap("JACKALTIGG", b.Err)
		}

		o := &objectStub{}
		if err := ke.UnmarshalUntyped(localContext, b.Bytes, o); err != nil {
			return kerr.Wrap("HCYGNBDFFA", err)
		}
		if o.Type == nil {
			return kerr.New("NUKWIHYFMQ", "%s has no type", b.File)
		}
		switch *o.Type {
		case *system.NewReference("kego.io/system", "type"):
			if err := ProcessTypeSourceBytes(ctx, env, b.Bytes, cache, hash); err != nil {
				return kerr.Wrap("IVEFDDSKHE", err)
			}
		case *system.NewReference("kego.io/system", "package"):
			cache.PackageBytes = b.Bytes
		default:
			if o.Id == nil {
				return kerr.New("DLLMKTDYFW", "%s has no id", b.File)
			}
			relativeFile, err := filepath.Rel(env.Dir, b.File)
			if err != nil {
				return kerr.Wrap("AWYRJSCYQS", err)
			}
			cache.Globals.Set(o.Id.Name, sysctx.GlobalInfo{
				File: relativeFile,
				Name: o.Id.Name,
			})
		}

	}
	return nil
}

func ProcessTypeSourceBytes(ctx context.Context, env *envctx.Env, bytes []byte, cache *sysctx.PackageInfo, hash *PackageHasher) error {
	var object interface{}
	err := json.Unmarshal(envctx.NewContext(ctx, env), bytes, &object)
	if err != nil {
		switch kerr.Source(err).(type) {
		case json.UnknownPackageError, json.UnknownTypeError:
			// don't return error
		default:
			return kerr.Wrap("NLRRVIDVWM", err)
		}
	}
	t, ok := object.(*system.Type)
	if !ok {
		return kerr.New("IVIFIOFGVK", "Should be *system.Type")
	}
	if hash != nil {
		hash.Types[t.Id.Name] = cityhash.CityHash64(bytes, uint32(len(bytes)))
	}
	cache.Types.Set(t.Id.Name, t)
	cache.TypeSource.Set(t.Id.Name, bytes)
	if t.Rule != nil {
		id := system.NewReference(t.Id.Package, fmt.Sprint("@", t.Id.Name))
		if t.Rule.Id != nil && *t.Rule.Id != *id {
			return kerr.New("JKARKEDTIW", "Incorrect id for %v - it should be %v", t.Rule.Id.String(), id.String())
		}
		t.Rule.Id = id

		// Check that the rule embeds system:rule
		found := false
		for _, em := range t.Rule.Embed {
			if *em == *system.NewReference("kego.io/system", "rule") {
				found = true
			}
		}
		if !found {
			return kerr.New("LMALEMKFDI", "%s does not embed system:rule", id.String())
		}

		cache.Types.Set(id.Name, t.Rule)
	} else {
		// If the rule is missing, automatically create a default.
		id := system.NewReference(t.Id.Package, fmt.Sprint("@", t.Id.Name))
		rule := &system.Type{
			Object: &system.Object{
				Description: fmt.Sprintf("Automatically created basic rule for %s", t.Id.Name),
				Type:        system.NewReference("kego.io/system", "type"),
				Id:          id,
			},
			Embed:     []*system.Reference{system.NewReference("kego.io/system", "rule")},
			Native:    system.NewString("object"),
			Interface: false,
		}
		cache.Types.Set(id.Name, rule)
	}

	return nil
}

func scanForPackage(ctx context.Context, env *envctx.Env) (*system.Package, error) {
	localContext := envctx.NewContext(ctx, env)
	files := scanner.ScanDirToFiles(ctx, env.Dir, false)
	bytes := scanner.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		if b.Err != nil {
			return nil, kerr.Wrap("GATNNQKNHY", b.Err)
		}
		o := &objectStub{}
		if err := ke.UnmarshalUntyped(localContext, b.Bytes, o); err != nil {
			return nil, kerr.Wrap("MTDCXBYBEJ", err)
		}
		if o.Type == nil {
			return nil, kerr.New("MSNIGTIDIO", "%s has no type", b.File)
		}
		switch *o.Type {
		case *system.NewReference("kego.io/system", "package"):
			var i interface{}
			err := json.Unmarshal(localContext, b.Bytes, &i)
			if err != nil {
				switch kerr.Source(err).(type) {
				case json.UnknownPackageError, json.UnknownTypeError:
				// don't return error
				default:
					return nil, kerr.Wrap("XTEQCAYQJP", err)
				}
			}
			if pkg, ok := i.(*system.Package); ok {
				return pkg, nil
			}
		}
	}
	return nil, nil
}
