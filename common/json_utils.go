package common

import (
	"encoding/json"
	"strings"
)

func MarshalStr(v any) (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// support struct, struct pointer (e.g. *image.Pointer), map, slice, etc.
func JsonToValue[T any](s string) (T, error) {
	var result T
	err := json.NewDecoder(strings.NewReader(s)).Decode(&result)
	return result, err
}
