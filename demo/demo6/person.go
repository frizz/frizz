package demo6

import "context"

func (f *Person) Label(ctx context.Context) string {
	if f == nil || f.Name == nil {
		return ""
	}
	return f.Name.Value()
}
