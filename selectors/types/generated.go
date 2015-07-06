package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728870016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image", Rules: []system.Rule(nil)}
	ptr8728508096 := &system.Type{Base: ptr8728870016, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728869248 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil)}
	ptr8728507312 := &system.Type{Base: ptr8728869248, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729019264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil)}
	ptr8728508992 := &system.Type{Base: ptr8729019264, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728870784 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c", Rules: []system.Rule(nil)}
	ptr8728506528 := &system.Type{Base: ptr8728870784, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728870656 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil)}
	ptr8728508544 := &system.Type{Base: ptr8728870656, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728868224 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728868352 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190784 := &system.Number_rule{Base: ptr8728868352, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728868480 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190880 := &system.Number_rule{Base: ptr8728868480, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728868608 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652192 := &system.String_rule{Base: ptr8728868608, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728868736 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652368 := &system.String_rule{Base: ptr8728868736, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728868864 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652544 := &system.String_rule{Base: ptr8728868864, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728868992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728839264 := &system.Bool_rule{Base: ptr8728868992, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728869120 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728839552 := &system.Bool_rule{Base: ptr8728869120, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728507200 := &system.Type{Base: ptr8728868224, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"int": ptr8729190784, "float": ptr8729190880, "string": ptr8728652192, "string2": ptr8728652368, "null": ptr8728652544, "true": ptr8728839264, "false": ptr8728839552}, Rule: (*system.Type)(nil)}
	ptr8729018752 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8728508768 := &system.Type{Base: ptr8729018752, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728867200 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728864896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728865152 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728649728 := &system.String_rule{Base: ptr8728865152, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729026560 := &system.Map_rule{Base: ptr8728864896, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728865280 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728649904 := &system.String_rule{Base: ptr8728865280, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728865408 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728865536 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728865664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651136 := &system.String_rule{Base: ptr8728865664, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729026624 := &system.Map_rule{Base: ptr8728865536, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651136, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729223168 := &system.Array_rule{Base: ptr8728865408, RuleBase: (*system.RuleBase)(nil), Items: ptr8729026624, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728865792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728865920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651312 := &system.String_rule{Base: ptr8728865920, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729223248 := &system.Array_rule{Base: ptr8728865792, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651312, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728866048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728866176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651488 := &system.String_rule{Base: ptr8728866176, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729223328 := &system.Array_rule{Base: ptr8728866048, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651488, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728866304 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190400 := &system.Number_rule{Base: ptr8728866304, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728507424 := &system.Type{Base: ptr8728867200, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729026560, "favoriteColor": ptr8728649904, "languagesSpoken": ptr8729223168, "seatingPreference": ptr8729223248, "drinkPreference": ptr8729223328, "weight": ptr8729190400}, Rule: (*system.Type)(nil)}
	ptr8728867840 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728867968 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652016 := &system.String_rule{Base: ptr8728867968, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728506976 := &system.Type{Base: ptr8728867840, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728652016}, Rule: (*system.Type)(nil)}
	ptr8728869760 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8728507760 := &system.Type{Base: ptr8728869760, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729020416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil)}
	ptr8728509328 := &system.Type{Base: ptr8729020416, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728870144 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728870272 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652720 := &system.String_rule{Base: ptr8728870272, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728870400 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652896 := &system.String_rule{Base: ptr8728870400, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728870528 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728842016 := &system.RuleBase{Optional: true}
	ptr8728841856 := &system.Bool_rule{Base: ptr8728870528, RuleBase: ptr8728842016, Default: system.Bool{}}
	ptr8728508432 := &system.Type{Base: ptr8728870144, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728652720, "level": ptr8728652896, "preferred": ptr8728841856}, Rule: (*system.Type)(nil)}
	ptr8728867328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728867456 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728867584 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651840 := &system.String_rule{Base: ptr8728867584, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729027008 := &system.Map_rule{Base: ptr8728867456, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651840, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728506752 := &system.Type{Base: ptr8728867328, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729027008}, Rule: (*system.Type)(nil)}
	ptr8728866560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728866816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190592 := &system.Number_rule{Base: ptr8728866816, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728866944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728867072 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190688 := &system.Number_rule{Base: ptr8728867072, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729026944 := &system.Map_rule{Base: ptr8728866944, RuleBase: (*system.RuleBase)(nil), Items: ptr8729190688, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728866688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190496 := &system.Number_rule{Base: ptr8728866688, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728506416 := &system.Type{Base: ptr8728866560, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"b": ptr8729190592, "c": ptr8729026944, "a": ptr8729190496}, Rule: (*system.Type)(nil)}
	ptr8728666112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8729018624 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728653424 := &system.String_rule{Base: ptr8729018624, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8729018368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8728842976 := &system.RuleBase{Optional: true}
	ptr8728653072 := &system.String_rule{Base: ptr8729018368, RuleBase: ptr8728842976, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729018496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728653248 := &system.String_rule{Base: ptr8729018496, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728508656 := &system.Type{Base: ptr8728666112, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"path": ptr8728653424, "protocol": ptr8728653072, "server": ptr8728653248}, Rule: (*system.Type)(nil)}
	ptr8728868096 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil)}
	ptr8728507088 := &system.Type{Base: ptr8728868096, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729021824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil)}
	ptr8728509552 := &system.Type{Base: ptr8729021824, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728867712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil)}
	ptr8728506864 := &system.Type{Base: ptr8728867712, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729019392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729019904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729020032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729191168 := &system.Number_rule{Base: ptr8729020032, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729026752 := &system.Map_rule{Base: ptr8729019904, RuleBase: (*system.RuleBase)(nil), Items: ptr8729191168, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729020160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729020288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729191264 := &system.Number_rule{Base: ptr8729020288, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729026816 := &system.Map_rule{Base: ptr8729020160, RuleBase: (*system.RuleBase)(nil), Items: ptr8729191264, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729019520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729190976 := &system.Number_rule{Base: ptr8729019520, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729019648 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729191072 := &system.Number_rule{Base: ptr8729019648, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729019776 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728644992 := &selectors.C_rule{Base: ptr8729019776, RuleBase: (*system.RuleBase)(nil)}
	ptr8728509216 := &system.Type{Base: ptr8729019392, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"d": ptr8729026752, "e": ptr8729026816, "a": ptr8729190976, "b": ptr8729191072, "c": ptr8728644992}, Rule: (*system.Type)(nil)}
	ptr8728869376 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728869504 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728869632 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729188608 := &selectors.Image_rule{Base: ptr8728869632, RuleBase: (*system.RuleBase)(nil)}
	ptr8729027136 := &system.Map_rule{Base: ptr8728869504, RuleBase: (*system.RuleBase)(nil), Items: ptr8729188608, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728507536 := &system.Type{Base: ptr8728869376, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729027136}, Rule: (*system.Type)(nil)}
	ptr8728869888 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728507984 := &system.Type{Base: ptr8728869888, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729018880 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729019008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729019136 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728644144 := &selectors.Kid_rule{Base: ptr8729019136, RuleBase: (*system.RuleBase)(nil)}
	ptr8729223808 := &system.Array_rule{Base: ptr8729019008, RuleBase: (*system.RuleBase)(nil), Items: ptr8728644144, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728508880 := &system.Type{Base: ptr8729018880, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729223808}, Rule: (*system.Type)(nil)}
	ptr8729020544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729020672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729020800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653600 := &system.String_rule{Base: ptr8729020800, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729026880 := &system.Map_rule{Base: ptr8729020672, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729020928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653776 := &system.String_rule{Base: ptr8729020928, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729021056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729021184 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728647728 := &selectors.Kid_rule{Base: ptr8729021184, RuleBase: (*system.RuleBase)(nil)}
	ptr8729027904 := &system.Map_rule{Base: ptr8729021056, RuleBase: (*system.RuleBase)(nil), Items: ptr8728647728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729021312 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728647824 := &selectors.Image_rule{Base: ptr8729021312, RuleBase: (*system.RuleBase)(nil)}
	ptr8729021440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729021568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653952 := &system.String_rule{Base: ptr8729021568, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729224688 := &system.Array_rule{Base: ptr8729021440, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729021696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729191360 := &system.Number_rule{Base: ptr8729021696, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728509440 := &system.Type{Base: ptr8729020544, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729026880, "favoriteColor": ptr8728653776, "kids": ptr8729027904, "avatar": ptr8728647824, "drinkPreference": ptr8729224688, "weight": ptr8729191360}, Rule: (*system.Type)(nil)}
	ptr8728866432 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil)}
	ptr8728502384 := &system.Type{Base: ptr8728866432, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors", "@expr", ptr8728507312)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728508992)
	system.RegisterType("kego.io/selectors", "basic", ptr8728507424)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728507760)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728509552)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728508880)
	system.RegisterType("kego.io/selectors", "typed", ptr8728509440)
	system.RegisterType("kego.io/selectors", "@c", ptr8728506528)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728506976)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728509328)
	system.RegisterType("kego.io/selectors", "kid", ptr8728508432)
	system.RegisterType("kego.io/selectors", "c", ptr8728506416)
	system.RegisterType("kego.io/selectors", "photo", ptr8728508656)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728507088)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728506864)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728502384)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728508544)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728508768)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728509216)
	system.RegisterType("kego.io/selectors", "@image", ptr8728508096)
	system.RegisterType("kego.io/selectors", "expr", ptr8728507200)
	system.RegisterType("kego.io/selectors", "collision", ptr8728506752)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728507536)
	system.RegisterType("kego.io/selectors", "image", ptr8728507984)
}
