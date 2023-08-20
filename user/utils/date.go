package utils

import "time"

func DateStringToTime(date_string string) (time.Time, error) {
	return time.Parse("2006-12-29", date_string)
}
