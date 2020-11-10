package caliber

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func ShowElapsedTime(what string) func() {
	start := time.Now()
	return func() {
		log.Infof("ShowElapsedTime: %s took %v\n", what, time.Since(start))
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
