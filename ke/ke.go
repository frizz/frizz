package ke // import "kego.io/ke"

// ke: {"package": {"complete": true}}

import (
	"context"

	"encoding/json"

	"github.com/davelondon/kerr"
	"kego.io/process/scanner"
	"kego.io/system"
)

func Open(ctx context.Context, filename string) (value interface{}, err error) {
	bytes, err := scanner.ProcessFile(filename)
	if err != nil {
		return nil, kerr.Wrap("HPWXWFTKWA", err)
	}
	err = system.Unmarshal(ctx, bytes, &value)
	if err != nil {
		return nil, kerr.Wrap("CXIULJCEBE", err)
	}
	return
}

func Unmarshal(ctx context.Context, data []byte, v interface{}) error {
	if err := system.Unmarshal(ctx, data, v); err != nil {
		return kerr.Wrap("SVXYHJWMOC", err)
	}
	return nil
}

func Marshal(ctx context.Context, v interface{}) ([]byte, error) {
	i, err := system.Repack(ctx, v)
	if err != nil {
		return nil, kerr.Wrap("JVGOWMVMXN", err)
	}
	b, err := json.Marshal(i)
	if err != nil {
		return nil, kerr.Wrap("LXDTUOBQPD", err)
	}
	return b, nil
}

func MarshalIndent(ctx context.Context, v interface{}, prefix, indent string) ([]byte, error) {
	i, err := system.Repack(ctx, v)
	if err != nil {
		return nil, kerr.Wrap("QNHVVOBKFO", err)
	}
	b, err := json.MarshalIndent(i, prefix, indent)
	if err != nil {
		return nil, kerr.Wrap("QXXEBOMOFG", err)
	}
	return b, nil
}

func NewContext(ctx context.Context, path string, aliases map[string]string) context.Context {
	return system.NewContext(ctx, path, aliases)
}
