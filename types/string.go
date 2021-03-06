package types

import (
	"strings"
)

type String string

func (s String) HasSuffix(suffix String) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

func (s String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func (s String) Compare(other String) int {
	return strings.Compare(string(s), string(other))
}

func (s String) Bytes() Bytes {
	return Bytes(s)
}

func (s String) Runes() Runes {
	return Runes(s)
}

func (s String) String() string {
	return string(s)
}
