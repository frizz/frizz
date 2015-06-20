package selectors

import (
	"kego.io/system"
)

var ptr8729213568 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8729213952 = &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8729332624 = &Photo{Base: ptr8729213952, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729214208 = &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8730682912 = &Diagram{Base: ptr8729214208, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730215952 = &Gallery{Base: ptr8729213568, Images: map[string]Image{"dave": ptr8729332624, "foo": ptr8730682912}}

var Faces = ptr8730215952
