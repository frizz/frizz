package system

func NewBool(b bool) *Bool {
	out := Bool(b)
	return &out
}

func (b *Bool) Value() bool {
	return bool(*b)
}

var _ Unpacker = (*Bool)(nil)

//var _ Repacker = (*Bool)(nil)

func (b *Bool) String() string {
	if b == nil {
		return ""
	}
	return formatBool(b)
}

func formatBool(b *Bool) string {
	if *b {
		return "true"
	}
	return "false"
}

type NativeBool interface {
	NativeBool() bool
}

func (b Bool) NativeBool() bool {
	return bool(b)
}

var _ NativeBool = (*Bool)(nil)

func (r *BoolRule) GetDefault() interface{} {
	if r.Default == nil {
		return nil
	}
	return r.Default
}

var _ DefaultRule = (*BoolRule)(nil)
