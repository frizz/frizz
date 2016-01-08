package ke // import "kego.io/ke"

import (
	"io"
	"os"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/json"
	"kego.io/kerr"
)

func Open(ctx context.Context, filename string) (value interface{}, err error) {

	file, err := os.Open(filename)
	if err != nil {
		err = kerr.New("NDJKHCDCIW", err, "os.Open")
		return
	}
	defer file.Close()

	err = json.NewDecoder(ctx, file).Decode(&value)
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

func NewContext(ctx context.Context) context.Context {
	return jsonctx.NewContext(ctx)
}
