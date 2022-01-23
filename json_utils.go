package caliber

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func Marshal(v interface{}) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return buf
}

func MarshalStr(v interface{}) string {
	return string(Marshal(v))
}

func UnMarshalStr(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		log.Panic(err)
	}
}
