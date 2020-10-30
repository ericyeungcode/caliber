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
