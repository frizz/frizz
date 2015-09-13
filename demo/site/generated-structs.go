package site

import (
	"reflect"

	"kego.io/demo/common/images"
	"kego.io/demo/common/words"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for gallery1
type Gallery1_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery1a
type Gallery1a_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2
type Gallery2_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2a
type Gallery2a_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2b
type Gallery2b_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery3
type Gallery3_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery3a
type Gallery3a_rule struct {
	*system.Base
	*system.RuleBase
}

// This represents a gallery - it's just a list of images
type Gallery1 struct {
	*system.Base
	Title system.String `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery1a struct {
	*system.Base
	Title system.String `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery2 struct {
	*system.Base
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery2a struct {
	*system.Base
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery2b struct {
	*system.Base
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery3 struct {
	*system.Base
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}

// This represents a gallery - it's just a list of images
type Gallery3a struct {
	*system.Base
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}

func init() {
	json.RegisterType("kego.io/demo/site", "@gallery1", reflect.TypeOf(&Gallery1_rule{}), 0x8997ec74747e71d5)
	json.RegisterType("kego.io/demo/site", "@gallery1a", reflect.TypeOf(&Gallery1a_rule{}), 0xf6237017f97fc838)
	json.RegisterType("kego.io/demo/site", "@gallery2", reflect.TypeOf(&Gallery2_rule{}), 0x3ff7f11dfc3e88c0)
	json.RegisterType("kego.io/demo/site", "@gallery2a", reflect.TypeOf(&Gallery2a_rule{}), 0x4a0bb74542b31d47)
	json.RegisterType("kego.io/demo/site", "@gallery2b", reflect.TypeOf(&Gallery2b_rule{}), 0x8dc28cbd83066e3d)
	json.RegisterType("kego.io/demo/site", "@gallery3", reflect.TypeOf(&Gallery3_rule{}), 0x327c5e90f012bd49)
	json.RegisterType("kego.io/demo/site", "@gallery3a", reflect.TypeOf(&Gallery3a_rule{}), 0x89fb5c874cf0f786)
	json.RegisterType("kego.io/demo/site", "gallery1", reflect.TypeOf(&Gallery1{}), 0x8997ec74747e71d5)
	json.RegisterType("kego.io/demo/site", "gallery1a", reflect.TypeOf(&Gallery1a{}), 0xf6237017f97fc838)
	json.RegisterType("kego.io/demo/site", "gallery2", reflect.TypeOf(&Gallery2{}), 0x3ff7f11dfc3e88c0)
	json.RegisterType("kego.io/demo/site", "gallery2a", reflect.TypeOf(&Gallery2a{}), 0x4a0bb74542b31d47)
	json.RegisterType("kego.io/demo/site", "gallery2b", reflect.TypeOf(&Gallery2b{}), 0x8dc28cbd83066e3d)
	json.RegisterType("kego.io/demo/site", "gallery3", reflect.TypeOf(&Gallery3{}), 0x327c5e90f012bd49)
	json.RegisterType("kego.io/demo/site", "gallery3a", reflect.TypeOf(&Gallery3a{}), 0x89fb5c874cf0f786)
}
