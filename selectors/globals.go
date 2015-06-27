package selectors

import (
	"kego.io/system"
)

var ptr8729559040 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729559168 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8729305168 = &Photo{Base: ptr8729559168, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729559424 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8729846240 = &Diagram{Base: ptr8729559424, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8729477248 = &Gallery{Base: ptr8729559040, Images: map[string]Image{"dave": ptr8729305168, "foo": ptr8729846240}}

var Faces = ptr8729477248
