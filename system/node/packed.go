package node

import "kego.io/packer"

var _ packer.Packed = (*packed)(nil)

func Pack(n *Node) *packed {
	return &packed{n}
}

type packed struct {
	n *Node
}

func (p *packed) Type() packer.Type {
	if p == nil || p.n == nil {
		return packer.J_NULL
	}
	if p.n.JsonType == packer.J_OBJECT {
		// packer.Packed is just a dumb collection of json values with no type information, so maps
		// and objects are the same, so we only use J_MAP
		return packer.J_MAP
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
func (p *packed) Array() []packer.Packed {
	out := []packer.Packed{}
	for _, v := range p.n.Array {
		out = append(out, Pack(v))
	}
	return out
}
func (p *packed) Map() map[string]packer.Packed {
	out := map[string]packer.Packed{}
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
