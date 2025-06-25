package common

import (
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func MarshalStr(v any) (string, error) {
	buf, err := Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func MustMarshalStr(v any) string {
	str, err := MarshalStr(v)
	if err != nil {
		panic(err)
	}
	return str
}

func UnMarshalStr(data string, v any) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return nil
}

func JsonToStructPtr[T any](jsonStr string) (*T, error) {
	var v T
	err := json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
