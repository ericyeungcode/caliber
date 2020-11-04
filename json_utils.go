package caliber

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func StructToStr(v interface{}) string {
	buf, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return string(buf)
}
