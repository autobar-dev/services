package types

import "time"

type ActivationSession struct {
	Id                string    `json:"id"`
	UserId            string    `json:"user_id"`
	SerialNumber      string    `json:"serial_number"`
	ProductId         string    `json:"product_id"`
	Price             int       `json:"price"`
	AmountMillilitres int       `json:"amount_millilitres"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
