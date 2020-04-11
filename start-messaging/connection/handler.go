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

func (broker *Broker) SetQueue(queuename string) error {
	queue, err := broker.Channel.QueueDeclare(queuename, false, false, false, false, nil)

	if err != nil {
		return err
	}

	broker.Queue = queue
	return nil

}
