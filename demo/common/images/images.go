//go:generate frizz
package images

// notest

import (
	"fmt"

	"context"

	"strings"

	"github.com/dave/kerr"
)

type Image interface {
	GetUrl() string
}

func (i *Icon) GetUrl() string {
	return i.Url.Value()
}

func (i *Icon) Label(ctx context.Context) string {
	if i.Url == nil {
		return ""
	}
	return i.Url.Value()[strings.LastIndex(i.Url.Value(), "/")+1:]
}

func (i *Photo) Label(ctx context.Context) string {
	if i.Path == nil {
		return ""
	}
	return i.Path.Value()[strings.LastIndex(i.Path.Value(), "/")+1:]
}

func (t *Photo) GetUrl() string {
	return fmt.Sprintf("%s://%s%s", t.Protocol.Value(), t.Server.Value(), t.Path.Value())
}

func (r *ImageRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	i, ok := data.(Image)
	if !ok {
		return true, nil, kerr.New("OSKCXRKIKC", "ImageRule.Enforce", "data %T does not implement Image", data)
	}
	if r.Secure != nil {
		url := i.GetUrl()
		secure := strings.HasPrefix(url, "https://")
		if r.Secure.Value() && !secure {
			fail = true
			messages = append(messages, fmt.Sprintf("Url %s must start with 'https://'", url))
		} else if !r.Secure.Value() && secure {
			fail = true
			messages = append(messages, fmt.Sprintf("Url %s must not start with 'https://'", url))
		}
	}
	return
}
