package utils

import "time"

func DateStringToTime(date_string string) (time.Time, error) {
	return time.Parse("2006-01-02", date_string)
}

func TimeToDateString(date time.Time) string {
	return date.Format("2006-01-02")
}
