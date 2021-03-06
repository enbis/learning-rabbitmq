package producer

import (
	"github.com/enbis/learning-rabbitmq/durability/connection"
	"github.com/streadway/amqp"
)

// producer is the program that sends messages
func Publish(body string, broker connection.Broker) error {
	//create a message to be sent to the queue.
	message := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	}

	// Publish the message
	err := broker.Channel.Publish("", broker.Queue.Name, false, false, message)
	if err != nil {
		return err
	}

	return nil
}
