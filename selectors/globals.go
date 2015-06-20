package selectors

import (
	"kego.io/system"
)

var ptr8729531264 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729531392 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8728926304 = &Photo{Base: ptr8729531392, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729531520 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8730608576 = &Diagram{Base: ptr8729531520, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730436656 = &Gallery{Base: ptr8729531264, Images: map[string]Image{"dave": ptr8728926304, "foo": ptr8730608576}}

var Faces = ptr8730436656
