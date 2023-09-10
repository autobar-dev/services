package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ServiceCurrencyResponse struct {
	Status string           `json:"status"`
	Error  *string          `json:"error"`
	Data   *ServiceCurrency `json:"data"`
}

type ServiceCurrency struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ServiceRateResponse struct {
	Status string       `json:"status"`
	Error  *string      `json:"error"`
	Data   *ServiceRate `json:"data"`
}

type ServiceRate struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Rate      float64   `json:"rate"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CurrencyRepository struct {
	service_url string
}

func NewCurrencyRepository(service_url string) *CurrencyRepository {
	return &CurrencyRepository{
		service_url: service_url,
	}
}

func (cr CurrencyRepository) GetCurrency(currency_code string) (*ServiceCurrency, error) {
	url := fmt.Sprintf("%s/currency/?code=%s", cr.service_url, currency_code)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var scr ServiceCurrencyResponse
	if err := json.NewDecoder(response.Body).Decode(&scr); err != nil {
		return nil, err
	}

	if scr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing currency response: %s", *scr.Error))
	}

	return scr.Data, nil
}

func (cr CurrencyRepository) GetRate(from string, to string) (*ServiceRate, error) {
	url := fmt.Sprintf("%s/rate/?from=%s&to=%s", cr.service_url, from, to)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var srr ServiceRateResponse
	if err := json.NewDecoder(response.Body).Decode(&srr); err != nil {
		return nil, err
	}

	if srr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing rate response: %s", *srr.Error))
	}

	return srr.Data, nil
}
