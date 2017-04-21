package generate

import (
	"context"

	"bytes"

	"frizz.io/context/envctx"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/kerr"
)

func Editor(ctx context.Context, env *envctx.Env) (source []byte, err error) {

	f := NewFile("main")
	f.Anon("frizz.io/system", env.Path)
	for _, p := range env.Aliases {
		f.Anon(p)
	}
	/*
		func main() {
			if err := {frizz.io/editor/client}.Start(); err != nil {
				{fmt}.Println(err.Error())
			}
		}
	*/
	f.Func().Id("main").Params().Block(
		If(
			Err().Op(":=").Qual("frizz.io/editor/client", "Start").Call(),
			Err().Op("!=").Nil(),
		).Block(
			Qual("fmt", "Println").Call(
				Err().Dot("Error").Call(),
			),
		),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("NEUSQXAQVL", err)
	}
	return buf.Bytes(), nil
}
