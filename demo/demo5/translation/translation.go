package translation

// notest

type Localized interface {
	Localize(country string) string
}

func (s *Simple) Localize(country string) string {
	return s.Text.Value()
}

func (s *Smartling) Localize(country string) string {
	for code, translation := range s.Translations {
		if code == country {
			return translation.Value()
		}
	}
	return s.English.Value()
}
