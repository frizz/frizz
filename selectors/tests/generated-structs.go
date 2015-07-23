package tests

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for basic
type Basic_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for c
type C_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for collision
type Collision_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for diagram
type Diagram_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for expr
type Expr_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery
type Gallery_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for image
type Image_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for instance
type Instance_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for instanceItem
type InstanceItem_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for kid
type Kid_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for photo
type Photo_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for polykids
type Polykids_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for rectangle
type Rectangle_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for rightscale
type Rightscale_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for rightscaleLink
type RightscaleLink_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for rightscaleList
type RightscaleList_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for sibling
type Sibling_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for simple
type Simple_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for simpleItem
type SimpleItem_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for typed
type Typed_rule struct {
	*system.Base
	*system.RuleBase
}
type Basic struct {
	*system.Base
	DrinkPreference   []system.String
	FavoriteColor     system.String
	LanguagesSpoken   []map[string]system.String
	Name              map[string]system.String
	SeatingPreference []system.String
	Weight            system.Number
}
type C struct {
	*system.Base
	A system.Number
	B system.Number
	C map[string]system.Number
}
type Collision struct {
	*system.Base
	Number map[string]system.String
}
type Diagram struct {
	*system.Base
	Url system.String
}
type Expr struct {
	*system.Base
	False   system.Bool
	Float   system.Number
	Int     system.Number
	Null    system.String
	String  system.String
	String2 system.String
	True    system.Bool
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Base
	Images map[string]Image
}
type Instance struct {
	*system.Base
	Cloud_type   system.String
	Display_name system.String
	Links        []*RightscaleLink
	Name         system.String
}
type InstanceItem struct {
	*system.Base
	Name system.String
}
type Kid struct {
	*system.Base
	Language  system.String
	Level     system.String
	Preferred system.Bool
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
	Size   *Rectangle
}
type Polykids struct {
	*system.Base
	A []*Kid
}
type Rectangle struct {
	*system.Base
	Height system.Int
	Width  system.Int
}
type Rightscale struct {
	*system.Base
	Child map[string]*InstanceItem
	Name  system.String
}
type RightscaleLink struct {
	*system.Base
	Href system.String
	Rel  system.String
}
type RightscaleList struct {
	*system.Base
	Foo []*Rightscale
}
type Sibling struct {
	*system.Base
	A system.Number
	B system.Number
	C *C
	D map[string]system.Number
	E map[string]system.Number
}
type Simple struct {
	*system.Base
	A *SimpleItem
}
type SimpleItem struct {
	*system.Base
	B system.String
}
type Typed struct {
	*system.Base
	Avatar          Image
	DrinkPreference []system.String
	FavoriteColor   system.String
	Kids            map[string]*Kid
	Name            map[string]system.String
	Weight          system.Number
}

