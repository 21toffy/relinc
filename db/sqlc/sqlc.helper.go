package db

import (
	"time"
)

func GetTime() time.Time {
	// parse the input string and return a time.Time object
	specificTime, _ := time.ParseInLocation("2006-01-02 15:04:05 -0700", "1997-08-06 12:00:00 +0000", time.UTC)
	return specificTime
}
