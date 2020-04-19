package consumer

import (
	"flag"
	"fmt"
	"testing"

	"github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/enbis/learning-rabbitmq/global/models"
	"github.com/enbis/learning-rabbitmq/global/utils"
	"github.com/spf13/viper"
)

var configuration *models.Configurations
var qb = flag.Int("qbinding", -1, "binding selection")

func TestLaunchConsumer(t *testing.T) {
	initConf()

	fmt.Println("qbinding ", *qb)

	if *qb == -1 {
		panic("qbinding flag not setted")
	}

	broker, err := connection.Connect(configuration.ConnString)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetExchange(configuration.ExchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	qname := fmt.Sprintf("%s.%s", configuration.QueueName, utils.RandomStr())
	err = broker.SetQueue(qname, false, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	if *qb == 0 {
		fmt.Println("Binding to all and first")
		for i := 0; i < 2; i++ {
			err = broker.SetBinding(configuration.ExchangeName, models.IntToString[i], qname, false, nil)
			if err != nil {
				panic(err.Error())
			}
		}
	} else {
		fmt.Println("Binding to all and second")
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue
			}
			err = broker.SetBinding(configuration.ExchangeName, models.IntToString[i], qname, false, nil)
			if err != nil {
				panic(err.Error())
			}
		}
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
	fmt.Println("")
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
