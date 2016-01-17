package system

import (
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/kerr"
)

type unpackerFunc func(ctx context.Context, data []byte, i *interface{}) error

func unmarshalFunc(ctx context.Context, data []byte, i *interface{}) error {
	return json.Unmarshal(ctx, data, i)
}
func unpackFunc(ctx context.Context, data []byte, i *interface{}) error {
	var j interface{}
	if err := json.UnmarshalPlain(data, &j); err != nil {
		return kerr.Wrap("WGRENSGSFM", err)
	}
	return json.Unpack(ctx, json.Pack(j), i)
}
