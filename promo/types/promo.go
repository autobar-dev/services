package types

import "time"

type PromoType string

const (
	FixedType      PromoType = "FIXED"
	PercentageType PromoType = "PERCENTAGE"
)

type Promo struct {
	Id        string    `json:"id"`
	Type      PromoType `json:"type"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
