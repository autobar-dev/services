package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceRealtimeClientType string

const (
	ModuleServiceRealtimeClientType ServiceRealtimeClientType = "module"
	UserServiceRealtimeClientType   ServiceRealtimeClientType = "user"
)

type CommandName string // command names and args defined in `types/realtime_commands`

type ServiceCommandRequestBody struct {
	ClientType ServiceRealtimeClientType `json:"client_type"`
	Identifier string                    `json:"identifier"`
	Command    CommandName               `json:"command"`
	Args       string                    `json:"args"`
}

type RealtimeRepository struct {
	service_url string
}

func NewRealtimeRepository(url string) *RealtimeRepository {
	return &RealtimeRepository{
		service_url: url,
	}
}

func (rr RealtimeRepository) SendCommand(identifier string, srct ServiceRealtimeClientType, command CommandName, args string) error {
	url := fmt.Sprintf("%s/send-command", rr.service_url)

	body := &ServiceCommandRequestBody{
		ClientType: srct,
		Identifier: identifier,
		Command:    command,
		Args:       args,
	}
	body_json, _ := json.Marshal(body)

	_, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		return err
	}

	return nil
}
