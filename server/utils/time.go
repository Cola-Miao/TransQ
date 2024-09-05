package utils

import "time"

const (
	timeout = time.Minute * 0
)

func GetOutTime() time.Time {
	if timeout == 0 {
		return time.Time{}
	}
	return time.Now().Add(timeout)
}
