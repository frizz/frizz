package node

import "frizz.io/system"

var _ system.Packed = (*packed)(nil)

func Pack(n *Node) *packed {
	return &packed{n}
}

type packed struct {
	n *Node
}

func (p *packed) Type() system.JsonType {
	if p == nil || p.n == nil {
		return system.J_NULL
	}
	if p.n.JsonType == system.J_OBJECT {
		// system.Packed is just a dumb collection of json values with no type information, so maps
		// and objects are the same, so we only use J_MAP
		return system.J_MAP
	}
	return p.n.JsonType
}
func (p *packed) Number() float64 {
	return p.n.ValueNumber
}
func (p *packed) String() string {
	return p.n.ValueString
}
func (p *packed) Bool() bool {
	return p.n.ValueBool
}
func (p *packed) Array() []system.Packed {
	out := []system.Packed{}
	for _, v := range p.n.Array {
		out = append(out, Pack(v))
	}
	return out
}
func (p *packed) Map() map[string]system.Packed {
	out := map[string]system.Packed{}
	for n, v := range p.n.Map {
		// v.Missing will only be true for object fields, not map items, so we don't need to
		// make a special case here
		if !v.Missing {
			out[n] = Pack(v)
		}
	}
	return out
}
func (p *packed) Interface() interface{} {
	return p.n.Value
}
