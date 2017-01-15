package generate

import (
	"context"

	"bytes"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
)

func ValidateCommand(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	f := NewFilePathName(env.Path, "main")
	f.Anon(env.Path)
	for _, p := range env.Aliases {
		f.Anon(p)
	}
	f.Func().Id("main").Params().Block(
		Id("kego.io/process/validate/command.ValidateMain").Call(Lit(env.Path)),
	)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("IIIRBBXASR", err)
	}
	return buf.Bytes(), nil
}
