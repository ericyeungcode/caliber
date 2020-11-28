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

func GetMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func TimeToTimestamp(t *time.Time) int64 {
	return t.UnixNano() / 1000000
}

func TimestampToTime(ms int64) time.Time {
	return time.Unix(ms/1000, 0)
}
