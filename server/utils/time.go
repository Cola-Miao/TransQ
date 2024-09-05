package utils

import "time"

func GetOutTime(d time.Duration) time.Time {
	if d == 0 {
		return time.Time{}
	}
	return time.Now().Add(d)
}
