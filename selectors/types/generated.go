package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729571808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592192 := &system.Base{Context: ptr8729571808, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729539040 := &system.Type{Base: ptr8729592192, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729359392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285760 := &system.Base{Context: ptr8729359392, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729535344 := &system.Type{Base: ptr8729285760, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729568704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592704 := &system.Base{Context: ptr8729568704, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729539568 := &system.Type{Base: ptr8729592704, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729360896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286144 := &system.Base{Context: ptr8729360896, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729535872 := &system.Type{Base: ptr8729286144, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729355136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283072 := &system.Base{Context: ptr8729355136, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729164848 := &system.Type{Base: ptr8729283072, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729571392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729167104 := &system.Base{Context: ptr8729571392, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729570848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729591936 := &system.Base{Context: ptr8729570848, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729538688 := &system.String_rule{Base: ptr8729591936, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729571264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592064 := &system.Base{Context: ptr8729571264, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729538864 := &system.String_rule{Base: ptr8729592064, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729570496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729591808 := &system.Base{Context: ptr8729570496, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729570464 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729538512 := &system.String_rule{Base: ptr8729591808, RuleBase: ptr8729570464, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729538336 := &system.Type{Base: ptr8729167104, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"server": ptr8729538688, "path": ptr8729538864, "protocol": ptr8729538512}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729567520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287936 := &system.Base{Context: ptr8729567520, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729537280 := &system.Type{Base: ptr8729287936, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729360320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285888 := &system.Base{Context: ptr8729360320, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729360192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286016 := &system.Base{Context: ptr8729360192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729535696 := &system.String_rule{Base: ptr8729286016, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729535520 := &system.Type{Base: ptr8729285888, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729535696}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729569312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729288192 := &system.Base{Context: ptr8729569312, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729568480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729288320 := &system.Base{Context: ptr8729568480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729537808 := &system.String_rule{Base: ptr8729288320, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729568800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729288448 := &system.Base{Context: ptr8729568800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729537984 := &system.String_rule{Base: ptr8729288448, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729569184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729288576 := &system.Base{Context: ptr8729569184, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729569152 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729569024 := &system.Bool_rule{Base: ptr8729288576, RuleBase: ptr8729569152, Default: system.Bool{Value: false, Exists: false}}

	ptr8729537632 := &system.Type{Base: ptr8729288192, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729537808, "level": ptr8729537984, "preferred": ptr8729569024}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729568192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592320 := &system.Base{Context: ptr8729568192, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729567808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592448 := &system.Base{Context: ptr8729567808, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729567776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592576 := &system.Base{Context: ptr8729567776, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729600160 := &selectors.Kid_rule{Base: ptr8729592576, RuleBase: (*system.RuleBase)(nil)}

	ptr8729007664 := &system.Array_rule{Base: ptr8729592448, RuleBase: (*system.RuleBase)(nil), Items: ptr8729600160, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729539216 := &system.Type{Base: ptr8729592320, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729007664}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728821344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287424 := &system.Base{Context: ptr8728821344, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728821216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287552 := &system.Base{Context: ptr8728821216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728821184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287680 := &system.Base{Context: ptr8728821184, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8728979920 := &selectors.Image_rule{Base: ptr8729287680, RuleBase: (*system.RuleBase)(nil)}

	ptr8729199872 := &system.Map_rule{Base: ptr8729287552, RuleBase: (*system.RuleBase)(nil), Items: ptr8728979920, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729536928 := &system.Type{Base: ptr8729287424, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729199872}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729569728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729165824 := &system.Base{Context: ptr8729569728, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729538160 := &system.Type{Base: ptr8729165824, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729567232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287808 := &system.Base{Context: ptr8729567232, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729537104 := &system.Type{Base: ptr8729287808, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729354720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283712 := &system.Base{Context: ptr8729354720, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729360608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284864 := &system.Base{Context: ptr8729360608, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729360576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284992 := &system.Base{Context: ptr8729360576, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729534464 := &system.String_rule{Base: ptr8729284992, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729396160 := &system.Array_rule{Base: ptr8729284864, RuleBase: (*system.RuleBase)(nil), Items: ptr8729534464, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729354432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729282944 := &system.Base{Context: ptr8729354432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729426592 := &system.Number_rule{Base: ptr8729282944, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729356768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283840 := &system.Base{Context: ptr8729356768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729356736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283968 := &system.Base{Context: ptr8729356736, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729165024 := &system.String_rule{Base: ptr8729283968, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729199232 := &system.Map_rule{Base: ptr8729283840, RuleBase: (*system.RuleBase)(nil), Items: ptr8729165024, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729357088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284096 := &system.Base{Context: ptr8729357088, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729165200 := &system.String_rule{Base: ptr8729284096, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729358016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284224 := &system.Base{Context: ptr8729358016, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729357984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284352 := &system.Base{Context: ptr8729357984, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729357952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284480 := &system.Base{Context: ptr8729357952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729165376 := &system.String_rule{Base: ptr8729284480, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729199296 := &system.Map_rule{Base: ptr8729284352, RuleBase: (*system.RuleBase)(nil), Items: ptr8729165376, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729396000 := &system.Array_rule{Base: ptr8729284224, RuleBase: (*system.RuleBase)(nil), Items: ptr8729199296, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729359968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284608 := &system.Base{Context: ptr8729359968, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729359936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729284736 := &system.Base{Context: ptr8729359936, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729165552 := &system.String_rule{Base: ptr8729284736, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729396080 := &system.Array_rule{Base: ptr8729284608, RuleBase: (*system.RuleBase)(nil), Items: ptr8729165552, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729164672 := &system.Type{Base: ptr8729283712, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"drinkPreference": ptr8729396160, "weight": ptr8729426592, "name": ptr8729199232, "favoriteColor": ptr8729165200, "languagesSpoken": ptr8729396000, "seatingPreference": ptr8729396080}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729572800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592832 := &system.Base{Context: ptr8729572800, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729569504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729592960 := &system.Base{Context: ptr8729569504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420256 := &system.Number_rule{Base: ptr8729592960, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729570240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593088 := &system.Base{Context: ptr8729570240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420352 := &system.Number_rule{Base: ptr8729593088, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729570784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593216 := &system.Base{Context: ptr8729570784, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@c", Package: "kego.io/selectors", Type: "@c", Exists: true}}

	ptr8729601504 := &selectors.C_rule{Base: ptr8729593216, RuleBase: (*system.RuleBase)(nil)}

	ptr8729571680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593344 := &system.Base{Context: ptr8729571680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729571456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593472 := &system.Base{Context: ptr8729571456, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420448 := &system.Number_rule{Base: ptr8729593472, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729199104 := &system.Map_rule{Base: ptr8729593344, RuleBase: (*system.RuleBase)(nil), Items: ptr8729420448, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729572672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593600 := &system.Base{Context: ptr8729572672, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729572512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593728 := &system.Base{Context: ptr8729572512, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420544 := &system.Number_rule{Base: ptr8729593728, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729200448 := &system.Map_rule{Base: ptr8729593600, RuleBase: (*system.RuleBase)(nil), Items: ptr8729420544, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729539744 := &system.Type{Base: ptr8729592832, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729420256, "b": ptr8729420352, "c": ptr8729601504, "d": ptr8729199104, "e": ptr8729200448}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729358080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285248 := &system.Base{Context: ptr8729358080, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729534816 := &system.Type{Base: ptr8729285248, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729358976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285376 := &system.Base{Context: ptr8729358976, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729358848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285504 := &system.Base{Context: ptr8729358848, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729358816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285632 := &system.Base{Context: ptr8729358816, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729535168 := &system.String_rule{Base: ptr8729285632, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729199488 := &system.Map_rule{Base: ptr8729285504, RuleBase: (*system.RuleBase)(nil), Items: ptr8729535168, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729534992 := &system.Type{Base: ptr8729285376, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729199488}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728820000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286272 := &system.Base{Context: ptr8728820000, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728819872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287168 := &system.Base{Context: ptr8728819872, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728819744 := &system.Bool_rule{Base: ptr8729287168, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729361504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286400 := &system.Base{Context: ptr8729361504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420064 := &system.Number_rule{Base: ptr8729286400, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729361952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286528 := &system.Base{Context: ptr8729361952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420160 := &system.Number_rule{Base: ptr8729286528, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729362400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286656 := &system.Base{Context: ptr8729362400, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729536224 := &system.String_rule{Base: ptr8729286656, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728814848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286784 := &system.Base{Context: ptr8728814848, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729536400 := &system.String_rule{Base: ptr8729286784, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728815360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729286912 := &system.Base{Context: ptr8728815360, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729536576 := &system.String_rule{Base: ptr8729286912, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728819488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287040 := &system.Base{Context: ptr8728819488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728819296 := &system.Bool_rule{Base: ptr8729287040, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729536048 := &system.Type{Base: ptr8729286272, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"false": ptr8728819744, "int": ptr8729420064, "float": ptr8729420160, "string": ptr8729536224, "string2": ptr8729536400, "null": ptr8729536576, "true": ptr8728819296}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729355456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593984 := &system.Base{Context: ptr8729355456, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729354496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594880 := &system.Base{Context: ptr8729354496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729354368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729595008 := &system.Base{Context: ptr8729354368, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729540624 := &system.String_rule{Base: ptr8729595008, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729003184 := &system.Array_rule{Base: ptr8729594880, RuleBase: (*system.RuleBase)(nil), Items: ptr8729540624, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729354976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729595136 := &system.Base{Context: ptr8729354976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729420640 := &system.Number_rule{Base: ptr8729595136, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729573984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594112 := &system.Base{Context: ptr8729573984, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729573952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594240 := &system.Base{Context: ptr8729573952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729540272 := &system.String_rule{Base: ptr8729594240, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729200640 := &system.Map_rule{Base: ptr8729594112, RuleBase: (*system.RuleBase)(nil), Items: ptr8729540272, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729574304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594368 := &system.Base{Context: ptr8729574304, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729540448 := &system.String_rule{Base: ptr8729594368, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729574784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594496 := &system.Base{Context: ptr8729574784, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729574752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594624 := &system.Base{Context: ptr8729574752, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729603664 := &selectors.Kid_rule{Base: ptr8729594624, RuleBase: (*system.RuleBase)(nil)}

	ptr8729200704 := &system.Map_rule{Base: ptr8729594496, RuleBase: (*system.RuleBase)(nil), Items: ptr8729603664, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729575040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729594752 := &system.Base{Context: ptr8729575040, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8729603824 := &selectors.Image_rule{Base: ptr8729594752, RuleBase: (*system.RuleBase)(nil)}

	ptr8729540096 := &system.Type{Base: ptr8729593984, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"drinkPreference": ptr8729003184, "weight": ptr8729420640, "name": ptr8729200640, "favoriteColor": ptr8729540448, "kids": ptr8729200704, "avatar": ptr8729603824}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728820416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729287296 := &system.Base{Context: ptr8728820416, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729536752 := &system.Type{Base: ptr8729287296, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729356064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729595264 := &system.Base{Context: ptr8729356064, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729540800 := &system.Type{Base: ptr8729595264, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729567936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729288064 := &system.Base{Context: ptr8729567936, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729537456 := &system.Type{Base: ptr8729288064, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729573216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729593856 := &system.Base{Context: ptr8729573216, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729539920 := &system.Type{Base: ptr8729593856, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729357472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283200 := &system.Base{Context: ptr8729357472, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729355712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283328 := &system.Base{Context: ptr8729355712, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729419776 := &system.Number_rule{Base: ptr8729283328, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729356320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283456 := &system.Base{Context: ptr8729356320, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729419872 := &system.Number_rule{Base: ptr8729283456, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729357312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729283584 := &system.Base{Context: ptr8729357312, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729357152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729285120 := &system.Base{Context: ptr8729357152, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729419968 := &system.Number_rule{Base: ptr8729285120, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729199040 := &system.Map_rule{Base: ptr8729283584, RuleBase: (*system.RuleBase)(nil), Items: ptr8729419968, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729534640 := &system.Type{Base: ptr8729283200, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729419776, "b": ptr8729419872, "c": ptr8729199040}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors:@basic", ptr8729164848)

	system.RegisterType("kego.io/selectors:@c", ptr8729534816)

	system.RegisterType("kego.io/selectors:@collision", ptr8729535344)

	system.RegisterType("kego.io/selectors:@diagram", ptr8729535872)

	system.RegisterType("kego.io/selectors:@expr", ptr8729536752)

	system.RegisterType("kego.io/selectors:@gallery", ptr8729537104)

	system.RegisterType("kego.io/selectors:@image", ptr8729537456)

	system.RegisterType("kego.io/selectors:@kid", ptr8729538160)

	system.RegisterType("kego.io/selectors:@photo", ptr8729539040)

	system.RegisterType("kego.io/selectors:@polykids", ptr8729539568)

	system.RegisterType("kego.io/selectors:@sibling", ptr8729539920)

	system.RegisterType("kego.io/selectors:@typed", ptr8729540800)

	system.RegisterType("kego.io/selectors:basic", ptr8729164672)

	system.RegisterType("kego.io/selectors:c", ptr8729534640)

	system.RegisterType("kego.io/selectors:collision", ptr8729534992)

	system.RegisterType("kego.io/selectors:diagram", ptr8729535520)

	system.RegisterType("kego.io/selectors:expr", ptr8729536048)

	system.RegisterType("kego.io/selectors:gallery", ptr8729536928)

	system.RegisterType("kego.io/selectors:image", ptr8729537280)

	system.RegisterType("kego.io/selectors:kid", ptr8729537632)

	system.RegisterType("kego.io/selectors:photo", ptr8729538336)

	system.RegisterType("kego.io/selectors:polykids", ptr8729539216)

	system.RegisterType("kego.io/selectors:sibling", ptr8729539744)

	system.RegisterType("kego.io/selectors:typed", ptr8729540096)

	// If there's no references to the non types package in here we
	// won't be able to compile, so we force it by calling a noop.
	selectors.DoNothing()
}
