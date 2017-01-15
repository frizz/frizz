package generate

import (
	"context"

	"bytes"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
)

func Editor(ctx context.Context, env *envctx.Env) (source []byte, err error) {

	f := NewFile("main")
	f.Anon("kego.io/system", env.Path)
	for _, p := range env.Aliases {
		f.Anon(p)
	}
	/*
		func main() {
			if err := {kego.io/editor/client}.Start(); err != nil {
				{fmt}.Println(err.Error())
			}
		}
	*/
	f.Func().Id("main").Params().Block(
		If(
			Err().Op(":=").Id("kego.io/editor/client.Start").Call(),
			Err().Op("!=").Nil(),
		).Block(
			Id("fmt.Println").Call(
				Id("err", "Error").Call(),
			),
		),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, kerr.Wrap("NEUSQXAQVL", err)
	}
	return buf.Bytes(), nil
}
