package utils

import (
	"math/rand"
	"reflect"
)

var RefreshTokenCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(length int, character_set []rune) string {
	result := make([]rune, length)

	for i := range result {
		result[i] = character_set[rand.Intn(len(character_set))]
	}

	return string(result)
}

func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	if item == nil {
		return res
	}

	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field)
			} else {
				res[tag] = field
			}
		}
	}

	return res
}
