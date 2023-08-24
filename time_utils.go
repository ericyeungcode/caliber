package caliber

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func ShowElapsedTime(format string, args ...interface{}) func() {
	start := time.Now()
	return func() {
		msg := fmt.Sprintf(format, args...)
		log.Infof("ShowElapsedTime: msg:%v took %v\n", msg, time.Since(start))
	}
}
