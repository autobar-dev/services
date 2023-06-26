package utils

import (
	"math/rand"

	"go.a5r.dev/services/module/types"
)

func GenerateSerialNumber(length int32) string {
	bytes := make([]rune, length)

	for i := range bytes {
		bytes[i] = types.SerialNumberRunes[rand.Intn(len(types.SerialNumberRunes))]
	}

	return string(bytes)
}
