package util

import "time"

func ParseTime(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}
