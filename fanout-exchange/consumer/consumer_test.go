package consumer

import (
	"fmt"
	"testing"

	"github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/enbis/learning-rabbitmq/global/models"
	"github.com/enbis/learning-rabbitmq/global/utils"
	"github.com/spf13/viper"
)

var configuration *models.Configurations

func TestLaunchConsumer(t *testing.T) {
	initConf()

	broker, err := connection.Connect(configuration.ConnString)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetExchange(configuration.ExchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	qname := fmt.Sprintf("%s.%s", configuration.QueueName, utils.RandomStr())
	err = broker.SetQueue(qname, false, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetBinding(configuration.ExchangeName, "", qname, false, nil)
	if err != nil {
		panic(err.Error())
	}

	msgs, err := ConsumeMsg(*broker, configuration.QueueName, "")
	if err != nil {
		panic(err.Error())
	}
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Println("message received: " + string(msg.Body))
		}
	}()

	fmt.Println("Waiting messages")
	<-forever
}

func initConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../global/config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("connstring", "")
	viper.SetDefault("queuename", "")

	configuration = &models.Configurations{}
	err = viper.Unmarshal(configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

}
