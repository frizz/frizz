package tests

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for basic
type Basic_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for c
type C_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for collision
type Collision_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for diagram
type Diagram_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for expr
type Expr_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for gallery
type Gallery_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for image
type Image_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for instance
type Instance_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for instanceItem
type InstanceItem_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for kid
type Kid_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for photo
type Photo_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for polykids
type Polykids_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for rectangle
type Rectangle_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for rightscale
type Rightscale_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for rightscaleLink
type RightscaleLink_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for rightscaleList
type RightscaleList_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for sibling
type Sibling_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for simple
type Simple_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for simpleItem
type SimpleItem_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for typed
type Typed_rule struct {
	*system.Object_base
	*system.Rule_base
}
type Basic struct {
	*system.Object_base
	DrinkPreference   []system.String            `json:"drinkPreference"`
	FavoriteColor     system.String              `json:"favoriteColor"`
	LanguagesSpoken   []map[string]system.String `json:"languagesSpoken"`
	Name              map[string]system.String   `json:"name"`
	SeatingPreference []system.String            `json:"seatingPreference"`
	Weight            system.Number              `json:"weight"`
}
type C struct {
	*system.Object_base
	A system.Number            `json:"a"`
	B system.Number            `json:"b"`
	C map[string]system.Number `json:"c"`
}
type Collision struct {
	*system.Object_base
	Number map[string]system.String `json:"number"`
}
type Diagram struct {
	*system.Object_base
	Url system.String `json:"url"`
}
type Expr struct {
	*system.Object_base
	False   system.Bool   `json:"false"`
	Float   system.Number `json:"float"`
	Int     system.Number `json:"int"`
	Null    system.String `json:"null"`
	String  system.String `json:"string"`
	String2 system.String `json:"string2"`
	True    system.Bool   `json:"true"`
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Object_base
	Images map[string]Image `json:"images"`
}
type Instance struct {
	*system.Object_base
	Cloud_type   system.String     `json:"cloud_type"`
	Display_name system.String     `json:"display_name"`
	Links        []*RightscaleLink `json:"links"`
	Name         system.String     `json:"name"`
}
type InstanceItem struct {
	*system.Object_base
	Name system.String `json:"name"`
}
type Kid struct {
	*system.Object_base
	Language  system.String `json:"language"`
	Level     system.String `json:"level"`
	Preferred system.Bool   `json:"preferred"`
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Object_base
	// The path for the url - e.g. /foo/bar.jpg
	Path system.String `json:"path"`
	// The protocol for the url - e.g. http or https
	Protocol system.String `kego:"{\"default\":{\"value\":\"http\"}}" json:"protocol"`
	// The server for the url - e.g. www.google.com
	Server system.String `json:"server"`
	Size   *Rectangle    `json:"size"`
}
type Polykids struct {
	*system.Object_base
	A []*Kid `json:"a"`
}
type Rectangle struct {
	*system.Object_base
	Height system.Int `json:"height"`
	Width  system.Int `json:"width"`
}
type Rightscale struct {
	*system.Object_base
	Child map[string]*InstanceItem `json:"child"`
	Name  system.String            `json:"name"`
}
type RightscaleLink struct {
	*system.Object_base
	Href system.String `json:"href"`
	Rel  system.String `json:"rel"`
}
type RightscaleList struct {
	*system.Object_base
	Foo []*Rightscale `json:"foo"`
}
type Sibling struct {
	*system.Object_base
	A system.Number            `json:"a"`
	B system.Number            `json:"b"`
	C *C                       `json:"c"`
	D map[string]system.Number `json:"d"`
	E map[string]system.Number `json:"e"`
}
type Simple struct {
	*system.Object_base
	A *SimpleItem `json:"a"`
}
type SimpleItem struct {
	*system.Object_base
	B system.String `json:"b"`
}
type Typed struct {
	*system.Object_base
	Avatar          Image                    `json:"avatar"`
	DrinkPreference []system.String          `json:"drinkPreference"`
	FavoriteColor   system.String            `json:"favoriteColor"`
	Kids            map[string]*Kid          `json:"kids"`
	Name            map[string]system.String `json:"name"`
	Weight          system.Number            `json:"weight"`
}

