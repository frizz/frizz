package selectors

import (
	"kego.io/system"
)

var ptr8729275008 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729275136 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8730742448 = &Photo{Base: ptr8729275136, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729275264 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8730534688 = &Diagram{Base: ptr8729275264, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730862160 = &Gallery{Base: ptr8729275008, Images: map[string]Image{"dave": ptr8730742448, "foo": ptr8730534688}}

var Faces = ptr8730862160
