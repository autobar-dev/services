package utils

import "crypto/rand"

const LowercaseUppercaseNumbersSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int, character_set string) string {
	ll := len(character_set)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes

	for i := 0; i < length; i++ {
		b[i] = character_set[int(b[i])%ll]
	}

	return string(b)
}
