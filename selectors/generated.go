package selectors

import (
	"reflect"

	"kego.io/json"

	"kego.io/system"
)

// Automatically created basic rule for image
type Image_rule struct {
	*system.Base

	*system.RuleBase
}

type Typed struct {
	*system.Base

	Avatar Image

	DrinkPreference []system.String

	FavoriteColor system.String

	Kids map[string]*Kid

	Name map[string]system.String

	Weight system.Number
}

// Automatically created basic rule for photo
type Photo_rule struct {
	*system.Base

	*system.RuleBase
}

type Kid struct {
	*system.Base

	Language system.String

	Level system.String

	Preferred system.Bool
}

// Automatically created basic rule for c
type C_rule struct {
	*system.Base

	*system.RuleBase
}

type Expr struct {
	*system.Base

	False system.Bool

	Float system.Number

	Int system.Number

	Null system.String

	String system.String

	String2 system.String

	True system.Bool
}

type C struct {
	*system.Base

	A system.Number

	B system.Number

	C map[string]system.Number
}

// Automatically created basic rule for expr
type Expr_rule struct {
	*system.Base

	*system.RuleBase
}

type Collision struct {
	*system.Base

	Number map[string]system.String
}

type Diagram struct {
	*system.Base

	Url system.String
}

// Automatically created basic rule for kid
type Kid_rule struct {
	*system.Base

	*system.RuleBase
}

type Polykids struct {
	*system.Base

	A []*Kid
}

// Automatically created basic rule for typed
type Typed_rule struct {
	*system.Base

	*system.RuleBase
}

// Automatically created basic rule for collision
type Collision_rule struct {
	*system.Base

	*system.RuleBase
}

// Automatically created basic rule for sibling
type Sibling_rule struct {
	*system.Base

	*system.RuleBase
}

// Automatically created basic rule for polykids
type Polykids_rule struct {
	*system.Base

	*system.RuleBase
}

type Basic struct {
	*system.Base

	DrinkPreference []system.String

	FavoriteColor system.String

	LanguagesSpoken []map[string]system.String

	Name map[string]system.String

	SeatingPreference []system.String

	Weight system.Number
}

// Automatically created basic rule for diagram
type Diagram_rule struct {
	*system.Base

	*system.RuleBase
}

// Automatically created basic rule for gallery
type Gallery_rule struct {
	*system.Base

	*system.RuleBase
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Base

	// The path for the url - e.g. /foo/bar.jpg
	Path system.String

	// The protocol for the url - e.g. http or https
	Protocol system.String `kego:"{\"default\":{\"value\":\"http\"}}"`

	// The server for the url - e.g. www.google.com
	Server system.String
}

// Automatically created basic rule for basic
type Basic_rule struct {
	*system.Base

	*system.RuleBase
}

type Sibling struct {
	*system.Base

	A system.Number

	B system.Number

	C *C

	D map[string]system.Number

	E map[string]system.Number
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Base

	Images map[string]Image
}

func init() {

	json.RegisterType("kego.io/selectors", "@gallery", reflect.TypeOf(&Gallery_rule{}))

	json.RegisterType("kego.io/selectors", "photo", reflect.TypeOf(&Photo{}))

	json.RegisterType("kego.io/selectors", "@basic", reflect.TypeOf(&Basic_rule{}))

	json.RegisterType("kego.io/selectors", "sibling", reflect.TypeOf(&Sibling{}))

	json.RegisterType("kego.io/selectors", "gallery", reflect.TypeOf(&Gallery{}))

	json.RegisterType("kego.io/selectors", "@diagram", reflect.TypeOf(&Diagram_rule{}))

	json.RegisterType("kego.io/selectors", "kid", reflect.TypeOf(&Kid{}))

	json.RegisterType("kego.io/selectors", "@c", reflect.TypeOf(&C_rule{}))

	json.RegisterType("kego.io/selectors", "expr", reflect.TypeOf(&Expr{}))

	json.RegisterType("kego.io/selectors", "c", reflect.TypeOf(&C{}))

	json.RegisterType("kego.io/selectors", "@expr", reflect.TypeOf(&Expr_rule{}))

	json.RegisterType("kego.io/selectors", "@image", reflect.TypeOf(&Image_rule{}))

	json.RegisterType("kego.io/selectors", "typed", reflect.TypeOf(&Typed{}))

	json.RegisterType("kego.io/selectors", "@photo", reflect.TypeOf(&Photo_rule{}))

	json.RegisterType("kego.io/selectors", "diagram", reflect.TypeOf(&Diagram{}))

	json.RegisterType("kego.io/selectors", "@kid", reflect.TypeOf(&Kid_rule{}))

	json.RegisterType("kego.io/selectors", "collision", reflect.TypeOf(&Collision{}))

	json.RegisterType("kego.io/selectors", "polykids", reflect.TypeOf(&Polykids{}))

	json.RegisterType("kego.io/selectors", "@collision", reflect.TypeOf(&Collision_rule{}))

	json.RegisterType("kego.io/selectors", "@sibling", reflect.TypeOf(&Sibling_rule{}))

	json.RegisterType("kego.io/selectors", "@polykids", reflect.TypeOf(&Polykids_rule{}))

	json.RegisterType("kego.io/selectors", "basic", reflect.TypeOf(&Basic{}))

	json.RegisterType("kego.io/selectors", "@typed", reflect.TypeOf(&Typed_rule{}))

}

var _ system.Nothing

type Nothing struct{}
