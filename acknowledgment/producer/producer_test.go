package producer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/acknowledgment/connection"
)

func TestPublish(t *testing.T) {
	broker, err := connection.Connect("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err.Error())
	}
	err = broker.SetQueue("ack")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Publish")
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		m := fmt.Sprintf("Message %d #%d", i, randomInt(5, 10))
		fmt.Printf("%d: %s \n", i, m)
		Publish(m, *broker)
	}
}

func randomInt(min, max int) int {
	tn := time.Now().Unix()
	rand.Seed(tn)
	return rand.Intn(max-min) + min
}