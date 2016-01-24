package json

type Type string

const (
	J_NULL     Type = "null"
	J_STRING   Type = "string"
	J_BOOL     Type = "bool"
	J_NUMBER   Type = "number"
	J_ARRAY    Type = "array"
	J_MAP      Type = "map"
	J_OBJECT   Type = "object"
	J_OPERATOR Type = "operator" // special type user by selectors
)
