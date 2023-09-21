package controllers

import (
	"errors"
	"fmt"

	"github.com/autobar-dev/services/realtime/repositories"
	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
)

func Reply(
	app_context types.AppContext,
	client_type types.ClientType,
	identifier string,
	message_id string,
) error {
	mr := app_context.Repositories.Mq
	sr := app_context.Repositories.State

	// Client message exchange
	exchange_name := utils.ExchangeNameFromClientInfo(client_type, identifier)

	listeners, err := sr.GetListenersCountForExchange(exchange_name)
	if err != nil {
		return err
	}

	if *listeners == 0 {
		return errors.New("no listeners connected")
	}

	// Reply exchange
	reply_exchange_name := utils.ReplyExchangeNameFromClientInfo(client_type, identifier)

	fmt.Printf("received reply for #%s\n", message_id)

	return mr.PublishReply(reply_exchange_name, &repositories.MqReply{
		Id: message_id,
	})
}
