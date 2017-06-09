package frizz

import "fmt"

type Stack []stackItem

func (s Stack) Append(item stackItem) Stack {
	n := make(Stack, len(s), len(s)+1)
	copy(n, s)
	return append(n, item)
}

func (s Stack) String() {
	var str string
	for i, item := range s {
		if i != 0 {
			str += "."
		}
		str += item.String()
	}
}

type stackItem interface {
	String() string
	stackItem()
}

type RootItem string
type FieldItem string
type MapItem string
type ArrayItem int

func (RootItem) stackItem()  {}
func (FieldItem) stackItem() {}
func (MapItem) stackItem()   {}
func (ArrayItem) stackItem() {}

func (r RootItem) String() string {
	return fmt.Sprintf(`{%s}`, r)
}
func (f FieldItem) String() string {
	return string(f)
}
func (m MapItem) String() string {
	return fmt.Sprintf(`["%s"]`, m)
}
func (a ArrayItem) String() string {
	return fmt.Sprintf(`[%d]`, a)
}
