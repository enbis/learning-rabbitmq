package consumer

import (
	conn "github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/streadway/amqp"
)

func ConsumeMsg(broker conn.Broker, queuename string, consumername string) (<-chan amqp.Delivery, error) {
	msgs, err := broker.Channel.Consume(broker.Queue.Name, consumername, true, false, false, false, nil)
	return msgs, err
}
