//go:generate ke
package words

// ke: {"package": {"notest": true}}

import (
	"context"

	"kego.io/system"
)

func (s *Simple) GetString(ctx context.Context) *system.String {
	return s.String
}

func (t *Translation) GetString(ctx context.Context) *system.String {
	return t.English
}

type Localizer interface {
	Localize(fallbacks []string) string
}

func (t *Translation) Localize(fallbacks []string) string {
	for _, lang := range fallbacks {
		if t.Translations[lang] != nil {
			return t.Translations[lang].Value()
		}
	}
	return ""
}

func (s *Simple) Localize(fallbacks []string) string {
	return s.String.Value()
}
