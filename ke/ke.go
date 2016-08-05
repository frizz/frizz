package ke // import "kego.io/ke"

// ke: {"package": {"complete": true}}

import (
	"io"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
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
	return json.Unmarshal(ctx, data, v)
}
func UnmarshalUntyped(ctx context.Context, data []byte, i interface{}) error {
	return json.UnmarshalUntyped(ctx, data, i)
}

func NewDecoder(ctx context.Context, r io.Reader) *json.Decoder {
	return json.NewDecoder(ctx, r)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func MarshalContext(ctx context.Context, v interface{}) ([]byte, error) {
	return json.MarshalContext(ctx, v)
}
func MarshalIndentContext(ctx context.Context, v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndentContext(ctx, v, prefix, indent)
}

func NewContext(ctx context.Context, path string, aliases map[string]string) context.Context {
	ctx = envctx.NewContext(ctx, &envctx.Env{Path: path, Aliases: aliases})
	ctx = jsonctx.AutoContext(ctx)
	return ctx
}
