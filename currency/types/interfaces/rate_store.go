package interfaces

import "time"

type RateStoreRow struct {
	Id                  uint32    `json:"id"`
	BaseCurrency        string    `json:"base_currency"`
	DestinationCurrency string    `json:"destination_currency"`
	Rate                float32   `json:"rate"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type RateStore interface {
	GetRate(string, string) (*RateStoreRow, error)
	Upsert(string, string, float32) error
	Delete(string, string) error
}
