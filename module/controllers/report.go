package controllers

import (
	"context"
	"encoding/json"

	"github.com/autobar-dev/services/module/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ReportController(app_context *types.AppContext, queue_name string, msr types.ModuleSentReport) error {
	amqp_channel := app_context.AmqpChannel

	message_json_bytes, _ := json.Marshal(msr)

	ctx := context.Background()
	err := amqp_channel.PublishWithContext(
		ctx,
		"",         // exchange
		queue_name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "encoding/json",
			Body:        message_json_bytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
