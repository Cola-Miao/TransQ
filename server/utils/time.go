package utils

import "time"

func GetOutTime(scd int) time.Time {
	if scd == 0 {
		return time.Time{}
	}
	return time.Now().Add(time.Second * time.Duration(scd))
}
