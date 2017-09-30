package scanner

import (
	"go/build"
	"io/ioutil"

	"os"

	"strings"

	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

type Scanner struct {
	env  vos.Env
	pkg  string
	prog *loader.Program
}

func New(pkg string, env vos.Env) *Scanner {
	return &Scanner{
		pkg: pkg,
		env: env,
	}
}

func (s *Scanner) Scan() error {
	conf := loader.Config{}
	conf.Import(s.pkg)
	conf.Build = &build.Default
	conf.Build.GOPATH = s.env.Getenv("GOPATH")

	// We must exclude all generated files from the scan. This is because when a user makes changes
	// to the signature of a type, the generated files would become invalid and go/types would
	// throw an error when scanning. Additionally, performance is increased by preventing the
	// scanning of the complex generated code. The results of scanning this code would never be
	// needed during the generation process.
	conf.Build.ReadDir = func(dir string) ([]os.FileInfo, error) {

		raw, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		var list []os.FileInfo
		for _, fi := range raw {
			if !strings.HasSuffix(fi.Name(), "generated.frizz.go") {
				list = append(list, fi)
			}
		}
		return list, nil
	}

	// Allow errors to ensure any references to the excluded generated file don't cause errors.
	conf.AllowErrors = true
	conf.TypeChecker.Error = func(e error) {}

	prog, err := conf.Load()
	if err != nil {
		return errors.WithStack(err)
	}
	s.prog = prog
	return nil
}
