package connection

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Broker struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func Connect(connString string) (*Broker, error) {

	//Dial accepts a string in the AMQP URI format and returns a new Connection over TCP using PlainAuth.
	fmt.Println("Connection string ", connString)
	connection, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}

	//Channel opens a unique, concurrent server channel to process the bulk of AMQP messages.
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &Broker{Channel: channel, Connection: connection}, nil
}

func (broker *Broker) Close() {
	broker.Connection.Close()
	broker.Channel.Close()
}

func (broker *Broker) SetExchange(topicname string) error {

	fmt.Println()
	//ExchangeDeclare declares an exchange on the server. If the exchange does not already exist, the server will create it.
	//The common types are "direct", "fanout", "topic" and "headers".
	//Durable and Non-Auto-Deleted exchanges will survive server restarts and remain declared when there are no remaining bindings. This is the best lifetime for long-lived exchange configurations
	//Non-Durable and Non-Auto-deleted exchanges will remain as long as the server is running including when there are no remaining bindings.
	//Exchanges declared as `internal` do not accept accept publishings. Internal exchanges are useful when you wish to implement inter-exchange topologies that should not be exposed to users of the broker.
	//When noWait is true, declare without waiting for a confirmation from the server.

	err := broker.Channel.ExchangeDeclare(topicname, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil

}

func (broker *Broker) SetQueue(queuename string) error {
	queue, err := broker.Channel.QueueDeclare(queuename, false, false, false, false, nil)

	if err != nil {
		return err
	}

	broker.Queue = queue
	return nil

}

func (broker *Broker) SetBinding(topicname, queuename string) error {
	err := broker.Channel.QueueBind(queuename, "#", topicname, false, nil)

	if err != nil {
		return err
	}
	return nil
}
