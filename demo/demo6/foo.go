package demo6

import "context"

func (f *Foo) Label(ctx context.Context) string {
	if f == nil || f.Foo == nil {
		return ""
	}
	return f.Foo.Value()
}
