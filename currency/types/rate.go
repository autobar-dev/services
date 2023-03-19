package types

import "time"

type Rate struct {
	BaseCurrency        string    `json:"base"`
	DestinationCurrency string    `json:"destination"`
	Rate                float64   `json:"rate"`
	UpdatedAt           time.Time `json:"updated_at"`
}
