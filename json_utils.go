package caliber

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func Marshal(v any) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return buf
}

func MarshalStr(v any) string {
	return string(Marshal(v))
}

func UnMarshalStr(data string, v any) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		log.Panic(err)
	}
}