func init() {
	json.RegisterType("kego.io/selectors/tests", "@basic", reflect.TypeOf(&Basic_rule{}), 0xe2f725e9d48b69d4)
	json.RegisterType("kego.io/selectors/tests", "@c", reflect.TypeOf(&C_rule{}), 0x6a4bf0b03046c68a)
	json.RegisterType("kego.io/selectors/tests", "@collision", reflect.TypeOf(&Collision_rule{}), 0x2f6382cf84b27ee7)
	json.RegisterType("kego.io/selectors/tests", "@diagram", reflect.TypeOf(&Diagram_rule{}), 0x2076b0eaf334855b)
	json.RegisterType("kego.io/selectors/tests", "@expr", reflect.TypeOf(&Expr_rule{}), 0x6214e678b1df35e3)
	json.RegisterType("kego.io/selectors/tests", "@gallery", reflect.TypeOf(&Gallery_rule{}), 0xa2261dbb985b3d3)
	json.RegisterType("kego.io/selectors/tests", "@image", reflect.TypeOf(&Image_rule{}), 0x2a2a6f416ac31013)
	json.RegisterType("kego.io/selectors/tests", "@instance", reflect.TypeOf(&Instance_rule{}), 0x91954502a09aa42a)
	json.RegisterType("kego.io/selectors/tests", "@instanceItem", reflect.TypeOf(&InstanceItem_rule{}), 0x2dca71ec3918c621)
	json.RegisterType("kego.io/selectors/tests", "@kid", reflect.TypeOf(&Kid_rule{}), 0x3b576fd38b04a2eb)
	json.RegisterType("kego.io/selectors/tests", "@photo", reflect.TypeOf(&Photo_rule{}), 0x9b14a4b75f8931a8)
	json.RegisterType("kego.io/selectors/tests", "@polykids", reflect.TypeOf(&Polykids_rule{}), 0xebaffecb552df4d1)
	json.RegisterType("kego.io/selectors/tests", "@rectangle", reflect.TypeOf(&Rectangle_rule{}), 0x21539507256190ab)
	json.RegisterType("kego.io/selectors/tests", "@rightscale", reflect.TypeOf(&Rightscale_rule{}), 0x19db6e12b53f9744)
	json.RegisterType("kego.io/selectors/tests", "@rightscaleLink", reflect.TypeOf(&RightscaleLink_rule{}), 0x127d9489c45697ea)
	json.RegisterType("kego.io/selectors/tests", "@rightscaleList", reflect.TypeOf(&RightscaleList_rule{}), 0x6a9527067852e20b)
	json.RegisterType("kego.io/selectors/tests", "@sibling", reflect.TypeOf(&Sibling_rule{}), 0xac837d672ff74ed6)
	json.RegisterType("kego.io/selectors/tests", "@simple", reflect.TypeOf(&Simple_rule{}), 0x7054dc6469f4db46)
	json.RegisterType("kego.io/selectors/tests", "@simpleItem", reflect.TypeOf(&SimpleItem_rule{}), 0xac167d73687ccb1f)
	json.RegisterType("kego.io/selectors/tests", "@typed", reflect.TypeOf(&Typed_rule{}), 0x216d0e1d4c22bb2d)
	json.RegisterType("kego.io/selectors/tests", "basic", reflect.TypeOf(&Basic{}), 0xe2f725e9d48b69d4)
	json.RegisterType("kego.io/selectors/tests", "c", reflect.TypeOf(&C{}), 0x6a4bf0b03046c68a)
	json.RegisterType("kego.io/selectors/tests", "collision", reflect.TypeOf(&Collision{}), 0x2f6382cf84b27ee7)
	json.RegisterType("kego.io/selectors/tests", "diagram", reflect.TypeOf(&Diagram{}), 0x2076b0eaf334855b)
	json.RegisterType("kego.io/selectors/tests", "expr", reflect.TypeOf(&Expr{}), 0x6214e678b1df35e3)
	json.RegisterType("kego.io/selectors/tests", "gallery", reflect.TypeOf(&Gallery{}), 0xa2261dbb985b3d3)
	json.RegisterType("kego.io/selectors/tests", "instance", reflect.TypeOf(&Instance{}), 0x91954502a09aa42a)
	json.RegisterType("kego.io/selectors/tests", "instanceItem", reflect.TypeOf(&InstanceItem{}), 0x2dca71ec3918c621)
	json.RegisterType("kego.io/selectors/tests", "kid", reflect.TypeOf(&Kid{}), 0x3b576fd38b04a2eb)
	json.RegisterType("kego.io/selectors/tests", "photo", reflect.TypeOf(&Photo{}), 0x9b14a4b75f8931a8)
	json.RegisterType("kego.io/selectors/tests", "polykids", reflect.TypeOf(&Polykids{}), 0xebaffecb552df4d1)
	json.RegisterType("kego.io/selectors/tests", "rectangle", reflect.TypeOf(&Rectangle{}), 0x21539507256190ab)
	json.RegisterType("kego.io/selectors/tests", "rightscale", reflect.TypeOf(&Rightscale{}), 0x19db6e12b53f9744)
	json.RegisterType("kego.io/selectors/tests", "rightscaleLink", reflect.TypeOf(&RightscaleLink{}), 0x127d9489c45697ea)
	json.RegisterType("kego.io/selectors/tests", "rightscaleList", reflect.TypeOf(&RightscaleList{}), 0x6a9527067852e20b)
	json.RegisterType("kego.io/selectors/tests", "sibling", reflect.TypeOf(&Sibling{}), 0xac837d672ff74ed6)
	json.RegisterType("kego.io/selectors/tests", "simple", reflect.TypeOf(&Simple{}), 0x7054dc6469f4db46)
	json.RegisterType("kego.io/selectors/tests", "simpleItem", reflect.TypeOf(&SimpleItem{}), 0xac167d73687ccb1f)
	json.RegisterType("kego.io/selectors/tests", "typed", reflect.TypeOf(&Typed{}), 0x216d0e1d4c22bb2d)
}