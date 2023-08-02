package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

func Reply(app_context types.AppContext, client_type types.ClientType, identifier string, message_id string) error {
	mr := app_context.Repositories.Mq
	rr := app_context.Repositories.Redis

	// Client message exchange
	exchange_name := utils.ExchangeNameFromClientInfo(client_type, identifier)

	listeners, err := rr.GetListenersCountForExchange(exchange_name)
	if err != nil {
		return err
	}

	if *listeners == 0 {
		return errors.New("no listeners connected")
	}

	// Reply exchange
	reply_exchange_name := utils.ReplyExchangeNameFromClientInfo(client_type, identifier)
	reply := &types.Reply{
		Id: message_id,
	}
	reply_json, _ := json.Marshal(reply)

	fmt.Printf("received reply for #%s\n", reply.Id)

	return mr.Publish(reply_exchange_name, string(reply_json))
}
