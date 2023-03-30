package types

import "time"

// SupportedCurrency is a currency that is supported by the service.
//
// swagger:model SupportedCurrency
type SupportedCurrency struct {
	// 3-letter ISO 4217 currency code
	// example: CAD
	Code string `json:"code"`

	// Descriptive name
	// example: Canadian Dollar
	Name string `json:"name"`

	// Whether the currency is enabled
	// example: true
	Enabled bool `json:"enabled"`

	// When the currency was last enabled/disabled
	// example: 2019-01-01T12:34:56Z
	UpdatedAt time.Time `json:"updated_at"`
}
