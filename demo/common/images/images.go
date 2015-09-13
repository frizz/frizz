//go:generate kego -v
package images

import (
	"fmt"

	"strings"

	"kego.io/kerr"
)

type Image interface {
	GetUrl() string
}

func (i *Icon) GetUrl() string {
	return i.Url.Value
}

func (t *Photo) GetUrl() string {
	return fmt.Sprintf("%s://%s%s", t.Protocol.Value, t.Server.Value, t.Path.Value)
}

func (r *Image_rule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {
	i, ok := data.(Image)
	if !ok {
		return false, "", kerr.New("OSKCXRKIKC", nil, "Image_rule.Enforce", "data %T does not implement Image", data)
	}
	if r.Secure.Exists {
		url := i.GetUrl()
		secure := strings.HasPrefix(url, "https://")
		if r.Secure.Value && !secure {
			return false, fmt.Sprintf("Url %s must start with 'https://'", url), nil
		} else if !r.Secure.Value && secure {
			return false, fmt.Sprintf("Url %s must not start with 'https://'", url), nil
		}
	}
	return true, "", nil
}
