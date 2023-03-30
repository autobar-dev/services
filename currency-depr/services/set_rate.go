package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func SetRate(base_code string, dest_code string, rate float64, rs *interfaces.RateStore) error {
	return (*rs).Upsert(base_code, dest_code, rate)
}
