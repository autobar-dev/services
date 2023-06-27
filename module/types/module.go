package types

import "time"

type Module struct {
	Id           int32  `json:"id"`
	SerialNumber string `json:"serial_number"`

	StationSlug *string `json:"station_slug"`
	ProductSlug *string `json:"product_slug"`

	Prices map[string]int `json:"prices"`

	CreatedAt time.Time `json:"created_at"`
}

type CreateModuleResponse struct {
	Id           int32  `json:"id"`
	SerialNumber string `json:"serial_number"`

	StationSlug *string `json:"station_slug"`
	ProductSlug *string `json:"product_slug"`

	Prices map[string]int `json:"prices"`

	CreatedAt time.Time `json:"created_at"`

	PrivateKey string `json:"private_key"`
}

type ModuleSentReport struct {
	Status string `json:"status"`
}

type ModuleReport struct {
	Status       string `json:"status"`
	ResponseTime int    `json:"response_time"`
}
