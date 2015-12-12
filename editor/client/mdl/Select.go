package mdl

import "honnef.co/go/js/dom"

type SelectStruct struct {
	*dom.HTMLSelectElement
}

func Select(options map[string]string) *SelectStruct {
	s := &SelectStruct{}
	s.HTMLSelectElement = get("select").(*dom.HTMLSelectElement)
	for k, v := range options {
		o := get("option").(*dom.HTMLOptionElement)
		o.Value = k
		o.SetTextContent(v)
		s.AppendChild(o)
	}
	return s
}
