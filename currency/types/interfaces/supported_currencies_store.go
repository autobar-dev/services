package interfaces

import "time"

type SupportedCurrenciesStoreRow struct {
	Id        uint32
	Code      string
	Name      string
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SupportedCurrenciesStore interface {
	Get(string) (*SupportedCurrenciesStoreRow, error)
	IsSupported(string) (bool, error)
	GetAll() (*[]SupportedCurrenciesStoreRow, error)
}
