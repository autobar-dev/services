package utils

import (
	"fmt"
	"regexp"
)

func ValidateFirstName(fn string) bool {
	return len(fn) > 0
}

func ValidateLastName(ln string) bool {
	return len(ln) > 0
}

func ValidateDateOfBirth(dob string) bool {
	res, err := regexp.Match(`^\d{4}-([0]\d|1[0-2])-([0-2]\d|3[01])$`, []byte(dob))
	if err != nil {
		fmt.Printf("ERROR: incorrect dob regexp: %+v\n", err)
	}

	return res
}

func ValidateNationality(nat string) bool {
	return true
}

func ValidateLocale(loc string) bool {
	return true
}
