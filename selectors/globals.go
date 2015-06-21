package selectors

import (
	"kego.io/system"
)

var ptr8729250432 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729250560 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8729485392 = &Photo{Base: ptr8729250560, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729250688 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8729772512 = &Diagram{Base: ptr8729250688, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730132592 = &Gallery{Base: ptr8729250432, Images: map[string]Image{"dave": ptr8729485392, "foo": ptr8729772512}}

var Faces = ptr8730132592
