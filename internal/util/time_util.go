package util

import "time"

func TruncateDate(date time.Time) time.Time {
	return date.Truncate(24 * time.Hour)
}

func DefaultOrProvidedTime(createdAt *time.Time) time.Time {
	if createdAt != nil {
		return *createdAt
	}
	return time.Now()
}
