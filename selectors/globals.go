package selectors

import (
	"kego.io/system"
)

var ptr8730865248 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8731207760 = &system.Base{Context: ptr8730865248, Description: "This is an instance of the gallery type containing two nested images", Id: "faces", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "gallery", Exists: true}}

var ptr8730864800 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8731207872 = &system.Base{Context: ptr8730864800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "photo", Exists: true}}

var ptr8729206944 = &Photo{Base: ptr8731207872, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8730865120 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8731207984 = &system.Base{Context: ptr8730865120, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "diagram", Exists: true}}

var ptr8730864992 = &Diagram{Base: ptr8731207984, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730718528 = &Gallery{Base: ptr8731207760, Images: map[string]Image{"dave": ptr8729206944, "foo": ptr8730864992}}

var Faces = ptr8730718528
