package consumer

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/global/connection"
)

func TestConsumeMsgUnfair(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("unfair", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Consuming unfair")
	msgs, err := ConsumeMsg(*broker, "unfair", "")
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

func TestConsumeMsgFair(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("fair", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQos()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Consuming fair")
	msgs, err := ConsumeMsg(*broker, "fair", "")
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
