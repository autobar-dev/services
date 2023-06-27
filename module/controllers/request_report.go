package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/types"
)

func RequestReportController(app_context *types.AppContext, serial_number string) (*types.ModuleReport, error) {
	rr := app_context.Repositories.Realtime

	amqp_channel := app_context.AmqpChannel

	queue, err := amqp_channel.QueueDeclare(
		"",
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	queue_name := queue.Name

	args := &types.RequestReportCommandArgs{
		Channel: queue.Name,
	}

	args_json_bytes, _ := json.Marshal(args)
	args_json := string(args_json_bytes)

	start_time := time.Now()

	err = rr.SendCommand(serial_number, repositories.ModuleServiceRealtimeClientType, types.RequestReportCommandName, args_json)
	if err != nil {
		_, _ = amqp_channel.QueueDelete(queue_name, false, false, true)
		return nil, err
	}

	messages, err := amqp_channel.Consume(
		queue_name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		_, _ = amqp_channel.QueueDelete(queue_name, false, false, true)
		return nil, err
	}

	timer_duration := time.Second * time.Duration(app_context.Config.ModuleReportTimeoutSeconds)
	timer := time.NewTimer(timer_duration)

	go func() {
		<-timer.C
		_, _ = amqp_channel.QueueDelete(queue_name, false, false, true)
	}()

	var received_message amqp.Delivery

	for message := range messages {
		received_message = message
		break
	}

	received_time := time.Now()

	_, _ = amqp_channel.QueueDelete(queue_name, false, false, true)

	var module_sent_report types.ModuleSentReport

	err = json.Unmarshal(received_message.Body, &module_sent_report)
	if err != nil {
		return nil, err
	}

	// time_delta :=
	// subtract start_time from received_time

}
