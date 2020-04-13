package consumer

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/acknowledgment/connection"
)

func TestConsumeMsg(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("ack")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Consuming")
	msgs, err := ConsumeMsg(*broker, "ack", "")
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for msg := range msgs {
			ss := strings.Split(string(msg.Body), "#")
			seconds, err := strconv.Atoi(ss[len(ss)-1])
			if err != nil {
				panic(err.Error())
			}
			time.Sleep(time.Duration(seconds) * time.Second)
			msg.Ack(false)
			fmt.Printf("message received and processed on Consumer: %s \n", string(msg.Body))
		}
	}()

	select {}
}
