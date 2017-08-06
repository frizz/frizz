package global

import (
	"fmt"

	"github.com/pkg/errors"
)

func NewStack(name string) Stack {
	return Stack{RootItem(name)}
}

type Stack []stackItem

func (s Stack) Child(item stackItem) Location {
	n := make(Stack, len(s), len(s)+1)
	copy(n, s)
	return append(n, item)
}

func (s Stack) String() string {
	var str string
	for _, item := range s {
		str += item.String()
	}
	return str
}

func (s Stack) Wrap(err error) error {
	return errors.Wrapf(err, "%s:", s)
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
	return fmt.Sprintf(`%s`, string(r))
}
func (f FieldItem) String() string {
	return fmt.Sprintf(`.%s`, string(f))
}
func (m MapItem) String() string {
	return fmt.Sprintf(`["%s"]`, string(m))
}
func (a ArrayItem) String() string {
	return fmt.Sprintf(`[%d]`, a)
}
