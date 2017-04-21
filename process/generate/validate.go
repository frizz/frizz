package generate

import (
	"context"

	"bytes"

	"frizz.io/context/envctx"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/kerr"
)

func ValidateCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	f := NewFilePathName(env.Path, "main")
	f.Anon(env.Path)
	f.Anon("frizz.io/system")
	for _, p := range env.Aliases {
		f.Anon(p)
	}
	f.Func().Id("main").Params().Block(
		Qual("frizz.io/process/validate/command", "ValidateMain").Call(Lit(env.Path)),
	)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("IIIRBBXASR", err)
	}
	return buf.Bytes(), nil
}
