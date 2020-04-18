package producer

import (
	"fmt"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/durability/connection"
	"github.com/enbis/learning-rabbitmq/global/utils"
)

func TestPublish(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("durability")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Publish")
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		m := fmt.Sprintf("Message %d #%d", i, utils.RandomInt(5, 10))
		fmt.Printf("%d: %s \n", i, m)
		Publish(m, *broker)
	}
}
