package selectors

import (
	"kego.io/system"
)

var ptr8731168864 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8730125184 = &system.Base{Context: ptr8731168864, Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}

var ptr8731168448 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8730125312 = &system.Base{Context: ptr8731168448, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}

var ptr8729048944 = &Photo{Base: ptr8730125312, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}}

var ptr8731168736 = &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

var ptr8730125440 = &system.Base{Context: ptr8731168736, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}

var ptr8731168608 = &Diagram{Base: ptr8730125440, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}

var ptr8730956080 = &Gallery{Base: ptr8730125184, Images: map[string]Image{"dave": ptr8729048944, "foo": ptr8731168608}}

var Faces = ptr8730956080
