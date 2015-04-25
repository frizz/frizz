package system

import (
	"fmt"
	"strconv"
)

func (p *Property) GoName(name string) string {
	return IdToGoName(name)
}

// GoTypeDescriptor returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func (p *Property) GoTypeDescriptor(localImports map[string]string, localPackagePath string) (string, error) {
	prefix, rule := collectionPrefix("", p.Item)
	_, isNative := rule.(NativeType)
	_, isInterface := rule.(InterfaceType)
	pointer := ""
	if !isNative && !isInterface {
		pointer = "*"
	}
	var typeReference Reference
	b, ok := rule.(Basic)
	if ok {
		typeReference = b.Base().Type
	} else {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: rule %#v is not Basic\n", rule)
		/*
			// I've disabled this code until we're less reliant on code and more reliant on
			// unmarshaled json type definitions

			// Looks like we unmarshaled a type that isn't registered. We should give out
			// best guess at a type name.
			i, ok := rule.(map[string]interface{})
			if !ok {
				return "", fmt.Errorf("Error in Property.GoTypeDescriptor: unknown type rule is not map[string]interface{}: %T\n", rule)
			}
			t, ok := i["type"]
			if !ok {
				return "", fmt.Errorf("Error in Property.GoTypeDescriptor: unknown type rule has no type\n")
			}
			s, ok := t.(string)
			if !ok {
				return "", fmt.Errorf("Error in Property.GoTypeDescriptor: unknown type rule type is not string: %T\n", t)
			}
			r := &Reference{}
			fmt.Printf("s: %s\n", s)
			err := r.UnmarshalJSON([]byte(strconv.Quote(s)), localPackagePath, localImports)
			if err != nil {
				return "", fmt.Errorf("Error in Property.GoTypeDescriptor: r.UnmarshalJSON returned an error:\n%v\n", err)
			}
			typeReference = *r
		*/
	}
	parentTypeReference, err := typeReference.RuleToParentType()
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: typeReference.RuleToParentType returned an error:\n%v\n", err)
	}
	name, err := parentTypeReference.GoReference(localImports, localPackagePath)
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: parentTypeReference.GoReference returned an error:\n%v\n", err)
	}
	return fmt.Sprintf("%s%s%s", prefix, pointer, name), nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array).
func collectionPrefix(prefix string, rule Rule) (string, Rule) {
	if c, ok := rule.(CollectionType); ok {
		prefix += c.CollectionTypePrefix()
		items := c.CollectionItemsRule()
		return collectionPrefix(prefix, items)
	} else {
		return prefix, rule
	}
}

func (p *Property) GoTag() (string, error) {
	// collectionPrefix returns the nested type of any collection rules
	_, rule := collectionPrefix("", p.Item)
	if d, ok := rule.(Defaulter); ok {
		if d.HasDefault() {
			defaultBytes, err := d.GetDefault()
			if err != nil {
				return "", fmt.Errorf("Error in Property.GoTag: d.GetDefault returned an error:\n%v\n", err)
			}
			jsonString := fmt.Sprintf("{\"default\": %s}", defaultBytes)
			return fmt.Sprintf("`kego:%s`", strconv.Quote(jsonString)), nil
		}
	}
	return "", nil
}
