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

func (a *Aljb2) GetAljb(ctx context.Context) *Aljb {
	out := new(Aljb)
	*out = Aljb(*a)
	return out
}

func (a *Aljn2) GetAljn(ctx context.Context) *Aljn {
	out := new(Aljn)
	*out = Aljn(*a)
	return out
}

func (a *Aljs2) GetAljs(ctx context.Context) *Aljs {
	out := new(Aljs)
	*out = Aljs(*a)
	return out
}
