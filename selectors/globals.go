package selectors

import (
	"kego.io/system"
)

var ptr8730314560 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8729417984 = &system.Base{Context: ptr8730314560, Description: "This is an instance of the gallery type containing two nested images", Id: "faces", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:gallery", Package: "kego.io/selectors", Type: "gallery", Exists: true}}

var ptr8730314432 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8729419008 = &system.Base{Context: ptr8730314432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:diagram", Package: "kego.io/selectors", Type: "diagram", Exists: true}}

var ptr8730314272 = &Diagram{Base: ptr8729419008, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730314080 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8729418112 = &system.Base{Context: ptr8730314080, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:photo", Package: "kego.io/selectors", Type: "photo", Exists: true}}

var ptr8730330128 = &Photo{Base: ptr8729418112, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8729398448 = &Gallery{Base: ptr8729417984, Images: map[string]Image{"foo": ptr8730314272, "dave": ptr8730330128}}

var Faces = ptr8729398448