func init() {
	json.Register("kego.io/selectors/tests", "@basic", reflect.TypeOf(&Basic_rule{}), 0xe2f725e9d48b69d4)
	json.Register("kego.io/selectors/tests", "@c", reflect.TypeOf(&C_rule{}), 0x6a4bf0b03046c68a)
	json.Register("kego.io/selectors/tests", "@collision", reflect.TypeOf(&Collision_rule{}), 0x2f6382cf84b27ee7)
	json.Register("kego.io/selectors/tests", "@diagram", reflect.TypeOf(&Diagram_rule{}), 0x2076b0eaf334855b)
	json.Register("kego.io/selectors/tests", "@expr", reflect.TypeOf(&Expr_rule{}), 0x6214e678b1df35e3)
	json.Register("kego.io/selectors/tests", "@gallery", reflect.TypeOf(&Gallery_rule{}), 0xa2261dbb985b3d3)
	json.Register("kego.io/selectors/tests", "@image", reflect.TypeOf(&Image_rule{}), 0x2a2a6f416ac31013)
	json.Register("kego.io/selectors/tests", "@instance", reflect.TypeOf(&Instance_rule{}), 0x91954502a09aa42a)
	json.Register("kego.io/selectors/tests", "@instanceItem", reflect.TypeOf(&InstanceItem_rule{}), 0x2dca71ec3918c621)
	json.Register("kego.io/selectors/tests", "@kid", reflect.TypeOf(&Kid_rule{}), 0x3b576fd38b04a2eb)
	json.Register("kego.io/selectors/tests", "@photo", reflect.TypeOf(&Photo_rule{}), 0x9b14a4b75f8931a8)
	json.Register("kego.io/selectors/tests", "@polykids", reflect.TypeOf(&Polykids_rule{}), 0xebaffecb552df4d1)
	json.Register("kego.io/selectors/tests", "@rectangle", reflect.TypeOf(&Rectangle_rule{}), 0x21539507256190ab)
	json.Register("kego.io/selectors/tests", "@rightscale", reflect.TypeOf(&Rightscale_rule{}), 0x19db6e12b53f9744)
	json.Register("kego.io/selectors/tests", "@rightscaleLink", reflect.TypeOf(&RightscaleLink_rule{}), 0x127d9489c45697ea)
	json.Register("kego.io/selectors/tests", "@rightscaleList", reflect.TypeOf(&RightscaleList_rule{}), 0x6a9527067852e20b)
	json.Register("kego.io/selectors/tests", "@sibling", reflect.TypeOf(&Sibling_rule{}), 0xac837d672ff74ed6)
	json.Register("kego.io/selectors/tests", "@simple", reflect.TypeOf(&Simple_rule{}), 0x7054dc6469f4db46)
	json.Register("kego.io/selectors/tests", "@simpleItem", reflect.TypeOf(&SimpleItem_rule{}), 0xac167d73687ccb1f)
	json.Register("kego.io/selectors/tests", "@typed", reflect.TypeOf(&Typed_rule{}), 0x216d0e1d4c22bb2d)
	json.Register("kego.io/selectors/tests", "basic", reflect.TypeOf(&Basic{}), 0xe2f725e9d48b69d4)
	json.Register("kego.io/selectors/tests", "c", reflect.TypeOf(&C{}), 0x6a4bf0b03046c68a)
	json.Register("kego.io/selectors/tests", "collision", reflect.TypeOf(&Collision{}), 0x2f6382cf84b27ee7)
	json.Register("kego.io/selectors/tests", "diagram", reflect.TypeOf(&Diagram{}), 0x2076b0eaf334855b)
	json.Register("kego.io/selectors/tests", "expr", reflect.TypeOf(&Expr{}), 0x6214e678b1df35e3)
	json.Register("kego.io/selectors/tests", "gallery", reflect.TypeOf(&Gallery{}), 0xa2261dbb985b3d3)
	json.Register("kego.io/selectors/tests", "image", reflect.TypeOf((*Image)(nil)).Elem(), 0x2a2a6f416ac31013)
	json.Register("kego.io/selectors/tests", "instance", reflect.TypeOf(&Instance{}), 0x91954502a09aa42a)
	json.Register("kego.io/selectors/tests", "instanceItem", reflect.TypeOf(&InstanceItem{}), 0x2dca71ec3918c621)
	json.Register("kego.io/selectors/tests", "kid", reflect.TypeOf(&Kid{}), 0x3b576fd38b04a2eb)
	json.Register("kego.io/selectors/tests", "photo", reflect.TypeOf(&Photo{}), 0x9b14a4b75f8931a8)
	json.Register("kego.io/selectors/tests", "polykids", reflect.TypeOf(&Polykids{}), 0xebaffecb552df4d1)
	json.Register("kego.io/selectors/tests", "rectangle", reflect.TypeOf(&Rectangle{}), 0x21539507256190ab)
	json.Register("kego.io/selectors/tests", "rightscale", reflect.TypeOf(&Rightscale{}), 0x19db6e12b53f9744)
	json.Register("kego.io/selectors/tests", "rightscaleLink", reflect.TypeOf(&RightscaleLink{}), 0x127d9489c45697ea)
	json.Register("kego.io/selectors/tests", "rightscaleList", reflect.TypeOf(&RightscaleList{}), 0x6a9527067852e20b)
	json.Register("kego.io/selectors/tests", "sibling", reflect.TypeOf(&Sibling{}), 0xac837d672ff74ed6)
	json.Register("kego.io/selectors/tests", "simple", reflect.TypeOf(&Simple{}), 0x7054dc6469f4db46)
	json.Register("kego.io/selectors/tests", "simpleItem", reflect.TypeOf(&SimpleItem{}), 0xac167d73687ccb1f)
	json.Register("kego.io/selectors/tests", "typed", reflect.TypeOf(&Typed{}), 0x216d0e1d4c22bb2d)
}
