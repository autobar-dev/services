package utils

import "testing"

func TestRandomString(t *testing.T) {
	result := RandomString(10, LowercaseUppercaseNumbersSet)
	if len(result) != 10 {
		t.Errorf("Expected length of 10, got %d", len(result))
	}
}
