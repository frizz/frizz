// info:{"Path":"kego.io/demo/demo6","Hash":7465480149826331644}
package demo6

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}

func (v *PageRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
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
func (v *PageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo6", "@page", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/demo6", "@page", system.J_OBJECT, nil
}

// Automatically created basic rule for person
type PersonRule struct {
	*system.Object
	*system.Rule
}

func (v *PersonRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
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
func (v *PersonRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo6", "@person", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/demo/demo6", "@person", system.J_OBJECT, nil
}

type Page struct {
	*system.Object
	Heading *system.String `json:"heading"`
	People  []*Person      `json:"people"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func UnpackPageInterface(ctx context.Context, in system.Packed) (PageInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo6", "page")
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
func (v *Page) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["heading"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Heading = ob0
	}
	if field, ok := in.Map()["people"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Person{}
		for i0 := range field.Array() {
			ob1 := new(Person)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.People = ob0
	}
	return nil
}
func (v *Page) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo6", "page", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Heading != nil {
		ob0, _, _, _, err := v.Heading.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["heading"] = ob0
	}
	if v.People != nil {
		ob0 := []interface{}{}
		for i0 := range v.People {
			ob1, _, _, _, err := v.People[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["people"] = ob0
	}
	return m, "kego.io/demo/demo6", "page", system.J_OBJECT, nil
}

type Person struct {
	*system.Object
	Age  *system.Int    `json:"age"`
	Name *system.String `json:"name"`
}
type PersonInterface interface {
	GetPerson(ctx context.Context) *Person
}

func (o *Person) GetPerson(ctx context.Context) *Person {
	return o
}
func UnpackPersonInterface(ctx context.Context, in system.Packed) (PersonInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo6", "person")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PersonInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PersonInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PersonInterface.", in.Type())
	}
}
func (v *Person) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	if field, ok := in.Map()["age"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Age = ob0
	}
	return nil
}
func (v *Person) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo6", "person", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Age != nil {
		ob0, _, _, _, err := v.Age.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["age"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "kego.io/demo/demo6", "person", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo6")
	pkg.SetHash(7465480149826331644)
	pkg.Init("page", func() interface{} { return new(Page) }, func() interface{} { return new(PageRule) }, func() reflect.Type { return reflect.TypeOf((*PageInterface)(nil)).Elem() })
	pkg.Init("person", func() interface{} { return new(Person) }, func() interface{} { return new(PersonRule) }, func() reflect.Type { return reflect.TypeOf((*PersonInterface)(nil)).Elem() })
}
