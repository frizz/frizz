package pkghelp

import (
	"path/filepath"

	"context"

	"strings"

	"frizz.io/context/envctx"
	"frizz.io/context/sysctx"
	"frizz.io/process/parser"
	"frizz.io/process/scanner"
	"frizz.io/system"
	"github.com/dave/kerr"
)

type Package struct {
	Env   *envctx.Env
	Files []*File
}

type File struct {
	Package          *Package
	Id               *system.Reference // Id extracted from the id json property
	Type             *system.Reference // Type extracted from the type json property
	Filename         string            // Filename e.g. baz.yaml
	Extension        string            // Extension including dot e.g. .yaml
	Yaml             bool              // Is the extension yml or yaml?
	RelativeFilepath string            // Filename relative to the path root e.g. bar/baz.yaml
	AbsoluteFilepath string            // Absolute filepath e.g. /Users/foo/go/src/github.com/foo/bar/baz.yaml
	Directory        string            // Absolute directory path e.g. /Users/foo/go/src/github.com/foo/bar/
	IsInRoot         bool              // True if the file in the root of the package. False if it's in a subdir.
}

func Scan(ctx context.Context, path string) (*Package, error) {
	p := new(Package)
	// use a new system context for the duration of the scan
	ctx = sysctx.NewContext(ctx)
	env, err := parser.ScanForEnv(ctx, path)
	if err != nil {
		return nil, kerr.Wrap("SFPSTLPKMX", err)
	}
	p.Env = env
	files := scanner.ScanDirToFiles(ctx, env.Dir, env.Recursive)
	bytes := scanner.ScanFilesToBytes(ctx, files)
	localContext := envctx.NewContext(ctx, env)
	for b := range bytes {
		f := new(File)
		f.Package = p
		if b.Err != nil {
			return nil, kerr.Wrap("IPUHPBBWEA", b.Err)
		}
		o := &system.ObjectStub{}
		if err := system.Unmarshal(localContext, b.Bytes, o); err != nil {
			return nil, kerr.Wrap("DCAGIDLXRT", err)
		}
		if o.Type == nil {
			return nil, kerr.New("AJPCEQTWPS", "%s has no type", b.File)
		}
		if o.Id == nil && *o.Type != *system.NewReference("frizz.io/system", "package") {
			// we tolerate missing ID only for system:package
			return nil, kerr.New("YDKYLXTGYC", "%s has no id", b.File)
		}
		f.AbsoluteFilepath = b.File
		f.Type = o.Type
		f.Id = o.Id
		if f.RelativeFilepath, err = filepath.Rel(env.Dir, b.File); err != nil {
			return nil, kerr.Wrap("QDAEGOWTWP", err)
		}
		f.Directory, f.Filename = filepath.Split(b.File)
		f.IsInRoot = DirEqual(f.Directory, env.Dir)
		f.Extension = filepath.Ext(b.File)
		f.Yaml = f.Extension == ".yaml" || f.Extension == ".yml"
		p.Files = append(p.Files, f)
	}
	return p, nil
}

// Dir can be expressed as "/foo/" and also "/foo".
func DirEqual(a, b string) bool {
	var err error
	a = filepath.Clean(a)
	b = filepath.Clean(b)
	if a, err = filepath.Abs(a); err != nil {
		return false
	}
	if b, err = filepath.Abs(b); err != nil {
		return false
	}
	aDirs := strings.Split(a, string(filepath.Separator))
	bDirs := strings.Split(b, string(filepath.Separator))
	if len(aDirs) != len(bDirs) {
		return false
	}
	for i, aVal := range aDirs {
		if aVal != bDirs[i] {
			return false
		}
	}
	return true
}

func (p *Package) File(rel string) *File {
	for _, f := range p.Files {
		if rel == f.RelativeFilepath {
			return f
		}
	}
	return nil
}

func Check(root, rel string, recursive bool) (string, error) {
	rel = filepath.Clean(rel)

	// Join the relative filename to the root to get the full filepath.
	full := filepath.Join(root, rel)

	// This may have ".." parts, so we must check it hasn't escaped
	rel, err := filepath.Rel(root, full)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", kerr.New("CNMFTLYQFG", "Error in filename %s", rel)
	}

	dir, _ := filepath.Split(full)
	if !recursive && !DirEqual(dir, root) {
		return "", kerr.New("OABDMRQSJO", "%s should be in the root dir if package is not recursive", rel)
	}

	return full, nil
}
