package ke // import "kego.io/ke"

// ke: {"package": {"complete": true}}

import (
	"io"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/json"
	"kego.io/process/scanner"
)

func Open(ctx context.Context, filename string) (value interface{}, err error) {
	bytes, err := scanner.ProcessFile(filename)
	if err != nil {
		return nil, kerr.Wrap("HPWXWFTKWA", err)
	}
	err = json.Unmarshal(ctx, bytes, &value)
	if err != nil {
		return nil, kerr.Wrap("CXIULJCEBE", err)
	}
	return
}

func Unmarshal(ctx context.Context, data []byte, v *interface{}) error {
	if err := json.Unmarshal(ctx, data, v); err != nil {
		return kerr.Wrap("SVXYHJWMOC", err)
	}
	return nil
}
func UnmarshalUntyped(ctx context.Context, data []byte, i interface{}) error {
	if err := json.UnmarshalUntyped(ctx, data, i); err != nil {
		return kerr.Wrap("HOPCKQEJFM", err)
	}
	return nil
}

func NewDecoder(ctx context.Context, r io.Reader) *json.Decoder {
	return json.NewDecoder(ctx, r)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func Marshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, kerr.Wrap("LXDTUOBQPD", err)
	}
	return b, nil
}
func MarshalContext(ctx context.Context, v interface{}) ([]byte, error) {
	b, err := json.MarshalContext(ctx, v)
	if err != nil {
		return nil, kerr.Wrap("XMHXROGFXM", err)
	}
	return b, nil
}
func MarshalIndentContext(ctx context.Context, v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndentContext(ctx, v, prefix, indent)
}

func NewContext(ctx context.Context, path string, aliases map[string]string) context.Context {
	ctx = envctx.NewContext(ctx, &envctx.Env{Path: path, Aliases: aliases})
	ctx = jsonctx.AutoContext(ctx)
	return ctx
}
