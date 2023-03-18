package types

import "time"

type SupportedCurrency struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	UpdatedAt time.Time `json:"updated_at"`
}
