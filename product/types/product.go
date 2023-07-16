package types

import "time"

type Product struct {
	Id           string            `json:"id"`
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Cover        *string           `json:"cover"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
