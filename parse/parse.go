package parse // import "kego.io/parse"

import (
	"fmt"

	"github.com/surge/cityhash"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/process/pkgutils"
	"kego.io/process/scanutils"
	"kego.io/system"
)

func Parse(ctx context.Context, path string, queue []string) (*envctx.Env, error) {

	cache := cachectx.FromContext(ctx)

	for _, q := range queue {
		if q == path {
			return nil, kerr.New("SCSCFJPPHD", nil, "Circular import %v -> %v", queue, path)
		}
	}

	if _, found := cache.Get("kego.io/json"); !found {
		system.RegisterJsonTypes(ctx)
	}

	hash := &PackageHasher{Path: path, Aliases: map[string]string{}, Types: map[string]uint64{}}

	importPackage := func(importPath string, importAlias string) error {
		//if aliasInfo, found := cache.Get(importPath); !found {
		if _, found := cache.Get(importPath); !found {
			//aliasEnv, err := Parse(ctx, importPath, append(queue, path))
			_, err := Parse(ctx, importPath, append(queue, path))
			if err != nil {
				return kerr.New("RIARRSCMVE", err, "Parse (%v)", importPath)
			}
			//hash.Aliases[importPath] = AliasInfo{Alias: importAlias, Hash: aliasEnv.Hash}
		} else {
			//hash.Aliases[importPath] = AliasInfo{Alias: importAlias, Hash: aliasInfo.Environment.Hash}
		}
		hash.Aliases[importPath] = importAlias
		return nil
	}

	// Always scan the system package first if we don't have it already
	if path != "kego.io/system" {
		if err := importPackage("kego.io/system", "system"); err != nil {
			return nil, kerr.New("ORRCDNUPOX", err, "importPackage")
		}
	}

	// Scan the local directory for as system:package object
	dir, err := pkgutils.GetDirFromPackage(path)
	if err != nil {
		return nil, kerr.New("LASRFKILIH", err, "pkgutils.GetDirFromPackage")
	}

	pkg, err := scanForPackage(ctx, path, dir)
	if err != nil {
		return nil, kerr.New("DTAEUSHJTQ", err, "scanForPackage")
	}

	env := &envctx.Env{Path: path, Aliases: map[string]string{}, Recursive: false}
	if pkg != nil {
		for aliasPath, aliasName := range pkg.Aliases {
			if aliasPath == "kego.io/system" || aliasName == "system" {
				return nil, kerr.New("EWMLNJDXKC", nil, "Illegal import %s", aliasName)
			}
			if err := importPackage(aliasPath, aliasName); err != nil {
				return nil, kerr.New("NOVMGYKHHI", err, "importPackage")
			}
		}
		env.Aliases = pkg.Aliases
		env.Recursive = pkg.Recursive
	}

	pcache := cache.Set(env)

	if err := scanForTypes(ctx, path, dir, env, pcache, hash); err != nil {
		return nil, kerr.New("VFUNPHUFHD", err, "scanForTypes")
	}

	h, err := hash.Hash()
	if err != nil {
		return nil, kerr.New("MIODRYNEJQ", err, "hash.Hash()")
	}
	env.Hash = h

	return env, nil
}

func scanForTypes(ctx context.Context, path string, dir string, env *envctx.Env, cache *cachectx.PackageInfo, hash *PackageHasher) error {

	// While we're scanning for types, we should use a custom unpacking env, because the env from
	// the context is the one of the local package.

	files := scanutils.ScanDirToFiles(ctx, dir, env.Recursive)
	bytes := scanutils.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		if b.Err != nil {
			return kerr.New("JACKALTIGG", b.Err, "ScanFiles")
		}
		var object interface{}
		err := json.Unmarshal(envctx.NewContext(ctx, env), b.Bytes, &object)
		if err != nil {
			switch err.(type) {
			case json.UnknownPackageError, json.UnknownTypeError:
				// don't return error
			default:
				return kerr.New("NLRRVIDVWM", err, "UnmarshalPlain")
			}
		}
		if t, ok := object.(*system.Type); ok {
			hash.Types[t.Id.Name] = cityhash.CityHash64(b.Bytes, uint32(len(b.Bytes)))
			cache.Types.Set(t.Id.Name, t)
			if t.Rule != nil {
				id := system.NewReference(t.Id.Package, fmt.Sprint("@", t.Id.Name))
				if t.Rule.Id != nil && *t.Rule.Id != *id {
					return kerr.New("JKARKEDTIW", nil, "Incorrect id for %v - it should be %v", t.Rule.Id.String(), id.String())
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
					return kerr.New("LMALEMKFDI", nil, "%s does not embed system:rule", id.String())
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
		}
	}
	return nil
}

/*
func scanForGlobals(ctx context.Context, path string, dir string, env *envctx.Env, cache *cachectx.PackageInfo) error {
	files := scanutils.ScanDirToFiles(ctx, dir, env.Recursive)
	bytes := scanutils.ScanFilesToBytes(ctx, nil, files)
	for b := range bytes {
		if b.Err != nil {
			return kerr.New("XAHQAPYRIO", b.Err, "ScanFiles (%s)", b.File)
		}
		n, err := node.Unmarshal(envctx.NewContext(ctx, env), b.Bytes)
		if err != nil {
			return kerr.New("YGYMJMLYED", err, "node.Unmarshal (%s)", b.File)
		}
		for _, sub := range n.Flatten(true) {
			if sub.Type.Native.Value() != "object" {
				continue
			}
			idi, ok := sub.Map["id"]
			if !ok {
				continue
			}
			id := idi.GetNode()
			if id.Missing || id.Null {
				continue
			}
			if *id.Origin != *system.NewReference("kego.io/system", "object") {
				continue
			}
			cache.Nodes.Set(id.ValueString, sub)
		}
	}
	return nil
}
*/

func scanForPackage(ctx context.Context, path string, dir string) (*system.Package, error) {
	files := scanutils.ScanDirToFiles(ctx, dir, false)
	bytes := scanutils.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		if b.Err != nil {
			return nil, kerr.New("GATNNQKNHY", b.Err, "ScanFiles")
		}
		var object interface{}
		err := json.Unmarshal(envctx.Dummy(path, map[string]string{}), b.Bytes, &object)
		if err != nil {
			switch err.(type) {
			case json.UnknownPackageError, json.UnknownTypeError:
				// don't return error
			default:
				return nil, kerr.New("XTEQCAYQJP", err, "UnmarshalPlain")
			}
		}
		if pkg, ok := object.(*system.Package); ok {
			return pkg, nil
		}
	}
	return nil, nil
}

func HasSourceFiles(ctx context.Context) bool {
	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)
	files := scanutils.ScanDirToFiles(ctx, cmd.Dir, env.Recursive)
	bytes := scanutils.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		if b.Err != nil {
			continue
		}
		return true
	}
	return false

}

type InfoStruct struct {
	Path string
	Hash uint64
}
