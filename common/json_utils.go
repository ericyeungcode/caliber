package common

import (
	"encoding/json"
	"log"
	"strings"
)

func MustMarshalStr(v any) string {
	buf, err := json.Marshal(v)
	if err != nil {
		log.Panicf("json marshal error: %v\n", err)
		return ""
	}
	return string(buf)
}

// support struct, struct pointer (e.g. *image.Pointer), map, slice, etc.
func JsonToValue[T any](s string) (T, error) {
	var result T
	err := json.NewDecoder(strings.NewReader(s)).Decode(&result)
	return result, err
}
