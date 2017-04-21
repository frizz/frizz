package data

import (
	"context"

	"sort"

	"frizz.io/system"
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

func (b *Aljb2) GetAljb(ctx context.Context) *Aljb {
	out := new(Aljb)
	*out = Aljb(*b)
	return out
}

func (b *Aljb2) GetBool(ctx context.Context) *system.Bool {
	return system.NewBool(bool(*b))
}

func (n *Aljn2) GetAljn(ctx context.Context) *Aljn {
	out := new(Aljn)
	*out = Aljn(*n)
	return out
}

func (n *Aljn2) GetNumber(ctx context.Context) *system.Number {
	return system.NewNumber(float64(*n))
}

func (s *Aljs2) GetAljs(ctx context.Context) *Aljs {
	out := new(Aljs)
	*out = Aljs(*s)
	return out
}

func (s *Aljs2) GetString(ctx context.Context) *system.String {
	return system.NewString(string(*s))
}
