package utils

import "testing"

func TestDateStringToTime(t *testing.T) {
	result, err := DateStringToTime("2023-08-19")
	if err != nil {
		t.Errorf("DateStringToTime() error = %v", err)
	}

	if result.Year() != 2023 && result.Month() != 8 && result.Day() != 19 {
		t.Errorf("DateStringToTime() = %v", result)
	}
}
