package controllers

import (
	"encoding/json"
	"errors"

	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

func SendCommandMessage(app_context types.AppContext, client_type types.ClientType, identifier string, command_name string, args string) error {
	mr := app_context.Repositories.Mq
	rr := app_context.Repositories.Redis

	exchange_name := utils.ExchangeNameFromClientInfo(client_type, identifier)

	listeners, err := rr.GetListenersCountForExchange(exchange_name)
	if err != nil {
		return err
	}

	if *listeners == 0 {
		return errors.New("no listeners connected")
	}

	message_bytes, err := json.Marshal(struct {
		MessageType types.MqMessageType `json:"type"`
		Command     string              `json:"command"`
		Args        string              `json:"args"`
	}{
		MessageType: types.CommandMessageType,
		Command:     command_name,
		Args:        args,
	})
	if err != nil {
		return err
	}

	return mr.Publish(exchange_name, string(message_bytes))
}
