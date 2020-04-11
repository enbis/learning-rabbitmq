package main

import (
	"fmt"
	"time"

	conn "github.com/enbis/rabbit_test/connection"
	"github.com/enbis/rabbit_test/consumer"
	models "github.com/enbis/rabbit_test/models"
	"github.com/enbis/rabbit_test/producer"
	"github.com/spf13/viper"
)

var configuration models.Configurations

func main() {

	Init()
	broker, err := conn.Connect(configuration.ConnString)
	if err != nil {
		errorHandler("Connect", err.Error())
	}
	defer broker.Close()

	err = broker.SetQueue(configuration.QueueName)
	if err != nil {
		errorHandler("SetQueue", err.Error())
	}

	msgs, err := consumer.ConsumeMsg(*broker, configuration.QueueName)
	if err != nil {
		errorHandler("ConsumeMsg", err.Error())
	}

	waiting := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Println("message received: " + string(msg.Body))
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(1000 * time.Millisecond)
			m := fmt.Sprintf("Message #%d", i)
			producer.Publish(m, *broker)
		}
		waiting <- true
	}()

	fmt.Println("Waiting for messages")
	<-waiting
	fmt.Println("Closing")

}

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("connstring", "amqp://guest:guest@localhost:5672")
	viper.SetDefault("topicname", "events")
	viper.SetDefault("queuename", "testqueue")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

}

func errorHandler(funcdef, err string) {
	e := fmt.Sprintf("Func %s, error %s", funcdef, err)
	panic(e)
}
