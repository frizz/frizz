package mdl

import "honnef.co/go/js/dom"

type ImageStruct struct {
	*dom.HTMLImageElement
}

func Image(url string) *ImageStruct {
	i := &ImageStruct{}
	i.HTMLImageElement = get("img").(*dom.HTMLImageElement)
	i.Src = url
	return i
}

func (i *ImageStruct) Visibility(state bool) {
	if state {
		i.Style().Set("display", "block")
	} else {
		i.Style().Set("display", "block")
	}
}
