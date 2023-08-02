package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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

	message_id := uuid.New().String()

	message_bytes, err := json.Marshal(struct {
		Id      string `json:"id"`
		Command string `json:"command"`
		Args    string `json:"args"`
	}{
		Id:      message_id,
		Command: command_name,
		Args:    args,
	})
	if err != nil {
		return err
	}

	// Declare replies exchange
	reply_queue, err := mr.CreatePubSub(utils.ReplyExchangeNameFromClientInfo(client_type, identifier))
	if err != nil {
		return err
	}

	message_consumer_name := utils.MessageConsumerName(message_id)
	replies, err := mr.Consume(*reply_queue, message_consumer_name)
	if err != nil {
		return err
	}

	fmt.Printf("listening for reply for #%s on queue %s\n", message_id, *reply_queue)

	err = mr.Publish(exchange_name, string(message_bytes))
	if err != nil {
		_ = mr.CancelConsumer(message_consumer_name)
		return err
	}

	// Wait for reply
	for true {
		select {
		case reply_delivery, ok := <-replies:
			if ok == false {
				_ = mr.CancelConsumer(message_consumer_name)
				return errors.New("replies channel closed")
			}

			var reply types.Reply
			err = json.Unmarshal(reply_delivery.Body, &reply)
			if err != nil {
				_ = mr.CancelConsumer(message_consumer_name)
				return err
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

	// should never reach this
	return nil
}
