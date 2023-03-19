package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func GetRate(base_code string, dest_code string, rs *interfaces.RateStore) (*interfaces.RateStoreRow, error) {
	return (*rs).GetRate(base_code, dest_code)
}
