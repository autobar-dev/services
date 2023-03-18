package interfaces

import (
	"time"

	"github.com/autobar-dev/services/currency/types/inputs"
)

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
	GetAll() (*[]SupportedCurrenciesStoreRow, error)
	SetEnabled(string, bool) error
	Insert(*inputs.Currency) error
	Delete(string) error
}
