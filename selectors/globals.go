package selectors

import (
	"kego.io/system"
)

var ptr8729350528 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729350656 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8729425376 = &Photo{Base: ptr8729350656, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729350784 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8730534816 = &Diagram{Base: ptr8729350784, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730674416 = &Gallery{Base: ptr8729350528, Images: map[string]Image{"dave": ptr8729425376, "foo": ptr8730534816}}

var Faces = ptr8730674416
