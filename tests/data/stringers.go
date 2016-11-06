package data

import (
	"context"

	"sort"

	"kego.io/system"
)

func (a Alajs) GetString(ctx context.Context) *system.String {
	out := ""
	for _, v := range a {
		out += v
	}
	return system.NewString(out)
}

func (m Almjs) GetString(ctx context.Context) *system.String {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := ""
	for _, k := range keys {
		out += k
		out += m[k]
	}
	return system.NewString(out)
}
