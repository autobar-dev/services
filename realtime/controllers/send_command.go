package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/autobar-dev/services/realtime/repositories"
	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
	"github.com/google/uuid"
)

func SendCommandMessage(
	app_context types.AppContext,
	client_type types.ClientType,
	identifier string,
	command_name string,
	args map[string]interface{},
) error {
	mr := app_context.Repositories.Mq
	rr := app_context.Repositories.State

	exchange_name := utils.ExchangeNameFromClientInfo(client_type, identifier)

	listeners, err := rr.GetListenersCountForExchange(exchange_name)
	if err != nil {
		return err
	}

	if *listeners == 0 {
		return errors.New("no listeners connected")
	}

	message_id := uuid.New().String()

	// Declare replies exchange
	reply_queue, err := mr.CreatePubSub(utils.ReplyExchangeNameFromClientInfo(client_type, identifier))
	if err != nil {
		return err
	}

	message_consumer_name := utils.MessageConsumerName(message_id)
	replies, err := mr.ConsumeReplies(*reply_queue, message_consumer_name)
	if err != nil {
		return err
	}

	fmt.Printf("listening for reply for #%s on queue %s\n", message_id, *reply_queue)

	err = mr.PublishCommand(exchange_name, &repositories.MqCommand{
		Id:      message_id,
		Command: command_name,
		Args:    args,
	})
	if err != nil {
		_ = mr.CancelConsumer(message_consumer_name)
		return err
	}

	// Wait for reply
	for {
		select {
		case reply, ok := <-replies:
			if !ok {
				_ = mr.CancelConsumer(message_consumer_name)
				return errors.New("replies channel closed")
			}

			fmt.Printf("received reply for %s\n", reply.Id)

			if reply.Id == message_id {
				_ = mr.CancelConsumer(message_consumer_name)
				return nil
			}
		case <-time.After(time.Second * time.Duration(types.SendTimeoutSeconds)):
			_ = mr.CancelConsumer(message_consumer_name)
			return errors.New("timeout")
		}
	}
}
