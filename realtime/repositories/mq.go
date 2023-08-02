package repositories

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type MqRepository struct {
	channel *amqp.Channel
}

func NewMqRepository(channel *amqp.Channel) *MqRepository {
	return &MqRepository{
		channel,
	}
}

func (mr MqRepository) CreatePubSub(exchange_name string) (*string, error) {
	// Declare fanout exchange for a specific client (no-op if exists)
	err := mr.channel.ExchangeDeclare(exchange_name, "fanout", false, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	// Create a listener queue
	queue, err := mr.channel.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		return nil, err
	}

	// Bind it to the client-specific exchange
	err = mr.channel.QueueBind(queue.Name, "#", exchange_name, false, nil)
	if err != nil {
		return nil, err
	}

	return &queue.Name, nil
}

func (mr MqRepository) Consume(queue_name string, consumer_name string) (<-chan amqp.Delivery, error) {
	messages, err := mr.channel.Consume(queue_name, consumer_name, true, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (mr MqRepository) CancelConsumer(consumer_name string) error {
	return mr.channel.Cancel(consumer_name, false)
}

func (mr MqRepository) Publish(exchange_name string, body string) error {
	err := mr.channel.ExchangeDeclare(exchange_name, "fanout", false, true, false, false, nil)
	if err != nil {
		return err
	}

	err = mr.channel.Publish(exchange_name, "", false, false, amqp.Publishing{
		ContentType: "encoding/json",
		Body:        []byte(body),
	})
	if err != nil {
		return err
	}

	return nil
}
