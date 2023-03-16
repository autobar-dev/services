package interfaces

import "time"

type RateStoreRow struct {
	Id                  uint32
	BaseCurrency        string
	DestinationCurrency string
	Rate                float32
	UpdatedAt           time.Time
}

type RateStore interface {
	GetRate(string, string) (*RateStoreRow, error)
	Upsert(string, string, float32) error
	Delete(string, string) error
}
