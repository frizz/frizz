package json

type Type string

const (
	J_NULL     Type = "null"
	J_STRING        = "string"
	J_BOOL          = "bool"
	J_NUMBER        = "number"
	J_ARRAY         = "array"
	J_MAP           = "map"
	J_OBJECT        = "object"
	J_OPERATOR      = "operator" // special type user by selectors
)
