package common

import (
	"fmt"
	"strconv"
	"time"
)

const (
	TimeFormat_StdDate = "20060102"
)

func ShowElapsedTime(format string, args ...any) func() {
	start := time.Now()
	return func() {
		msg := fmt.Sprintf(format, args...)
		fmt.Printf("ShowElapsedTime: msg:%v took %v\n", msg, time.Since(start))
	}
}

func FormatStdDate(x time.Time) string {
	return x.Format(TimeFormat_StdDate)
}

func DurToMs(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

func MsToDur(durMs int64) time.Duration {
	return time.Duration(durMs) * time.Millisecond
}

func YYYYMMDDToDate(intDate int) (time.Time, error) {
	return time.Parse(TimeFormat_StdDate, strconv.Itoa(intDate))
}

func DateToYYYYMMDD(tm time.Time) int {
	return tm.Year()*10000 + int(tm.Month())*100 + tm.Day()
}

func GetDateOfTime(t time.Time) time.Time {
	year, month, day := t.Date()
	// Create a new time.Time object with time set to 00:00:00
	currentDate := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return currentDate
}

func GetCurrentDateUTC() time.Time {
	return GetDateOfTime(time.Now().UTC())
}
