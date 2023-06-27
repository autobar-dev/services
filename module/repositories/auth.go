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

type ServiceVerifySessionResponse struct {
	Status string              `json:"status"`
	Error  *string             `json:"error"`
	Data   *ServiceSessionData `json:"data"`
}

type ServiceAuthClientType string

const (
	ModuleServiceAuthClientType ServiceAuthClientType = "module"
	UserServiceAuthClientType   ServiceAuthClientType = "user"
)

type ServiceSessionData struct {
	ClientIdentifier string                `json:"client_identifier"`
	ClientType       ServiceAuthClientType `json:"client_type"`
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

func (ar AuthRepository) VerifySession(session string) (*ServiceSessionData, error) {
	url := fmt.Sprintf("%s/session/verify?session_id=%s", ar.service_url, session)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var svsr ServiceVerifySessionResponse
	if err := json.NewDecoder(response.Body).Decode(&svsr); err != nil {
		return nil, err
	}

	if svsr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing session verify response: %s", *svsr.Error))
	}

	return svsr.Data, nil
}
