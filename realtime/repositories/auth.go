package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
	http_client *http.Client
}

func NewAuthRepository(service_url string) *AuthRepository {
	client := &http.Client{}

	return &AuthRepository{
		service_url: service_url,
		http_client: client,
	}
}

func (ar AuthRepository) VerifySession(session string) (*ServiceSessionData, error) {
	url := fmt.Sprintf("%s/session/verify?session_id=%s", ar.service_url, session)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", "realtime")

	response, err := ar.http_client.Do(req)
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
