package node

import "kego.io/json"

var _ json.Packed = (*packed)(nil)

func Pack(n *Node) *packed {
	return &packed{n}
}

type packed struct {
	n *Node
}

func (p *packed) Type() json.Type {
	if p == nil || p.n == nil {
		return json.J_NULL
	}
	if p.n.JsonType == json.J_OBJECT {
		// unpackables are just a dumb collection of json values with no type
		// information, so maps and objects are the same, so we only use J_MAP
		return json.J_MAP
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
func (p *packed) Array() []json.Packed {
	out := []json.Packed{}
	for _, v := range p.n.Array {
		out = append(out, Pack(v))
	}
	return out
}
func (p *packed) Map() map[string]json.Packed {
	out := map[string]json.Packed{}
	for n, v := range p.n.Map {
		// v.Missing will only be true for object fields, not map items, so we don't need to
		// make a special cae here
		if !v.Missing {
			out[n] = Pack(v)
		}
	}
	return out
}
func (p *packed) Interface() interface{} {
	return p.n.Value
}
