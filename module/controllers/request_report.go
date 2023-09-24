package controllers

import (
	"encoding/json"
	"time"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
	amqp "github.com/rabbitmq/amqp091-go"
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

	args := &types.RequestReportCommandArgs{
		Queue: queue.Name,
	}

	args_map, err := utils.StructToJsonMap(args)
	if err != nil {
		return nil, err
	}

	start_time := time.Now()

	err = rr.SendCommand(
		serial_number,
		repositories.ModuleServiceRealtimeClientType,
		types.RequestReportCommandName,
		args_map,
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

	time_delta := received_time.Sub(start_time)

	module_report_response := &types.ModuleReport{
		Status:       module_sent_report.Status,
		ResponseTime: int(time_delta.Milliseconds()),
	}
	return module_report_response, nil
}
