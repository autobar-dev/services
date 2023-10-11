package types

import "time"

type Module struct {
	Id           int32          `json:"id"`
	SerialNumber string         `json:"serial_number"`
	StationId    *string        `json:"station_id"`
	ProductId    *string        `json:"product_id"`
	Enabled      bool           `json:"enabled"`
	Prices       map[string]int `json:"prices"`
	DisplayUnit  DisplayUnit    `json:"display_unit"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type CreateModuleResponse struct {
	Module     *Module `json:"module"`
	PrivateKey string  `json:"private_key"`
}

type ModuleSentReport struct {
	Status string `json:"status"`
}

type ModuleReport struct {
	Status       string `json:"status"`
	ResponseTime int    `json:"response_time"`
}
