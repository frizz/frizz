// info:{"Path":"kego.io/demo/site","Hash":3726678967244446909}
package site

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/demo/common/images"
	"kego.io/packer"
	"kego.io/system"
)

// Automatically created basic rule for body
type BodyRule struct {
	*system.Object
	*system.Rule
}

func (v *BodyRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *BodyRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "@body", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "@body", nil
}

// Automatically created basic rule for columns
type ColumnsRule struct {
	*system.Object
	*system.Rule
}

func (v *ColumnsRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *ColumnsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "@columns", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "@columns", nil
}

// Automatically created basic rule for hero
type HeroRule struct {
	*system.Object
	*system.Rule
}

func (v *HeroRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *HeroRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "@hero", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "@hero", nil
}

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}

func (v *PageRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *PageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "@page", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "@page", nil
}

// Automatically created basic rule for section
type SectionRule struct {
	*system.Object
	*system.Rule
}

func (v *SectionRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *SectionRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "@section", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "@section", nil
}

type Body struct {
	*system.Object
	Align *system.String         `json:"align"`
	Copy  system.StringInterface `json:"copy"`
	Title system.StringInterface `json:"title"`
}
type BodyInterface interface {
	GetBody(ctx context.Context) *Body
}

func (o *Body) GetBody(ctx context.Context) *Body {
	return o
}
func UnpackBodyInterface(ctx context.Context, in packer.Packed) (BodyInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/site", "body")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(BodyInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement BodyInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into BodyInterface.", in.Type())
	}
}
func (v *Body) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["align"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Align = ob0
	}
	if field, ok := in.Map()["copy"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Copy = ob0
	}
	if field, ok := in.Map()["title"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Title = ob0
	}
	return nil
}
func (v *Body) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "body", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Align != nil {
		ob0, _, _, err := v.Align.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["align"] = ob0
	}
	return m, "kego.io/demo/site", "body", nil
}

type Columns struct {
	*system.Object
	Columns []Section `json:"columns"`
}
type ColumnsInterface interface {
	GetColumns(ctx context.Context) *Columns
}

func (o *Columns) GetColumns(ctx context.Context) *Columns {
	return o
}
func UnpackColumnsInterface(ctx context.Context, in packer.Packed) (ColumnsInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/site", "columns")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ColumnsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ColumnsInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into ColumnsInterface.", in.Type())
	}
}
func (v *Columns) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["columns"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Section{}
		for i0 := range field.Array() {
			ob1, err := UnpackSection(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Columns = ob0
	}
	return nil
}
func (v *Columns) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "columns", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Columns != nil {
		ob0 := []interface{}{}
		for i0 := range v.Columns {
			ob1_value, pkg, name, err := v.Columns[i0].(packer.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", err
			}
			typRef := system.NewReference(pkg, name)
			typeVal, err := typRef.ValueContext(ctx)
			if err != nil {
				return nil, "", "", err
			}
			ob1 := map[string]interface{}{}
			ob1["type"] = typeVal
			ob1["value"] = ob1_value
			ob0 = append(ob0, ob1)
		}
		m["columns"] = ob0
	}
	return m, "kego.io/demo/site", "columns", nil
}

type Hero struct {
	*system.Object
	Head    system.StringInterface `json:"head"`
	Image   images.Image           `json:"image"`
	Subhead system.StringInterface `json:"subhead"`
}
type HeroInterface interface {
	GetHero(ctx context.Context) *Hero
}

func (o *Hero) GetHero(ctx context.Context) *Hero {
	return o
}
func UnpackHeroInterface(ctx context.Context, in packer.Packed) (HeroInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/site", "hero")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(HeroInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement HeroInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into HeroInterface.", in.Type())
	}
}
func (v *Hero) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["image"]; ok && field.Type() != packer.J_NULL {
		ob0, err := images.UnpackImage(ctx, field)
		if err != nil {
			return err
		}
		v.Image = ob0
	}
	if field, ok := in.Map()["subhead"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Subhead = ob0
	}
	if field, ok := in.Map()["head"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Head = ob0
	}
	return nil
}
func (v *Hero) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "hero", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/site", "hero", nil
}

type Page struct {
	*system.Object
	Sections []Section              `json:"sections"`
	Title    system.StringInterface `json:"title"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func UnpackPageInterface(ctx context.Context, in packer.Packed) (PageInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/site", "page")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PageInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PageInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PageInterface.", in.Type())
	}
}
func (v *Page) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["title"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Title = ob0
	}
	if field, ok := in.Map()["sections"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Section{}
		for i0 := range field.Array() {
			ob1, err := UnpackSection(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Sections = ob0
	}
	return nil
}
func (v *Page) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/site", "page", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Sections != nil {
		ob0 := []interface{}{}
		for i0 := range v.Sections {
			ob1_value, pkg, name, err := v.Sections[i0].(packer.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", err
			}
			typRef := system.NewReference(pkg, name)
			typeVal, err := typRef.ValueContext(ctx)
			if err != nil {
				return nil, "", "", err
			}
			ob1 := map[string]interface{}{}
			ob1["type"] = typeVal
			ob1["value"] = ob1_value
			ob0 = append(ob0, ob1)
		}
		m["sections"] = ob0
	}
	return m, "kego.io/demo/site", "page", nil
}
func UnpackSection(ctx context.Context, in packer.Packed) (Section, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/site", "section")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(Section)
		if !ok {
			return nil, fmt.Errorf("%T does not implement Section", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into Section.", in.Type())
	}
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/site")
	pkg.SetHash(3726678967244446909)
	pkg.Init("body", func() interface{} { return new(Body) }, func() interface{} { return new(BodyRule) }, func() reflect.Type { return reflect.TypeOf((*BodyInterface)(nil)).Elem() })
	pkg.Init("columns", func() interface{} { return new(Columns) }, func() interface{} { return new(ColumnsRule) }, func() reflect.Type { return reflect.TypeOf((*ColumnsInterface)(nil)).Elem() })
	pkg.Init("hero", func() interface{} { return new(Hero) }, func() interface{} { return new(HeroRule) }, func() reflect.Type { return reflect.TypeOf((*HeroInterface)(nil)).Elem() })
	pkg.Init("page", func() interface{} { return new(Page) }, func() interface{} { return new(PageRule) }, func() reflect.Type { return reflect.TypeOf((*PageInterface)(nil)).Elem() })
	pkg.Init("section", func() interface{} { return new(Section) }, func() interface{} { return new(SectionRule) }, func() reflect.Type { return reflect.TypeOf((*Section)(nil)).Elem() })
}
