package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/enbis/learning-rabbitmq/global/utils"
	conn "github.com/enbis/learning-rabbitmq/round-robin-queuing/connection"
	"github.com/enbis/learning-rabbitmq/round-robin-queuing/consumer"
	models "github.com/enbis/learning-rabbitmq/round-robin-queuing/models"
	"github.com/enbis/learning-rabbitmq/round-robin-queuing/producer"
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

	// err = broker.SetExchange(configuration.TopicName)
	// if err != nil {
	// 	errorHandler("SetExchange", err.Error())
	// }

	err = broker.SetQueue(configuration.QueueName)
	if err != nil {
		errorHandler("SetQueue", err.Error())
	}

	// err = broker.SetBinding(configuration.TopicName, configuration.QueueName)
	// if err != nil {
	// 	errorHandler("SetBinding", err.Error())
	// }

	waiting := make(chan bool)
	for i := 0; i < 2; i++ {
		fmt.Println("Consumer #", i)
		msgs, err := consumer.ConsumeMsg(*broker, configuration.QueueName)
		if err != nil {
			errorHandler("ConsumeMsg", err.Error())
		}

		go func(x int) {
			for msg := range msgs {
				ss := strings.Split(string(msg.Body), "#")
				seconds, err := strconv.Atoi(ss[len(ss)-1])
				if err != nil {
					errorHandler("strconv", err.Error())
				}
				time.Sleep(time.Duration(seconds) * time.Second)
				fmt.Printf("message received on Consumer %d: %s \n", x, string(msg.Body))
			}
		}(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			m := fmt.Sprintf("Message #%d", utils.RandomInt(5, 10))
			fmt.Printf("%d: %s \n", i, m)
			producer.Publish(m, *broker)
		}
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
