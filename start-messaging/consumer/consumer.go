package consumer

import (
	conn "github.com/enbis/rabbit_test/connection"
	"github.com/streadway/amqp"
)

func ConsumeMsg(broker conn.Broker, queuename string) (<-chan amqp.Delivery, error) {
	msgs, err := broker.Channel.Consume(broker.Queue.Name, "", true, false, false, false, nil)
	return msgs, err
}
