package producer

import (
	"github.com/enbis/rabbit_test/connection"
	"github.com/streadway/amqp"
)

// producer is the program that sends messages
func Publish(body string, broker connection.Broker) error {
	//create a message to be sent to the queue.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	// Publish the message
	err := broker.Channel.Publish("", broker.Queue.Name, false, false, message)
	if err != nil {
		return err
	}

	return nil
}
