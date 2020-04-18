package producer

import (
	"fmt"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/global/connection"
)

func TestPublishFair(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("fair")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Publish")
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		m := ""
		if i%2 == 0 {
			m = fmt.Sprintf("Message %d #%d", i, 8)
		} else {
			m = fmt.Sprintf("Message %d #%d", i, 3)
		}
		fmt.Printf("%d: %s \n", i, m)
		Publish(m, *broker)
	}
}
func TestPublishUnfair(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("unfair")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Publish")
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		m := ""
		if i%2 == 0 {
			m = fmt.Sprintf("Message %d #%d", i, 8)
		} else {
			m = fmt.Sprintf("Message %d #%d", i, 3)
		}
		fmt.Printf("%d: %s \n", i, m)
		Publish(m, *broker)
	}
}
