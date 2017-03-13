package generate

import (
	"context"

	"bytes"

	. "github.com/dave/jennifer/jen"
	"github.com/dave/kerr"
	"kego.io/context/envctx"
)

func ValidateCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	f := NewFilePathName(env.Path, "main")
	f.Anon(env.Path)
	f.Anon("kego.io/system")
	for _, p := range env.Aliases {
		f.Anon(p)
	}
	f.Func().Id("main").Params().Block(
		Qual("kego.io/process/validate/command", "ValidateMain").Call(Lit(env.Path)),
	)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("IIIRBBXASR", err)
	}
	return buf.Bytes(), nil
}
