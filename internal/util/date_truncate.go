package util

import "time"

func TruncateDate(date time.Time) time.Time {
	return date.Truncate(24 * time.Hour)
}
