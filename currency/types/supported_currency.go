package types

import "time"

type SupportedCurrency struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}
