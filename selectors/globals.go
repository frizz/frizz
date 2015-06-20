package selectors

import (
	"kego.io/system"
)

var ptr8729406848 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729406976 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8730887104 = &Photo{Base: ptr8729406976, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729407104 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8730445984 = &Diagram{Base: ptr8729407104, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8729250096 = &Gallery{Base: ptr8729406848, Images: map[string]Image{"dave": ptr8730887104, "foo": ptr8730445984}}

var Faces = ptr8729250096
