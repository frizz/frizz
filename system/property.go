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
	b, ok := rule.(Basic)
	if !ok {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: rule %#v is not Basic\n", rule)
	}
	parentType, err := b.Base().Type.RuleToParentType()
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: b.Base().Type.RuleToParentType returned an error:\n%v\n", err)
	}
	name, err := parentType.GoReference(localImports, localPackagePath)
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: parentType.GoReference returned an error:\n%v\n", err)
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
