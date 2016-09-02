package common

import "time"

const (
	EditorKeyboardDebounceShort = time.Millisecond * 50
	EditorKeyboardDebounceLong  = time.Millisecond * 500
	ClientConnectionTimeout     = time.Millisecond * 500
)
