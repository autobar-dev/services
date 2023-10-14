package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MqRepository struct{}

type MqCommand struct {
	Id      string                 `json:"id"`
	Command string                 `json:"command"`
	Args    map[string]interface{} `json:"args"`
}

type MqReply struct {
	Id string `json:"id"`
}

func NewMqRepository() *MqRepository {
	return &MqRepository{}
}

func (mr MqRepository) CreatePubSub(channel *amqp.Channel, exchange_name string) (*string, error) {
	// Declare fanout exchange for a specific client (no-op if exists)
	err := channel.ExchangeDeclare(exchange_name, "fanout", false, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	// Create a listener queue
	queue, err := channel.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		return nil, err
	}

	// Bind it to the client-specific exchange
	err = channel.QueueBind(queue.Name, "#", exchange_name, false, nil)
	if err != nil {
		return nil, err
	}

	return &queue.Name, nil
}

func (mr MqRepository) ConsumeCommands(
	channel *amqp.Channel,
	queue_name string,
	consumer_name string,
) (<-chan MqCommand, error) {
	messages, err := channel.Consume(queue_name, consumer_name, true, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	parsed_messages := make(chan MqCommand)

	go func() {
		for message := range messages {
			var parsed_message MqCommand
			err := json.Unmarshal(message.Body, &parsed_message)
			if err != nil {
				fmt.Printf("failed to parse message: %s\n", err)
				continue
			}

			parsed_messages <- parsed_message
		}

		close(parsed_messages)
	}()

	return parsed_messages, nil
}

func (mr MqRepository) ConsumeReplies(channel *amqp.Channel, queue_name string, consumer_name string) (<-chan MqReply, error) {
	messages, err := channel.Consume(queue_name, consumer_name, true, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	parsed_messages := make(chan MqReply)

	go func() {
		for message := range messages {
			var parsed_message MqReply
			err := json.Unmarshal(message.Body, &parsed_message)
			if err != nil {
				fmt.Printf("failed to parse message: %s\n", err)
				continue
			}

			parsed_messages <- parsed_message
		}

		close(parsed_messages)
	}()

	return parsed_messages, nil
}

func (mr MqRepository) CancelConsumer(channel *amqp.Channel, consumer_name string) error {
	return channel.Cancel(consumer_name, false)
}

func (mr MqRepository) PublishCommand(channel *amqp.Channel, exchange_name string, message *MqCommand) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(exchange_name, "fanout", false, true, false, false, nil)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = channel.PublishWithContext(ctx, exchange_name, "", false, false, amqp.Publishing{
		ContentType: "encoding/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (mr MqRepository) PublishReply(channel *amqp.Channel, exchange_name string, message *MqReply) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(exchange_name, "fanout", false, true, false, false, nil)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = channel.PublishWithContext(ctx, exchange_name, "", false, false, amqp.Publishing{
		ContentType: "encoding/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}
