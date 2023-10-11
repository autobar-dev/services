package types

import "time"

type DisplayUnit struct {
	Id                     int32     `json:"id"`
	Amount                 float64   `json:"amount"`
	Symbol                 string    `json:"symbol"`
	DivisorFromMillilitres float64   `json:"divisor_from_millilitres"`
	DecimalsDisplayed      int32     `json:"decimals_displayed"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
