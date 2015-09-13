//go:generate kego -v
package words

type Localizer interface {
	Localize(fallbacks []string) string
}

func (t *Translation) Localize(fallbacks []string) string {
	for _, lang := range fallbacks {
		if t.Translations[lang].Exists {
			return t.Translations[lang].Value
		}
	}
	return ""
}

func (s *Simple) Localize(fallbacks []string) string {
	return s.String.Value
}
