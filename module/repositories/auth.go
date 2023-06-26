package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ServiceModuleRegisterRequestBody struct {
	SerialNumber string `json:"serial_number"`
}

type ServiceModuleRegisterResponse struct {
	Status string         `json:"status"`
	Error  *string        `json:"error"`
	Data   *ServiceModule `json:"data"`
}

type ServiceModule struct {
	Id           int       `json:"id"`
	SerialNumber string    `json:"serial_number"`
	PrivateKey   string    `json:"private_key"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthRepository struct {
	service_url string
}

func NewAuthRepository(service_url string) *AuthRepository {
	return &AuthRepository{
		service_url: service_url,
	}
}

func (ar AuthRepository) Create(serial_number string) (*ServiceModule, error) {
	url := fmt.Sprintf("%s/module/register", ar.service_url)

	body := &ServiceModuleRegisterRequestBody{
		SerialNumber: serial_number,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var smrr ServiceModuleRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&smrr); err != nil {
		return nil, err
	}

	if smrr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing module register response: %s", *smrr.Error))
	}

	return smrr.Data, nil
}
