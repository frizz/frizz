package helper

import (
	"encoding/json"

	"frizz.io/global"
)

type packable global.Packable
type unmarshaler json.Unmarshaler
type marshaler json.Marshaler
