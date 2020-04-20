package producer

import (
	"github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/streadway/amqp"
)

// producer is the program that sends messages
func Publish(body string, broker connection.Broker, exchangeName, routingKey string) error {
	//create a message to be sent to the queue.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	// Publish the message
	err := broker.Channel.Publish(exchangeName, routingKey, false, false, message)
	if err != nil {
		return err
	}

	return nil
}
