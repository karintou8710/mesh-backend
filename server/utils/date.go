package utils

import (
	"time"
)

func IsPastTime(t time.Time) bool {
	return t.Before(time.Now())
}

func ParseISO8601(isoStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, isoStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
