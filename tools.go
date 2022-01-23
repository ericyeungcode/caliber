package caliber

import "log"

func Assert(cond bool, format string, args ...interface{}) {
	if !cond {
		log.Panicf(format, args...)
	}
}
