package controllers

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.a5r.dev/services/module/types"
)

func ReportController(app_context *types.AppContext, queue_name string, msr types.ModuleSentReport) error {
	amqp_channel := app_context.AmqpChannel

	message_json_bytes, _ := json.Marshal(msr)

	err := amqp_channel.Publish(
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
